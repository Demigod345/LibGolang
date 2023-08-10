package controller

import (
	"LibGolang/pkg/views"
	// "encoding/json"
	"fmt"
	// "log"
	"net/http"
)

func YoPage(writer http.ResponseWriter, request *http.Request) {
	t := views.YoPage()
	writer.WriteHeader(http.StatusOK)
	// booksList := models.FetchBooks()
	t.Execute(writer, "")
}

func Yo(writer http.ResponseWriter, request *http.Request){

	// var x string;
	// err := json.NewDecoder(request.Body).Decode(&x)
	// if err != nil {
	// 	fmt.Println("There was an error decoding the request body into the struct", err)
	// }

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(x)
	fmt.Println(request.FormValue("firstname"))
	fmt.Println(request.FormValue("lastname"))

}