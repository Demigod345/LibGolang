package models

import (
	"LibGolang/pkg/types"
	"fmt"
)

func FetchBooks() types.BookList {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}

	selectSql := "SELECT * FROM books;"
	rows, err := db.Query(selectSql)
	db.Close()

	if err != nil {
		fmt.Printf("error %s querying the database", err)
	}

	var fetchBooks []types.CompleteBook
	for rows.Next() {
		var book types.CompleteBook
		err := rows.Scan(&book.BookId, &book.Title, &book.TotalQuantity, &book.Available)
		if err != nil {
			fmt.Printf("error %s scanning the row", err)
		}
		fetchBooks = append(fetchBooks, book)
	}

	var listBooks types.BookList
	listBooks.Books = fetchBooks
	return listBooks

}