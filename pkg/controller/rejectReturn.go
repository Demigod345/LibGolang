package controller

import (
	"LibGolang/pkg/models"
	"fmt"
	"net/http"
	"strconv"
)

func RejectReturn(writer http.ResponseWriter, request *http.Request) {
	userIdStr := request.FormValue("userId")
	bookIdStr := request.FormValue("bookId")
	requestIdStr := request.FormValue("requestId")

	userId, _ := strconv.Atoi(userIdStr);
	bookId, _ := strconv.Atoi(bookIdStr);
	requestId, _ := strconv.Atoi(requestIdStr);

	requestExists, rejectReturn := models.RequestUserExists(bookId,userId)

	if requestExists && rejectReturn.State == "checkedIn" && rejectReturn.RequestId == requestId{
				models.RejectReturn(requestId)
			
	}else{
		fmt.Println("Invalid Request.")
		return
	}

	fmt.Printf("Rejecting Return req to the database \n")
	http.Redirect(writer,request,"/admin/adminRequests", http.StatusSeeOther)

}
