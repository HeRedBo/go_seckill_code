package main

import "Go/rabbitMQ/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "imoocSimle")
	rabbitmq.ConsumeSimple()
}
