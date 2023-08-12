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

	selectSql := "SELECT * FROM requests where state= '" + state + "' ;"
	
	if state == "issued" {
		selectSql = "SELECT * FROM requests where state in ( 'issued' , 'checkedIn' );"
	}
	
	rows, err := db.Query(selectSql)
	db.Close()

	if err != nil {
		return requestList, err
	}

	var fetchRequests []types.CompleteRequest
	for rows.Next() {
		var request types.CompleteRequest
		err := rows.Scan(&request.RequestId, &request.BookId , &request.UserId , &request.State , &request.CreatedAt , &request.UpdatedAt)
		if err != nil {
			return requestList, err
		}
		fetchRequests = append(fetchRequests, request)
	}

	requestList.Requests = fetchRequests
	return requestList, nil
}