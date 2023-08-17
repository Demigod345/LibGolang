package models

import ()

func RejectRequest(requestId int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}
	state:= "requested"
	updateSql := `DELETE FROM requests WHERE requestId= ? AND state = ? ;`
	_, err = db.Exec(updateSql, requestId,state)
	if err != nil {
		return "", err
	}
	return "Sucessfully Rejected Issue Request", nil
}