#  Setup Instructions

## 1. Set up RabbitMQ on Docker
#### Pull and run the RabbitMQ container with the management plugin
``` golang
docker pull rabbitmq:management
docker run -d -p 5672:5672 -p 15672:15672 --name rabbitmq rabbitmq:management
```

## Dependencies:
go get github.com/streadway/amqp


## Code Structure
startlearning/rabbitmq.go:
Contains the connection setup, queue declaration, and reusable methods.

producer/main.go:
Sends messages to the RabbitMQ queue.

consumer/main.go:
Receives and processes messages from the RabbitMQ queue.

## Run Producer 
go run Rabbitmq_producer/main.go
## Run Consumer
go run Rabbitmqc/main.go     

## Output_Image

<img width="482" alt="image" src="https://github.com/user-attachments/assets/3f38ad94-a1a5-4636-8132-ff29934c7f9f" />
