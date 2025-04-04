package repository

import (
	"gorm.io/gorm"
	"rest-project/internal/models"
)

type BookRepositoryImpl struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepositoryImpl {
	return &BookRepositoryImpl{db: db}
}

func (r BookRepositoryImpl) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r BookRepositoryImpl) GetById(id int) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error
	return &book, err
}

func (r BookRepositoryImpl) Create(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r BookRepositoryImpl) Update(id int, book *models.BookEdit) error {
	return r.db.Model(&models.Book{}).Where("id = ?", id).Omit("id").Updates(book).Error
}

func (r BookRepositoryImpl) Delete(bookID int) error {
	return r.db.Delete(&models.Book{}, bookID).Error
}
