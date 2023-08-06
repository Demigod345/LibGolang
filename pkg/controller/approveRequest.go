package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ApproveRequest(writer http.ResponseWriter, request *http.Request) {
	var req types.CompleteRequest
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
	}

	requestExists, approveRequest := models.RequestUserExists(req.BookId, req.UserId)

	if requestExists && approveRequest.State == "requested" && approveRequest.RequestId == req.RequestId{
		bookExists, err, book := models.BookIdExists(req.BookId)

		if err != nil {
			log.Fatal(err)
		}
		
		if bookExists {
			if book.Available != 0{
				models.ApproveRequest(req.RequestId, book.BookId)
			} else {
				fmt.Println("Book unavailable.")
			}
		}
	}else{
		fmt.Println("Invalid Request.")
		return
	}

	fmt.Printf("Approving Book req to the database \n")
	
}
