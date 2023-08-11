package controller

import (
	"LibGolang/pkg/models"
	"fmt"
	"net/http"
	"strconv"
)

func RequestAdmin(writer http.ResponseWriter, request *http.Request){
	bookIdStr := request.FormValue("bookId")
	userId := request.Context().Value(userIdContextKey).(int)
	bookId, _ := strconv.Atoi(bookIdStr)

	if bookId == -1 {
		fmt.Printf("Adding req to the database \n")
		models.RequestAdmin( userId)
	} else {
		// wrong request
		fmt.Println("Wrong request")
		

	}

	
	
}	