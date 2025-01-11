package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

// FailOnError handles errors and logs a message.
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// ConnectRabbitMQ establishes a connection to RabbitMQ.
func ConnectRabbitMQ() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")

	return conn, ch
}

// DeclareQueue declares a queue and ensures it exists.
func DeclareQueue(ch *amqp.Channel, name string) amqp.Queue {
	queue, err := ch.QueueDeclare(
		name,  // Queue name
		false, // Durable
		false, // Auto-delete
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)
	FailOnError(err, "Failed to declare a queue")
	return queue
}
