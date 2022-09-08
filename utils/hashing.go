package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password *string) error {
	bytePass := []byte(*password)
	hPass, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	*password = string(hPass)
	return err
}
