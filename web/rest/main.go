package main

import (
	"Learn-Go/web/rest/app"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuth)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Println("Error on Listen and Serve", err)
	}
}
