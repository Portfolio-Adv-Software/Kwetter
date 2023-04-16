package TweetService

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"github.com/Portfolio-Adv-Software/Kwetter/KwetterGateway/grpc/TweetService/proto"
)

func InitClient() {
	// Set up a gRPC client connection to your backend service
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()
	client := pbtweet.
}
