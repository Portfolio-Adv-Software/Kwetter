package main

import (
	"fmt"
	. "github.com/Portfolio-Adv-Software/Kwetter/AccountService/grpc"
	. "github.com/Portfolio-Adv-Software/Kwetter/AccountService/rabbitmq"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

// port 50054
func main() {
	loadEnv()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		InitGRPC(&wg)
	}()

	go func() {
		defer wg.Done()
		ConsumeMessage("user_queue", &wg)
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
