package model

import "time"

type Book struct {
	ID          string    `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Author      string    `json:"author" binding:"required"`
	ISBN        string    `json:"isbn" binding:"required"`
	Description string    `json:"description"`
	Price       float64   `json:"price" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
