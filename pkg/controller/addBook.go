package controller

import (
	"LibGolang/pkg/models"
	"fmt"
	"net/http"
	"strconv"
)

func AddBook(writer http.ResponseWriter, request *http.Request){
	title := request.FormValue("title")
	quantityStr := request.FormValue("quantity")


	fmt.Printf("Adding the book %s to the database, quantity: %s\n", title ,quantityStr)
	quantity, _ := strconv.Atoi(quantityStr);
	models.AddBook(title, quantity)
	http.Redirect(writer,request,"/adminHome", http.StatusSeeOther)
}	