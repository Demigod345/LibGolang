package controller

import (
	"LibGolang/pkg/models"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

func SetFlash(writer http.ResponseWriter, request *http.Request, message string) {
	key, err := models.GetJWTSecretKey()
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}
	store := sessions.NewCookieStore([]byte(key))

	session, err := store.Get(request, "flash-session")
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	session.AddFlash(message, "message")
	session.Save(request, writer)
}

func GetFlash(writer http.ResponseWriter, request *http.Request) (interface{}, error) {
	key, err := models.GetJWTSecretKey()
	if err != nil {
		return "", err
	}
	store := sessions.NewCookieStore([]byte(key))

	session, err := store.Get(request, "flash-session")
	if err != nil {
		return "", err
	}

	fm := session.Flashes("message")
	if fm == nil {
		return "", nil
	}

	session.Save(request, writer)
	message := fm[0]
	return message, err
}
