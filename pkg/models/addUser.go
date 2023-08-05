package models

import (
	"fmt"
)

func AddUser(username string, hash string) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}

	insertSql := "INSERT INTO users (userName, hash, isAdmin) VALUES(?,?, false)"
	_, err = db.Exec(insertSql, username, hash)
	if err != nil {
		fmt.Printf("error %s inserting into the database", err)
	} else {
		fmt.Printf("successfully inserted %s into the database", username)
	}
}
