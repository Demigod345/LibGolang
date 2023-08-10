package views

import (
	"html/template"
)

func YoPage() *template.Template {
	
	temp := template.Must(template.ParseFiles("templates/abc.html"))
	return temp
}
