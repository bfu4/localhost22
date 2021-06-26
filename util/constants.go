package util

import (
	"errors"
	"os"
)

func GetEnvironment(key string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		err := errors.New("No environment variable specified for " + key)
		Fatal(err.Error())
	}

	return value
}

func GetJWTSecret() []byte {
	return []byte(GetEnvironment("JWT_SECRET"))
}

// DefaultFormMaxMem 32 MB, default
var DefaultFormMaxMem int64 = 32 << 20
