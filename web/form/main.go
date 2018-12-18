package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

func printValidInputs(form url.Values) {
	fmt.Println("Valid Fruit(s):", form["fruit"])
	fmt.Println("Valid Username:", form["username"])
	fmt.Println("Valid Sport:", form["sport"])
	fmt.Println("Valid Gender:", form["gender"])
}

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

func inputs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("inputs.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()

		isValidUser := isValidString(r.Form.Get("username"))
		isValidFruit := isFruit(r.Form.Get("fruit"))
		isValidSport := isValidString(r.Form.Get("sport"))
		isValidGender := isValidString(r.Form.Get("gender"))

		if isValidFruit && isValidGender && isValidSport && isValidUser {
			printValidInputs(r.Form)
		} else {
			fmt.Println("Something is invalid")
		}
	}
}

func main() {
	http.HandleFunc("/", login)
	http.HandleFunc("/inputs", inputs)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server", err)
	}
}
