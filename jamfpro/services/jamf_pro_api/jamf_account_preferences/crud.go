package jamf_account_preferences

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// JamfAccountPreferencesServiceInterface defines the interface for Jamf Pro account preferences operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-account-preferences
	JamfAccountPreferencesServiceInterface interface {
		// GetV3 returns the current Jamf Pro account preferences.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-account-preferences
		GetV3(ctx context.Context) (*ResourceAccountPreferences, *interfaces.Response, error)

		// UpdateV3 updates Jamf Pro account preferences (PATCH).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v3-account-preferences
		UpdateV3(ctx context.Context, request *ResourceAccountPreferences) (*ResourceAccountPreferences, *interfaces.Response, error)
	}

	// Service handles communication with the Jamf Pro account preferences methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-account-preferences
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ JamfAccountPreferencesServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Account Preferences Operations
// -----------------------------------------------------------------------------

// GetV3 returns the current Jamf Pro account preferences.
// URL: GET /api/v3/account-preferences
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-account-preferences
func (s *Service) GetV3(ctx context.Context) (*ResourceAccountPreferences, *interfaces.Response, error) {
	var result ResourceAccountPreferences

	endpoint := EndpointAccountPreferencesV3

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateV3 updates Jamf Pro account preferences.
// URL: PATCH /api/v3/account-preferences
// Returns 204 No Content on success; the API may not return the updated resource.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v3-account-preferences
func (s *Service) UpdateV3(ctx context.Context, request *ResourceAccountPreferences) (*ResourceAccountPreferences, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceAccountPreferences

	endpoint := EndpointAccountPreferencesV3

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	// PATCH returns 204 No Content; if no body, return the request as the "updated" value
	if resp != nil && resp.StatusCode == 204 && len(resp.Body) == 0 {
		return request, resp, nil
	}

	return &result, resp, nil
}
