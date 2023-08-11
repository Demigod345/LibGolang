package models

import (
	"LibGolang/pkg/types"
	"fmt"
)

func fetchUserRequestsList(state string, userId int) types.RequestList {
	db, err := Connection()
	if err != nil {
		fmt.Printf("error %s connecting to the database", err)
	}

	selectSql := "SELECT * FROM requests where state= '" + state + "' and userId = ? ;"
	rows, err := db.Query(selectSql, userId)
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

func FetchUserRequests(userId int) types.AdminRequests {
	var completeList types.AdminRequests
	completeList.ApproveRequests = fetchUserRequestsList("requested", userId);
	completeList.IssuedBooks = fetchUserRequestsList("issued", userId);
	completeList.ReturnRequests = fetchUserRequestsList("checkedIn", userId);
	completeList.AdminRequests = fetchUserRequestsList("AdminRequest", userId);
	return completeList
}