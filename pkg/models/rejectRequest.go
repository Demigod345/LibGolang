package models

import ()

func RejectRequest(requestId int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}

	updateSql := `DELETE FROM requests WHERE requestId= ? AND state = 'requested';`
	_, err = db.Exec(updateSql, requestId)
	if err != nil {
		return "", err
	}
	return "Sucessfully Rejected Issue Request", nil
}