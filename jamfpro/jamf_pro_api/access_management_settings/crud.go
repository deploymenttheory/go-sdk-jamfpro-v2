package access_management_settings

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the access management settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment-access-management
	AccessManagementSettings struct {
		client client.Client
	}
)

func NewAccessManagementSettings(client client.Client) *AccessManagementSettings {
	return &AccessManagementSettings{client: client}
}

// GetV4 retrieves the current access management settings.
// URL: GET /api/v4/enrollment/access-management
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment-access-management
func (s *AccessManagementSettings) GetV4(ctx context.Context) (*ResourceAccessManagementSettings, *resty.Response, error) {
	var result ResourceAccessManagementSettings

	endpoint := constants.EndpointJamfProAccessManagementSettingsV4

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateV4 configures the access management settings.
// URL: POST /api/v4/enrollment/access-management
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v4-enrollment-access-management
func (s *AccessManagementSettings) CreateV4(ctx context.Context, request *ResourceAccessManagementSettings) (*ResourceAccessManagementSettings, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceAccessManagementSettings

	endpoint := constants.EndpointJamfProAccessManagementSettingsV4

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
