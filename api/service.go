package api

type MovieService interface {
	GetMovie(id string) (Movie, error)
	GetMovies() ([]Movie, error)
	PostMovie(movie Movie) error
	PutMovie(movie Movie) (Movie, error)
	DeleteMovie(id string) error
}
