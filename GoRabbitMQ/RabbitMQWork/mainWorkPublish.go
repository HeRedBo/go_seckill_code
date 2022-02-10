package main

import (
	"Go/rabbitMQ/RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "imoocSimle")
	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("Hello rabbitMQ!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
