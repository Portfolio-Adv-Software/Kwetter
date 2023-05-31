package authserver

import (
	"fmt"
	pbauth "github.com/Portfolio-Adv-Software/Kwetter/AuthService/proto"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
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
	"time"
)

type AuthServiceServer struct {
	pbauth.UnimplementedAuthServiceServer
}

type User struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

func (a AuthServiceServer) Register(ctx context.Context, req *pbauth.RegisterReq) (*pbauth.RegisterRes, error) {
	data := req.GetEmail()
	user := &User{}
	err := authdb.FindOne(ctx, bson.M{"email": data}).Decode(user)
	if err != nil {
		return &pbauth.RegisterRes{Status: "Email is already registered"}, nil
	}

	newUser := &pbauth.RegisterReq{
		Email:    req.GetEmail(),
		Password: HashPassword(req.GetPassword()),
	}

	_, err = authdb.InsertOne(ctx, newUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unable to register user: %v", err))
	}
	return &pbauth.RegisterRes{
		Status: "Registration successful",
	}, nil
}

func (a AuthServiceServer) Login(ctx context.Context, req *pbauth.LoginReq) (*pbauth.LoginRes, error) {
	email := req.Email
	user := &User{}
	err := authdb.FindOne(ctx, bson.M{"email": email}).Decode(user)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Login credentials invalid"))
	}

	if !verifyPassword(req.GetPassword(), user.Password) {
		return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Login credentials invalid"))
	}
	token, _ := generateJWTToken(user.Email)
	return &pbauth.LoginRes{
		Token:  token,
		Status: "Login succesful",
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

	// Token is valid
	return &pbauth.ValidateRes{Status: "VALID"}, nil
}

var secretKey = os.Getenv("SECRET_KEY")

func generateJWTToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email":   email,
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

func InitGRPC() {
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
