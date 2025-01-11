package main

import (
	"log"
	"startlearning/rabbitmq"

	"github.com/streadway/amqp"
)

func main() {
	// Connect to RabbitMQ and declare the queue
	conn, ch := rabbitmq.ConnectRabbitMQ()
	defer conn.Close()
	defer ch.Close()

	queue := rabbitmq.DeclareQueue(ch, "hello")

	// Publish a message
	body := "Hello, prateek!"
	err := ch.Publish(
		"",         // Exchange
		queue.Name, // Routing key (queue name)
		false,      // Mandatory
		false,      // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	rabbitmq.FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}
