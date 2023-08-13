package controller

import (
	"LibGolang/pkg/models"
	"log"
	"net/http"
	"strconv"
)

func ReturnBook(writer http.ResponseWriter, request *http.Request){
	userIdStr := request.FormValue("userId")
	bookIdStr := request.FormValue("bookId")

	userId, err := strconv.Atoi(userIdStr)
	bookId, err := strconv.Atoi(bookIdStr)
	
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	message, err := models.ReturnBook(bookId, userId)
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}
	SetFlash(writer, request, message)

	http.Redirect(writer, request, "/user/userRequests/issued", http.StatusSeeOther)

}	