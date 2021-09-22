package models

type dress struct {
	ID       string  `json:"id"`
	NAME     string  `json:"name"`
	PRICE    float64 `json:"price"`
	MATERIAL string  `json:material`
}

var Dresses = []dress{
	{ID: "1", NAME: "Sweater", PRICE: 1000, MATERIAL: "wool"},
	{ID: "2", NAME: "Shirt", PRICE: 500, MATERIAL: "cotton"},
	{ID: "3", NAME: "Jacket", PRICE: 1500, MATERIAL: "Silk"},
}
