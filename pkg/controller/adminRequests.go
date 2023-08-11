package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/views"
	"fmt"
	"net/http"
	"strings"
)

func AdminRequestsPage(writer http.ResponseWriter, request *http.Request) {
	// http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	// writer.WriteHeader(http.StatusOK)

	state := strings.Split(request.URL.Path, "/")[3]

	if state == "requested" || state == "issued" || state == "checkedIn" || state == "AdminRequest" {
		t := views.AdminRequestsPage(state)

		requestList := models.FetchAllRequestsList(state)
		fmt.Println(requestList)
		t.Execute(writer, requestList)
	} else {
		http.Redirect(writer, request, "/admin/adminHome", http.StatusSeeOther)
		fmt.Println("Bad Request")
	}
}