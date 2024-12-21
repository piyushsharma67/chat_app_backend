package utils

import (
	"errors"
	"fmt"
	"mainserver/schema"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func ValidateToken(tokenString string) (schema.User, error) {
	secretKey := []byte(os.Getenv("secret"))

	// Remove "Bearer " prefix if it exists
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[len("Bearer "):]
	}

	// Parse the token
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return schema.User{}, err
	}

	// Check if the token is valid and extract claims
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		// Optionally check for expiration or other conditions
		expirationTime, ok := claims["exp"].(float64)
		if !ok || time.Now().Unix() > int64(expirationTime) {
			return schema.User{}, errors.New("token has expired")
		}

		// Type assertion for userID and email (or other claims)
		userID, ok := claims["userID"].(float64)
		if !ok {
			return schema.User{}, errors.New("userID claim is missing or invalid")
		}

		email, ok := claims["email"].(string)
		if !ok {
			return schema.User{}, errors.New("email claim is missing or invalid")
		}

		// Return the user
		user := schema.User{
			ID:    int64(userID), // Convert from float64 to int
			Email: email,
		}

		return user, nil
	}
	fmt.Println("yay4")
	return schema.User{}, errors.New("invalid token")
}
