package TweetService

import (
	"context"
	"fmt"
	pbtweet "github.com/Portfolio-Adv-Software/Kwetter/KwetterGateway/grpc/TweetService/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

func InitClient() (pbtweet.TweetServiceClient, error) {
	// Set up a gRPC client connection to your backend service
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	c := pbtweet.NewTweetServiceClient(conn)
	return c, nil
}

func ReturnAll(c pbtweet.TweetServiceClient) ([]*pbtweet.Tweet, error) {
	var tweets []*pbtweet.Tweet

	stream, err := c.ReturnAll(context.Background(), &pbtweet.ReturnAllReq{})
	if err != nil {
		return nil, fmt.Errorf("failed to call ReturnAll: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			// End of stream
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to receive tweet: %v", err)
		}

		tweet := res.GetTweet()
		if tweet != nil {
			tweets = append(tweets, tweet)
		}
	}
	return tweets, nil
}

func ReturnTweet(c pbtweet.TweetServiceClient, tweetID string) (*pbtweet.Tweet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pbtweet.ReturnTweetReq{TweetID: tweetID}
	res, err := c.ReturnTweet(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get tweet: %v", err)
	}
	return res.Tweet, nil
}

func PostTweet(c pbtweet.TweetServiceClient, tweet *pbtweet.Tweet) (*pbtweet.Tweet, error) {
	res, err := c.PostTweet(context.Background(), &pbtweet.PostTweetReq{Tweet: tweet})
	if err != nil {
		return nil, fmt.Errorf("failed to call PostTweet: %v", err)
	}

	return res.GetTweet(), nil
}
