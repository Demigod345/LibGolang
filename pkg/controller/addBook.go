package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/types"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func AddBook(writer http.ResponseWriter, request *http.Request){
	var body types.Book
	err :=  json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
	}

	fmt.Printf("Adding the book %s to the database, quantity: %s\n", body.Title, body.Quantity)
	quantity, _ := strconv.Atoi(body.Quantity);
	models.AddBook(body.Title, quantity)
}	