package database

type Plate struct {
	Id         string  `json:"id"`
	Color      string  `json:"color"`
	Price      float64 `json:"price"`
	Dimensions string  `json:"dimensions"`
}

var Plates []Plate