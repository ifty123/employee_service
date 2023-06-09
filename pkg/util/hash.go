package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(plainPw string) (hashedPw string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPw), 10)
	return string(bytes), err
}

func CompareHashPassword(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
