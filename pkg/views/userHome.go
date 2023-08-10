package views

import (
	"html/template"
)

func UserHomePage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/userHome.html"))
	return temp
}
