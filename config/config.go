package config

import (
	"github.com/crit/critical-go/config/file"
	"github.com/subosito/gotenv"
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

func Dotfile(path string) Config {
	gotenv.Load(path)
	return basic{}
}
