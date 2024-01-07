package util

import (
	"errors"
	"log"
	"os"
)

// Getenv returns the value of the environment variable named by the key.
func Getenv(key string, defaultValue ...string) string {
	value := os.Getenv(key)

	if value != "" {
		return value
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	log.Fatal(errors.New("Environment variable " + key + " is not set"))
	return ""
}
