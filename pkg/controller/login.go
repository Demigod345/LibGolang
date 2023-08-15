package controller

import (
	"LibGolang/pkg/models"
	"LibGolang/pkg/types"
	"LibGolang/pkg/views"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"strings"
	"time"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	role := strings.TrimPrefix(request.URL.String(), "/login")

	isAdminReq := false
	if role == "Admin" {
		isAdminReq = true
	}

	username := request.FormValue("username")
	password := request.FormValue("password")

	db, err := models.Connection()
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	userExists, user, err := models.UserExists(db, username)
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	if !userExists {
		SetFlash(writer, request, "User doesn't Exist.")
		http.Redirect(writer, request, "/login", http.StatusSeeOther)
		return
	} else {
		if isAdminReq != user.IsAdmin {
			if isAdminReq {
				SetFlash(writer, request, "Try logging in as User.")
				http.Redirect(writer, request, "/login", http.StatusSeeOther)
				return
			} else {
				SetFlash(writer, request, "Try logging in as Admin.")
				http.Redirect(writer, request, "/login", http.StatusSeeOther)
				return
			}
		}

		loginSuccesssful := models.DoPasswordsMatch(user.Hash, password)
		if loginSuccesssful {
			expirationTime := time.Now().Add(30 * time.Minute)
			claims := &types.Claims{
				Username: username,
				UserId:   user.UserId,
				IsAdmin:  user.IsAdmin,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(expirationTime),
				},
			}

			key, err := models.GetJWTSecretKey()
			jwtKey := []byte(key)
			if err != nil {
				log.Println(err)
				http.Redirect(writer, request, "/500", http.StatusSeeOther)
				return
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(jwtKey)
			if err != nil {
				http.Redirect(writer, request, "/500", http.StatusSeeOther)
				return
			}

			http.SetCookie(writer, &http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: expirationTime,
			})

			SetFlash(writer, request, "Login Successful!!")

			isAdmin := user.IsAdmin
			if isAdmin {
				http.Redirect(writer, request, "/admin/adminHome", http.StatusSeeOther)
			} else {
				http.Redirect(writer, request, "/user/userHome", http.StatusSeeOther)
			}
		} else {
			SetFlash(writer, request, "Unsucessful Login, Wrong Password.")
			http.Redirect(writer, request, "/login", http.StatusSeeOther)
			return
		}
	}
}

func LoginPage(writer http.ResponseWriter, request *http.Request) {
	var message types.PageMessage
	var err error
	message.Message, err = GetFlash(writer, request)
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}
	files := views.ViewFileNames()
	template := views.ViewPage(files.Login)
	template.Execute(writer, message)
}
