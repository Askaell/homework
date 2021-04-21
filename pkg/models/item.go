package models

type Item struct {
	Id          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Price       float32 `json:"price" db:"price"`
}
