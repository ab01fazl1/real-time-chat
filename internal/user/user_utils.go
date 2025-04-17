package user

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPass(pasword string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pasword), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}

func CheckPassHash(hashedpassword string, passowrd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(passowrd))
	if err != nil {
		return err
	}
	return nil
}

