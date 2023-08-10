package views

import (
	"html/template"
)

func AdminRequestsPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/adminRequests.html"))
	return temp
}
