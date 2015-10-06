package password

import (
	"code.google.com/p/go.crypto/bcrypt"
)

func Hash(in string) (string, error) {
	password := []byte(in)

	hashedPassword, err := bcrypt.GenerateFromPassword(password, 10)

	return string(hashedPassword), err
}

func Unequal(current string, request string) bool {
	a := []byte(current)
	b := []byte(request)

	err := bcrypt.CompareHashAndPassword(a, b)

	return err == nil
}
