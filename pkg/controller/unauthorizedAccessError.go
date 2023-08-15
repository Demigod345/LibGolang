package controller

import (
	"LibGolang/pkg/views"
	"net/http"
)

func UnauthorizedAccessError(writer http.ResponseWriter, request *http.Request) {
	files := views.ViewFileNames()
	t := views.ViewPage(files.UnauthorizedAccess)
	t.Execute(writer, "")
}