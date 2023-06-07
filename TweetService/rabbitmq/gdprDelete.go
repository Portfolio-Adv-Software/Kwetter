package rabbitmq

import (
	"fmt"
	pbtweet "github.com/Portfolio-Adv-Software/Kwetter/TweetService/proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
	"sync"
)

func DeleteGDPRUser(wg *sync.WaitGroup) {
	rMQUrl := os.Getenv("RMQ_KEY")
	defer wg.Done()
	conn, err := amqp.Dial(rMQUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	deleteQueue, err := ch.QueueDeclare(
		"delete_tweet", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")

	deleteMsgs, err := ch.Consume(
		deleteQueue.Name, // queue
		"",               // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range deleteMsgs {
			fmt.Printf("Message Body: %s\n", string(d.Body))
			req := &pbtweet.DeleteDataReq{}
			err := proto.Unmarshal(d.Body, req)
			if err != nil {
				log.Printf("failed to unmarshal delete req: %v", err)
				continue
			}
			log.Printf("Received a message to delete everything regarding user: %+v", req.GetUserId())
			c, _ := InitClient()
			deleteData(c, req)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func InitClient() (pbtweet.TweetServiceClient, error) {
	// Set up a gRPC client connection to your backend service
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	c := pbtweet.NewTweetServiceClient(conn)
	return c, nil
}

func deleteData(c pbtweet.TweetServiceClient, req *pbtweet.DeleteDataReq) (*pbtweet.DeleteDataRes, error) {
	res, err := c.DeleteData(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to call deleteData: %v", err)
	}

	return res, nil
}
