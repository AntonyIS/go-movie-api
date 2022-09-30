package api

import "errors"

var (
	ErrInvalidItem  = errors.New("invalid item")
	ErrItemNotFound = errors.New("item not found")
)

type movieRedirectService struct {
	movieRepo MovieRepository
}

func NewMovieService(movieRepo MovieRepository) MovieService {
	return &movieRedirectService{
		movieRepo,
	}
}

func (ms *movieRedirectService) GetMovie(id string) (Movie, error) {
	return ms.movieRepo.GetMovie(id)
}

func (ms *movieRedirectService) GetMovies() ([]Movie, error) {
	return ms.movieRepo.GetMovies()
}

func (ms *movieRedirectService) PostMovie(movie Movie) error {
	return ms.movieRepo.PostMovie(movie)
}

func (ms *movieRedirectService) PutMovie(movie Movie) (Movie, error) {
	return ms.movieRepo.PutMovie(movie)
}

func (ms *movieRedirectService) DeleteMovie(id string) error {
	return ms.movieRepo.DeleteMovie(id)
}
