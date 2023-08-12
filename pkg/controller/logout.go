package controller

import (
	"net/http"
	"time"
)

func Logout(writer http.ResponseWriter, request *http.Request) {
	expirationTime := time.Now().Add(5 * time.Minute)
	
	http.SetCookie(writer, &http.Cookie{
		Name:    "token",
		Expires: expirationTime,
	})
	SetFlash(writer, request, "Logged Out Successfullly.")
	http.Redirect(writer, request, "/login", http.StatusSeeOther)
}