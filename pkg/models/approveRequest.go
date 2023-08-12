package models

import ()

func ApproveRequest(requestId int, bookId int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}

	updateSql := `UPDATE books SET available = available -1 WHERE bookId = ?;`
	_, err = db.Exec(updateSql, bookId)
	if err != nil {
		return "", err
	}

	update2Sql := `UPDATE requests SET state = 'issued' WHERE requestId= ? AND state = 'requested';`
	_, err = db.Exec(update2Sql, requestId)
	if err != nil {
		return "", err
	}

	return "Successfully Approved Issue Request.", err
}
