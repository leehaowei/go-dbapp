package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/leehaowei/go-dbapp/src/routes"
	"log"
	"net/http"
)

func main() {
	fmt.Println("First line of the program")

	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	fmt.Printf("The server is running on: 8080\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
