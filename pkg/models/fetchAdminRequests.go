package models

import (
	"LibGolang/pkg/types"
	"fmt"
)

func FetchAllRequestsList(state string) types.RequestList {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}

	selectSql := "SELECT * FROM requests where state= '" + state + "' ;"
	rows, err := db.Query(selectSql)
	db.Close()

	if err != nil {
		fmt.Printf("error %s querying the database", err)
	}

	var fetchRequests []types.CompleteRequest
	for rows.Next() {
		var request types.CompleteRequest
		err := rows.Scan(&request.RequestId, &request.BookId , &request.UserId , &request.State , &request.CreatedAt , &request.UpdatedAt)
		if err != nil {
			fmt.Printf("error %s scanning the row", err)
		}
		fetchRequests = append(fetchRequests, request)
	}

	var requestList types.RequestList
	requestList.Requests = fetchRequests
	return requestList
}