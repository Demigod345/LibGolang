package controller

import (
	"LibGolang/pkg/models"
	"log"
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

	requestExists, approveReturn, err := models.RequestUserExists(bookId, userId)
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}
	if requestExists && approveReturn.State == "checkedIn" {
		message, err := models.ApproveReturn(requestId, bookId)
		if err != nil {
			log.Println(err)
			http.Redirect(writer, request, "/500", http.StatusSeeOther)
			return
		}
		SetFlash(writer, request, message)
	}else{
		SetFlash(writer, request, "Invalid Request.")
		return
	}

	http.Redirect(writer,request,"/admin/adminRequests/checkedIn", http.StatusSeeOther)

	
}
