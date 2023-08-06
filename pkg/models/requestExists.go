package models

import (
	"LibGolang/pkg/types"
	"database/sql"
	"fmt"
	"log"
)

func RequestExists(bookId int) (bool, types.CompleteRequest) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}

	var request types.CompleteRequest

	sqlStmt := `SELECT * FROM requests where bookId=?`
	err = db.QueryRow(sqlStmt, bookId).Scan(&request.RequestId, &request.BookId, &request.UserId,
		&request.State, &request.CreatedAt, &request.UpdatedAt)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			// to "(bool, error)" and return "false, err" here
			log.Print(err)
		}
		return false, request
	}
	return true, request
}

func RequestUserExists(bookId int, userId int) (bool, types.CompleteRequest) {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}

	var request types.CompleteRequest

	sqlStmt := `SELECT * FROM requests where bookId=? and userId= ?`
	err = db.QueryRow(sqlStmt, bookId, userId).Scan(&request.RequestId, &request.BookId, &request.UserId,
		&request.State, &request.CreatedAt, &request.UpdatedAt)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			// to "(bool, error)" and return "false, err" here
			log.Print(err)
		}
		return false, request
	}
	return true, request
}