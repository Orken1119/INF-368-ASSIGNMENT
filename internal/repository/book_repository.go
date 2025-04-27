package repository

import (
	"INF368/internal/models"
)

type BookRepository struct {
	books map[int]string
}

func NewBookRepository() *BookRepository {
	return &BookRepository{
		books: make(map[int]string), 
	}
}

func (r *BookRepository) AddBook(book models.Book) {
	r.books[book.ID] = book.Title
}

func (r *BookRepository) GetAllBooks() []models.Book {
	books := make([]models.Book, 0, len(r.books)) 
	for id, title := range r.books {
		books = append(books, models.Book{ID: id, Title: title}) 
	}
	return books
}

func (r *BookRepository) DeleteBook(id int) bool {
	if _, exists := r.books[id]; exists {
		delete(r.books, id) 
		return true
	}
	return false
}
