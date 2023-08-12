package views

import (
	"html/template"
)

func LoginPage() *template.Template {
	file := "templates/login.html"
	temp := template.Must(template.ParseFiles(file))
	return temp
}
