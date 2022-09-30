package api

type Cast struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
type Movie struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Rating      string `json:"rating"`
	Casts       []Cast `json:"casts"`
}
