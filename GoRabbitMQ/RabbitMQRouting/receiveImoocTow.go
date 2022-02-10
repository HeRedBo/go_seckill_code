package main

import "Go/rabbitMQ/RabbitMQ"

func main() {
	imoocTwo := RabbitMQ.NewRabbitMQRouting("exImooc", "imooc_two")
	imoocTwo.ReceiveRouting()
}
