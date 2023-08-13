package controller

import (
	"LibGolang/pkg/middleware"
	"LibGolang/pkg/models"
	"log"
	"net/http"
	"strconv"
)

func RequestBook(writer http.ResponseWriter, request *http.Request){
	bookIdString := request.FormValue("bookId")
	userId := request.Context().Value(middleware.UserIdContextKey).(int)
	bookId, err := strconv.Atoi(bookIdString)
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}
	
	message, err := models.RequestBook(bookId, userId)
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	SetFlash(writer, request, message)
	http.Redirect(writer, request, "/user/userHome", http.StatusSeeOther)
}	