package jamf_pro_system_initialization

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Jamf Pro system initialization API.
	JamfProSystemInitialization struct {
		client client.Client
	}
)

func NewJamfProSystemInitialization(client client.Client) *JamfProSystemInitialization {
	return &JamfProSystemInitialization{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - System Initialization Operations
// -----------------------------------------------------------------------------

// Initialize initializes a fresh Jamf Pro Server installation.
// URL: POST /api/v1/system/initialize
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-system-initialize
func (s *JamfProSystemInitialization) Initialize(ctx context.Context, request *ResourceSystemInitialize) (*resty.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := constants.EndpointJamfProSystemInitialize

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to initialize Jamf Pro system: %w", err)
	}

	return resp, nil
}

// InitializeDatabaseConnection sets up the database password during startup.
// URL: POST /api/v1/system/initialize-database-connection
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-system-initialize-database-connection
func (s *JamfProSystemInitialization) InitializeDatabaseConnection(ctx context.Context, password string) (*resty.Response, error) {
	endpoint := constants.EndpointJamfProInitializeDatabaseConnection

	request := &ResourceDatabasePassword{
		Password: password,
	}

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to initialize database connection: %w", err)
	}

	return resp, nil
}

// PlatformInitialize sets up Jamf Pro Server with OIDC SSO and a federated user (no password required).
// URL: POST /api/v1/system/platform-initialize
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-system-platform-initialize
func (s *JamfProSystemInitialization) PlatformInitialize(ctx context.Context, request *ResourcePlatformInitialize) (*resty.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := constants.EndpointJamfProSystemPlatformInitializeV1

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to platform initialize Jamf Pro system: %w", err)
	}

	return resp, nil
}
