package controller

import (
	// "encoding/json"
	"fmt"
	"log"

	"LibGolang/pkg/models"
	"LibGolang/pkg/views"
	"net/http"
)

func SignupPage(writer http.ResponseWriter, request *http.Request) {
	t := views.SignupPage()
	writer.WriteHeader(http.StatusOK)
	// booksList := models.FetchBooks()
	t.Execute(writer, "")
}

func AddUser(writer http.ResponseWriter, request *http.Request) {

	username := request.FormValue("username");
	password := request.FormValue("password");
	passwordC := request.FormValue("passwordC");


	if password != passwordC {
		fmt.Println("Passwords don't match.")
		writer.WriteHeader(400)
		return
	} else if password == "" {
		fmt.Println("Passwords empty.")
		writer.WriteHeader(400)
		return
	} else {
		userExists, user := models.UserExists(username)
		fmt.Println(userExists)
		fmt.Println(user)
		if userExists {
			fmt.Printf("%s Already Exists.", user.UserName)
			writer.WriteHeader(400)
			return
		} else {
			pass, err := models.HashPassword(password)
			if err != nil {
				log.Fatal(err)
			} else {
				http.Redirect(writer, request, "/login", http.StatusSeeOther)

				fmt.Println("Password Hash: ", pass)
				models.AddUser(username, pass)
			}
		}

	}

	fmt.Printf("Signing Up %s", username)
}
