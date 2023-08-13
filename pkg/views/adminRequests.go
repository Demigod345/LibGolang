package views

import (
	"html/template"
)

func AdminRequestsPage(state string) *template.Template {
	file := "templates/adminRequests_" + state + ".html"
	temp := template.Must(template.ParseFiles(file))
	return temp
}
