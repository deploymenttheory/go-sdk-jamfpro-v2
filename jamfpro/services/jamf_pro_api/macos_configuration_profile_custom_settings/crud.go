package macos_configuration_profile_custom_settings

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

// ServiceInterface defines the interface for macOS configuration profile custom settings operations.
//
// This service provides access to custom settings schemas and configuration profiles
// for macOS devices. API reference: Undocumented.
type ServiceInterface interface {
	// GetSchemaList retrieves the list of custom settings schemas.
	//
	// Returns the schema list organized by buckets and domains.
	GetSchemaList(ctx context.Context) (*ResponseCustomSettingsSchemaList, *interfaces.Response, error)

	// GetByPayloadUUID retrieves a macOS configuration profile by payload UUID.
	//
	// id is the payload UUID of the configuration profile.
	GetByPayloadUUID(ctx context.Context, id string) (*ResourceConfigProfile, *interfaces.Response, error)

	// Create creates a new macOS configuration profile with custom settings schema.
	//
	// profile is the configuration profile to create.
	Create(ctx context.Context, profile *ResourceConfigProfile) (*ResponseConfigProfileCreate, *interfaces.Response, error)
}

type (
	// Service handles communication with the macOS configuration profile custom settings
	// methods of the Jamf Pro API.
	//
	// API reference: Undocumented
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ServiceInterface = (*Service)(nil)

// NewService creates a new macOS configuration profile custom settings service.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetSchemaList retrieves the list of custom settings schemas.
// URL: GET /api/config-profiles/macos/custom-settings/v1/schema-list
func (s *Service) GetSchemaList(ctx context.Context) (*ResponseCustomSettingsSchemaList, *interfaces.Response, error) {
	endpoint := EndpointCustomSettingsSchemaList

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	var result ResponseCustomSettingsSchemaList
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get custom settings schema list: %w", err)
	}

	return &result, resp, nil
}

// GetByPayloadUUID retrieves a macOS configuration profile by payload UUID.
// URL: GET /api/config-profiles/macos/{id}
func (s *Service) GetByPayloadUUID(ctx context.Context, id string) (*ResourceConfigProfile, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("payload UUID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointConfigProfilesMacOS, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	var result ResourceConfigProfile
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get macOS configuration profile with ID %s: %w", id, err)
	}

	return &result, resp, nil
}

// Create creates a new macOS configuration profile with custom settings schema.
// URL: POST /api/config-profiles/macos
func (s *Service) Create(ctx context.Context, profile *ResourceConfigProfile) (*ResponseConfigProfileCreate, *interfaces.Response, error) {
	if profile == nil {
		return nil, nil, fmt.Errorf("profile is required")
	}

	endpoint := EndpointConfigProfilesMacOS

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	var result ResponseConfigProfileCreate
	resp, err := s.client.Post(ctx, endpoint, profile, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create macOS configuration profile: %w", err)
	}

	return &result, resp, nil
}
