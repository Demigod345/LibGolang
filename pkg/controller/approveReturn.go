package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/types"
	"encoding/json"
	"fmt"
	"net/http"
)

func ApproveReturn(writer http.ResponseWriter, request *http.Request) {
	var req types.CompleteRequest
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
	}

	requestExists, approveReturn := models.RequestUserExists(req.BookId, req.UserId)

	if requestExists && approveReturn.State == "checkedIn" {
		models.ApproveReturn(req.RequestId, req.BookId)
	}else{
		fmt.Println("Invalid Request.")
		return
	}

	fmt.Printf("Approving Book req to the database \n")
	
}
