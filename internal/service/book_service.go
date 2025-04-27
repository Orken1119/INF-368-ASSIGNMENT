package service

import (
	"INF368/internal/repository"
	"INF368/internal/models"
)

type BookService struct {
	repo *repository.BookRepository
	lastID int
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) AddBook(title string) {
	s.lastID++
	book := models.Book{
		ID:     s.lastID,
		Title:  title,
	}
	s.repo.AddBook(book)
}

func (s *BookService) GetAllBooks() []models.Book {
	return s.repo.GetAllBooks()
}

func (s *BookService) DeleteBook(id int) bool {
	return s.repo.DeleteBook(id)
}
