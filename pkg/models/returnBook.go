package models

func ReturnBook(bookId int, userId int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}

	bookExists, _, err := BookIdExists(bookId)
	if err != nil {
		return "", err
	}
	if !bookExists {
		return "Book doesn't Exists.", nil
	} else {

		requestExists, request, err := RequestUserExists(bookId, userId)
		if err != nil {
			return "", err
		}
		if !requestExists {
			return "Invalid Request.", nil
		} else {
			if request.State != "issued" {
				return "Book not issued.", nil
			} else {
				updateSql := `UPDATE requests SET state = 'checkedIn' WHERE bookId= ? AND userId= ? AND state = 'issued';`
				_, err = db.Exec(updateSql, bookId, userId)
				if err != nil {
					return "", err
				} else {
					return "Successfully returned book.", nil
				}
			}
		}
	}
}
