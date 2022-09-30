package serializer

import (
	"encoding/json"
	"go-movie-app/api"

	"github.com/pkg/errors"
)

type Movie struct{}

func (m *Movie) Decode(input []byte) (*api.Movie, error) {
	movie := &api.Movie{}
	if err := json.Unmarshal(input, movie); err != nil {
		return nil, errors.Wrap(err, "Serializer.Movie.Decode")
	}
	return movie, nil
}

func (m *Movie) Encode(movie *Movie) ([]byte, error) {
	bs, err := json.Marshal(movie)
	if err != nil {
		return nil, errors.Wrap(err, "serlializer.Movie.Encode")
	}
	return bs, nil
}
