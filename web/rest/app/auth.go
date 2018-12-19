package app

import (
	"Learn-Go/web/rest/models"
	u "Learn-Go/web/rest/utils"
	"context"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		unauthPages := []string{"/api/user/new", "/api/user/login"}
		requestPath := r.URL.Path

		// check to see if hitting an unauthed route
		// if so, move along
		for _, path := range unauthPages {
			if path == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization") // get token header

		// if token is empty, send error message
		if tokenHeader == "" {
			response = u.Message(false, "Auth token not provided")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		splitToken := strings.Split(tokenHeader, " ") // split token into ["Bearer", "token"]

		// invalid header shape
		if len(splitToken) != 2 {
			response = u.Message(false, "Invalid/Malformed token provided")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		tokenVal := splitToken[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenVal, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			response = u.Message(false, "Invalid/Malformed token provided")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		if !token.Valid {
			response = u.Message(false, "Invalid/Malformed token provided")
			w.WriteHeader(http.StatusForbidden)
			u.Respond(w, response)
			return
		}

		// Everything went well
		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
