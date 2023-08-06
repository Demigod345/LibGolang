package models

import (
	"fmt"
	"log"
)

func ReturnBook(bookId int, userId int) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}

	bookExists, err, _ := BookIdExists(bookId)
	if err != nil {
		log.Fatal(err)
	}
	if !bookExists {
		fmt.Println("Book doesn't Exists.")
		return
	} else {

		requestExists, request := RequestUserExists(bookId, userId)
		if !requestExists {
			fmt.Println("Invalid Request.")
			return
		} else {
			if request.State != "issued" {
				fmt.Println("Book not issued.")
				return
			} else {
				updateSql := `UPDATE requests SET state = 'checkedIn' WHERE bookId= ? AND userId= ? AND state = 'issued';`
				_, err = db.Exec(updateSql, bookId, userId)
				if err != nil {
					fmt.Printf("error %s updating the database", err)
				} else {
					fmt.Printf("Successfully returned book.")
				}
			}
		}
	}

	fmt.Println("Models ReturnBook() Function")
}
