package main

import (
	"IrisMovies/datasource"
	"IrisMovies/repositories"
	"IrisMovies/services"
	"IrisMovies/web/controllers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	// 加载视图模板地址
	app.RegisterView(iris.HTML("./web/views",".html"))

	// 注册控制器  前端 通过 http://localhost:8081/hello 访问即可
	//mvc.New(app.Party("hello")).Handle(new(controllers.MovieController))
	//你也可以使用  `mvc.Configure` 方法拆分编写 MVC 应用程序的配置。
	// 如下所示：
	mvc.Configure(app.Party("/movies"), movies)

	app.Run(
		// Start the web server at localhost:8080
		iris.Addr("localhost:8081"),
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,

		)
}

func movies(app *mvc.Application) {
	// 中间件
	// app.Router.Use()

	repo := repositories.NewMovieRepository(datasource.Movies)
	movieService := services.NewMovieService(repo)
	app.Register(movieService)

	//初始化控制器
	// 注意，你可以初始化多个控制器
	// 你也可以 使用 `movies.Party(relativePath)` 或者 `movies.Clone(app.Party(...))` 创建子应用。s
	app.Handle(new(controllers.MovieController))
}