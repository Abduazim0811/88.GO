package postgres

import (
	"Book/internal/models"
	"Book/internal/repository"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
)

type PostgresBookRepository struct {
	Db *sql.DB
}
func NewPostgresBookRepository(db *sql.DB) repository.BookRepository{
	return &PostgresBookRepository{Db: db}
}

func (p *PostgresBookRepository) Create(book models.Book)(*models.Book, error){
	var Book models.Book
	sql, args, err := squirrel.
		Insert("book").
		Columns("title", "author", "publishedyear").
		Values(book.Title, book.Author, book.PublishedYear).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil{
		return nil, fmt.Errorf("unable to insert book")
	}

	row := p.Db.QueryRow(sql, args...)
	
	if err := row.Scan(&Book.Id); err != nil {
		return nil, fmt.Errorf("unable to scan book")
	}

	return &Book, nil
}

func (p *PostgresBookRepository) GetbyId(id int32)(*models.Book, error){
	var Book models.Book

	sql, args, err := squirrel.
		Select("*").
		From("book").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("unable to get book")
	}

	row := p.Db.QueryRow(sql, args...)
	if err := row.Scan(&Book.Id, &Book.Title, &Book.Author, &Book.PublishedYear); err != nil{
		return nil, fmt.Errorf("to scan error")
	}

	return &Book, nil
}

func (p *PostgresBookRepository) GetAll()(*[]models.Book, error){
	var Books []models.Book

	sql, args, err := squirrel.
		Select("*").
		From("Book").
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("unable to get book")
	}

	rows, err := p.Db.Query(sql, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next(){
		var book models.Book

		if err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.PublishedYear); err != nil {
			return nil, fmt.Errorf("to scan error")
		}
		Books = append(Books, book)
	}

	return &Books, nil
}

func (p *PostgresBookRepository) Update(book models.Book) error{
	sql, args, err := squirrel.
		Update("book").
		Set("title", book.Title).
		Set("author", book.Author).
		Set("publishedyear", book.PublishedYear).
		Where(squirrel.Eq{"id": book.Id}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil{
		return fmt.Errorf("unable to update book")
	}
	
	_, err = p.Db.Exec(sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresBookRepository) Delete(id int32) error{
	sql, args, err := squirrel.
		Delete("book").
		Where(squirrel.Eq{"id": id}).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = p.Db.Exec(sql, args...)
	return err
}