package controller

import (
	"LibGolang/pkg/views"
	"net/http"
)

func PageNotFound(writer http.ResponseWriter, request *http.Request) {
	t := views.PageNotFound()
	t.Execute(writer, "")
}