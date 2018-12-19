package main

import (
	"Learn-Go/web/rest/app"
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
}
