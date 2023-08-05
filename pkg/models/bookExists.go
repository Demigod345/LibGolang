package models

import (
	"LibGolang/pkg/types"
	"database/sql"
	"fmt"
)

func BookTitleExists(bookTitle string) (bool, error, types.CompleteBook) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}

	var book types.CompleteBook

	sqlStmt := `SELECT * FROM books WHERE title = ?`
	err = db.QueryRow(sqlStmt, bookTitle).Scan(&book.BookId, &book.Title, &book.TotalQuantity, &book.Available)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			// to "(bool, error)" and return "false, err" here
			return false, err, book
		}

		return false, nil, book
	}

	return true, nil, book
}

func BookIdExists(bookId int) (bool, error, types.CompleteBook) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}

	var book types.CompleteBook

	sqlStmt := `SELECT * FROM books WHERE bookId = ?`
	err = db.QueryRow(sqlStmt, bookId).Scan(&book.BookId, &book.Title, &book.TotalQuantity, &book.Available)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			// to "(bool, error)" and return "false, err" here
			return false, err, book
		}

		return false, nil, book
	}

	return true, nil, book
}
