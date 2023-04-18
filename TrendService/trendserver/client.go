package trendserver

import (
	"context"
	"fmt"
	pbtrend "github.com/Portfolio-Adv-Software/Kwetter/TrendService/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func InitClient() (pbtrend.TrendServiceClient, error) {
	// Set up a gRPC client connection to your backend service
	conn, err := grpc.Dial(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	c := pbtrend.NewTrendServiceClient(conn)
	return c, nil
}

func PostTrend(c pbtrend.TrendServiceClient, tweet *pbtrend.Tweet) (*pbtrend.Tweet, error) {
	res, err := c.PostTrend(context.Background(), &pbtrend.PostTrendReq{Tweet: tweet})
	if err != nil {
		return nil, fmt.Errorf("failed to call PostTweet: %v", err)
	}

	return res.GetTweet(), nil
}
