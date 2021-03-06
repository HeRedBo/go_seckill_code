package main

import (
	"ImoocIrisShop/common"
	"ImoocIrisShop/fronted/middleware"
	controllers2 "ImoocIrisShop/fronted/web/controllers"
	"ImoocIrisShop/rabbitmq"
	"ImoocIrisShop/repositories"
	"ImoocIrisShop/services"
	"context"
	"log"
	"time"

	"github.com/kataras/iris/v12/sessions"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	//1.创建iris 实例
	app := iris.New()
	//2.设置错误模式，在mvc模式下提示错误
	app.Logger().SetLevel("debug")
	//3.注册模板
	template := iris.HTML("./fronted/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	//4.设置模板
	app.HandleDir("/public", "./fronted/web/public")
	//访问生成好的html静态文件
	app.HandleDir("/html", "./fronted/web/views/htmlProductShow/")
	//出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	// 连接数据库
	db, err := common.NewGormMysqlConn()
	if err != nil {
		log.Fatalf("%s", err)
	}

	sess := sessions.New(sessions.Config{
		Cookie:  "AdminCookie",
		Expires: 600 * time.Minute,
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// 注册控制器
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userParty := app.Party("/user")
	user := mvc.New(userParty)
	user.Register(ctx, userService, sess)
	app.UseRouter(sess.Handler())
	user.Handle(new(controllers2.UserController))

	// 注册 product 控制器
	//db2, err := common.NewMysqlConn()
	//if err != nil {
	//
	//}

	rabbitmq := rabbitmq.NewRabbitMQSimple("imoocProduct")
	product := repositories.NewProductRepository("product", db)
	productService := services.NewProductService(product)
	order := repositories.NewOrderRepository("orders", db)
	orderService := services.NewOrderService(order)

	proProduct := app.Party("/product")
	pro := mvc.New(proProduct)
	proProduct.Use(middleware.AuthCheck)
	pro.Register(productService, orderService, rabbitmq)

	pro.Handle(new(controllers2.ProductController))

	// 6、启动服务
	app.Run(
		// 启动服务在8080端口
		iris.Addr("localhost:8081"),
		// 启动时禁止检测框架版本差异
		//iris.WithoutVersionChecker,
		//忽略服务器错误
		iris.WithoutServerError(iris.ErrServerClosed),
		//让程序自身尽可能的优化
		iris.WithOptimizations,
	)

}
