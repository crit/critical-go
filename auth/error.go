package auth

import (
	"errors"
)

var (
	ErrorLoggedOut = errors.New("already logged out")
)
