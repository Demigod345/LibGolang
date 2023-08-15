package controller

import (
	"LibGolang/pkg/views"
	"net/http"
)

func InternalServerError(writer http.ResponseWriter, request *http.Request) {
	files := views.ViewFileNames()
	t := views.ViewPage(files.InternalServerError)
	t.Execute(writer, "")
}