package main

import (
	"Go/rabbitMQ/RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	imoocOne := RabbitMQ.NewRabbitMQRouting("exImoocTopic", "imooc.topic.one")
	imoocTwo := RabbitMQ.NewRabbitMQRouting("exImoocTopic", "imooc.topic.two")
	for i := 0; i <= 10; i++ {
		imoocOne.PublishTopic("Hello imooc one!" + strconv.Itoa(i))
		imoocTwo.PublishTopic("Hello imooc Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
