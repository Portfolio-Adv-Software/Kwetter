package main

import (
	"fmt"
	. "github.com/Portfolio-Adv-Software/Kwetter/AuthService/internal/authserver"
	"github.com/Portfolio-Adv-Software/Kwetter/AuthService/internal/rabbitmq"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

// port 50053
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
