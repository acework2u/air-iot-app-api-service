package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Could not hash password %w", err)
	}
	return string(hashPassword), nil
}
func VerifyPassword(hashPassword string, candidatePassword string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(candidatePassword))
}
