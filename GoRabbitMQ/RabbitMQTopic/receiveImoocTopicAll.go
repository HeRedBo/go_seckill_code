package main

import "Go/rabbitMQ/RabbitMQ"

func main() {
	imoocAll := RabbitMQ.NewRabbitMQTopic("exImoocTopic", "#")
	imoocAll.ReceiveTopic()
}
