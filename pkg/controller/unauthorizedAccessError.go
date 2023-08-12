package controller

import (
	"LibGolang/pkg/views"
	"net/http"
)

func UnauthorizedAccessError(writer http.ResponseWriter, request *http.Request) {
	t := views.UnauthorizedAccessError()
	t.Execute(writer, "")
}