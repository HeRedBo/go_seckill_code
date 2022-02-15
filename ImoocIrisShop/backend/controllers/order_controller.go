package controllers

import (
	"ImoocIrisShop/common"
	"ImoocIrisShop/datamodels"
	"ImoocIrisShop/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"log"
	"strconv"
)

type OrderController struct {
	Ctx iris.Context
	OrderService services.IOrderService
}

func (o *OrderController) Get() mvc.View {
	orderArray, err := o.OrderService.GetAllOrderInfo()
	if err != nil {
		o.Ctx.Application().Logger().Debug("查询订单信息失败")
	}
	log.Println(orderArray)
	return mvc.View {
		Name : "order/view.html",
		Data : iris.Map{
			"order" : orderArray,
		},
	}
}


func (o *OrderController) GetEdit()  mvc.View {
	idString:=o.Ctx.URLParam("id")
	id ,err := strconv.ParseInt(idString,10,64)
	if err !=nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	orderInfo, err := o.OrderService.GetOrderInfoBy(id)
	if err !=nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	return mvc.View {
		Name : "order/edit.html",
		Data : iris.Map{
			"order" : orderInfo,
		},
	}
	//return order
}


func (o *OrderController) PostUpdate() {
	order  := &datamodels.Order{}
	o.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName: "imooc"})
	if err:= dec.Decode(o.Ctx.Request().Form,order);err!=nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	err:=o.OrderService.UpdateOrder(order)
	if err !=nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	o.Ctx.Redirect("/order")
}


// 订单删除
func(o *OrderController) GetDelete() {
	idString:=o.Ctx.URLParam("id")
	id ,err := strconv.ParseInt(idString,10,64)
	if err !=nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	isOk:=o.OrderService.DeleteByID(id)
	if isOk{
		o.Ctx.Application().Logger().Debug("删除商品成功，ID为："+idString)
	} else {
		o.Ctx.Application().Logger().Debug("删除商品失败，ID为："+idString)
	}
	o.Ctx.Redirect("/order/")
}