package models

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
}

type Input struct {
	Title  string `json:"title" binding:"required"`
}