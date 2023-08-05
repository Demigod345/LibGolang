package models

import (
	"fmt"
	"log"
)

func RequestBook(bookId int, userID int) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}

	bookExists, err, book := BookIdExists(bookId)
	if err != nil {
		log.Fatal(err)
	}
	if !bookExists {
		fmt.Println("Book doesn't Exists.")
		return
	} else {
		if book.Available == 0 {
			fmt.Println("Book Out of Stock.")
			return
		} else {
			requestExists := RequestExists(bookId)
			if requestExists {
				fmt.Println("Already Requested.")
				return
			} else {
				insertSql := "INSERT INTO requests (bookId, userId) VALUES (?,?)"
				_, err = db.Exec(insertSql, bookId, userID)
				if err != nil {
					fmt.Printf("error %s inserting into the database", err)
				} else {
					fmt.Println("successfully inserted request into the database")
				}
			}
		}
	}
}
