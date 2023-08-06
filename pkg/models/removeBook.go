package models

import (
	"fmt"
	"log"
)


func RemoveBook(bookTitle string, Quantity int) {
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
			newTotalQuantity := book.TotalQuantity - Quantity
			newAvailable := book.Available - Quantity
			bookId := book.BookId

			if Quantity <= book.Available &&
				newAvailable >= 0 &&
				newTotalQuantity > 0 {

				updateSql := `UPDATE books SET totalQuantity = ?, available = ? WHERE title= ?;`
				_, err = db.Exec(updateSql, newTotalQuantity, newAvailable, book.Title)
				if err != nil {
					fmt.Printf("error %s updating the database", err)
				} else {
					fmt.Printf("successfully updated the database : %s", bookTitle)
				}

				fmt.Println("Records Updated.")

			} else if newAvailable == newTotalQuantity &&
				newTotalQuantity == 0 {
					requestExists,_ := RequestExists(bookId)

				if requestExists {
					fmt.Println("Book can't be deleted as there are pending requests")
				} else {
					deleteSql := `DELETE from books 
					WHERE bookId=?;`
					_, err = db.Exec(deleteSql, bookId)
					if err != nil {
						fmt.Printf("error %s updating the database", err)
					} else {
						fmt.Printf("successfully deleted the book : %s", bookTitle)
					}
				}

			}

		} else {
			fmt.Println("Book doesn't Exist.")
		}
	} else {
		fmt.Println("Invalid Quantity.")
	}
}
