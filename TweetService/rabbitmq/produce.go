package rabbitmq

import (
	"context"
	pb "github.com/Portfolio-Adv-Software/Kwetter/TweetService/proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
	"log"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func ProduceMessage(name string, tweet *pb.Tweet) {
	conn, err := amqp.Dial("amqps://ctltdklj:qV9vx5HIf7JyfDDA0fRto3Disk-T57CF@goose.rmq2.cloudamqp.com/ctltdklj")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := proto.Marshal(tweet)
	if err != nil {
		log.Panicf("Failed to marshal tweet: %v", err)
	}

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/protobuf",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %v\n", body)
}
