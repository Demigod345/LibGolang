package models

import (
	"database/sql"
)

func CheckAdmin(userId int) (bool, error) {
	db, err := Connection()
	if err != nil {
		return false, err
	}

	var isAdmin bool

	sqlSelect := `SELECT isAdmin FROM users WHERE userId = ?`
	err = db.QueryRow(sqlSelect, userId).Scan(&isAdmin)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			// to "(bool, error)" and return "false, err" here
			return false , err
		}

		return isAdmin, nil
	}

	return isAdmin, nil
}