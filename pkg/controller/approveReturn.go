package controller

import (
	"LibGolang/pkg/models"
	"fmt"
	"net/http"
	"strconv"
)

func ApproveReturn(writer http.ResponseWriter, request *http.Request) {
	userIdStr := request.FormValue("userId")
	bookIdStr := request.FormValue("bookId")
	requestIdStr := request.FormValue("requestId")

	userId, _ := strconv.Atoi(userIdStr);
	bookId, _ := strconv.Atoi(bookIdStr);
	requestId, _ := strconv.Atoi(requestIdStr);

	requestExists, approveReturn := models.RequestUserExists(bookId, userId)

	if requestExists && approveReturn.State == "checkedIn" {
		models.ApproveReturn(requestId, bookId)
	}else{
		fmt.Println("Invalid Request.")
		return
	}

	fmt.Printf("Approving Book req to the database \n")
	http.Redirect(writer,request,"/admin/adminRequests/issued", http.StatusSeeOther)

	
}
