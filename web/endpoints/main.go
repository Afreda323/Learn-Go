// Dead simple REST api
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type person struct {
	ID        string   `json:"id"`
	FirstName string   `json:"firstName,omitempty"`
	LastName  string   `json:"lastName,omitempty"`
	Address   *address `json:"address,omitempty"`
}

type address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []person

func setHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;") // send JSON to user
		h.ServeHTTP(w, r)
	})
}

// getPeople: GET /people
// Get a list of all the people
func getPeople(w http.ResponseWriter, r *http.Request) { // 200 response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(people) // send list of people to user
}

// getPerson: GET /people/{id}
// get a person by ID
func getPerson(w http.ResponseWriter, r *http.Request) {
	found := false

	vars := mux.Vars(r) // decode url vars
	for _, person := range people {
		// loop through people
		// when desired user is found. return user
		if person.ID == vars["id"] {
			found = true
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(person)
			return
		}
	}
	if !found {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found")
	}
}

// addPerson: POST /people
// Adds a new person to slice of people
func addPerson(w http.ResponseWriter, r *http.Request) {
	var p person                           // new instance of person struct
	_ = json.NewDecoder(r.Body).Decode(&p) // pump req body into person struct
	p.ID = uuid.New().String()             // generate a uuid
	people = append(people, p)             // add person to list

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(people) // send updated list of people to user

}
func deletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // decode url vars
	for i, person := range people {
		if vars["id"] == person.ID {
			people = append(people[:i], people[i+1:]...)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(people) // send updated list of people to user
}

func main() {
	router := mux.NewRouter()
	router.Use(setHeaders)
	router.HandleFunc("/people", getPeople).Methods("GET")
	router.HandleFunc("/people/{id}", getPerson).Methods("GET")
	router.HandleFunc("/people", addPerson).Methods("POST")
	router.HandleFunc("/people/{id}", deletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
