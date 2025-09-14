package service

import (
	"fiber-golang-json-api/internal/models"
	"fiber-golang-json-api/internal/repository"
)

type BookService struct {
	bookRepo *repository.BookRepository
}

func NewBookService(bookRepo *repository.BookRepository) *BookService {
	return &BookService{bookRepo: bookRepo}
}

func (s *BookService) CreateBook(book *models.Book) error {
	return s.bookRepo.Create(book)
}

func (s *BookService) GetBookByID(id uint, userID uint) (*models.Book, error) {
	return s.bookRepo.GetByID(id, userID)
}

func (s *BookService) GetBooksByUserID(userID uint) ([]models.Book, error) {
	return s.bookRepo.GetByUserID(userID)
}

func (s *BookService) GetBooksByUserIDWithFilters(userID uint, filters map[string]string) ([]models.Book, error) {
	return s.bookRepo.GetByUserIDWithFilters(userID, filters)
}

func (s *BookService) UpdateBook(book *models.Book) error {
	return s.bookRepo.Update(book)
}

func (s *BookService) DeleteBook(id uint, userID uint) error {
	return s.bookRepo.Delete(id, userID)
}

func (s *BookService) GetAllBooks() ([]models.Book, error) {
	return s.bookRepo.GetAll()
}
