package jamf_account_preferences

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Jamf Pro account preferences methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-account-preferences
	JamfAccountPreferences struct {
		client transport.HTTPClient
	}
)

func NewJamfAccountPreferences(client transport.HTTPClient) *JamfAccountPreferences {
	return &JamfAccountPreferences{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Account Preferences Operations
// -----------------------------------------------------------------------------

// GetV3 returns the current Jamf Pro account preferences.
// URL: GET /api/v3/account-preferences
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-account-preferences
func (s *JamfAccountPreferences) GetV3(ctx context.Context) (*ResourceAccountPreferences, *resty.Response, error) {
	var result ResourceAccountPreferences

	endpoint := constants.EndpointJamfProAccountPreferencesV3

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *JamfAccountPreferences) UpdateV3(ctx context.Context, request *ResourceAccountPreferences) (*ResourceAccountPreferences, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceAccountPreferences

	endpoint := constants.EndpointJamfProAccountPreferencesV3

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	// PATCH returns 204 No Content; if no body, return the request as the "updated" value
	bodyBytes := resp.Bytes()
	if resp != nil && resp.StatusCode() == 204 && len(bodyBytes) == 0 {
		return request, resp, nil
	}

	return &result, resp, nil
}
