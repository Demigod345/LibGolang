package views

import (
	"html/template"
)

func UserRequestsPage(state string) *template.Template {

	file := "templates/userRequests_" + state + ".html"
	temp := template.Must(template.ParseFiles(file))
	return temp
}
