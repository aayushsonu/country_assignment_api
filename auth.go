package main

// swagger:meta
import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("myakhere_xyzabc1234")

// AuthHandler godoc
// @Summary Authenticate user and generate access token
// @Description Authenticate user with credentials and generate access token
// @Tags authentication
// @Accept  json
// @Produce  json
// @Param credentials body Credentials true "User credentials"
// @Success 200 {object} map[string]string "JWT token"
// @Failure 400 {object} ErrorResponse "Invalid request payload"
// @Failure 401 {object} ErrorResponse "Invalid credentials"
// @Failure 500 {object} ErrorResponse "Error generating token"
// @Router /auth [post]
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Replace with your authentication logic
	if creds.Username == "snifyak" && creds.Password == "123@snifyak@123" {
		tokenString, err := generateToken()
		if err != nil {
			writeJSONError(w, http.StatusInternalServerError, "Error generating token")
			return
		}

		response := map[string]string{"token": tokenString}
		writeJSONResponse(w, http.StatusOK, response)
	} else {
		writeJSONError(w, http.StatusUnauthorized, "Invalid credentials")
	}
}

// writeJSONResponse writes a JSON response with the specified status code and data.
func writeJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			writeJSONError(w, http.StatusUnauthorized, "Missing auth token")
			return
		}

		authToken := strings.TrimPrefix(authHeader, "Bearer ")

		_, err := validateToken(authToken)
		if err != nil {
			writeJSONError(w, http.StatusUnauthorized, "Invalid auth token")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func generateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		// Return a JSON response for the error
		return "", errors.New("error generating token")
	}

	return tokenString, nil
}

func validateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		// If the token is expired, return a specific JSON response
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token has expired")
			}
		}

		// For other errors, return a general JSON response
		return nil, errors.New("invalid token")
	}

	return token, nil
}
