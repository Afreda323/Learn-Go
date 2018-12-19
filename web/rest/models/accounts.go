package models

import (
	"fmt"
	"os"
	"strings"

	u "Learn-Go/web/rest/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Token struct
type Token struct {
	UserID uint
	jwt.StandardClaims
}

// Account struct
type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

// Validate user details
func (account *Account) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Invalid email provided"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Invalid password provided"), false
	}

	temp := &Account{}

	// Check to see if email exists already
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error

	// Make sure db is working properly
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("Err:", err)
		return u.Message(false, "Connection error. Please retry"), false
	}

	if temp.Email != "" {
		return u.Message(false, "Email is already taken"), false
	}

	// Email is not taken, password is valid
	return u.Message(false, "Requirement passed"), true
}

// Create a user account
func (account *Account) Create() map[string]interface{} {
	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPass)

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	tk := &Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))

	account.Token = tokenString
	account.Password = ""

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}

// Login a user by their email and password
func Login(email string, password string) map[string]interface{} {
	account := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email not found")
		}

		return u.Message(false, "DB error, please try again")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Invalid login creds. Please try again")
	}

	account.Password = ""

	tk := &Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

// GetUser by id
func GetUser(id uint) *Account {
	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", id).First(acc)

	if acc.Email == "" {
		return nil
	}

	acc.Password = ""
	return acc
}
