package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(pasword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(pasword))
	return err == nil
}
