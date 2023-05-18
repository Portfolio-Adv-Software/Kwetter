package grpc

import (
	"fmt"
	pbuser "github.com/Portfolio-Adv-Software/Kwetter/AccountService/proto"
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
)

type UserServiceServer struct {
	pbuser.UnimplementedUserServiceServer
}

func (u UserServiceServer) CreateUser(ctx context.Context, req *pbuser.CreateUserReq) (*pbuser.CreateUserRes, error) {
	data := req.GetUser()

	user := &pbuser.User{
		UserID:   data.GetUserID(),
		Username: data.GetUsername(),
		Email:    data.GetEmail(),
		Password: data.GetPassword(),
	}

	_, err := accountdb.InsertOne(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error inserting account into db: %v", err))
	}
	res := &pbuser.CreateUserRes{User: user}
	return res, nil
}

func (u UserServiceServer) GetUser(ctx context.Context, req *pbuser.GetUserReq) (*pbuser.GetUserRes, error) {
	userID := req.GetUserID()
	data := &pbuser.User{}
	err := accountdb.FindOne(ctx, bson.M{"userID": userID}).Decode(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error finding user: %v", err))
	}

	user := &pbuser.User{
		UserID:   data.GetUserID(),
		Username: data.GetUsername(),
		Email:    data.GetEmail(),
		Password: data.GetPassword(),
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
		"username": user.GetUsername(),
		"email":    user.GetEmail(),
		"password": user.GetPassword(),
	}}
	err := accountdb.FindOneAndUpdate(ctx, user.GetUserID(), update).Decode(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error updating user: %v", err))
	}
	updatedUser := &pbuser.User{
		UserID:   data.GetUserID(),
		Username: data.GetUsername(),
		Email:    data.GetEmail(),
		Password: data.GetPassword(),
	}
	res := &pbuser.UpdateUserRes{User: updatedUser}
	return res, nil
}

func (u UserServiceServer) DeleteUser(ctx context.Context, req *pbuser.DeleteUserReq) (*pbuser.DeleteUserRes, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

var db *mongo.Client
var accountdb *mongo.Collection
var mongoCtx context.Context

var mongoUser = os.Getenv("MONGO_USERNAME")
var mongoPwd = os.Getenv("MONGO_PASSWORD")
var dbconn = "mongodb+srv://" + mongoUser + ":" + mongoPwd + "@kwetter.vduy1tl.mongodb.net/test"

func InitGRPC() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port: 50054")

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
	accountdb = db.Database("AccountTest").Collection("kwetterAccounts")

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
