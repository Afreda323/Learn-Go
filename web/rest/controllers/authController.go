package controllers

import (
	"Learn-Go/web/rest/models"
	u "Learn-Go/web/rest/utils"
	"encoding/json"
	"net/http"
)

// CreateAccount - controller for /api/user/new
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(account)

	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
	}

	resp := account.Create()
	u.Respond(w, resp)
}

// Authenticate - controller for /api/user/login
func Authenticate(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)

	if err != nil {
		u.Respond(w, u.Message(false, "Login failed"))
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}
