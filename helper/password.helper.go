package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plainPassword string) string {
	bytesString, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 14)
	if err != nil {
		log.Fatal(err)
	}

	return string(bytesString)
}

func VerifyPassword(plainPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	check := true

	if err != nil {
		check = false
	}

	return check
}
