package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func SetupRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial("amqps://ctltdklj:qV9vx5HIf7JyfDDA0fRto3Disk-T57CF@goose.rmq2.cloudamqp.com/ctltdklj")
	failOnError(err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		failOnError(err, "Failed to open channel")
	}

	// Set up the acknowledgment queue
	_, err = ch.QueueDeclare(
		"auth_deletion_ack", // Queue name for acknowledgments
		false,               // Durable
		false,               // Delete when unused
		false,               // Exclusive
		false,               // No-wait
		nil,                 // Arguments
	)
	if err != nil {
		return nil, nil, err
	}
	return conn, ch, err
}
