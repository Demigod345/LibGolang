package models

import (
	"database/sql"
	"fmt"
)

func CheckAdmin(userId int) (bool, error) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}

	var isAdmin bool

	sqlStmt := `SELECT isAdmin FROM users WHERE userId = ?`
	err = db.QueryRow(sqlStmt, userId).Scan(&isAdmin)
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