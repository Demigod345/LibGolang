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

	r.Use(controller.TokenMiddleware)

	adminRouter := r.PathPrefix("/admin").Subrouter()
	userRouter := r.PathPrefix("/user").Subrouter()
	adminRouter.Use(controller.RoleMiddleware (true))
	userRouter.Use(controller.RoleMiddleware (false))


	// http.HandleFunc("/user", controller.VerifyJWT(controller.UserHomePage))
	r.HandleFunc("/", controller.SignupPage).Methods("GET")
	r.HandleFunc("/login", controller.LoginPage).Methods("GET")
	userRouter.HandleFunc("/userHome", controller.UserHomePage).Methods("GET")
	adminRouter.HandleFunc("/adminHome", controller.AdminHomePage).Methods("GET")
	adminRouter.HandleFunc("/admin/adminRequests", controller.AdminRequestsPage).Methods("GET")



	r.HandleFunc("/signup", controller.AddUser).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("POST")
	adminRouter.HandleFunc("/addBook", controller.AddBook).Methods("POST")
	adminRouter.HandleFunc("/removeBook", controller.RemoveBook).Methods("POST")
	userRouter.HandleFunc("/user/requestBook", controller.RequestBook).Methods("POST")
	userRouter.HandleFunc("/user/returnBook", controller.ReturnBook).Methods("POST")
	adminRouter.HandleFunc("/approveRequest", controller.ApproveRequest).Methods("POST")
	adminRouter.HandleFunc("/rejectRequest", controller.RejectRequest).Methods("POST")
	adminRouter.HandleFunc("/rejectReturn", controller.RejectReturn).Methods("POST")
	adminRouter.HandleFunc("/approveReturn", controller.ApproveReturn).Methods("POST")

	http.ListenAndServe(":8000", r)
}
