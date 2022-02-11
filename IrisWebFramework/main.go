package main

import (
	"IrisWebFramework/web/controllers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./web/views",".html"))
	// 注册控制器  前端 通过 http://localhost:8081/hello 访问即可
	mvc.New(app.Party("hello")).Handle(new(controllers.MovieController))
	app.Run(iris.Addr("localhost:8081"), )
}
