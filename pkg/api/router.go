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

	router.Use(controller.TokenMiddleware)

	adminRouter := router.PathPrefix("/admin").Subrouter()
	userRouter := router.PathPrefix("/user").Subrouter()
	adminRouter.Use(controller.RoleMiddleware (true))
	userRouter.Use(controller.RoleMiddleware (false))


	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		http.Redirect(w,r, "/login", http.StatusSeeOther)
	})
	router.HandleFunc("/signup", controller.SignupPage).Methods("GET")
	router.HandleFunc("/login", controller.LoginPage).Methods("GET")
	router.HandleFunc("/500", controller.InternalServerError).Methods("GET")
	router.HandleFunc("/403", controller.UnauthorizedAccessError).Methods("GET")
	
	router.HandleFunc("/signup", controller.AddUser).Methods("POST")
	router.HandleFunc("/loginUser", controller.Login).Methods("POST")
	router.HandleFunc("/loginAdmin", controller.Login).Methods("POST")
	router.HandleFunc("/logout", controller.Logout).Methods("POST")


	userRouter.HandleFunc("/userHome", controller.UserHomePage).Methods("GET")
	userRouter.HandleFunc("/userRequests/{state}", controller.UserRequestsPage).Methods("GET")

	userRouter.HandleFunc("/requestBook", controller.RequestBook).Methods("POST")
	userRouter.HandleFunc("/returnBook", controller.ReturnBook).Methods("POST")
	userRouter.HandleFunc("/requestAdmin", controller.RequestAdmin).Methods("POST")


	adminRouter.HandleFunc("/adminHome", controller.AdminHomePage).Methods("GET")
	adminRouter.HandleFunc("/adminRequests/{state}", controller.AdminRequestsPage).Methods("GET")

	adminRouter.HandleFunc("/addBook", controller.AddBook).Methods("POST")
	adminRouter.HandleFunc("/removeBook", controller.RemoveBook).Methods("POST")
	adminRouter.HandleFunc("/deleteBook", controller.DeleteBook).Methods("POST")		
	adminRouter.HandleFunc("/approveRequest", controller.ApproveRequest).Methods("POST")
	adminRouter.HandleFunc("/rejectRequest", controller.RejectRequest).Methods("POST")
	adminRouter.HandleFunc("/rejectReturn", controller.RejectReturn).Methods("POST")
	adminRouter.HandleFunc("/approveReturn", controller.ApproveReturn).Methods("POST")
	adminRouter.HandleFunc("/approveAdmin", controller.ApproveAdmin).Methods("POST")
	adminRouter.HandleFunc("/rejectAdmin", controller.RejectAdmin).Methods("POST")

	router.NotFoundHandler = http.HandlerFunc(controller.PageNotFound)
	http.ListenAndServe(":8000", router)
}
