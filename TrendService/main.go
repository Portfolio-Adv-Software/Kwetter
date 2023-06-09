package main

import (
	"fmt"
	. "github.com/Portfolio-Adv-Software/Kwetter/TrendService/internal/rabbitmq"
	. "github.com/Portfolio-Adv-Software/Kwetter/TrendService/internal/trendserver"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(4)

	loadEnv()

	go func() {
		defer wg.Done()
		ConsumeMessage("tweet_queue", &wg)
	}()
	go func() {
		defer wg.Done()
		DeleteGDPRUser(&wg)
	}()
	go func() {
		defer wg.Done()
		go InitGRPC(&wg)
	}()
	wg.Wait()
}

// port 50052
func loadEnv() {
	fmt.Println("loading env")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
