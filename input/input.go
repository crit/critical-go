package input

import (
	"net/http"
)

// Input proxies the http.Request.Form values as a simple map[string]string. It allows
// convenience methods on the values. More complex usage of http.Request.Form should use
// the built in net/http package directly.
type Input interface {
	// All returns the map[string]string representation of http.Request.Form that
	// the Input was initialized with.
	All() map[string]string

	// Get returns either the key's value or an empty string. Note that
	// Input.Has should be used to check for whether the key was even passed
	// to the server.
	Get(key string) string

	// Has returns whether a key was passed to the server. It does not check for
	// whether a value was set.
	Has(key string) bool

	// Only provides a filtered version of Input.All with only the keys listed in the return.
	Only(include ...string) map[string]string

	// Except provides a filtered version of Input.All with all the keys listed in the return
	// except the passed in keys.
	Except(exclude ...string) map[string]string
}

// New returns a basic Input provider. Best used in a middleware.
func New(req *http.Request) Input {
	b := basic{}
	b.load(req)

	return b
}
