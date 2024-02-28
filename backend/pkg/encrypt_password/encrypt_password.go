package encrypt_password

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPassword(password, hash string) (isValid bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
