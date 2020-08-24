package errors

import (
	"fmt"
	"net/http"
)

// Err is the base error structure.
type Err struct {
	Code    int
	Message string
}

// Error returns error code and message in a formatted string. Satisfies
// the error interface.
func (e Err) Error() string {
	return fmt.Sprintf("[%d] - %s", e.Code, e.Message)
}

// Code attempts to cast the error to Err and determines the code. Nil is
// 200; Non-Err is 500.
func Code(err error) int {
	if err == nil {
		return http.StatusOK
	}

	e, ok := err.(Err)

	if !ok {
		return http.StatusInternalServerError
	}

	return e.Code
}

// Message attempts to cast the error to Err and determines the message. Nil is an empty
// string; Non-Err is the error interface output.
func Message(err error) string {
	if err == nil {
		return ""
	}

	e, ok := err.(Err)

	if !ok {
		return err.Error()
	}

	return e.Message
}

// Error is a convenience function that creates a new Err with code and
// the Message of the error.
func Error(code int, err error) error {
	return Err{code, Message(err)}
}

// ErrorString is a convenience function that creates a new Err with code and a message
func ErrorString(code int, msg string) error {
	return Err{code, msg}
}

// ErrorPackage is a convenience function to output a standard response for
// an HTTP handler of our specific API Service.
func ErrorPackage(err error) map[string]interface{} {
	if err == nil {
		return nil
	}

	return map[string]interface{}{
		"code":    Code(err),
		"message": Message(err),
	}
}
