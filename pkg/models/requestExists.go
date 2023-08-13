package models

import (
	"LibGolang/pkg/types"
	"database/sql"
)

func RequestExists(bookId int) (bool, types.CompleteRequest, error) {

	var request types.CompleteRequest

	db, err := Connection()
	if err != nil {
		return false, request, err
	}

	sqlStmt := `SELECT requestId, bookId, userId, state FROM requests where bookId=?`
	err = db.QueryRow(sqlStmt, bookId).Scan(&request.RequestId, &request.BookId, &request.UserId,
		&request.State)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			// to "(bool, error)" and return "false, err" here
			return false, request, err
		}
		return false, request, nil
	}
	return true, request, nil
}

func RequestUserExists(bookId int, userId int) (bool, types.CompleteRequest, error) {
	var request types.CompleteRequest

	db, err := Connection()
	if err != nil {
		return false, request, err
	}

	sqlStmt := `SELECT requestId, bookId, userId, state FROM requests where bookId=? and userId= ?`
	err = db.QueryRow(sqlStmt, bookId, userId).Scan(&request.RequestId, &request.BookId, &request.UserId,
		&request.State)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			// to "(bool, error)" and return "false, err" here
			return false, request, err
		}
		return false, request, nil
	}
	return true, request, nil
}
