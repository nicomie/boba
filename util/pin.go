package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPin(pin string) (string, error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(pin), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("failed to hash pin: %w", err)
	}

	return string(hashed), nil

}

func CheckPin(pin string, hashedPin string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPin), []byte(pin))
}
