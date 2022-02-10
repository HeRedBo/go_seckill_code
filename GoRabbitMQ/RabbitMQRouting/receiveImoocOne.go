package main

import "Go/rabbitMQ/RabbitMQ"

func main() {
	imoocOne := RabbitMQ.NewRabbitMQRouting("exImooc", "imooc_one")
	imoocOne.ReceiveRouting()
}
