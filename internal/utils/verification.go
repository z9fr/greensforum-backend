package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHashSecret(secret, token string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(token+secret), 14)
	return string(bytes), err
}

func ValidateHash(value, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
	return err == nil
}
