package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/leehaowei/go-dbapp/src/controllers"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go World")
}

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.ListBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
}
