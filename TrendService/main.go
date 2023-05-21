package main

import (
	"fmt"
	. "github.com/Portfolio-Adv-Software/Kwetter/TrendService/rabbitmq"
	. "github.com/Portfolio-Adv-Software/Kwetter/TrendService/trendserver"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

var wg sync.WaitGroup

func init() {
	fmt.Println("loading env")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		ConsumeMessage("tweet_queue")
	}()

	go func() {
		defer wg.Done()
		InitGRPC()
	}()
	wg.Wait()
}
