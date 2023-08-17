package controller

import (
	"LibGolang/pkg/models"
	"log"
	"net/http"
	"strconv"
)

func RejectRequest(writer http.ResponseWriter, request *http.Request) {
	userIdString := request.FormValue("userId")
	bookIdString := request.FormValue("bookId")
	requestIdString := request.FormValue("requestId")

	userId, err := strconv.Atoi(userIdString);
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	bookId, err := strconv.Atoi(bookIdString);
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	requestId, err := strconv.Atoi(requestIdString);
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	requestExists, approveRequest, err := models.RequestUserExists(bookId, userId)
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}
	if requestExists && approveRequest.State == "requested" {
		message, err := models.RejectRequest(requestId)
		if err != nil {
			log.Println(err)
			http.Redirect(writer, request, "/500", http.StatusSeeOther)
			return
		}
		SetFlash(writer, request, message)	
	} else {
		SetFlash(writer, request, "Invalid Request.")
		return
	}

	http.Redirect(writer,request,"/admin/adminRequests/requested", http.StatusSeeOther)
}
