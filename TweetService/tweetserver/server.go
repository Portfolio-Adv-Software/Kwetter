package tweetserver

import (
	"context"
	"fmt"
	pbtweet "github.com/Portfolio-Adv-Software/Kwetter/TweetService/proto"
	"github.com/Portfolio-Adv-Software/Kwetter/TweetService/rabbitmq"
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
	"regexp"
	"sync"
	"time"
)

type TweetServiceServer struct {
	pbtweet.UnimplementedTweetServiceServer
}

func (t TweetServiceServer) DeleteData(ctx context.Context, req *pbtweet.DeleteDataReq) (*pbtweet.DeleteDataRes, error) {
	filter := bson.M{"userid": req.GetUserId()}
	maxRetries := 3
	retryCount := 0
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	for retryCount < maxRetries {
		count, err := tweetdb.CountDocuments(ctx, filter)
		if err != nil {
			return nil, err
		}
		if count == 0 {
			res := &pbtweet.DeleteDataRes{Status: "No documents found to delete"}
			return res, nil
		}
		deleteResult, err := tweetdb.DeleteMany(ctx, filter)
		if err != nil {
			return nil, err
		}
		if deleteResult.DeletedCount == count {
			res := &pbtweet.DeleteDataRes{Status: "All found documents deleted"}
			return res, nil
		}
		retryCount++
	}
	res := &pbtweet.DeleteDataRes{Status: "Failed to delete records"}
	return res, nil
}

func (t TweetServiceServer) ReturnAll(ctx context.Context, req *pbtweet.ReturnAllReq) (*pbtweet.ReturnAllRes, error) {
	cursor, err := tweetdb.Find(ctx, bson.M{"userid": req.GetUserId()})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Error finding tweets: %v", err))
	}
	defer cursor.Close(ctx)
	var tweets []*pbtweet.Tweet
	for cursor.Next(ctx) {
		data := &pbtweet.Tweet{}
		err := cursor.Decode(data)
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("error decoding data: %v", err))
		}
		tweets = append(tweets, data)
	}
	res := &pbtweet.ReturnAllRes{Tweet: tweets}
	return res, nil
}

func (t TweetServiceServer) ReturnTweet(ctx context.Context, req *pbtweet.ReturnTweetReq) (*pbtweet.ReturnTweetRes, error) {
	tweetID := req.GetTweetID()
	data := &pbtweet.Tweet{}
	err := tweetdb.FindOne(ctx, bson.M{"_id": tweetID}).Decode(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("unknown internal error: %v", err))
	}
	tweet := &pbtweet.Tweet{
		UserID:   data.UserID,
		Username: data.Username,
		Body:     data.Body,
	}
	res := &pbtweet.ReturnTweetRes{Tweet: tweet}
	return res, nil
}

func (t TweetServiceServer) PostTweet(ctx context.Context, req *pbtweet.PostTweetReq) (*pbtweet.PostTweetRes, error) {
	data := req.GetTweet()

	tweet := &pbtweet.Tweet{
		UserID:   data.UserID,
		Username: data.Username,
		Body:     data.Body,
	}

	re := regexp.MustCompile(`#\w+`)
	hashtags := re.FindAllString(tweet.Body, -1)

	if len(hashtags) > 0 {
		rabbitmq.ProduceMessage("tweet_queue", tweet)
	}

	_, err := tweetdb.InsertOne(ctx, tweet)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("error inserting tweet into db: %v", err))
	}
	res := &pbtweet.PostTweetRes{Tweet: tweet}
	return res, nil
}

var db *mongo.Client
var tweetdb *mongo.Collection
var mongoCtx context.Context

func InitGRPC(wg *sync.WaitGroup) {
	defer wg.Done()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port: 50051")

	var mongoUser = os.Getenv("MONGO_USERNAME")
	var mongoPwd = os.Getenv("MONGO_PASSWORD")
	var dbconn = "mongodb+srv://" + mongoUser + ":" + mongoPwd + "@kwetter.vduy1tl.mongodb.net/test"

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Unable to listen on port :50051: %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)

	srv := &TweetServiceServer{}
	pbtweet.RegisterTweetServiceServer(s, srv)
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
	tweetdb = db.Database("TweetTest").Collection("KwetterTweets")

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :50051")

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
