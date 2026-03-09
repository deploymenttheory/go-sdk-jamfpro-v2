package account_preferences

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the account preferences-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-account-preferences
	AccountPreferences struct {
		client transport.HTTPClient
	}
)

func NewAccountPreferences(client transport.HTTPClient) *AccountPreferences {
	return &AccountPreferences{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Account Preferences Operations
// -----------------------------------------------------------------------------

// GetV3 returns the current account preferences.
// URL: GET /api/v3/account-preferences
// https://developer.jamf.com/jamf-pro/reference/get_v3-account-preferences
func (s *AccountPreferences) GetV3(ctx context.Context) (*ResourceAccountPreferencesV2, *resty.Response, error) {
	var result ResourceAccountPreferencesV2

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

// UpdateV3 updates account preferences.
// URL: PATCH /api/v3/account-preferences
// https://developer.jamf.com/jamf-pro/reference/patch_v3-account-preferences
func (s *AccountPreferences) UpdateV3(ctx context.Context, request *ResourceAccountPreferencesV2) (*ResourceAccountPreferencesV2, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceAccountPreferencesV2

	endpoint := constants.EndpointJamfProAccountPreferencesV3

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
