package controller

import (
	"LibGolang/pkg/models"
	"log"
	"net/http"
	"strconv"
)

func RemoveBook(writer http.ResponseWriter, request *http.Request) {
	quantityString := request.FormValue("Quantity")
	title := request.FormValue("title")

	quantity, err := strconv.Atoi(quantityString)
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	message, err := models.RemoveBook(title, quantity)
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}
	
	SetFlash(writer, request, message)
	http.Redirect(writer, request, "/admin/adminHome", http.StatusSeeOther)
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	title := request.FormValue("title")

	message, err := models.DeleteBook(title)
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	SetFlash(writer, request, message)
	http.Redirect(writer, request, "/admin/adminHome", http.StatusSeeOther)
}