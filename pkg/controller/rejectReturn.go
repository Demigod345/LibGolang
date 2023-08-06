package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/types"
	"encoding/json"
	"fmt"
	"net/http"
)

func RejectReturn(writer http.ResponseWriter, request *http.Request) {
	var req types.CompleteRequest
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
	}

	requestExists, rejectReturn := models.RequestUserExists(req.BookId, req.UserId)

	if requestExists && rejectReturn.State == "checkedIn" && rejectReturn.RequestId == req.RequestId{
				models.RejectReturn(req.RequestId)
			
	}else{
		fmt.Println("Invalid Request.")
		return
	}

	fmt.Printf("Rejecting Return req to the database \n")
	
}
