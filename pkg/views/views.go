package views

import (
	"LibGolang/pkg/types"
	"html/template"
)

func ViewFileNames() types.FileName {
	return types.FileName{
		AdminHome:           "adminHome",
		UserHome:            "userHome",
		Login:               "login",
		PageNotFound:        "pageNotFound",
		UnauthorizedAccess:  "unauthorizedAccessError",
		InternalServerError: "internalServerError",
		Signup:              "signup",
	}
}

func ViewPage(page string) *template.Template {
	file := "templates/" + page + ".html"
	temp := template.Must(template.ParseFiles(file))
	return temp
}