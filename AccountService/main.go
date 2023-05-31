package main

import (
	"fmt"
	. "github.com/Portfolio-Adv-Software/Kwetter/AccountService/grpc"
	. "github.com/Portfolio-Adv-Software/Kwetter/AccountService/rabbitmq"
	"github.com/joho/godotenv"
	"log"
)

// port 50054
func main() {
	loadEnv()
	go InitGRPC()
	go ConsumeMessage("user_queue")
}

func loadEnv() {
	fmt.Println("loading env")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
