package util

import (
	"cdn/router/functions"
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

func GetJWTWrapper() functions.JwtWrapper {
	key := []byte(GetEnvironment("JWT_SECRET"))

	return functions.JwtWrapper{
		SecretKey:       key,
		Issuer:          "localhost22",
		ExpirationHours: 24 * 7,
	}
}

// DefaultFormMaxMem 32 MB, default
var DefaultFormMaxMem int64 = 32 << 20
