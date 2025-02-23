package booksrepository

import (
	"database/sql"
	"fmt"

	"github.com/sunzhqr/book_library/internal/books/model"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Get() ([]model.Book, error) {
	rows, err := r.db.Query("SELECT id, title, author, available FROM books")
	if err != nil {
		return nil, fmt.Errorf("bookRepository.Get: %w", err)
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Available)
		if err != nil {
			continue
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *Repository) Add(req model.Book) error {
	_, err := r.db.Exec("INSERT INTO books (title, author, available) VALUES ($1, $2, $3)",
		req.Title, req.Author, true)
	if err != nil {
		return fmt.Errorf("bookRepository.Add: %w", err)
	}

	return nil
}
