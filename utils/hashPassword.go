package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPasword(pasword string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(pasword), 10)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
