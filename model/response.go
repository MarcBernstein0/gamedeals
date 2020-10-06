package model

type Response struct {
	Success bool    `json:"success"`
	Title   string  `json:"title"`
	Price   float64 `json:"Price,string"`
}
