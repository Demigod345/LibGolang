package controller

import (
	"LibGolang/pkg/models"
	"log"
	"net/http"
	"strconv"
)

func RequestAdmin(writer http.ResponseWriter, request *http.Request){
	bookIdStr := request.FormValue("bookId")
	userId := request.Context().Value(userIdContextKey).(int)
	bookId, _ := strconv.Atoi(bookIdStr)

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