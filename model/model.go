package model

type Album struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Artist   string  `json:"artist"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

var Albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99, Quantity: 12},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99, Quantity: 3},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99, Quantity: 99},
}
