package controller

import (
	"LibGolang/pkg/models"
	"log"
	"net/http"
	"strconv"
)

func ApproveAdmin(writer http.ResponseWriter, request *http.Request) {
	userIdStr := request.FormValue("userId")
	bookIdStr := request.FormValue("bookId")
	requestIdStr := request.FormValue("requestId")

	userId, err := strconv.Atoi(userIdStr)
	bookId, err := strconv.Atoi(bookIdStr)
	requestId, err := strconv.Atoi(requestIdStr)

	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	if bookId == -1 {
		requestExists, approveRequest, err := models.RequestUserExists(bookId, userId)
		if err != nil {
			log.Println(err)
			http.Redirect(writer, request, "/500", http.StatusSeeOther)
		}

		if requestExists && approveRequest.State == "AdminRequest" && approveRequest.RequestId == requestId {
			message, err := models.ApproveAdmin(requestId, userId)
			if err != nil {
				log.Println(err)
				http.Redirect(writer, request, "/500", http.StatusSeeOther)
			}
			SetFlash(writer, request, message)
		} else {
			SetFlash(writer, request, "Invalid Request.")
		}
	} else {
		SetFlash(writer, request, "Invalid Request.")
	}

	http.Redirect(writer, request, "/admin/adminRequests/AdminRequest", http.StatusSeeOther)

}
