package models

import (
	"LibGolang/pkg/types"
)

func FetchBooks() (types.BookList, error) {
	var listBooks types.BookList

	db, err := Connection()
	if err != nil {
		return listBooks, err
	}

	selectSql := "SELECT * FROM books where bookId > 0;"
	rows, err := db.Query(selectSql)
	db.Close()
	if err != nil {
		return listBooks, err
	}

	var fetchBooks []types.CompleteBook
	for rows.Next() {
		var book types.CompleteBook
		err := rows.Scan(&book.BookId, &book.Title, &book.TotalQuantity, &book.Available)
		if err != nil {
			return listBooks, err
		}
		fetchBooks = append(fetchBooks, book)
	}

	listBooks.Books = fetchBooks
	return listBooks, nil
}