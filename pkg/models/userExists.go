package models

import (
	"LibGolang/pkg/types"
	"database/sql"
	"fmt"
	"log"
)

func UserExists(userName string) (bool, types.User) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}

	var user types.User

	sqlStmt := `SELECT * FROM users where userName=?`
	err = db.QueryRow(sqlStmt, userName).Scan(&user.UserId , &user.UserName , &user.Hash , &user.IsAdmin )
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			// to "(bool, error)" and return "false, err" here
			log.Print(err)
		
		}
		return false, user
	}
	return true, user
}