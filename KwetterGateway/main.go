package main

import (
	"github.com/Portfolio-Adv-Software/Kwetter/KwetterGateway/config"
	"github.com/Portfolio-Adv-Software/Kwetter/KwetterGateway/gatewayserver"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"sync"
)

// port 50055
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	mux := runtime.NewServeMux()
	config.SetConfig()
	go func() {
		defer wg.Done()
		gatewayserver.InitGRPC(&wg, mux)
	}()

	go func() {
		defer wg.Done()
		gatewayserver.InitMux(&wg, mux)
	}()
	wg.Wait()
}
