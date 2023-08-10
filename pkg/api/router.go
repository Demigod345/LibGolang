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
	// r.Use(controller.ValidateJWT)

	// Serving the static files
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/static/").Handler(s)

	// http.HandleFunc("/user", controller.VerifyJWT(controller.UserHomePage))
	r.HandleFunc("/", controller.SignupPage).Methods("GET")
	r.HandleFunc("/login", controller.LoginPage).Methods("GET")
	r.HandleFunc("/userHome", controller.UserHomePage).Methods("GET")
	r.HandleFunc("/adminHome", controller.AdminHomePage).Methods("GET")
	r.HandleFunc("/adminRequests", controller.AdminRequestsPage).Methods("GET")


	// r.Use(controller.Authenticate)

	r.HandleFunc("/signup", controller.AddUser).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/addBook", controller.AddBook).Methods("POST")
	r.HandleFunc("/removeBook", controller.RemoveBook).Methods("POST")
	r.HandleFunc("/requestBook", controller.RequestBook).Methods("POST")
	r.HandleFunc("/returnBook", controller.ReturnBook).Methods("POST")
	r.HandleFunc("/approveRequest", controller.ApproveRequest).Methods("POST")
	r.HandleFunc("/rejectRequest", controller.RejectRequest).Methods("POST")
	r.HandleFunc("/rejectReturn", controller.RejectReturn).Methods("POST")
	r.HandleFunc("/approveReturn", controller.ApproveReturn).Methods("POST")

	http.ListenAndServe(":8000", r)
}
