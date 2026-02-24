package account_preferences

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// AccountPreferencesServiceInterface defines the interface for account preferences operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-account-preferences
	AccountPreferencesServiceInterface interface {
		// GetV3 returns the current account preferences (Get Jamf Pro Account Preferences).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-account-preferences
		GetV3(ctx context.Context) (*ResourceAccountPreferencesV2, *interfaces.Response, error)

		// UpdateV3 updates account preferences (Update Jamf Pro Account Preferences / PATCH).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v3-account-preferences
		UpdateV3(ctx context.Context, request *ResourceAccountPreferencesV2) (*ResourceAccountPreferencesV2, *interfaces.Response, error)
	}

	// Service handles communication with the account preferences-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-account-preferences
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ AccountPreferencesServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Account Preferences Operations
// -----------------------------------------------------------------------------

// GetV3 returns the current account preferences.
// URL: GET /api/v3/account-preferences
// https://developer.jamf.com/jamf-pro/reference/get_v3-account-preferences
func (s *Service) GetV3(ctx context.Context) (*ResourceAccountPreferencesV2, *interfaces.Response, error) {
	var result ResourceAccountPreferencesV2

	endpoint := EndpointAccountPreferencesV3

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateV3 updates account preferences.
// URL: PATCH /api/v3/account-preferences
// https://developer.jamf.com/jamf-pro/reference/patch_v3-account-preferences
func (s *Service) UpdateV3(ctx context.Context, request *ResourceAccountPreferencesV2) (*ResourceAccountPreferencesV2, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceAccountPreferencesV2

	endpoint := EndpointAccountPreferencesV3

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
