package controller

import (
	"LibGolang/pkg/models"
	"fmt"
	"net/http"
	"strconv"
)

func ReturnBook(writer http.ResponseWriter, request *http.Request){
	userIdStr := request.FormValue("userId")
	bookIdStr := request.FormValue("bookId")

	userId, _ := strconv.Atoi(userIdStr)
	bookId, _ := strconv.Atoi(bookIdStr)

	fmt.Printf("Adding req to the database \n")
	models.ReturnBook(bookId, userId)
	http.Redirect(writer, request, "/user/userRequests/issued", http.StatusSeeOther)

}	