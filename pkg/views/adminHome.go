package views

import (
	"html/template"
)

func AdminHomePage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/adminHome.html"))
	return temp
}
