package app

import (
	"bookmarks-api/models"
	"bookmarks-api/utils"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// List endpoints for which we don't need authentication
		notAuth := []string{"/api/user/new", "/api/user/login"}

		// Current request path
		requestPath := r.URL.Path

		// Servicing request - if it doesn't require authentication
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})

		// Get token
		tokenHeader := r.Header.Get("Authorization")

		// If token absent / return 403 http - Unauthorized
		if tokenHeader == "" {
			response = utils.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		// Check if token is in format `Bearer {token-body}`
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response = utils.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		// Get seconds token part
		tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		// Wrong token return 403-http
		if err != nil {
			response = utils.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
		}

		// Invalid token
		if !token.Valid {
			response = utils.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		// Success, continue request execution
		fmt.Sprintf("User %", tk.UserId)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)

		// Pass manager to next handler
		next.ServeHTTP(w, r)
	})
}
