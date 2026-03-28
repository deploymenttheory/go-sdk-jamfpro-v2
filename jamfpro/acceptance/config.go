package acceptance

import (
	"fmt"
	"log"
	"time"

	jamfpro "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/environment"
)

// TestConfig holds configuration for acceptance tests driven by environment variables.
// All credential variables mirror the names read by jamfpro.AuthConfigFromEnv().
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
		InstanceDomain: environment.GetEnv("INSTANCE_DOMAIN", ""),
		AuthMethod:     environment.GetEnv("AUTH_METHOD", ""),
		ClientID:       environment.GetEnv("CLIENT_ID", ""),
		ClientSecret:   environment.GetEnv("CLIENT_SECRET", ""),
		Username:       environment.GetEnv("BASIC_AUTH_USERNAME", ""),
		Password:       environment.GetEnv("BASIC_AUTH_PASSWORD", ""),
		RequestTimeout: environment.GetDurationEnv("JAMF_REQUEST_TIMEOUT", 30*time.Second),
		SkipCleanup:    environment.GetEnvAsBool("JAMF_SKIP_CLEANUP", false),
		Verbose:        environment.GetEnvAsBool("JAMF_VERBOSE", false),
	}
}

// InitClient creates the shared Jamf Pro client from environment variables.
// Returns an error if required credentials are absent.
func InitClient() error {
	authConfig := jamfpro.AuthConfigFromEnv()
	if err := authConfig.Validate(); err != nil {
		return fmt.Errorf("invalid acceptance test credentials: %w", err)
	}

	var err error
	transportTimeout := environment.GetDurationEnv("JAMF_TRANSPORT_TIMEOUT", 10*time.Minute)
	Client, err = jamfpro.NewClient(authConfig, jamfpro.WithTimeout(transportTimeout))
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
