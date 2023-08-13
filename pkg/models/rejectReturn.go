package models

import ()

func RejectReturn(requestId int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}
	updateSql := `UPDATE requests SET state = 'issued' WHERE requestId= ? AND state = 'checkedIn';`
	_, err = db.Exec(updateSql, requestId)
	if err != nil {
		return "", err
	}
	return "Sucessfully Rejected Return Request", nil
}
