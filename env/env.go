package env

import "os"

func GetString(key, def string) string {
	value := os.Getenv(key)

	if value == "" {
		return def
	}

	return value
}
