package jamf_pro_user_account_settings

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// JamfProUserAccountSettingsServiceInterface defines the interface for Jamf Pro user account settings operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-user-preferences-settings-keyid
	JamfProUserAccountSettingsServiceInterface interface {
		// GetSettingsV1 returns the user preferences for the authenticated user and key (username, key, values).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-user-preferences-settings-keyid
		GetSettingsV1(ctx context.Context, keyID string) (*ResourceUserPreferencesSettings, *resty.Response, error)

		// GetV1 returns the user setting value for the authenticated user and key (string value).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-user-preferences-keyid
		GetV1(ctx context.Context, keyID string) (string, *resty.Response, error)

		// PutV1 persists the user setting for the authenticated user and key.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-user-preferences-keyid
		PutV1(ctx context.Context, keyID string, values RequestUserPreferences) (*resty.Response, error)

		// DeleteV1 removes the specified setting for the authenticated user.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-user-preferences-keyid
		DeleteV1(ctx context.Context, keyID string) (*resty.Response, error)
	}

	// Service handles communication with the Jamf Pro user account settings methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-user-preferences-settings-keyid
	JamfProUserAccountSettings struct {
		client transport.HTTPClient
	}
)

var _ JamfProUserAccountSettingsServiceInterface = (*JamfProUserAccountSettings)(nil)

func NewJamfProUserAccountSettings(client transport.HTTPClient) *JamfProUserAccountSettings {
	return &JamfProUserAccountSettings{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - User Account Settings Operations
// -----------------------------------------------------------------------------

// GetSettingsV1 returns the user preferences for the authenticated user and key.
// URL: GET /api/v1/user/preferences/settings/{keyId}
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-user-preferences-settings-keyid
func (s *JamfProUserAccountSettings) GetSettingsV1(ctx context.Context, keyID string) (*ResourceUserPreferencesSettings, *resty.Response, error) {
	if keyID == "" {
		return nil, nil, fmt.Errorf("keyId is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointUserPreferencesSettingsV1, keyID)

	var result ResourceUserPreferencesSettings

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetV1 returns the user setting value for the authenticated user and key.
// URL: GET /api/v1/user/preferences/{keyId}
// Returns the raw string value (API may return JSON string or plain string).
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-user-preferences-keyid
func (s *JamfProUserAccountSettings) GetV1(ctx context.Context, keyID string) (string, *resty.Response, error) {
	if keyID == "" {
		return "", nil, fmt.Errorf("keyId is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointUserPreferencesV1, keyID)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, body, err := s.client.GetBytes(ctx, endpoint, nil, headers)
	if err != nil {
		return "", resp, err
	}

	// API may return JSON-encoded string (e.g. "value") or plain string
	var raw string
	if err := json.Unmarshal(body, &raw); err != nil {
		// If not valid JSON string, treat as raw body
		return string(body), resp, nil
	}
	return raw, resp, nil
}

// PutV1 persists the user setting for the authenticated user and key.
// URL: PUT /api/v1/user/preferences/{keyId}
// Body: map of key-value pairs to persist.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-user-preferences-keyid
func (s *JamfProUserAccountSettings) PutV1(ctx context.Context, keyID string, values RequestUserPreferences) (*resty.Response, error) {
	if keyID == "" {
		return nil, fmt.Errorf("keyId is required")
	}
	if values == nil || len(values) == 0 {
		return nil, fmt.Errorf("values is required and must not be empty")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointUserPreferencesV1, keyID)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, values, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteV1 removes the specified setting for the authenticated user.
// URL: DELETE /api/v1/user/preferences/{keyId}
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-user-preferences-keyid
func (s *JamfProUserAccountSettings) DeleteV1(ctx context.Context, keyID string) (*resty.Response, error) {
	if keyID == "" {
		return nil, fmt.Errorf("keyId is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointUserPreferencesV1, keyID)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
