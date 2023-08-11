package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/views"
	"fmt"
	"net/http"
)

func AdminRequestsPage(writer http.ResponseWriter, request *http.Request) {
	// http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	// writer.WriteHeader(http.StatusOK)
	t := views.AdminRequestsPage()
	
	RequestList := models.FetchAdminRequests()
	fmt.Println(RequestList)
	t.Execute(writer, RequestList)
}