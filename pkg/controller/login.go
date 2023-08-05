package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	var loginRequest types.SignupRequest
	err := json.NewDecoder(request.Body).Decode(&loginRequest)
	if err != nil {
		fmt.Println("There was an error decoding the request body into the struct", err)
	}

	if err != nil {
		log.Fatal(err)
	}

	userExists, user := models.UserExists(loginRequest.Username)

	if !userExists {
		fmt.Println("User doesn't Exist.")
		return
	} else {
		loginSuccesssful := models.DoPasswordsMatch(user.Hash, loginRequest.Password)

		if loginSuccesssful {
			fmt.Println("Logging in user ", loginRequest.Username)
		} else {
			fmt.Println("Unsucessful Login, password doesn't match")
		}
	}
}
