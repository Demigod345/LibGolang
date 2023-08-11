package controller

import (
	"LibGolang/pkg/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func ApproveRequest(writer http.ResponseWriter, request *http.Request) {
	userIdStr := request.FormValue("userId")
	bookIdStr := request.FormValue("bookId")
	requestIdStr := request.FormValue("requestId")

	userId, _ := strconv.Atoi(userIdStr);
	bookId, _ := strconv.Atoi(bookIdStr);
	requestId, _ := strconv.Atoi(requestIdStr);

	fmt.Println(userId, bookId, requestId)


	requestExists, approveRequest := models.RequestUserExists(bookId, userId)

	if requestExists && approveRequest.State == "requested" && approveRequest.RequestId == requestId{
		bookExists, err, book := models.BookIdExists(bookId)

		if err != nil {
			log.Fatal(err)
		}
		
		if bookExists {
			if book.Available != 0{
				models.ApproveRequest(requestId, book.BookId)
			} else {
				fmt.Println("Book unavailable.")
			}
		}
	}else{
		fmt.Println("Invalid Request.")
		return
	}

	fmt.Printf("Approving Book req to the database \n")
	http.Redirect(writer,request,"/admin/adminRequests", http.StatusSeeOther)
	
}
