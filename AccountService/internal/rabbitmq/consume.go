package rabbitmq

import (
	"github.com/Portfolio-Adv-Software/Kwetter/AccountService/internal/proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
	"sync"
)

func ConsumeMessage(queue string, wg *sync.WaitGroup) {
	rMQUrl := os.Getenv("RMQ_KEY")
	defer wg.Done()
	conn, err := amqp.Dial(rMQUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			user := &__.User{}
			err := proto.Unmarshal(d.Body, user)
			if err != nil {
				log.Printf("failed to unmarshal user: %v", err)
				continue
			}
			log.Printf("received user: %v", user)
			c := InitClient()
			CreateUser(c, user)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func InitClient() __.UserServiceClient {
	// Set up a gRPC client connection to your backend service
	conn, err := grpc.Dial(":50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	c := __.NewUserServiceClient(conn)
	return c
}

func CreateUser(c __.UserServiceClient, user *__.User) *__.User {
	res, err := c.CreateUser(context.Background(), &__.CreateUserReq{User: user})
	if err != nil {
		return nil
	}

	return res.GetUser()
}
