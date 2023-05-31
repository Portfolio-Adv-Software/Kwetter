package main

import (
	"fmt"
	. "github.com/Portfolio-Adv-Software/Kwetter/TrendService/rabbitmq"
	. "github.com/Portfolio-Adv-Software/Kwetter/TrendService/trendserver"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	loadEnv()
	go ConsumeMessage("tweet_queue")
	go InitGRPC()
}

// port 50052
func loadEnv() {
	fmt.Println("loading env")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
