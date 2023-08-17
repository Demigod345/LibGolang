package models

import (
	"LibGolang/pkg/types"
	"database/sql"
)

func UserExists(db *sql.DB ,username string) (bool, types.User, error) {
	var user types.User

	selectSql := `SELECT * FROM users where username = ?`
	err := db.QueryRow(selectSql, username).Scan(&user.UserId , &user.UserName , &user.Hash , &user.IsAdmin )
	if err != nil {
		if err != sql.ErrNoRows {
			return false, user, err		
		}
		return false, user, nil
	}
	return true, user, nil
}