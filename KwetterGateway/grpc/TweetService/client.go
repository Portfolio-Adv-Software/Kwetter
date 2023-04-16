package TweetService

import (
	"context"
	"fmt"
	pbtweet "github.com/Portfolio-Adv-Software/Kwetter/KwetterGateway/grpc/TweetService/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func InitClient() {
	// Set up a gRPC client connection to your backend service
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()
	c := pbtweet.NewTweetServiceClient(conn)

}

func ReturnAll(c pbtweet.TweetServiceClient) (response pbtweet.TweetService_ReturnAllClient) {
	response, err := c.ReturnAll(context.Background(), &pbtweet.ReturnAllReq{})
	if err != nil {
		log.Fatalf("Failed to call ReturnAll: %v", err)
	}
	fmt.Print(response)
	return response
}
