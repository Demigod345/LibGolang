package api

import (
	"LibGolang/pkg/controller"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	fmt.Println("Server Started at post http://localhost:8000/")
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("/templates"))
	http.Handle("/", fs)
	r.HandleFunc("/", controller.Signup).Methods("GET")



	r.HandleFunc("/returnBook", controller.ReturnBook).Methods("POST")
	r.HandleFunc("/requestBook", controller.RequestBook).Methods("POST")
	r.HandleFunc("/", controller.AddUser).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/add", controller.AddBook).Methods("POST")
	r.HandleFunc("/remove", controller.RemoveBook).Methods("POST")
	http.ListenAndServe(":8000", r)
}