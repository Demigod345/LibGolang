package controller

import (
	"LibGolang/pkg/models"
	"log"
	"net/http"
	"strconv"
)

func AddBook(writer http.ResponseWriter, request *http.Request){
	title := request.FormValue("title")
	quantityString := request.FormValue("quantity")

	quantity, err := strconv.Atoi(quantityString);
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}
	
	message, err := models.AddBook(title, quantity)
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}
	SetFlash(writer, request, message)
	http.Redirect(writer,request,"/admin/adminHome", http.StatusSeeOther)
}	