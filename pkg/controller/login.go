package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/views"
	"fmt"
	"net/http"
)

func Login(writer http.ResponseWriter, request *http.Request) {

	username := request.FormValue("username");
	password := request.FormValue("password");


	userExists, user := models.UserExists(username)

	if !userExists {
		fmt.Println("User doesn't Exist.")
		return
	} else {
		loginSuccesssful := models.DoPasswordsMatch(user.Hash, password)

		if loginSuccesssful {
			cookie := getJWTCookie(writer, request, user.UserId, user.UserName, user.IsAdmin)
			http.SetCookie(writer, &cookie)
			http.Redirect(writer, request, "/userHome", http.StatusSeeOther)

			fmt.Println("Logging in user ", username)
		} else {
			fmt.Println("Unsucessful Login, password doesn't match")
		}
	}
}

func LoginPage(writer http.ResponseWriter, request *http.Request) {

	t := views.LoginPage()
	writer.WriteHeader(http.StatusOK)
	// booksList := models.FetchBooks()
	t.Execute(writer, "")
}

