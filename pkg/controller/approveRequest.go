package controller

import (
	"LibGolang/pkg/models"
	"log"
	"net/http"
	"strconv"
)

func ApproveRequest(writer http.ResponseWriter, request *http.Request) {
	userIdStr := request.FormValue("userId")
	bookIdStr := request.FormValue("bookId")
	requestIdStr := request.FormValue("requestId")

	userId, _ := strconv.Atoi(userIdStr)
	bookId, _ := strconv.Atoi(bookIdStr)
	requestId, _ := strconv.Atoi(requestIdStr)

	requestExists, approveRequest, err := models.RequestUserExists(bookId, userId)
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	if requestExists && approveRequest.State == "requested" && approveRequest.RequestId == requestId {
		bookExists, book, err := models.BookIdExists(bookId)

		if err != nil {
			log.Println(err)
			http.Redirect(writer, request, "/500", http.StatusSeeOther)
			return
		}

		if bookExists {
			if book.Available != 0 {
				message, err := models.ApproveRequest(requestId, book.BookId)
				if err != nil {
					log.Println(err)
					http.Redirect(writer, request, "/500", http.StatusSeeOther)
					return
				}
				SetFlash(writer, request, message)
			} else {
				SetFlash(writer, request, "Book Unavailable.")
			}
		}
	} else {
		SetFlash(writer, request, "Invalid Request.")
		return
	}

	http.Redirect(writer, request, "/admin/adminRequests/requested", http.StatusSeeOther)

}
