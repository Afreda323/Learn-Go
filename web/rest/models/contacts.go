package models

import (
	u "Learn-Go/web/rest/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

// Contact struct
type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserID uint   `json:"user_id"`
}

// Validate that all required fields are preset
func (contact *Contact) Validate() (map[string]interface{}, bool) {
	if contact.Name == "" {
		return u.Message(false, "Name required"), false
	}
	if contact.Phone == "" {
		return u.Message(false, "Phone required"), false
	}
	if contact.UserID <= 0 {
		return u.Message(false, "User not recognized"), false
	}

	return u.Message(true, "success"), true
}

// Create and save contact to db
func (contact *Contact) Create() map[string]interface{} {
	if resp, ok := contact.Validate(); !ok {
		return resp
	}

	GetDB().Create(contact)

	resp := u.Message(true, "success")
	resp["contact"] = contact
	return resp
}

// GetContact - get contact from db by ID
func GetContact(id uint) *Contact {
	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user uint) []*Contact {
	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contacts
}
