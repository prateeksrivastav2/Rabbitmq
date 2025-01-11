package main

import (
	"log"
	"startlearning/rabbitmq" // Import the RabbitMQ package
)

func main() {
	// Connect to RabbitMQ and declare the queue
	conn, ch := rabbitmq.ConnectRabbitMQ()
	defer conn.Close()
	defer ch.Close()

	queue := rabbitmq.DeclareQueue(ch, "hello")

	// Consume messages
	msgs, err := ch.Consume(
		queue.Name, // Queue name
		"",         // Consumer tag
		true,       // Auto-acknowledge
		false,      // Exclusive
		false,      // No-local
		false,      // No-wait
		nil,        // Arguments
	)
	rabbitmq.FailOnError(err, "Failed to register a consumer")

	// Process messages
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf(" [x] Received %s", d.Body)
		}
	}()
	log.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
