package views

import (
	"html/template"
)

func SignupPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/signup.html"))
	return temp
}
