package repository

import (
	"fiber-golang-json-api/internal/models"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Create(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *BookRepository) GetByID(id uint, userID uint) (*models.Book, error) {
	var book models.Book
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) GetByUserID(userID uint) ([]models.Book, error) {
	var books []models.Book
	err := r.db.Where("user_id = ?", userID).Find(&books).Error
	return books, err
}

func (r *BookRepository) GetByUserIDWithFilters(userID uint, filters map[string]string) ([]models.Book, error) {
	var books []models.Book
	query := r.db.Where("user_id = ?", userID)

	if status, ok := filters["status"]; ok && status != "" {
		query = query.Where("status = ?", status)
	}
	if author, ok := filters["author"]; ok && author != "" {
		query = query.Where("author LIKE ?", "%"+author+"%")
	}
	if title, ok := filters["title"]; ok && title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if year, ok := filters["year"]; ok && year != "" {
		query = query.Where("year = ?", year)
	}

	err := query.Find(&books).Error
	return books, err
}

func (r *BookRepository) Update(book *models.Book) error {
	return r.db.Save(book).Error
}

func (r *BookRepository) Delete(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Book{}).Error
}

func (r *BookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}
