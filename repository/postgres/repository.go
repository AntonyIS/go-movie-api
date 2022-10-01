package postgres

import (
	"fmt"
	"go-movie-app/api"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type movieRepository struct {
	db        *gorm.DB
	tableName string
}

func newPostgresClient() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password))

	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(false)
	db.AutoMigrate(api.Movie{})

	return db, nil
}

func NewMovieRepository(tableName string) (api.MovieRepository, error) {
	repo := movieRepository{
		tableName: tableName,
	}

	db, err := newPostgresClient()

	if err != nil {
		return nil, err
	}
	repo.db = db
	return repo, nil
}

func (m movieRepository) GetMovie(id string) (*api.Movie, error) {
	var movie api.Movie
	m.db.First(&movie, id)
	return &movie, nil
}

func (m movieRepository) GetMovies() ([]*api.Movie, error) {
	movies := []*api.Movie{}
	m.db.Find(&movies)
	return movies, nil
}

func (m movieRepository) PostMovie(movie *api.Movie) error {
	if err := m.db.Create(&movie).Error; err != nil {
		return err
	}
	return nil
}

func (m movieRepository) PutMovie(movie *api.Movie) (*api.Movie, error) {
	if err := m.db.Save(movie).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func (m movieRepository) DeleteMovie(id string) error {
	if err := m.db.Delete(id).Error; err != nil {
		return err
	}
	return nil
}
