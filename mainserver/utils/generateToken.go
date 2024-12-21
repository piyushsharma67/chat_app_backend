package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var TokenSecret = []byte(os.Getenv(("secret")))

func GenerateToken(userID int, email string) (string, error) {

	claims := jwt.MapClaims{
		"userID": userID,
		"email":  email,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expiry: 24 hours
	}

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the secret key
	signedToken, err := token.SignedString(TokenSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
