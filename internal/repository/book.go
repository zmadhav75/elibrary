package repository

import (
	"database/sql"
	"elibrary/internal/models"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) CreateBook(book *models.BookRequest) error {
	_, err := r.db.Exec(
		`INSERT INTO books (title, author, isbn, available) 
		VALUES ($1, $2, $3, $4)`,
		book.Title, book.Author, book.ISBN, book.Available,
	)
	return err
}

func (r *BookRepository) GetBookByID(id int) (*models.Book, error) {
	var book models.Book
	err := r.db.QueryRow(
		`SELECT id, title, author, isbn, available 
		FROM books WHERE id = $1`, id,
	).Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Available)

	return &book, err
}

func (r *BookRepository) GetAllBooks() ([]models.Book, error) {
	rows, err := r.db.Query(`
		SELECT id, title, author, isbn, available 
		FROM books
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.ISBN,
			&book.Available,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
