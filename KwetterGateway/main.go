package main

import (
	"context"
	. "github.com/Portfolio-Adv-Software/Kwetter/KwetterGateway/grpc/TweetService"
	pb "github.com/Portfolio-Adv-Software/Kwetter/KwetterGateway/grpc/TweetService/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	//InitRouter()
	c, _ := InitClient()
	c.PostTweet(context.Background(), demoTweet())
}

func demoTweet() *pb.PostTweetReq {
	tweet := &pb.Tweet{
		UserID:   "1",
		Username: "John",
		TweetID:  "1",
		Body:     "This is a test tweet #Test",
		Created:  timestamppb.Now(),
	}

	req := &pb.PostTweetReq{
		Tweet: tweet,
	}
	return req
}
