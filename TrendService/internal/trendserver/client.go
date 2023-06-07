package trendserver

import (
	"context"
	"fmt"
	"github.com/Portfolio-Adv-Software/Kwetter/TrendService/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func InitClient() (__.TrendServiceClient, error) {
	// Set up a gRPC client connection to your backend service
	conn, err := grpc.Dial(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	c := __.NewTrendServiceClient(conn)
	return c, nil
}

func PostTrend(c __.TrendServiceClient, tweet *__.Tweet) (*__.Tweet, error) {
	res, err := c.PostTrend(context.Background(), &__.PostTrendReq{Tweet: tweet})
	if err != nil {
		return nil, fmt.Errorf("failed to call PostTweet: %v", err)
	}

	return res.GetTweet(), nil
}
