package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/types"
	"LibGolang/pkg/views"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

//jwt

func Login(writer http.ResponseWriter, request *http.Request) {

	username := request.FormValue("username")
	password := request.FormValue("password")

	userExists, user := models.UserExists(username)

	if !userExists {
		fmt.Println("User doesn't Exist.")
		return
	} else {
		loginSuccesssful := models.DoPasswordsMatch(user.Hash, password)

		if loginSuccesssful {
			// cookie := getJWTCookie(writer, request, user.UserId, user.UserName, user.IsAdmin)
			// http.SetCookie(writer, &cookie)

			// implement jwt

			expirationTime := time.Now().Add(15 * time.Minute)
			claims := &types.Claims{
				Username: username,
				UserId:   user.UserId,
				IsAdmin:  user.IsAdmin,
				RegisteredClaims: jwt.RegisteredClaims{
					// In JWT, the expiry time is expressed as unix milliseconds
					ExpiresAt: jwt.NewNumericDate(expirationTime),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			// Create the JWT string
			tokenString, err := token.SignedString(jwtKey)
			if err != nil {
				// If there is an error in creating the JWT return an internal server error
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			http.SetCookie(writer, &http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: expirationTime,
			})

			isAdmin := user.IsAdmin

			if isAdmin {
				http.Redirect(writer, request, "/admin/adminHome", http.StatusSeeOther)
				fmt.Println("Logging in admin ", username)
			}else {
				http.Redirect(writer, request, "/user/userHome", http.StatusSeeOther)
				fmt.Println("Logging in user ", username)
			}			
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
