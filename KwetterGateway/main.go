package main

import (
	"bufio"
	"context"
	"fmt"
	. "github.com/Portfolio-Adv-Software/Kwetter/KwetterGateway/grpc/TweetService"
	pb "github.com/Portfolio-Adv-Software/Kwetter/KwetterGateway/grpc/TweetService/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
)

func main() {
	//InitRouter()
	c, _ := InitClient()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("UserID: ")
		scanner.Scan()
		userID := scanner.Text()

		fmt.Print("UserName: ")
		scanner.Scan()
		username := scanner.Text()

		fmt.Print("tweetID: ")
		scanner.Scan()
		tweetID := scanner.Text()

		fmt.Print("Body: ")
		scanner.Scan()
		body := scanner.Text()

		c.PostTweet(context.Background(), demoTweet(userID, username, tweetID, body))

		fmt.Print("Post another tweet? (y/n): ")
		scanner.Scan()
		choice := scanner.Text()
		if choice != "y" {
			break
		}
	}
}

func demoTweet(userID string, username string, tweetID string, body string) *pb.PostTweetReq {
	tweet := &pb.Tweet{
		//UserID:   "1",
		//Username: "John",
		//TweetID:  "1",
		//Body:     "This is a test tweet #Test",
		UserID:   userID,
		Username: username,
		TweetID:  tweetID,
		Body:     body,
		Created:  timestamppb.Now(),
	}

	req := &pb.PostTweetReq{
		Tweet: tweet,
	}
	return req
}
