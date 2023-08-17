package models

func AddBook(bookTitle string, quantity int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}

	if quantity > 0 {
		bookExists, book, err := BookTitleExists(bookTitle)
		if err != nil {
			return "", nil
		}
		if bookExists {
			newTotalQuantity := quantity + book.TotalQuantity
			newAvailable := quantity + book.Available

			updateSql := `UPDATE books SET totalQuantity = ?, available = ? WHERE title= ?;`
			_, err = db.Exec(updateSql, newTotalQuantity, newAvailable, book.Title)
			if err != nil {
				return "", err
			}
			return "Successfully updated the quantity of book " + bookTitle, nil
		} else {
			insertSql := "INSERT INTO books(title, totalQuantity, available) VALUES (?,?,?)"
			_, err = db.Exec(insertSql, bookTitle, quantity, quantity)
			if err != nil {
				return "", err
			}
			return "Successfully added " + bookTitle + " to the library.", nil
		}
	} else {
		return "Invalid Quantity.", nil
	}
}
