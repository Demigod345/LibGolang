package controller

import (
	"LibGolang/pkg/models"

	"log"
	"net/http"
	"strconv"
)

func RejectRequest(writer http.ResponseWriter, request *http.Request) {
	userIdStr := request.FormValue("userId")
	bookIdStr := request.FormValue("bookId")
	requestIdStr := request.FormValue("requestId")

	userId, _ := strconv.Atoi(userIdStr);
	bookId, _ := strconv.Atoi(bookIdStr);
	requestId, _ := strconv.Atoi(requestIdStr);

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
