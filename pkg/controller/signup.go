package controller

import (
	"encoding/json"
	"fmt"
	"log"

	"LibGolang/pkg/models"
	"LibGolang/pkg/types"
	"LibGolang/pkg/views"
	"net/http"
)

func Signup(writer http.ResponseWriter, request *http.Request) {
	// http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	// writer.WriteHeader(http.StatusOK)
	t := views.SignupPage()
	writer.WriteHeader(http.StatusOK)
	// booksList := models.FetchBooks()
	t.Execute(writer, "")
}

func AddUser(writer http.ResponseWriter, request *http.Request) {
	var signupRequest types.SignupRequest

	err := json.NewDecoder(request.Body).Decode(&signupRequest)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
	}

	if signupRequest.Password != signupRequest.PasswordC {
		fmt.Println("Passwords don't match.")
		return
	} else if signupRequest.Password == "" {
		fmt.Println("Passwords empty.")
		return
	} else {
		userExists, user := models.UserExists(signupRequest.Username)
		fmt.Println(userExists);
		fmt.Println(user)
		if userExists {
			fmt.Printf("%s Already Exists.", user.UserName)
			return
		} else {
			pass, err := models.HashPassword(signupRequest.Password)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("Password Hash: ", pass)
				models.AddUser(signupRequest.Username, pass)
			}
		}

	}

	fmt.Printf("Signing Up %s", signupRequest.Username)
}
