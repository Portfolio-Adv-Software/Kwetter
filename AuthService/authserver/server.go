package authserver

import (
	"fmt"
	pbauth "github.com/Portfolio-Adv-Software/Kwetter/AuthService/proto"
	"github.com/Portfolio-Adv-Software/Kwetter/AuthService/rabbitmq"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
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

type AuthServiceServer struct {
	pbauth.UnimplementedAuthServiceServer
}

func (a AuthServiceServer) GetData(ctx context.Context, req *pbauth.GetDataReq) (*pbauth.GetDataRes, error) {
	userID, err := primitive.ObjectIDFromHex(req.GetUserId())
	if err != nil {
		return nil, err
	}
	data := &pbauth.AuthData{}
	err = authdb.FindOne(ctx, bson.M{"_id": userID}).Decode(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error finding user: %v", err))
	}

	authData := &pbauth.AuthData{
		Id:       req.GetUserId(),
		Email:    data.GetEmail(),
		Password: data.GetPassword(),
	}
	res := &pbauth.GetDataRes{AuthData: authData}
	return res, nil
}

func (a AuthServiceServer) DeleteData(ctx context.Context, req *pbauth.DeleteDataReq) (*pbauth.DeleteDataRes, error) {
	objectID, err := primitive.ObjectIDFromHex(req.GetUserId())
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}
	maxRetries := 3
	retryCount := 0
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	for retryCount < maxRetries {
		count, err := authdb.CountDocuments(ctx, filter)
		if err != nil {
			return nil, err
		}

		if count == 0 {
			res := &pbauth.DeleteDataRes{Status: "No documents found to delete"}
			return res, nil
		}
		deleteResult, err := authdb.DeleteMany(ctx, filter)
		if err != nil {
			return nil, err
		}
		if deleteResult.DeletedCount == count {
			res := &pbauth.DeleteDataRes{Status: "All found documents deleted"}
			return res, nil
		}
		retryCount++
	}
	res := &pbauth.DeleteDataRes{Status: "Failed to delete records"}
	return res, nil
}

func (a AuthServiceServer) Register(ctx context.Context, req *pbauth.RegisterReq) (*pbauth.RegisterRes, error) {
	if !req.DataPermission {
		return nil, status.Error(codes.Aborted, "No permission to store data")
	}
	data := req.GetEmail()
	user := &pbauth.AuthData{}
	err := authdb.FindOne(ctx, bson.M{"email": data}).Decode(user)
	if err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "Email is already registered")
	}

	newUser := &pbauth.RegisterReq{
		Email:          req.GetEmail(),
		Password:       HashPassword(req.GetPassword()),
		DataPermission: req.GetDataPermission(),
	}

	insertResult, err := authdb.InsertOne(ctx, newUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unable to register user: %v", err))
	}

	insertedID, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, "Invalid type for InsertedID")
	}

	registeredUser := &pbauth.AuthData{
		Id:       insertedID.Hex(),
		Email:    newUser.GetEmail(),
		Password: "",
	}
	rabbitmq.ProduceMessage("user_queue", registeredUser)
	return &pbauth.RegisterRes{
		Status: "Registration successful",
	}, nil
}

func (a AuthServiceServer) Login(ctx context.Context, req *pbauth.LoginReq) (*pbauth.LoginRes, error) {
	user := &pbauth.AuthData{}
	err := authdb.FindOne(ctx, bson.M{"email": req.Email}).Decode(user)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Login credentials invalid"))
	}

	if !verifyPassword(req.GetPassword(), user.Password) {
		return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Login credentials invalid"))
	}
	token, _ := generateJWTToken(user)
	return &pbauth.LoginRes{
		Status: "Login successful",
		Token:  token,
	}, nil
}

func (a AuthServiceServer) Validate(ctx context.Context, req *pbauth.ValidateReq) (*pbauth.ValidateRes, error) {
	tokenString := req.Token
	// Parse and validate the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		// Return the secret key used for signing the token
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	// Verify the token is valid
	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		return &pbauth.ValidateRes{Status: "INVALID"}, nil
	}

	type User struct {
		Id    string `bson:"_id"`
		Email string `bson:"email"`
	}
	user := &User{}
	emailClaim := token.Claims.(jwt.MapClaims)
	email := emailClaim["email"].(string)
	err = authdb.FindOne(ctx, bson.M{"email": email}).Decode(user)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}
	// Token is valid
	return &pbauth.ValidateRes{
		Status: "Token Valid",
		Userid: user.Id,
	}, nil
}

var secretKey = os.Getenv("SECRET_KEY")

// enum for roles
type Role int

const (
	user Role = iota
	moderator
	admin
)

func generateJWTToken(user *pbauth.AuthData) (string, error) {
	claims := jwt.MapClaims{
		"id":      user.GetId(),
		"email":   user.GetEmail(),
		"role":    user,
		"expires": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes)
}

func verifyPassword(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

var db *mongo.Client
var authdb *mongo.Collection
var mongoCtx context.Context

func InitGRPC(wg *sync.WaitGroup) {
	defer wg.Done()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port: 50053")

	var mongoUser = os.Getenv("MONGO_USERNAME")
	var mongoPwd = os.Getenv("MONGO_PASSWORD")
	var dbconn = "mongodb+srv://" + mongoUser + ":" + mongoPwd + "@kwetter.vduy1tl.mongodb.net/test"

	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Unable to listen on port :50053: %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	srv := &AuthServiceServer{}
	pbauth.RegisterAuthServiceServer(s, srv)
	reflection.Register(s)

	fmt.Println("Connecting to MongoDB...")
	mongoCtx = context.Background()
	db, err = mongo.Connect(mongoCtx, options.Client().ApplyURI(dbconn))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to Mongodb")
	}
	authdb = db.Database("AuthTest").Collection("KwetterAuth")

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :50053")

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
