package controller

import (
	"LibGolang/pkg/middleware"
	"LibGolang/pkg/models"
	"LibGolang/pkg/views"
	"log"
	"net/http"
)

func UserHomePage(writer http.ResponseWriter, request *http.Request) {
	template := views.UserHomePage()
	
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