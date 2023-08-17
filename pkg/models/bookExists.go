package models

import (
	"LibGolang/pkg/types"
	"database/sql"
)

func BookTitleExists(bookTitle string) (bool, types.CompleteBook, error) {
	var book types.CompleteBook

	db, err := Connection()
	if err != nil {
		return false, book, err
	}

	selectSql := `SELECT * FROM books WHERE title = ?`
	err = db.QueryRow(selectSql, bookTitle).Scan(&book.BookId, &book.Title, &book.TotalQuantity, &book.Available)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, book, err
		}
		return false, book, nil
	}
	return true, book, nil
}

func BookIdExists(bookId int) (bool, types.CompleteBook, error) {
	var book types.CompleteBook

	db, err := Connection()
	if err != nil {
		return false, book, err
	}

	selectSql := `SELECT * FROM books WHERE bookId = ?`
	err = db.QueryRow(selectSql, bookId).Scan(&book.BookId, &book.Title, &book.TotalQuantity, &book.Available)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, book, err
		}
		return false, book, nil
	}
	return true, book, nil
}
