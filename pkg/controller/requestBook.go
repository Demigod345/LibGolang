package controller

import (
	"LibGolang/pkg/models"
	"fmt"
	"net/http"
	"strconv"
)

func RequestBook(writer http.ResponseWriter, request *http.Request){
	bookIdStr := request.FormValue("bookId")
	userId := request.Context().Value(userIdContextKey).(int)
	bookId, _ := strconv.Atoi(bookIdStr)

	
	fmt.Printf("Adding req to the database \n")
	models.RequestBook(bookId, userId)
	http.Redirect(writer, request, "/user/userHome", http.StatusSeeOther)

}	