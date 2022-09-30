package api

type Cast struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Rating      string `json:"rating"`
	Casts       []Cast `json:"casts"`
}
