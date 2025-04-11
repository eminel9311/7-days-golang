package service

import (
	"go-bookstore/internal/model"
	"go-bookstore/internal/repository"
)

// BookService interface
type BookService interface {
	GetAllBooks() ([]model.Book, error)
	GetBookByID(id string) (model.Book, error)
	CreateBook(book model.Book) (model.Book, error)
	UpdateBook(id string, book model.Book) (model.Book, error)
	DeleteBook(id string) error
}

// bookService implements BookService
type bookService struct {
	repo repository.BookRepository
}

// NewBookService creates a new instance of BookService
func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{
		repo: repo,
	}
}

// GetAllBooks retrieves all books from the repository
func (s *bookService) GetAllBooks() ([]model.Book, error) {
	books, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return books, nil
}

// GetBookByID retrieves a book by its ID from the repository
func (s *bookService) GetBookByID(id string) (model.Book, error) {
	book, err := s.repo.FindByID(id)
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

// CreateBook creates a new book in the repository
func (s *bookService) CreateBook(book model.Book) (model.Book, error) {
	book, err := s.repo.Create(book)
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

// UpdateBook updates an existing book in the repository
func (s *bookService) UpdateBook(id string, book model.Book) (model.Book, error) {
	book, err := s.repo.Update(id, book)
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

// DeleteBook deletes a book by its ID from the repository
func (s *bookService) DeleteBook(id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
