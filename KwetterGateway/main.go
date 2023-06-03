package main

import (
	"github.com/Portfolio-Adv-Software/Kwetter/KwetterGateway/gatewayserver"
	"sync"
)

// port 50055
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var config gatewayserver.ServiceConfig
	setConfig(&config)
	go func() {
		defer wg.Done()
		gatewayserver.InitGRPC(&wg, &config)
	}()

	go func() {
		defer wg.Done()
		gatewayserver.InitMux(&wg, &config)
	}()
	wg.Wait()
}

func setConfig(config *gatewayserver.ServiceConfig) {
	config.AuthServiceAddr = "localhost:50053"
	config.UserServiceAddr = "localhost:50054"
	config.TrendServiceAddr = "localhost:50052"
	config.TweetServiceAddr = "localhost:50051"
}
