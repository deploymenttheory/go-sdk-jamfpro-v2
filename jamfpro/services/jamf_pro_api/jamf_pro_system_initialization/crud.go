package jamf_pro_system_initialization

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ServiceInterface defines the interface for Jamf Pro system initialization operations.
	//
	// These endpoints are used during initial Jamf Pro setup: system initialization
	// and database connection configuration.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-system-initialize
	ServiceInterface interface {
		// Initialize initializes a fresh Jamf Pro Server installation.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-system-initialize
		Initialize(ctx context.Context, request *ResourceSystemInitialize) (*interfaces.Response, error)

		// InitializeDatabaseConnection sets up the database password during startup.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-system-initialize-database-connection
		InitializeDatabaseConnection(ctx context.Context, password string) (*interfaces.Response, error)
	}

	// Service handles communication with the Jamf Pro system initialization API.
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - System Initialization Operations
// -----------------------------------------------------------------------------

// Initialize initializes a fresh Jamf Pro Server installation.
// URL: POST /api/v1/system/initialize
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-system-initialize
func (s *Service) Initialize(ctx context.Context, request *ResourceSystemInitialize) (*interfaces.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	endpoint := EndpointSystemInitialize

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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
func (s *Service) InitializeDatabaseConnection(ctx context.Context, password string) (*interfaces.Response, error) {
	endpoint := EndpointInitializeDatabaseConnection

	request := &ResourceDatabasePassword{
		Password: password,
	}

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to initialize database connection: %w", err)
	}

	return resp, nil
}
