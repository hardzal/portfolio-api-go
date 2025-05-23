package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Generate(raw string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}

func Verify(hash string, raw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
}
