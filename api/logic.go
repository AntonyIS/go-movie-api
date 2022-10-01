package api

import (
	"errors"

	"github.com/teris-io/shortid"
)

var (
	ErrInvalidMovie  = errors.New("invalid movie")
	ErrMovieNotFound = errors.New("movie not found")
)

type movieService struct {
	movieRepo MovieRepository
}

func NewMovieService(movieRepo MovieRepository) MovieService {
	return &movieService{
		movieRepo,
	}
}

func (ms *movieService) GetMovie(id string) (*Movie, error) {
	return ms.movieRepo.GetMovie(id)
}

func (ms *movieService) GetMovies() ([]*Movie, error) {
	return ms.movieRepo.GetMovies()
}

func (ms *movieService) PostMovie(movie *Movie) error {
	movie.ID = shortid.MustGenerate()
	casts := []Cast{}
	for _, cast := range movie.Casts {
		cast.ID = shortid.MustGenerate()
	}
	movie.Casts = casts
	return ms.movieRepo.PostMovie(movie)
}

func (ms movieService) PutMovie(movie *Movie) (*Movie, error) {
	return ms.movieRepo.PutMovie(movie)
}

func (ms *movieService) DeleteMovie(id string) error {
	return ms.movieRepo.DeleteMovie(id)
}
