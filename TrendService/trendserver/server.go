package trendserver

import (
	"context"
	"fmt"
	pbtrend "github.com/Portfolio-Adv-Software/Kwetter/TrendService/proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

type TrendServiceServer struct {
	pbtrend.UnimplementedTrendServiceServer
	ch *amqp.Channel
}

func (t TrendServiceServer) DeleteData(ctx context.Context, req *pbtrend.DeleteDataReq) (*pbtrend.DeleteDataRes, error) {
	filter := bson.M{"userid": req.GetUserId()}
	maxRetries := 3
	retryCount := 0
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	for retryCount < maxRetries {
		count, err := trenddb.CountDocuments(ctx, filter)
		if err != nil {
			return nil, err
		}

		if count == 0 {
			res := &pbtrend.DeleteDataRes{Status: "No documents found to delete"}
			return res, nil
		}
		deleteResult, err := trenddb.DeleteMany(ctx, filter)
		if err != nil {
			return nil, err
		}
		if deleteResult.DeletedCount == count {
			res := &pbtrend.DeleteDataRes{Status: "All found documents deleted"}
			return res, nil
		}
		retryCount++
	}
	res := &pbtrend.DeleteDataRes{Status: "Failed to delete records"}
	return res, nil
}

func (t TrendServiceServer) GetTrend(ctx context.Context, req *pbtrend.GetTrendReq) (*pbtrend.GetTrendRes, error) {
	//TODO implement me
	panic("implement me")
}

func (t TrendServiceServer) PostTrend(ctx context.Context, req *pbtrend.PostTrendReq) (*pbtrend.PostTrendRes, error) {
	data := req.GetTweet()

	tweet := &pbtrend.Tweet{
		UserID:   data.UserID,
		Username: data.Username,
		Body:     data.Body,
		Trend:    data.Trend,
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

func InitGRPC(wg *sync.WaitGroup) {
	defer wg.Done()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port: 50052")

	var mongoUser = os.Getenv("MONGO_USERNAME")
	var mongoPwd = os.Getenv("MONGO_PASSWORD")
	var dbconn = "mongodb+srv://" + mongoUser + ":" + mongoPwd + "@kwetter.vduy1tl.mongodb.net/test"

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Unable to listen on port :50052: %v", err)
	}

	// Set options, here we can configure things like TLS support
	var opts []grpc.ServerOption
	// Create new gRPC server with (blank) options
	s := grpc.NewServer(opts...)
	srv := &TrendServiceServer{}
	// Register the service with the server
	pbtrend.RegisterTrendServiceServer(s, srv)
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
