package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/views"
	"net/http"
	"log"
)

func AdminHomePage(writer http.ResponseWriter, request *http.Request) {
	// http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	// writer.WriteHeader(http.StatusOK)
	template := views.AdminHomePage()
	
	booksList, err := models.FetchBooks()
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}
	booksList.Message, _ = GetFlash(writer, request)
	booksList.Username = request.Context().Value(usernameContextKey).(string)
	template.Execute(writer, booksList)
}