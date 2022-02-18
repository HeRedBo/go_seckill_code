package main

import (
	"ImoocIrisShop/common"
	"ImoocIrisShop/rabbitmq"
	"ImoocIrisShop/repositories"
	"ImoocIrisShop/services"
	"fmt"
)

func main() {
	db, err := common.NewMysqlConn()
	if err != nil {
		fmt.Println(err)
	}

	product := repositories.NewProductRepository("product", db)
	productService := services.NewProductService(product)

	order := repositories.NewOrderRepository("orders", db)
	orderService := services.NewOrderService(order)

	rabbitmqConsumeSimple := rabbitmq.NewRabbitMQSimple("imoocProduct")
	rabbitmqConsumeSimple.ConsumeSimple(orderService, productService)
}
