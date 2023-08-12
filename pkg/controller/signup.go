package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/types"
	"LibGolang/pkg/views"
	"log"
	"net/http"
)

func SignupPage(writer http.ResponseWriter, request *http.Request) {
	var message types.PageMessage
	var err error
	message.Message, err = GetFlash(writer, request)
	if err != nil {
		log.Fatal(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	t := views.SignupPage()
	t.Execute(writer, message)
}

func AddUser(writer http.ResponseWriter, request *http.Request) {

	username := request.FormValue("username")
	password := request.FormValue("password")
	passwordC := request.FormValue("passwordC")

	if password != passwordC {
		SetFlash(writer, request, "Passwords don't match.")
		http.Redirect(writer, request, "/signup", http.StatusSeeOther)
		return
	} else if password == "" {
		SetFlash(writer, request, "Passwords can't empty.")
		http.Redirect(writer, request, "/signup", http.StatusSeeOther)
		return
	} else {
		userExists, _, err := models.UserExists(username)

		if err != nil {
			log.Println(err)
			http.Redirect(writer, request, "/500", http.StatusSeeOther)
			return
		}

		if userExists {
			SetFlash(writer, request, "User Already Exists.")
			http.Redirect(writer, request, "/signup", http.StatusSeeOther)
			return
		} else {
			pass, err := models.HashPassword(password)
			if err != nil {
				log.Println(err)
				http.Redirect(writer, request, "/500", http.StatusSeeOther)
				return
			} else {
				err := models.AddUser(username, pass)
				if err != nil {
					log.Fatal(err)
					http.Redirect(writer, request, "/500", http.StatusSeeOther)
					return
				}
			}
		}

	}

	SetFlash(writer, request, "User Registered Successfully.")
	http.Redirect(writer, request, "/login", http.StatusSeeOther)
}
