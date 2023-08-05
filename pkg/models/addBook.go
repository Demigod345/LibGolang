package models

import (
	"fmt"
	"log"
)

func AddBook(bookTitle string, Quantity int) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}

	if Quantity > 0 {
		bookExists, err, book := BookTitleExists(bookTitle)
		if err != nil {
			log.Fatal(err)
		}
		if bookExists {

			newTotalQuantity := Quantity + book.TotalQuantity
			newAvailable := Quantity + book.Available

			updateSql := `UPDATE books SET totalQuantity = ?, available = ? WHERE title= ?;`
			_, err = db.Exec(updateSql, newTotalQuantity, newAvailable, book.Title)
			if err != nil {
				fmt.Printf("error %s updating the database", err)
			} else {
				fmt.Printf("successfully updated the database : %s", bookTitle)
			}

		} else {
			insertSql := "INSERT INTO books(title, totalQuantity, available) VALUES (?,?,?)"
			_, err = db.Exec(insertSql, bookTitle, Quantity, Quantity)
			if err != nil {
				fmt.Printf("error %s inserting into the database", err)
			} else {
				fmt.Printf("successfully inserted %s into the database", bookTitle)
			}
		}
	} else {
		fmt.Println("Invalid Quantity.")
	}
}
