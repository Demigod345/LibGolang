package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/types"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func RemoveBook(writer http.ResponseWriter, request *http.Request){
	var body types.Book
	err :=  json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
	}

	fmt.Printf("Removing the book %s from the database, quantity: %s\n", body.Title, body.Quantity)
	quantity, _ := strconv.Atoi(body.Quantity);
	models.RemoveBook(body.Title, quantity)
}	