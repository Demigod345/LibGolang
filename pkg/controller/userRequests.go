package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/views"
	"fmt"
	"net/http"
	"strings"
)

func UserRequestsPage(writer http.ResponseWriter, request *http.Request) {

	state := strings.Split(request.URL.Path, "/")[3]

	if state == "requested" || state == "issued" || state == "checkedIn" || state == "AdminRequest" {
		t := views.UserRequestsPage(state)
		userId := request.Context().Value(userIdContextKey).(int)

		requestList := models.FetchUserRequestsList(state, userId)
		fmt.Println(requestList)
		t.Execute(writer, requestList)
	} else {
		http.Redirect(writer, request, "/user/userHome", http.StatusSeeOther)
		fmt.Println("Bad Request")
	}

}
