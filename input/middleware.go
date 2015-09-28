package input

import (
	"github.com/codegangsta/inject"
	"net/http"
)

// Context matches the martini.Context interface.
type Context interface {
	inject.Injector
	Next()
	Written() bool
}

// MartiniMiddleware is a helper method for martini based projects.
// Martini: https://github.com/go-martini/martini
func MartiniMiddleware() interface{} {
	return func(req *http.Request, c Context) {
		c.Map(New(req))
	}
}
