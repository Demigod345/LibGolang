package models

import (
	"fmt"
	// "log"
)

func RequestAdmin(userID int) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}

	requestExists, _ := RequestUserExists(-1, userID)

	if requestExists {
		fmt.Println("Already Requested.")
		return
	} else {
		insertSql := "INSERT INTO requests (bookId, userId, state) VALUES (?,?, 'AdminRequest')"
		_, err = db.Exec(insertSql, -1, userID)
		if err != nil {
			fmt.Printf("error %s inserting into the database", err)
		} else {
			fmt.Println("successfully inserted request into the database")
		}
		
	}
}
