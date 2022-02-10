package main

import "Go/rabbitMQ/RabbitMQ"

func main() {
	imoocTwo := RabbitMQ.NewRabbitMQTopic("exImoocTopic", "imooc.*.two")
	imoocTwo.ReceiveTopic()
}
