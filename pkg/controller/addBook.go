package controller

import (
	"LibGolang/pkg/models"
	"log"
	"net/http"
	"strconv"
)

func AddBook(writer http.ResponseWriter, request *http.Request){
	title := request.FormValue("title")
	quantityStr := request.FormValue("quantity")

	quantity, _ := strconv.Atoi(quantityStr);
	message, err := models.AddBook(title, quantity)
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}
	SetFlash(writer, request, message)
	http.Redirect(writer,request,"/admin/adminHome", http.StatusSeeOther)
}	