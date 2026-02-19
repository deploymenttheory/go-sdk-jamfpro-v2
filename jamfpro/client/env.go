package client

import (
	"os"
	"strconv"
	"time"
)

// AuthConfigFromEnv builds AuthConfig from environment variables.
// Required: INSTANCE_DOMAIN, AUTH_METHOD; for oauth2: CLIENT_ID, CLIENT_SECRET; for basic: BASIC_AUTH_USERNAME, BASIC_AUTH_PASSWORD.
// Optional: TOKEN_REFRESH_BUFFER_SECONDS (default 300), HIDE_SENSITIVE_DATA (default false).
func AuthConfigFromEnv() *AuthConfig {
	bufferSec := getEnvAsInt("TOKEN_REFRESH_BUFFER_SECONDS", 300)
	return &AuthConfig{
		InstanceDomain:           getEnv("INSTANCE_DOMAIN", ""),
		AuthMethod:               getEnv("AUTH_METHOD", ""),
		ClientID:                 getEnv("CLIENT_ID", ""),
		ClientSecret:             getEnv("CLIENT_SECRET", ""),
		Username:                 getEnv("BASIC_AUTH_USERNAME", ""),
		Password:                 getEnv("BASIC_AUTH_PASSWORD", ""),
		TokenRefreshBufferPeriod: time.Duration(bufferSec) * time.Second,
		HideSensitiveData:        getEnvAsBool("HIDE_SENSITIVE_DATA", false),
	}
}

func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}

func getEnvAsInt(key string, def int) int {
	v := getEnv(key, "")
	if v == "" {
		return def
	}
	i, _ := strconv.Atoi(v)
	return i
}

func getEnvAsBool(key string, def bool) bool {
	v := getEnv(key, "")
	if v == "" {
		return def
	}
	b, _ := strconv.ParseBool(v)
	return b
}
