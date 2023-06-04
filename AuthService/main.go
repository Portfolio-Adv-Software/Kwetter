package main

import (
	"fmt"
	. "github.com/Portfolio-Adv-Software/Kwetter/AuthService/authserver"
	"github.com/joho/godotenv"
	"log"
)

// port 50053
func main() {
	loadEnv()
	InitGRPC()
}

func loadEnv() {
	fmt.Println("loading env")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

//test
