package main

import (
	"github.com/Portfolio-Adv-Software/Kwetter/KwetterGateway/internal/config"
	"github.com/Portfolio-Adv-Software/Kwetter/KwetterGateway/internal/gatewayserver"
	"sync"
)

// port 50055
func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	//mux := runtime.NewServeMux()
	config.SetConfig()
	go func() {
		defer wg.Done()
		gatewayserver.InitGRPC(&wg, nil) //add mux
	}()

	//go func() {
	//	defer wg.Done()
	//	gatewayserver.InitMux(&wg, mux)
	//}()
	wg.Wait()
}
