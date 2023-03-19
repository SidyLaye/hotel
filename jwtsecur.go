package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Secret key used to sign JWT tokens
var secretKey = []byte("mySecretKey")

// Claims struct used to create a JWT token
type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

// Function that generates a JWT token for a given user ID
func generateToken(userId int64) (string, error) {
	// Create the payload (content) of the token
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // expiration in 24 hours
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Encode the token as a string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Function that verifies a JWT token and returns the user ID if it is valid
func verifyToken(tokenString string) (int64, error) {
	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the signature method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return 0, err
	}

	// Check the validity of the token and get the payload
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := int64(claims["user_id"].(float64))
		return userId, nil
	} else {
		return 0, fmt.Errorf("Invalid token")
	}
}
