package controller

import (
	"LibGolang/pkg/middleware"
	"LibGolang/pkg/models"
	"LibGolang/pkg/views"
	"log"
	"net/http"
	"strings"
)

func UserRequestsPage(writer http.ResponseWriter, request *http.Request) {
	state := strings.Split(request.URL.Path, "/")[3]
	if state == "requested" || state == "issued" || state == "checkedIn" || state == "AdminRequest" {
		template := views.UserRequestsPage(state)
		userId := request.Context().Value(middleware.UserIdContextKey).(int)

		requestList, err := models.FetchUserRequestsList(state, userId)
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
		requestList.Username = request.Context().Value(middleware.UsernameContextKey).(string)
		template.Execute(writer, requestList)
	} else {
		http.Redirect(writer, request, "/400", http.StatusSeeOther)
	}
}
