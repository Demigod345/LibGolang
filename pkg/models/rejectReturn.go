package models

import ()

func RejectReturn(requestId int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}
	state:= "checkedIn"
	updateSql := `UPDATE requests SET state = 'issued' WHERE requestId= ? AND state = ?;`
	_, err = db.Exec(updateSql, requestId, state)
	if err != nil {
		return "", err
	}
	return "Sucessfully Rejected Return Request", nil
}
