package controller

import (
	"LibGolang/pkg/views"
	"net/http"
)

func InternalServerError(writer http.ResponseWriter, request *http.Request) {
	t := views.InternalServerError()
	t.Execute(writer, "")
}