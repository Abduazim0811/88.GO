package service

import (
	"Book/internal/models"
	"Book/internal/repository"
)

type BookService struct {
	Repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{Repo: repo}
}

func (b *BookService) CreateBooks(book models.Book) (*models.Book, error) {
	return b.Repo.Create(book)
}

func (b *BookService) GetbyIdBooks(id int32) (*models.Book, error) {
	return b.Repo.GetbyId(id)
}

func (b *BookService) GetAllBooks()(*[]models.Book, error){
	return b.Repo.GetAll()
}

func (b *BookService) UpdateBooks(book models.Book) error{
	return b.Repo.Update(book)
}

func (b *BookService) DeleteBooks(id int32) error{
	return b.Repo.Delete(id)
}