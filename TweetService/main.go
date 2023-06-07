package main

import (
	"fmt"
	"github.com/Portfolio-Adv-Software/Kwetter/TweetService/rabbitmq"
	. "github.com/Portfolio-Adv-Software/Kwetter/TweetService/tweetserver"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

// port 50051
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	loadEnv()

	go func() {
		defer wg.Done()
		InitGRPC(&wg)
	}()

	go func() {
		defer wg.Done()
		rabbitmq.DeleteGDPRUser(&wg)
	}()
	wg.Wait()
}

func loadEnv() {
	fmt.Println("loading env")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
