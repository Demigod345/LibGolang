package models

import (
	"LibGolang/pkg/types"
	"database/sql"
	// "log"
)

func UserExists(db *sql.DB ,username string) (bool, types.User, error) {
	var user types.User

	sqlStmt := `SELECT * FROM users where username = ?`
	err := db.QueryRow(sqlStmt, username).Scan(&user.UserId , &user.UserName , &user.Hash , &user.IsAdmin )
	// log.Println(err)
	// log.Println(user)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, user, err		
		}
		return false, user, nil
	}
	return true, user, nil
}