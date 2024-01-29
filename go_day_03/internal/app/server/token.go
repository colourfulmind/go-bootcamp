package server

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

var secretKey = []byte("21school")

// TokenJSON is a struct for tokens
type TokenJSON struct {
	Token string `json:"token"`
}

// GetTokenHandler is a handler for `/api/get_token`
func (server *APIServer) GetTokenHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := GenerateToken()
		if err != nil {
			RequestError(w, err)
			return
		}

		err = GetToken(w, tokenString)
		if err != nil {
			RequestError(w, err)
			return
		}
	}
}

// GenerateToken generates a token
func GenerateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "JWT"
	return token.SignedString(secretKey)
}

// GetToken sends the token to the user as JSON
func GetToken(w http.ResponseWriter, tokenString string) error {
	var tokenJson TokenJSON
	tokenJson.Token = tokenString
	resp, err := json.MarshalIndent(tokenJson, "", "\t")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(resp)
	return err
}

// AuthorizeToken authorizes the token
func AuthorizeToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
