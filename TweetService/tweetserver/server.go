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
)

type TweetServiceServer struct {
	pbtweet.UnimplementedTweetServiceServer
}

func (t TweetServiceServer) ReturnAll(_ *pbtweet.ReturnAllReq, s pbtweet.TweetService_ReturnAllServer) error {
	data := &pbtweet.Tweet{}
	cursor, err := tweetdb.Find(context.Background(), bson.M{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		err := cursor.Decode(data)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		s.Send(&pbtweet.ReturnAllRes{
			Tweet: &pbtweet.Tweet{
				UserID:   data.UserID,
				Username: data.Username,
				TweetID:  data.TweetID,
				Body:     data.Body,
				Created:  data.Created,
			},
		})
	}
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown cursor error: %v", err))
	}
	return nil
}

func (t TweetServiceServer) ReturnTweet(_ context.Context, req *pbtweet.ReturnTweetReq) (*pbtweet.ReturnTweetRes, error) {
	tweetID := req.TweetID
	data := &pbtweet.Tweet{}
	err := tweetdb.FindOne(context.Background(), bson.M{"tweetid": tweetID}).Decode(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("unknown internal error: %v", err))
	}
	tweet := &pbtweet.Tweet{
		UserID:   data.UserID,
		Username: data.Username,
		TweetID:  data.TweetID,
		Body:     data.Body,
		Created:  data.Created,
	}
	res := &pbtweet.ReturnTweetRes{Tweet: tweet}
	return res, nil
}

func (t TweetServiceServer) PostTweet(ctx context.Context, req *pbtweet.PostTweetReq) (*pbtweet.PostTweetRes, error) {
	data := req.GetTweet()

	tweet := &pbtweet.Tweet{
		UserID:   data.UserID,
		Username: data.Username,
		TweetID:  data.TweetID,
		Body:     data.Body,
		Created:  data.Created,
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

func InitGRPC() {
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
