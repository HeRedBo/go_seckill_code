package controllers

import (
	"IrisMovies/repositories"
	"IrisMovies/services"
	"github.com/kataras/iris/v12/mvc"
)

type MovieController struct {

}

func (c *MovieController) Get() mvc.View {
	MovieRepository := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManger(MovieRepository)
	MovieResult := movieService.ShowMovieName()
	return mvc.View{
		Name:"movie/index.html",
		Data: MovieResult,
	}
}
