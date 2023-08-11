package controller

import (
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	expirationTime := time.Now().Add(5 * time.Minute)

	
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: expirationTime,
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)


}