package controller

import (
	"LibGolang/pkg/views"
	"net/http"
)

func PageNotFound(writer http.ResponseWriter, request *http.Request) {
	files := views.ViewFileNames()
	t := views.ViewPage(files.PageNotFound)
	t.Execute(writer, "")
}