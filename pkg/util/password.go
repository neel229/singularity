package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the bcrypt hash of the provided string
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error encrypting the password: %v", err)
	}

	return string(hashedPassword), nil
}

// VerifyPassword returns an error if the password and hashedPassword
// are not equal, and nil otherwise
func VerifyPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
