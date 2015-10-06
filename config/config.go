package config

import (
	"github.com/crit/critical-go/config/file"
)

type Config interface {
	Get(key string) string
	Put(key string, value string)
}

func Basic() Config {
	return basic{}
}

func File(path string, fileName string) Config {
	return file.New(path, fileName)
}
