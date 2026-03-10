package macos_configuration_profile_custom_settings

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the macOS configuration profile custom settings
	// methods of the Jamf Pro API.
	//
	// API reference: Undocumented
	MacosConfigurationProfileCustomSettings struct {
		client client.Client
	}
)

// NewService creates a new macOS configuration profile custom settings service.
func NewMacosConfigurationProfileCustomSettings(client client.Client) *MacosConfigurationProfileCustomSettings {
	return &MacosConfigurationProfileCustomSettings{client: client}
}

// GetSchemaList retrieves the list of custom settings schemas.
// URL: GET /api/config-profiles/macos/custom-settings/v1/schema-list
func (s *MacosConfigurationProfileCustomSettings) GetSchemaList(ctx context.Context) (*ResponseCustomSettingsSchemaList, *resty.Response, error) {
	endpoint := constants.EndpointJamfProCustomSettingsSchemaList

	var result ResponseCustomSettingsSchemaList

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
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

	var result ResourceConfigProfile

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
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

	var result ResponseConfigProfileCreate

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(profile).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create macOS configuration profile: %w", err)
	}

	return &result, resp, nil
}
