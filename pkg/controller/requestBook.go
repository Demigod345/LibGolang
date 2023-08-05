package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/types"
	"encoding/json"
	"fmt"
	"net/http"
)

func RequestBook(writer http.ResponseWriter, request *http.Request){
	var req types.CompleteRequest
	err :=  json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
	}

	fmt.Printf("Adding req to the database \n")
	models.RequestBook(req.BookId, req.UserId)
}	