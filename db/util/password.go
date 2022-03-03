package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hashed password: %w", err)
	}
	return string(h), nil
}

// CheckPassword will compare password p and the hashed h is equal if it equals return nil
func CheckPassword(p, h string) error {
	// h: hased password
	// p: password
	return bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
}