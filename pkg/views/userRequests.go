package views

import (
	"html/template"
)

func UserRequestsPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/userRequests.html"))
	return temp
}
