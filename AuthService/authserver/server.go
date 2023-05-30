package authserver

import (
	"fmt"
	pbauth "github.com/Portfolio-Adv-Software/Kwetter/AuthService/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

type AuthServiceServer struct {
	pbauth.UnimplementedAuthServiceServer
}

func (a AuthServiceServer) Register(ctx context2.Context, req *pbauth.RegisterReq) (*pbauth.RegisterRes, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthServiceServer) Login(ctx context2.Context, req *pbauth.LoginReq) (*pbauth.LoginRes, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthServiceServer) Validate(ctx context2.Context, req *pbauth.ValidateReq) (*pbauth.ValidateRes, error) {
	//TODO implement me
	panic("implement me")
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
