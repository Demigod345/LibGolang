package models

import (
	"LibGolang/pkg/types"
	"database/sql"
)

func UserExists(userName string) (bool, types.User, error) {
	var user types.User

	db, err := Connection()
	if err != nil {
		return false, user, err
	}


	sqlStmt := `SELECT * FROM users where userName=?`
	err = db.QueryRow(sqlStmt, userName).Scan(&user.UserId , &user.UserName , &user.Hash , &user.IsAdmin )
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			// to "(bool, error)" and return "false, err" here
			return false, user, err		
		}
		return false, user, nil
	}
	return true, user, nil
}