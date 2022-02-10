package main

import (
	"Go/rabbitMQ/RabbitMQ"
	"fmt"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "imoocSimle")
	rabbitmq.PublishSimple("Hello rabbitMQ 123123")
	fmt.Println("队列发送成功！")
}
