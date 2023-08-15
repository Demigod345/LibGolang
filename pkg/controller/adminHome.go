package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/views"
	"LibGolang/pkg/middleware"
	"net/http"
	"log"
)

func AdminHomePage(writer http.ResponseWriter, request *http.Request) {
	files := views.ViewFileNames()
	template := views.ViewPage(files.AdminHome)
	
	booksList, err := models.FetchBooks()
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	booksList.Message, err = GetFlash(writer, request)
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	booksList.Username = request.Context().Value(middleware.UsernameContextKey).(string)
	template.Execute(writer, booksList)
}