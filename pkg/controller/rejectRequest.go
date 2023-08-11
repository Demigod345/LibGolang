package controller

import (
	"LibGolang/pkg/models"

	"fmt"
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

	requestExists, approveRequest := models.RequestUserExists(bookId, userId)

	if requestExists && approveRequest.State == "requested" {
		models.RejectRequest(requestId)
	} else {
		fmt.Println("Invalid Request.")
		return
	}

	fmt.Printf("Removing Book req from the database \n")
	http.Redirect(writer,request,"/admin/adminRequests", http.StatusSeeOther)

}
