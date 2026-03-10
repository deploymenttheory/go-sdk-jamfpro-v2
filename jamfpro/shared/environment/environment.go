package environment

import (
	"log"
	"os"
	"strconv"
	"time"
)

// GetEnv retrieves an environment variable or returns a default value.
// Uses LookupEnv so that a variable explicitly set to "" is returned as "".
func GetEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}

// GetEnvAsInt retrieves an integer environment variable or returns a default value.
func GetEnvAsInt(key string, def int) int {
	v := GetEnv(key, "")
	if v == "" {
		return def
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		log.Printf("Warning: invalid int for %s=%q, using default %v", key, v, def)
		return def
	}
	return i
}

// GetEnvAsBool retrieves a boolean environment variable or returns a default value.
func GetEnvAsBool(key string, def bool) bool {
	v := GetEnv(key, "")
	if v == "" {
		return def
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.Printf("Warning: invalid bool for %s=%q, using default %v", key, v, def)
		return def
	}
	return b
}

// GetDurationEnv retrieves a duration environment variable or returns a default value.
func GetDurationEnv(key string, def time.Duration) time.Duration {
	v := GetEnv(key, "")
	if v == "" {
		return def
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		log.Printf("Warning: invalid duration for %s=%q, using default %v", key, v, def)
		return def
	}
	return d
}
