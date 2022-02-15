package main

import (
	"ImoocIrisShop/common"
	controllers2 "ImoocIrisShop/fronted/web/controllers"
	"ImoocIrisShop/repositories"
	"ImoocIrisShop/services"
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"log"
)

func main() {
	//1.创建iris 实例
	app := iris.New()
	//2.设置错误模式，在mvc模式下提示错误
	app.Logger().SetLevel("debug")
	//3.注册模板
	tmplate := iris.HTML("./fronted/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmplate)
	//4.设置模板
	app.HandleDir("/public", "./fronted/web/public")
	//访问生成好的html静态文件
	app.HandleDir("/html", "./fronted/web/htmlProductShow")
	//出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	// 连接数据库
	db,err := common.NewGormMysqlConn()
	if err != nil {
		log.Fatalf("%s", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// 注册控制器
	userRepository := repositories.NewUserRepository(db)
	userService    := services.NewUserService(userRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx,userService)
	product.Handle(new(controllers2.UserController))


	//db,err := common.NewMysqlConn()
	//if err != nil {
	//	log.Fatalf("%s", err)
	//}

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
