package api

type MovieSerializer interface {
	Decode(input []byte) (*Movie, error)
	Encode(input *Movie) ([]byte, error)
}
