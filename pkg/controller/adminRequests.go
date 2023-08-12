package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/views"
	"log"
	"net/http"
	"strings"
)

func AdminRequestsPage(writer http.ResponseWriter, request *http.Request) {

	state := strings.Split(request.URL.Path, "/")[3]

	if state == "requested" || state == "issued" || state == "checkedIn" || state == "AdminRequest" {
		template := views.AdminRequestsPage(state)

		requestList, err := models.FetchAllRequestsList(state)
		if err != nil {
			log.Println(err)
			http.Redirect(writer, request, "/500", http.StatusSeeOther)
			return
		}

		requestList.Message, err = GetFlash(writer, request)
		if err != nil {
			log.Println(err)
			http.Redirect(writer, request, "/500", http.StatusSeeOther)
			return
		}

		requestList.Username = request.Context().Value(usernameContextKey).(string)
		template.Execute(writer, requestList)
	} else {
		http.Redirect(writer, request, "/admin/adminHome", http.StatusSeeOther)
	}
}
