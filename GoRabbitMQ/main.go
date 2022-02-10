package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// 声明变量
var conn *amqp.Connection
var channel *amqp.Channel
var count = 0

const (
	// 队列名称
	queueName = "imooc"
	exchange  = ""
	mqurl     = "amqp://imoocuser:imoocuser@127.0.0.1:5672/imooc"
)

func Connect() {
	fmt.Println(mqurl)
	conn, err := amqp.Dial(mqurl)
	failOnErr(err, "failed to connect")
	channel, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
}

func close() {
	// 1、关闭 channel
	channel.Close()
	// 2、关闭链接
	conn.Close()
}

// 消息生产
func push() {
	// 1、判断是否存在 channel
	if channel == nil {
		Connect()
	}

	// 2、消息
	messages := "Hello simple imooc"

	// 3、声明队列
	q, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	// 4、错误判断
	if err != nil {
		fmt.Println(err)
	}

	//5、 生产消息
	channel.Publish(exchange, q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(messages),
	})
}

// 消费端
func receive() {
	// 1、判断channel 是否存在
	if channel == nil {
		Connect()
	}

	// 2、声明队列
	q, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	// 4、错误判断
	failOnErr(err, "声明队列")

	// 3、消费diam
	msg, err := channel.Consume(q.Name, "", true, false, false, false, nil)
	failOnErr(err, "获取消费信息异常")

	msgForver := make(chan bool)
	//消费逻辑
	go func() {
		for d := range msg {
			//相同效果，把[]byte类型转化为字符串类型
			//s := queue.BytesToString(&d.Body)
			s := string(d.Body)
			count++
			fmt.Printf("接收信息是%s-- %d\n", s, count)
		}
	}()

	fmt.Println("退出请按 CTRL+C\n")
	<-msgForver
}

//错误处理函数
func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

func main() {
	go func() {
		for {
			push()
			time.Sleep(5 * time.Second)
		}
	}()
	receive()
	fmt.Println("生产消费完成")
	close()

}
