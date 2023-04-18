package main

import (
	. "github.com/Portfolio-Adv-Software/Kwetter/TrendService/rabbitmq"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// create a channel to receive signals to stop the application
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// start the goroutine to receive messages from the queue
	go ConsumeMessage("tweet_queue")

	// wait for a signal to stop the application
	<-stop
}
