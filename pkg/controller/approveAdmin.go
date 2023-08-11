package controller

import (
	"LibGolang/pkg/models"
	"fmt"
	"net/http"
	"strconv"
)

func ApproveAdmin(writer http.ResponseWriter, request *http.Request) {
	userIdStr := request.FormValue("userId")
	bookIdStr := request.FormValue("bookId")
	requestIdStr := request.FormValue("requestId")

	userId, _ := strconv.Atoi(userIdStr)
	bookId, _ := strconv.Atoi(bookIdStr)
	requestId, _ := strconv.Atoi(requestIdStr)

	fmt.Println(userId, bookId, requestId)

	if bookId == -1 {
		requestExists, approveRequest := models.RequestUserExists(bookId, userId)

		if requestExists && approveRequest.State == "AdminRequest" && approveRequest.RequestId == requestId {
			models.ApproveAdmin(requestId, userId)

		} else {
			fmt.Println("Invalid Request.")
			return
		}
	} else {
		fmt.Println("Invalid Request.")
		return
	}

	fmt.Printf("Approving Book req to the database \n")
	http.Redirect(writer, request, "/admin/adminRequests/AdminRequest", http.StatusSeeOther)

}
