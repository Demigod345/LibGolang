package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/views"
	"net/http"
)

func UserHomePage(writer http.ResponseWriter, request *http.Request) {
	// http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	// writer.WriteHeader(http.StatusOK)
	t := views.UserHomePage()
	
	booksList := models.FetchBooks()
	t.Execute(writer, booksList)
}