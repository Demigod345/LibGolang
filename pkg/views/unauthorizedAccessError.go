package views

import (
	"html/template"
)

func UnauthorizedAccessError() *template.Template {
	
	temp := template.Must(template.ParseFiles("templates/unauthorizedAccessError.html"))
	return temp
}
