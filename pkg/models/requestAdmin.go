package models

import ()

func RequestAdmin(userID int) (string, error) {
	db, err := Connection()
	if err != nil {
		return "", err
	}
	requestExists, _, err := RequestUserExists(-1, userID)
	if err != nil {
		return "", err
	}
	if requestExists {
		return "Already Requested.", nil
	} else {
		insertSql := "INSERT INTO requests (bookId, userId, state) VALUES (?,?, 'AdminRequest')"
		_, err = db.Exec(insertSql, -1, userID)
		if err != nil {
			return "", err
		} else {
			return "Successfully Requested for Admin", err
		}
	}
}
