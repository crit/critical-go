package password

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash creates an encrypted string from the passed string value.
func Hash(in string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in), 4)

	return string(hashedPassword)
}

// Unequal compares a current hash string to a requested password string value for equality.
func Unequal(current string, request string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(current), []byte(request))

	return err != nil
}
