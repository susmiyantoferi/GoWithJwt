package utils

import "golang.org/x/crypto/bcrypt"

func ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
	return err == nil
}
