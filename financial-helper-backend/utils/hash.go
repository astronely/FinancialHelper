package util

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	log.Println("HASH: ", hash)

	return string(hash), err
}

func VerifyPassword(providedPassword, userPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
}
