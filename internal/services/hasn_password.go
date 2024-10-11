package services

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	HashPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	return string(HashPassword), err
}
