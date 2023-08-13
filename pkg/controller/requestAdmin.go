package controller

import (
	"LibGolang/pkg/middleware"
	"LibGolang/pkg/models"
	"log"
	"net/http"
	"strconv"
)

func RequestAdmin(writer http.ResponseWriter, request *http.Request){
	bookIdString := request.FormValue("bookId")
	userId := request.Context().Value(middleware.UserIdContextKey).(int)
	bookId, err := strconv.Atoi(bookIdString)
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	if bookId == -1 {
		message, err := models.RequestAdmin( userId)
		if err != nil {
			log.Println(err)
			http.Redirect(writer, request, "/500", http.StatusSeeOther)
			return
		}
		SetFlash(writer, request, message)
	} else {
		SetFlash(writer, request, "Wrong request")
	}

	http.Redirect(writer, request, "/user/userRequests/AdminRequest", http.StatusSeeOther)

	
}	