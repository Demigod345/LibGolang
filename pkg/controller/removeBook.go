package controller

import (
	"LibGolang/pkg/models"
	// "LibGolang/pkg/types"
	// "encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func RemoveBook(writer http.ResponseWriter, request *http.Request){
	// var body types.Book
	// err :=  json.NewDecoder(request.Body).Decode(&body)
	// if err != nil {
	// 	fmt.Print("There was an error decoding the request body into the struct")
	// }

	quantityStr := request.FormValue("Quantity")
	title := request.FormValue("title")

	fmt.Printf("Removing the book %s from the database, quantity: %s\n",  title, quantityStr)
	quantity, _ := strconv.Atoi(quantityStr);
	models.RemoveBook(title, quantity)
	http.Redirect(writer,request,"/adminHome", http.StatusSeeOther)
}	