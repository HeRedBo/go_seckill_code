package services

import (
	"IrisMovies/datamodels"
	"IrisMovies/repositories"
)

type MovieService interface {
	GetAll() []datamodels.Movie
	GetByID(id int64) (datamodels.Movie, bool)
	UpdatePosterAndGenreByID(id int64, poster , genre string) (datamodels.Movie, error)
	DeleteByID(id int64) bool
}

func NewMovieService(repo repositories.MovieRepository) MovieService {
	return &movieService{
		repo: repo,
	}
}


type movieService struct {
	 repo repositories.MovieRepository
}


//获取所有的 movies
func(s *movieService) GetAll() []datamodels.Movie{
	return s.repo.SelectMany(func(_ datamodels.Movie) bool {
		return true
	},-1)
}

/**
根据ID 返回一个 movie.
 */
func (s *movieService) GetByID(id int64) (datamodels.Movie, bool) {
	return s.repo.Select(func(movie datamodels.Movie) bool {
		return movie.ID == id
	})
}

// 更新一个 movie 的 poster 和 genre  字段
func (s *movieService) UpdatePosterAndGenreByID(id int64, poster , genre string) (datamodels.Movie, error){
	return s.repo.InsertOrUpdate(datamodels.Movie{
		ID : id,
		Poster: poster,
		Genre: genre,
	})
}

// 根据ID 删除一个 movie
func(s *movieService) DeleteByID(id int64) bool {
	return s.repo.Delete(func(m datamodels.Movie) bool {
		return m.ID == id
	},1)
}
