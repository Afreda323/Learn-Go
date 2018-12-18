package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("Username:", r.Form["username"])
		fmt.Println("Password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server", err)
	}
}
