package views

import (
	"fmt"
	"html/template"
)

func AdminRequestsPage(state string) *template.Template {

	file := "templates/adminRequests_" + state + ".html"
	fmt.Println(file)
	temp := template.Must(template.ParseFiles(file))
	return temp
}
