package repository

import "Book/internal/models"

type BookRepository interface {
	Create(book models.Book) (*models.Book, error)
	GetbyId(id int32)(*models.Book, error)
	GetAll()(*[]models.Book, error)
	Update(book models.Book)error
	Delete(id int32)error
}
