package input

import (
	"net/http"
	"strings"
)

// concrete type that implements the Input interface
type basic struct {
	data map[string]string
}

func (i *basic) load(req *http.Request) {
	i.data = map[string]string{}

	req.ParseForm()

	for key, values := range req.Form {
		i.data[key] = strings.Join(values, "")
	}
}

func (i basic) All() map[string]string {
	return i.data
}

func (i basic) Get(key string) string {
	for k, value := range i.All() {
		if key == k {
			return value
		}
	}

	return ""
}

func (i basic) Has(key string) bool {
	_, has := i.All()[key]
	return has
}

func (i basic) Only(include ...string) map[string]string {
	keep := map[string]bool{}
	out := map[string]string{}

	for _, key := range include {
		keep[key] = true
	}

	for key, value := range i.All() {
		if keep[key] {
			out[key] = value
		}
	}

	return out
}

func (i basic) Except(exclude ...string) map[string]string {
	remove := map[string]bool{}
	out := map[string]string{}

	for _, key := range exclude {
		remove[key] = true
	}

	for key, value := range i.All() {
		if !remove[key] {
			out[key] = value
		}
	}

	return out
}
