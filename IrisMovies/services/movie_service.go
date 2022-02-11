package services

import (
	"IrisMovies/repositories"
	"fmt"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManger struct {
	repo repositories.MovieRepository
}

func NewMovieServiceManger(repo repositories.MovieRepository)  MovieService {
	return &MovieServiceManger{ repo: repo}
}

func (m *MovieServiceManger) ShowMovieName() string {
	Name := m.repo.GetMovieName()
	fmt.Println("获取的电影名称为：" + Name)
	return Name
}

