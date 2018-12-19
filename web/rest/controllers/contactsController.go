package controllers

import (
	"Learn-Go/web/rest/models"
	u "Learn-Go/web/rest/utils"
	"encoding/json"
	"net/http"
)

// CreateContact - POST /api/me/contacts
func CreateContact(w http.ResponseWriter, r *http.Request) {
	contact := &models.Contact{}
	err := json.NewDecoder(r.Body).Decode(contact)

	if err != nil {
		u.Respond(w, u.Message(false, "Error parsing request"))
		return
	}

	user := r.Context().Value("user").(uint)
	contact.UserID = user
	resp := contact.Create()

	u.Respond(w, resp)
}

// GetContactsFor - GET /api/me/contacts
func GetContactsFor(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	data := models.GetContacts(user)
	resp := u.Message(true, "success")
	resp["data"] = data

	u.Respond(w, resp)
}
