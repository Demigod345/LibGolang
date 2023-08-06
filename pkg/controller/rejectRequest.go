package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/types"
	"encoding/json"
	"fmt"
	"net/http"
)

func RejectRequest(writer http.ResponseWriter, request *http.Request) {
	var req types.CompleteRequest
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
	}

	requestExists, approveRequest := models.RequestUserExists(req.BookId, req.UserId)

	if requestExists && approveRequest.State == "requested" {
		models.RejectRequest(req.RequestId)
	} else {
		fmt.Println("Invalid Request.")
		return
	}

	fmt.Printf("Removing Book req from the database \n")

}
