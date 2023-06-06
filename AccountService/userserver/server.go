package userserver

import (
	"fmt"
	pbuser "github.com/Portfolio-Adv-Software/Kwetter/AccountService/proto"
	"github.com/Portfolio-Adv-Software/Kwetter/AccountService/rabbitmq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"
)

type UserServiceServer struct {
	pbuser.UnimplementedUserServiceServer
}

func (u UserServiceServer) GetAllUserData(ctx context.Context, req *pbuser.GetAllUserDataReq) (*pbuser.GetAllUserDataRes, error) {
	//TODO implement me
	panic("implement me")
}
func (u UserServiceServer) CreateUser(ctx context.Context, req *pbuser.CreateUserReq) (*pbuser.CreateUserRes, error) {
	data := req.GetUser()

	user := &pbuser.User{
		UserID:   data.GetUserID(),
		Email:    data.GetEmail(),
		Password: data.GetPassword(),
		Username: data.GetUsername(),
	}

	_, err := accountdb.InsertOne(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error inserting account into db: %v", err))
	}
	res := &pbuser.CreateUserRes{User: user}
	return res, nil
}

func (u UserServiceServer) GetUser(ctx context.Context, req *pbuser.GetUserReq) (*pbuser.GetUserRes, error) {
	userID := req.GetUserid()
	data := &pbuser.User{}
	err := accountdb.FindOne(ctx, bson.M{"userid": userID}).Decode(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error finding user: %v", err))
	}

	user := &pbuser.User{
		UserID:   data.GetUserID(),
		Email:    data.GetEmail(),
		Password: data.GetPassword(),
		Username: data.GetUsername(),
	}
	res := &pbuser.GetUserRes{User: user}
	return res, nil
}

func (u UserServiceServer) GetALlUsers(ctx context.Context, _ *pbuser.GetAllUsersReq) (*pbuser.GetAllUsersRes, error) {
	cursor, err := accountdb.Find(ctx, bson.M{})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error finding users: %v", err))
	}
	defer cursor.Close(ctx)

	var users []*pbuser.User
	for cursor.Next(ctx) {
		data := &pbuser.User{}
		err := cursor.Decode(data)
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("error decoding data: %v", err))
		}
		users = append(users, data)
	}
	res := &pbuser.GetAllUsersRes{User: users}
	return res, nil
}

func (u UserServiceServer) UpdateUser(ctx context.Context, req *pbuser.UpdateUserReq) (*pbuser.UpdateUserRes, error) {
	user := req.GetUser()
	data := &pbuser.User{}
	update := bson.M{"$set": bson.M{
		"email":    user.GetEmail(),
		"password": user.GetPassword(),
		"username": user.GetUsername(),
	}}
	err := accountdb.FindOneAndUpdate(ctx, user.GetUserID(), update).Decode(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error updating user: %v", err))
	}
	updatedUser := &pbuser.User{
		UserID:   data.GetUserID(),
		Email:    data.GetEmail(),
		Password: data.GetPassword(),
		Username: data.GetUsername(),
	}
	res := &pbuser.UpdateUserRes{User: updatedUser}
	return res, nil
}

func (u UserServiceServer) DeleteUser(ctx context.Context, req *pbuser.DeleteUserReq) (*pbuser.DeleteUserRes, error) {
	rabbitmq.SendDeleteGDPRUser(req.GetUserId())
	filter := bson.M{"userid": req.GetUserId()}
	maxRetries := 3
	retryCount := 0
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	for retryCount < maxRetries {
		count, err := accountdb.CountDocuments(ctx, filter)
		if err != nil {
			return nil, err
		}
		if count == 0 {
			res := &pbuser.DeleteUserRes{Status: "No documents found to delete"}
			return res, nil
		}
		deleteResult, err := accountdb.DeleteMany(ctx, filter)
		if err != nil {
			return nil, err
		}
		if deleteResult.DeletedCount == count {
			res := &pbuser.DeleteUserRes{Status: "All found documents deleted"}
			return res, nil
		}
		retryCount++
	}
	res := &pbuser.DeleteUserRes{Status: "User data deleted successfully"}
	return res, nil
}

var db *mongo.Client
var accountdb *mongo.Collection
var mongoCtx context.Context

func InitGRPC(wg *sync.WaitGroup) {
	defer wg.Done()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port: 50054")

	var mongoUser = os.Getenv("MONGO_USERNAME")
	var mongoPwd = os.Getenv("MONGO_PASSWORD")
	var dbconn = "mongodb+srv://" + mongoUser + ":" + mongoPwd + "@kwetter.vduy1tl.mongodb.net/test"

	listener, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("Unable to listen on port :50054: %v", err)
	}

	// Set options, here we can configure things like TLS support
	var opts []grpc.ServerOption
	// Create new gRPC server with (blank) options
	s := grpc.NewServer(opts...)
	srv := &UserServiceServer{}
	// Register the service with the server
	pbuser.RegisterUserServiceServer(s, srv)
	reflection.Register(s)

	// Initialize MongoDb client
	fmt.Println("Connecting to MongoDB...")

	// non-nil empty context
	mongoCtx = context.Background()

	// Connect takes in a context and options, the connection URI is the only option we pass for now
	db, err = mongo.Connect(mongoCtx, options.Client().ApplyURI(dbconn))
	// Handle potential errors
	if err != nil {
		log.Fatal(err)
	}

	// Check whether the connection was succesful by pinging the MongoDB server
	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to Mongodb")
	}

	// Bind our collection to our global variable for use in other methods
	accountdb = db.Database("AccountTest").Collection("KwetterAccounts")

	// Start the server in a child routine
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :50054")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	// Block main routine until a signal is received
	// As long as user doesn't press CTRL+C a message is not passed and our main routine keeps running
	<-c
	// After receiving CTRL+C Properly stop the server
	fmt.Println("\nStopping the server...")
	s.Stop()
	listener.Close()
	fmt.Println("Closing MongoDB connection")
	db.Disconnect(mongoCtx)
	fmt.Println("Done.")
}
