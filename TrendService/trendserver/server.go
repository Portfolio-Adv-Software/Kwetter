package trendserver

import (
	"context"
	"fmt"
	pbtrend "github.com/Portfolio-Adv-Software/Kwetter/TrendService/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
)

type TrendServiceServer struct {
	pbtrend.UnimplementedTrendServiceServer
}

func (t TrendServiceServer) PostTrend(ctx context.Context, req *pbtrend.PostTrendReq) (*pbtrend.PostTrendRes, error) {
	data := req.GetTweet()

	tweet := &pbtrend.Tweet{
		UserID:   data.UserID,
		Username: data.Username,
		TweetID:  data.TweetID,
		Body:     data.Body,
		Trends:   data.Trends,
		Created:  data.Created,
	}

	_, err := trenddb.InsertOne(ctx, tweet)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error inserting tweet into db: %v", err))
	}
	res := &pbtrend.PostTrendRes{Tweet: tweet}
	return res, nil
}

var db *mongo.Client
var trenddb *mongo.Collection
var mongoCtx context.Context

func InitGRPC() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port: 50052")

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Unable to listen on port :50052: %v", err)
	}

	// Set options, here we can configure things like TLS support
	var opts []grpc.ServerOption
	// Create new gRPC server with (blank) options
	s := grpc.NewServer(opts...)
	// Create BlogService type
	srv := &TrendServiceServer{}
	// Register the service with the server
	pbtrend.RegisterTrendServiceServer(s, srv)

	// Initialize MongoDb client
	fmt.Println("Connecting to MongoDB...")

	// non-nil empty context
	mongoCtx = context.Background()

	// Connect takes in a context and options, the connection URI is the only option we pass for now
	db, err = mongo.Connect(mongoCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
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
	trenddb = db.Database("TrendTest").Collection("KwetterTrends")

	// Start the server in a child routine
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :50052")

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
