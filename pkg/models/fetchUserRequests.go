package models

import (
	"LibGolang/pkg/types"
)

func FetchUserRequestsList(state string, userId int) (types.RequestList, error) {
	var requestList types.RequestList

	db, err := Connection()
	if err != nil {
		return requestList, err
	}

	selectSql := "SELECT requests.requestId, requests.bookId, requests.userId, requests.state, books.title FROM requests, books WHERE state= '" + state + "' AND userId = ? AND requests.bookId = books.bookId ;"
	if state == "issued" {
		selectSql = "SELECT requests.requestId, requests.bookId, requests.userId, requests.state, books.title FROM requests, books WHERE state in ( 'issued' , 'checkedIn' ) AND userId = ? AND requests.bookId = books.bookId ;"
	}

	rows, err := db.Query(selectSql, userId)
	db.Close()

	if err != nil {
		return requestList, err
	}

	var fetchRequests []types.CompleteRequest
	for rows.Next() {
		var request types.CompleteRequest
		err := rows.Scan(&request.RequestId, &request.BookId, &request.UserId, &request.State, &request.Title)
		if err != nil {
			return requestList, err
		}
		fetchRequests = append(fetchRequests, request)
	}
	requestList.Requests = fetchRequests
	return requestList, nil
}
