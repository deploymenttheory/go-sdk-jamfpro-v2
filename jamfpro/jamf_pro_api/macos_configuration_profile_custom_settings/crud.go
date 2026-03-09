package macos_configuration_profile_custom_settings

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

// ServiceInterface defines the interface for macOS configuration profile custom settings operations.
//
// This service provides access to custom settings schemas and configuration profiles
// for macOS devices. API reference: Undocumented.
type ServiceInterface interface {
	// GetSchemaList retrieves the list of custom settings schemas.
	//
	// Returns the schema list organized by buckets and domains.
	GetSchemaList(ctx context.Context) (*ResponseCustomSettingsSchemaList, *resty.Response, error)

	// GetByPayloadUUID retrieves a macOS configuration profile by payload UUID.
	//
	// id is the payload UUID of the configuration profile.
	GetByPayloadUUID(ctx context.Context, id string) (*ResourceConfigProfile, *resty.Response, error)

	// Create creates a new macOS configuration profile with custom settings schema.
	//
	// profile is the configuration profile to create.
	Create(ctx context.Context, profile *ResourceConfigProfile) (*ResponseConfigProfileCreate, *resty.Response, error)
}

type (
	// Service handles communication with the macOS configuration profile custom settings
	// methods of the Jamf Pro API.
	//
	// API reference: Undocumented
	MacosConfigurationProfileCustomSettings struct {
		client transport.HTTPClient
	}
)

var _ ServiceInterface = (*MacosConfigurationProfileCustomSettings)(nil)

// NewService creates a new macOS configuration profile custom settings service.
func NewMacosConfigurationProfileCustomSettings(client transport.HTTPClient) *MacosConfigurationProfileCustomSettings {
	return &MacosConfigurationProfileCustomSettings{client: client}
}

// GetSchemaList retrieves the list of custom settings schemas.
// URL: GET /api/config-profiles/macos/custom-settings/v1/schema-list
func (s *MacosConfigurationProfileCustomSettings) GetSchemaList(ctx context.Context) (*ResponseCustomSettingsSchemaList, *resty.Response, error) {
	endpoint := constants.EndpointJamfProCustomSettingsSchemaList

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *MacosConfigurationProfileCustomSettings) GetByPayloadUUID(ctx context.Context, id string) (*ResourceConfigProfile, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("payload UUID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProConfigProfilesMacOS, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *MacosConfigurationProfileCustomSettings) Create(ctx context.Context, profile *ResourceConfigProfile) (*ResponseConfigProfileCreate, *resty.Response, error) {
	if profile == nil {
		return nil, nil, fmt.Errorf("profile is required")
	}

	endpoint := constants.EndpointJamfProConfigProfilesMacOS

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	var result ResponseConfigProfileCreate
	resp, err := s.client.Post(ctx, endpoint, profile, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create macOS configuration profile: %w", err)
	}

	return &result, resp, nil
}
