package rabbitmq

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/net/context"
	"log"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func SendDeleteGDPRUser(userid string) {
	queueNames := []string{"delete_auth", "delete_tweet", "delete_trend"}
	conn, err := amqp.Dial("amqps://ctltdklj:qV9vx5HIf7JyfDDA0fRto3Disk-T57CF@goose.rmq2.cloudamqp.com/ctltdklj")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(userid)
	if err != nil {
		log.Panicf("Failed to marshal userid: %s", err)
	}

	for _, queueName := range queueNames {
		_, err := ch.QueueDeclare(
			queueName, // name
			false,     // durable
			false,     // delete when unused
			false,     // exclusive
			false,     // no-wait
			nil,       // arguments
		)
		failOnError(err, "Failed to declare a queue")
		err = ch.PublishWithContext(ctx,
			"",        // exchange
			queueName, // routing key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			})
		failOnError(err, "Failed to publish a message to queue "+queueName)

		log.Printf(" [x] Sent %s to queue %s\n", body, queueName)
	}
}
