package api

import (
	"LibGolang/pkg/controller"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	fmt.Println("Server Started at post http://localhost:8000/")
	router := mux.NewRouter()
	// r.Use(controller.ValidateJWT)

	// Serving the static files
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/static/").Handler(s)

	router.Use(controller.TokenMiddleware)

	adminRouter := router.PathPrefix("/admin").Subrouter()
	userRouter := router.PathPrefix("/user").Subrouter()
	adminRouter.Use(controller.RoleMiddleware (true))
	userRouter.Use(controller.RoleMiddleware (false))


	// http.HandleFunc("/user", controller.VerifyJWT(controller.UserHomePage))
	router.HandleFunc("/", controller.SignupPage).Methods("GET")
	router.HandleFunc("/login", controller.LoginPage).Methods("GET")
	router.HandleFunc("/signup", controller.AddUser).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/logout", controller.Logout).Methods("POST")


	userRouter.HandleFunc("/userHome", controller.UserHomePage).Methods("GET")
	userRouter.HandleFunc("/userRequests", controller.UserRequestsPage).Methods("GET")

	userRouter.HandleFunc("/requestBook", controller.RequestBook).Methods("POST")
	userRouter.HandleFunc("/returnBook", controller.ReturnBook).Methods("POST")


	adminRouter.HandleFunc("/adminHome", controller.AdminHomePage).Methods("GET")
	adminRouter.HandleFunc("/adminRequests", controller.AdminRequestsPage).Methods("GET")

	adminRouter.HandleFunc("/addBook", controller.AddBook).Methods("POST")
	adminRouter.HandleFunc("/removeBook", controller.RemoveBook).Methods("POST")	
	adminRouter.HandleFunc("/approveRequest", controller.ApproveRequest).Methods("POST")
	adminRouter.HandleFunc("/rejectRequest", controller.RejectRequest).Methods("POST")
	adminRouter.HandleFunc("/rejectReturn", controller.RejectReturn).Methods("POST")
	adminRouter.HandleFunc("/approveReturn", controller.ApproveReturn).Methods("POST")

	http.ListenAndServe(":8000", router)
}
