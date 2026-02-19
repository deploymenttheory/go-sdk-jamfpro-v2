package jamfpro

import (
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/categories"
	"go.uber.org/zap"
)

// Client is the main entry point for the Jamf Pro API SDK.
type Client struct {
	transport *client.Transport

	Categories *categories.Service
}

// NewClient creates a new Jamf Pro API client.
func NewClient(authConfig *client.AuthConfig, options ...client.ClientOption) (*Client, error) {
	transport, err := client.NewTransport(authConfig, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create transport: %w", err)
	}
	c := &Client{
		transport:  transport,
		Categories: categories.NewService(transport),
	}
	return c, nil
}

// NewClientFromEnv creates a new client using environment variables.
// Required: INSTANCE_DOMAIN, AUTH_METHOD; for oauth2: CLIENT_ID, CLIENT_SECRET; for basic: BASIC_AUTH_USERNAME, BASIC_AUTH_PASSWORD.
func NewClientFromEnv(options ...client.ClientOption) (*Client, error) {
	authConfig := client.AuthConfigFromEnv()
	if err := authConfig.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config from env: %w", err)
	}
	return NewClient(authConfig, options...)
}

// GetLogger returns the configured zap logger.
func (c *Client) GetLogger() *zap.Logger {
	return c.transport.GetLogger()
}

// EnableTracing enables OpenTelemetry HTTP tracing on the client's transport.
// Pass nil to use default OTel config (global tracer, "jamfpro-client" service name).
func (c *Client) EnableTracing(config *client.OTelConfig) error {
	return c.transport.EnableTracing(config)
}
