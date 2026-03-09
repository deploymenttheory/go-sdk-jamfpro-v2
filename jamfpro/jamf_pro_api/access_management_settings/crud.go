package access_management_settings

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// AccessManagementSettingsServiceInterface defines the interface for access management settings operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment-access-management
	AccessManagementSettingsServiceInterface interface {
		// GetV4 retrieves the current access management settings (Get Access Management Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment-access-management
		GetV4(ctx context.Context) (*ResourceAccessManagementSettings, *resty.Response, error)

		// CreateV4 configures the access management settings (Create/Update Access Management Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v4-enrollment-access-management
		CreateV4(ctx context.Context, request *ResourceAccessManagementSettings) (*ResourceAccessManagementSettings, *resty.Response, error)
	}

	// Service handles communication with the access management settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment-access-management
	AccessManagementSettings struct {
		client interfaces.HTTPClient
	}
)

var _ AccessManagementSettingsServiceInterface = (*AccessManagementSettings)(nil)

func NewAccessManagementSettings(client interfaces.HTTPClient) *AccessManagementSettings {
	return &AccessManagementSettings{client: client}
}

// GetV4 retrieves the current access management settings.
// URL: GET /api/v4/enrollment/access-management
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment-access-management
func (s *AccessManagementSettings) GetV4(ctx context.Context) (*ResourceAccessManagementSettings, *resty.Response, error) {
	var result ResourceAccessManagementSettings
	endpoint := EndpointAccessManagementSettingsV4
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
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
	endpoint := EndpointAccessManagementSettingsV4
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
