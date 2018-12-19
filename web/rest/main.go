package main

import (
	"Learn-Go/web/rest/app"
	"Learn-Go/web/rest/controllers"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuth)
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Println("Error on Listen and Serve", err)
	}
}
