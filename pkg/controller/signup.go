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

	files := views.ViewFileNames()
	t := views.ViewPage(files.Signup)
	t.Execute(writer, message)
}

func AddUser(writer http.ResponseWriter, request *http.Request) {
	var RequestUser types.RequestUser
	RequestUser.UserName = request.FormValue("username")
	RequestUser.Password = request.FormValue("password")
	RequestUser.ConfirmPassword = request.FormValue("passwordC")

	if RequestUser.Password != RequestUser.ConfirmPassword {
		SetFlash(writer, request, "Passwords don't match.")
		http.Redirect(writer, request, "/signup", http.StatusSeeOther)
		return
	} else if len([]byte(RequestUser.Password)) == 0 {
		SetFlash(writer, request, "Passwords can't empty.")
		http.Redirect(writer, request, "/signup", http.StatusSeeOther)
		return
	} else {
		db, err := models.Connection()
		if err != nil {
			log.Println(err)
			http.Redirect(writer, request, "/500", http.StatusSeeOther)
			return
		}

		userExists, _, err := models.UserExists(db, RequestUser.UserName)
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
			pass, err := models.HashPassword(RequestUser.Password)
			if err != nil {
				log.Println(err)
				http.Redirect(writer, request, "/500", http.StatusSeeOther)
				return
			} else {
				err := models.AddUser(RequestUser.UserName, pass)
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
