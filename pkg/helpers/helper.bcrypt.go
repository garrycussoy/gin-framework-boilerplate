package helpers

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Function to generate hash from password
func GenerateHash(password string) (string, error) {
	// Ensure it's not empty
	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	// Generate hash
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// Function to validate hash
func ValidateHash(secret, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
	return err == nil
}
