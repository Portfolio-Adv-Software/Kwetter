package main

import (
	"fmt"
	"github.com/Portfolio-Adv-Software/Kwetter/AccountService/internal/rabbitmq"
	. "github.com/Portfolio-Adv-Software/Kwetter/AccountService/internal/userserver"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

// port 50054
func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	loadEnv()

	go func() {
		defer wg.Done()
		InitGRPC(&wg)
	}()
	go func() {
		defer wg.Done()
		rabbitmq.ConsumeMessage("user_queue", &wg)
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

//test
