package repositories

import "IrisWebFramework/datamodels"

type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct {

}


func NewMovieManager() MovieRepository {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName() string {
	movie := &datamodels.Movie{Name:"胖墩墩奥运奇遇记"}
	return movie.Name
}

