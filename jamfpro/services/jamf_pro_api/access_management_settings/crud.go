package access_management_settings

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// AccessManagementSettingsServiceInterface defines the interface for access management settings operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment-access-management
	AccessManagementSettingsServiceInterface interface {
		// GetV4 retrieves the current access management settings (Get Access Management Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment-access-management
		GetV4(ctx context.Context) (*ResourceAccessManagementSettings, *interfaces.Response, error)

		// CreateV4 configures the access management settings (Create/Update Access Management Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v4-enrollment-access-management
		CreateV4(ctx context.Context, request *ResourceAccessManagementSettings) (*ResourceAccessManagementSettings, *interfaces.Response, error)
	}

	// Service handles communication with the access management settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment-access-management
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ AccessManagementSettingsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetV4 retrieves the current access management settings.
// URL: GET /api/v4/enrollment/access-management
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment-access-management
func (s *Service) GetV4(ctx context.Context) (*ResourceAccessManagementSettings, *interfaces.Response, error) {
	var result ResourceAccessManagementSettings
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Get(ctx, EndpointAccessManagementSettingsV4, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// CreateV4 configures the access management settings.
// URL: POST /api/v4/enrollment/access-management
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v4-enrollment-access-management
func (s *Service) CreateV4(ctx context.Context, request *ResourceAccessManagementSettings) (*ResourceAccessManagementSettings, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	var result ResourceAccessManagementSettings
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Post(ctx, EndpointAccessManagementSettingsV4, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
