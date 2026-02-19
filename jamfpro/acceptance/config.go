package acceptance

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	jamfpro "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
)

// TestConfig holds configuration for acceptance tests driven by environment variables.
// All credential variables mirror the names read by client.AuthConfigFromEnv().
type TestConfig struct {
	// Auth (read by client.AuthConfigFromEnv — see jamfpro/client/env.go)
	InstanceDomain string
	AuthMethod     string

	// OAuth2
	ClientID     string
	ClientSecret string

	// Basic auth
	Username string
	Password string

	// Test behaviour
	RequestTimeout time.Duration
	SkipCleanup    bool
	Verbose        bool
}

var (
	// Config is the global acceptance test configuration, initialised from env.
	Config *TestConfig
	// Client is the shared Jamf Pro SDK client for acceptance tests.
	Client *jamfpro.Client
)

func init() {
	Config = &TestConfig{
		InstanceDomain: getEnv("INSTANCE_DOMAIN", ""),
		AuthMethod:     getEnv("AUTH_METHOD", ""),
		ClientID:       getEnv("CLIENT_ID", ""),
		ClientSecret:   getEnv("CLIENT_SECRET", ""),
		Username:       getEnv("BASIC_AUTH_USERNAME", ""),
		Password:       getEnv("BASIC_AUTH_PASSWORD", ""),
		RequestTimeout: getDurationEnv("JAMF_REQUEST_TIMEOUT", 30*time.Second),
		SkipCleanup:    getBoolEnv("JAMF_SKIP_CLEANUP", false),
		Verbose:        getBoolEnv("JAMF_VERBOSE", false),
	}
}

// InitClient creates the shared Jamf Pro client from environment variables.
// Returns an error if required credentials are absent.
func InitClient() error {
	authConfig := client.AuthConfigFromEnv()
	if err := authConfig.Validate(); err != nil {
		return fmt.Errorf("invalid acceptance test credentials: %w", err)
	}

	var err error
	Client, err = jamfpro.NewClient(authConfig)
	if err != nil {
		return fmt.Errorf("failed to create Jamf Pro client: %w", err)
	}

	if Config.Verbose {
		log.Printf("Acceptance test client initialised: %s (%s)", Config.InstanceDomain, Config.AuthMethod)
	}
	return nil
}

// IsConfigured returns true if the minimum required credentials are set.
func IsConfigured() bool {
	return Config.InstanceDomain != "" && Config.AuthMethod != ""
}

// Helper functions to get environment variables.

// getEnv retrieves an environment variable or returns a default value.
func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

// getBoolEnv retrieves a boolean environment variable or returns a default value.
func getBoolEnv(key string, def bool) bool {
	v := os.Getenv(key)
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

// getDurationEnv retrieves a duration environment variable or returns a default value.
func getDurationEnv(key string, def time.Duration) time.Duration {
	v := os.Getenv(key)
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
