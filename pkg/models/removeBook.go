package models

import ()

func RemoveBook(bookTitle string, Quantity int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}

	if Quantity > 0 {
		bookExists, book, err := BookTitleExists(bookTitle)
		if err != nil {
			return "", err
		}
		if bookExists {
			newTotalQuantity := book.TotalQuantity - Quantity
			newAvailable := book.Available - Quantity

			if Quantity <= book.Available &&
				newAvailable >= 0 &&
				newTotalQuantity >= 0 {

				updateSql := `UPDATE books SET totalQuantity = ?, available = ? WHERE title= ?;`
				_, err = db.Exec(updateSql, newTotalQuantity, newAvailable, book.Title)
				if err != nil {
					return "", err
				}
				return "Records Updated.", nil
			} else {
				return "Invalid Quantity.", nil
			}
		} else {
			return "Book doesn't Exist.", nil
		}
	} else {
		return "Invalid Quantity.", nil
	}
}

func DeleteBook(bookTitle string) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}

	bookExists, book, err := BookTitleExists(bookTitle)
	if err != nil {
		return "", err
	}
	if bookExists {
		totalQuantity := book.TotalQuantity
		available := book.Available
		bookId := book.BookId
		if available != totalQuantity {
			return "Book can't be deleted.", nil
		} else {
			requestExists, _, err := RequestExists(bookId)
			if err != nil {
				return "", err
			}
			if requestExists {
				return "Book can't be deleted as there are pending requests", nil
			} else {
				deleteSql := `DELETE from books
						WHERE bookId=?;`
				_, err = db.Exec(deleteSql, bookId)
				if err != nil {
					return "", err
				} else {
					return "Successfully deleted the book " + bookTitle, nil
				}
			}
		}
	} else {
		return "Book doesn't Exist.", nil
	}
}
