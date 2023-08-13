package models

import ()

func RequestBook(bookId int, userID int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}

	bookExists, book, err := BookIdExists(bookId)
	if err != nil {
		return "", err
	}
	if !bookExists {
		return "Book doesn't Exists.", err
	} else {
		if book.Available == 0 {
			return "Book Out of Stock.", err
		} else {
			requestExists,_, err := RequestUserExists(bookId, userID)
			if err != nil {
				return "", err
			}
			if requestExists {
				return "Already Requested.", nil
			} else {
				insertSql := "INSERT INTO requests (bookId, userId) VALUES (?,?)"
				_, err = db.Exec(insertSql, bookId, userID)
				if err != nil {
					return "", err
					} else {
					return "Successfully Requested the book.", err
				}
			}
		}
	}
}
