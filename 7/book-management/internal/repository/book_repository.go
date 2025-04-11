package repository

import (
	"fmt"
	"go-bookstore/internal/model"
	"sync"
	"time"

	"github.com/google/uuid"
)

// BookRepository interface defines the methods for interacting with the book data store.
type BookRepository interface {
	FindAll() ([]model.Book, error)
	FindByID(id string) (model.Book, error)
	Create(book model.Book) (model.Book, error)
	Update(id string, book model.Book) (model.Book, error)
	Delete(id string) error
}

// InMemoryBookRepository is an in-memory implementation of the BookRepository interface.
type InMemoryBookRepository struct {
	mutex sync.RWMutex
	books map[string]model.Book
}

// NewBookRepository creates a new instance of BookRepository
func NewBookRepository() BookRepository {

	// Pre-populate with some sample data
	books := make(map[string]model.Book)

	// Add sample books
	sampleBooks := []model.Book{
		{
			ID:          uuid.New().String(),
			Title:       "The Go Programming Language",
			Author:      "Alan A. A. Donovan, Brian W. Kernighan",
			ISBN:        "978-0134190440",
			Description: "The Go Programming Language is the authoritative resource for any programmer who wants to learn Go.",
			Price:       34.99,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          uuid.New().String(),
			Title:       "Go in Action",
			Author:      "William Kennedy, Brian Ketelsen, Erik St. Martin",
			ISBN:        "978-1617291784",
			Description: "Go in Action introduces the Go language, guiding you from inquisitive developer to Go guru.",
			Price:       29.99,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	for _, book := range sampleBooks {
		books[book.ID] = book
	}

	return &InMemoryBookRepository{
		books: books,
	}

}

// Create implements BookRepository.
func (i *InMemoryBookRepository) Create(book model.Book) (model.Book, error) {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	// Generate a new UUID for the book
	if book.ID == "" {
		book.ID = uuid.New().String()
	}

	// Set timestamps
	now := time.Now()
	book.CreatedAt = now
	book.UpdatedAt = now

	i.books[book.ID] = book

	return book, nil
}

// Delete implements BookRepository.
func (i *InMemoryBookRepository) Delete(id string) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	if _, exists := i.books[id]; !exists {
		return fmt.Errorf("book with ID %s not found", id)
	}
	delete(i.books, id)
	return nil
}

// FindAll implements BookRepository.
func (i *InMemoryBookRepository) FindAll() ([]model.Book, error) {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	books := make([]model.Book, 0, len(i.books))
	for _, book := range i.books {
		books = append(books, book)
	}

	return books, nil
}

// FindByID implements BookRepository.
func (i *InMemoryBookRepository) FindByID(id string) (model.Book, error) {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	book, exists := i.books[id]

	if !exists {
		return model.Book{}, fmt.Errorf("book with ID %s not found", id)
	}
	return book, nil
}

// Update implements BookRepository.
func (i *InMemoryBookRepository) Update(id string, book model.Book) (model.Book, error) {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	if _, exists := i.books[id]; !exists {
		return model.Book{}, fmt.Errorf("book with ID %s not found", id)
	}

	existingBook := i.books[id]
	book.ID = id
	book.CreatedAt = existingBook.CreatedAt
	book.UpdatedAt = time.Now()

	i.books[id] = book

	return book, nil
}
