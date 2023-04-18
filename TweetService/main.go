package main

import (
	"github.com/Portfolio-Adv-Software/Kwetter/TweetService/tweetserver"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go tweetserver.InitGRPC()

	<-stop
}
