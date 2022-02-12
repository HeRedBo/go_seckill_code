package controllers

import (
	"IrisMovies/datamodels"
	"IrisMovies/services"
	"errors"
	"github.com/kataras/iris/v12"
)

type MovieController struct {

	MovieService services.MovieService
}

/**
	Get 返回 movies 的列表
	// as
 */
func (c *MovieController) Get() (results []datamodels.Movie) {
	return c.MovieService.GetAll()
}

// 返回 一个 movie
func (c *MovieController) GetBy(id int64) (movie datamodels.Movie, found bool) {
	return c.MovieService.GetByID(id)
}


func (c *MovieController) PutBy(ctx iris.Context, id int64) (datamodels.Movie, error) {
	file, info, err := ctx.FormFile(`poster`)
	if err != nil {
		return datamodels.Movie{}, errors.New("failed due form file 'poster' missing")
	}

	// 关闭
	file.Close()
	//  文件上传
	poster := info.Filename
	genre := ctx.FormValue("genre")

	return c.MovieService.UpdatePosterAndGenreByID(id , poster, genre)
}

func(c *MovieController) DeleteBy(id int64) interface{} {
	wasDel := c.MovieService.DeleteByID(id)
	if wasDel {
		// 被删除的的movie 的ID
		return iris.Map{"delete":id}
	}
	//在这里，我们可以看到一个方法函数可以返回两种类型中的任何一种（map 或者 int）,
	// 我们不用指定特定的返回类型。
	return iris.StatusBadRequest
}




