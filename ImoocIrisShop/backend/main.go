package main

import (
	"ImoocIrisShop/backend/controllers"
	"ImoocIrisShop/common"
	"ImoocIrisShop/repositories"
	"ImoocIrisShop/services"
	"context"
	"log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {

	//1.创建iris 实例
	app := iris.New()
	//2.设置错误模式，在mvc模式下提示错误
	app.Logger().SetLevel("debug")
	//3.注册模板
	tmplate := iris.HTML("./backend/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmplate)

	//4.设置模板目标
	app.HandleDir("/assets", "./backend/web/assets")

	//出现异常跳转到指定页面
	//app.OnAnyErrorCode(func(ctx iris.Context) {
	//	ctx.ViewData("message",ctx.Values().GetStringDefault("message","访问的页面出错！"))
	//	ctx.ViewLayout("")
	//	ctx.View("shared/error.html")
	//})

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"status":  ctx.GetStatusCode(),
			"code":    ctx.Values().GetStringDefault("code", ""),
			"message": ctx.Values().GetStringDefault("message", ""),
			"trace":   ctx.Values().GetStringDefault("trace", "")})
	})

	// 连接数据库
	//db, err := common.NewMysqlConn()
	//if err != nil {
	//	log.Fatalf("%s", err)
	//}

	db2, err := common.NewGormMysqlConn()
	if err != nil {
		log.Fatalf("%s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 5、注册控制器
	productRepository := repositories.NewProductRepository("product", db2)
	productSerivce := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productSerivce)
	product.Handle(new(controllers.ProductController))

	// 注册 Order 控制器
	orderRepository := repositories.NewOrderRepository("orders", db2)
	orderService := services.NewOrderService(orderRepository)
	orderParty := app.Party("/order")
	order := mvc.New(orderParty)
	order.Register(ctx, orderService)
	order.Handle(new(controllers.OrderController))
	// 6、启动服务
	app.Run(
		// 启动服务在8080端口
		iris.Addr("localhost:8082"),
		// 启动时禁止检测框架版本差异
		//iris.WithoutVersionChecker,
		//忽略服务器错误
		iris.WithoutServerError(iris.ErrServerClosed),
		//让程序自身尽可能的优化
		iris.WithOptimizations,
	)

}
