package models

import (
	"LibGolang/pkg/types"
)

func FetchAllRequestsList(state string) (types.RequestList, error) {
	var requestList types.RequestList

	db, err := Connection()
	if err != nil {
		return requestList, err
	}

	selectSql := "SELECT requests.requestId, requests.bookId, requests.state, requests.userId, books.title, books.available, books.totalQuantity FROM requests, books WHERE state= '" + state + "' AND requests.bookId = books.bookId ;"
	
	if state == "issued" {
		selectSql = "SELECT requests.requestId, requests.bookId, requests.state, requests.userId, books.title, books.available, books.totalQuantity FROM requests, books WHERE state in ( 'issued' , 'checkedIn' ) AND requests.bookId = books.bookId ;"
	}
	
	rows, err := db.Query(selectSql)
	db.Close()

	if err != nil {
		return requestList, err
	}

	var fetchRequests []types.CompleteRequest
	for rows.Next() {
		var request types.CompleteRequest
		err := rows.Scan(&request.RequestId, &request.BookId , &request.State, &request.UserId, &request.Title , &request.Available , &request.TotalQuantity)
		if err != nil {
			return requestList, err
		}
		fetchRequests = append(fetchRequests, request)
	}

	requestList.Requests = fetchRequests
	return requestList, nil
}