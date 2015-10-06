package config

import (
	"os"
)

type basic struct{}

func (basic) Get(key string) string {
	return os.Getenv(key)
}

func (basic) Put(key, value string) {
	os.Setenv(key, value)
}
