package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/views"
	"fmt"
	"net/http"
)

func UserRequestsPage(writer http.ResponseWriter, request *http.Request) {
	// http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	// writer.WriteHeader(http.StatusOK)
	t := views.UserRequestsPage()
	userId := request.Context().Value(userIdContextKey).(int)

	requestList := models.FetchUserRequests(userId)
	fmt.Println(requestList)
	t.Execute(writer, requestList)
}