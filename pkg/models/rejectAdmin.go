package models

import ()

func RejectAdmin(requestId int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}
	state:= "AdminRequest"
	update2Sql := `DELETE FROM requests WHERE requestId= ? AND state = ?;`
	_, err = db.Exec(update2Sql, requestId, state)
	if err != nil {
		return "", err
	}
	return "Sucessfully Rejected Admin Request", nil
}
