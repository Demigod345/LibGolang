package views

import (
	"html/template"
)

func PageNotFound() *template.Template {
	temp := template.Must(template.ParseFiles("templates/pageNotFound.html"))
	return temp
}
