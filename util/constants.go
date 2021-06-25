package util

import (
	"errors"
	"log"
	"os"
)

func GetEnvironment(key string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		log.Fatalln(errors.New("No environment variable specified for " + key))
	}

	return value
}

func GetJWTSecret() []byte {
	return []byte(GetEnvironment("JWT_SECRET"))
}

// DefaultFormMaxMem 32 MB, default
var DefaultFormMaxMem int64 = 32 << 20
