package main

import "Go/rabbitMQ/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("" + "newProduct")
	rabbitmq.ReceiveSub()
}
