package local_admin_password

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the local admin password methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-pending-rotations
	LocalAdminPassword struct {
		client client.Client
	}
)

func NewLocalAdminPassword(client client.Client) *LocalAdminPassword {
	return &LocalAdminPassword{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Local Admin Password (LAPS) Operations
// -----------------------------------------------------------------------------

// GetPendingRotationsV2 retrieves a list of devices and usernames with pending LAPS rotations.
// URL: GET /api/v2/local-admin-password/pending-rotations
// https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-pending-rotations
func (s *LocalAdminPassword) GetPendingRotationsV2(ctx context.Context) (*PendingRotationsResponse, *resty.Response, error) {
	var result PendingRotationsResponse

	endpoint := fmt.Sprintf("%s/pending-rotations", constants.EndpointJamfProLocalAdminPasswordV2)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetSettingsV2 retrieves current Jamf Pro LAPS settings.
// URL: GET /api/v2/local-admin-password/settings
// https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-settings
func (s *LocalAdminPassword) GetSettingsV2(ctx context.Context) (*SettingsResource, *resty.Response, error) {
	var result SettingsResource

	endpoint := fmt.Sprintf("%s/settings", constants.EndpointJamfProLocalAdminPasswordV2)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateSettingsV2 updates the current Jamf Pro LAPS settings.
// URL: PUT /api/v2/local-admin-password/settings
// https://developer.jamf.com/jamf-pro/reference/put_v2-local-admin-password-settings
func (s *LocalAdminPassword) UpdateSettingsV2(ctx context.Context, settings *SettingsResource) (*resty.Response, error) {
	if settings == nil {
		return nil, fmt.Errorf("settings is required")
	}

	endpoint := fmt.Sprintf("%s/settings", constants.EndpointJamfProLocalAdminPasswordV2)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(settings).
		Put(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetPasswordHistoryByClientManagementIDV2 retrieves the password view history for a specific username on a target device.
// History will include password, who viewed the password and when it was viewed.
// URL: GET /api/v2/local-admin-password/{clientManagementId}/account/{username}/audit
// https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-account-username-audit
func (s *LocalAdminPassword) GetPasswordHistoryByClientManagementIDV2(ctx context.Context, clientManagementID string, username string) (*PasswordHistoryResponse, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}
	if username == "" {
		return nil, nil, fmt.Errorf("username is required")
	}

	endpoint := fmt.Sprintf("%s/%s/account/%s/audit", constants.EndpointJamfProLocalAdminPasswordV2, clientManagementID, username)

	var result PasswordHistoryResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetCurrentPasswordByClientManagementIDV2 retrieves the current LAPS password for a specific username on a target device.
// Note: Once viewed, the password will be rotated based on rotation time settings.
// URL: GET /api/v2/local-admin-password/{clientManagementId}/account/{username}/password
// https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-account-username-password
func (s *LocalAdminPassword) GetCurrentPasswordByClientManagementIDV2(ctx context.Context, clientManagementID string, username string) (*CurrentPasswordResponse, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}
	if username == "" {
		return nil, nil, fmt.Errorf("username is required")
	}

	endpoint := fmt.Sprintf("%s/%s/account/%s/password", constants.EndpointJamfProLocalAdminPasswordV2, clientManagementID, username)

	var result CurrentPasswordResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetHistoryByUsernameV2 retrieves the LAPS history for a specific username on a target device.
// History includes date created, date last seen, expiration time, and rotational status.
// URL: GET /api/v2/local-admin-password/{clientManagementId}/account/{username}/history
// https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-account-username-history
func (s *LocalAdminPassword) GetHistoryByUsernameV2(ctx context.Context, clientManagementID string, username string) (*AccountHistoryResponse, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}
	if username == "" {
		return nil, nil, fmt.Errorf("username is required")
	}

	endpoint := fmt.Sprintf("%s/%s/account/%s/history", constants.EndpointJamfProLocalAdminPasswordV2, clientManagementID, username)

	var result AccountHistoryResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetAuditByUsernameAndGUIDV2 retrieves the password view history for a specific username and guid on a target device.
// Use when multiple accounts share the same username; guid disambiguates them.
// URL: GET /api/v2/local-admin-password/{clientManagementId}/account/{username}/{guid}/audit
// https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-account-username-guid-audit
func (s *LocalAdminPassword) GetAuditByUsernameAndGUIDV2(ctx context.Context, clientManagementID string, username string, guid string) (*PasswordHistoryResponse, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}
	if username == "" {
		return nil, nil, fmt.Errorf("username is required")
	}
	if guid == "" {
		return nil, nil, fmt.Errorf("guid is required")
	}

	endpoint := fmt.Sprintf("%s/%s/account/%s/%s/audit", constants.EndpointJamfProLocalAdminPasswordV2, clientManagementID, username, guid)

	var result PasswordHistoryResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetHistoryByUsernameAndGUIDV2 retrieves the LAPS history for a specific username and guid on a target device.
// Use when multiple accounts share the same username; guid disambiguates them.
// URL: GET /api/v2/local-admin-password/{clientManagementId}/account/{username}/{guid}/history
// https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-account-username-guid-history
func (s *LocalAdminPassword) GetHistoryByUsernameAndGUIDV2(ctx context.Context, clientManagementID string, username string, guid string) (*AccountHistoryResponse, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}
	if username == "" {
		return nil, nil, fmt.Errorf("username is required")
	}
	if guid == "" {
		return nil, nil, fmt.Errorf("guid is required")
	}

	endpoint := fmt.Sprintf("%s/%s/account/%s/%s/history", constants.EndpointJamfProLocalAdminPasswordV2, clientManagementID, username, guid)

	var result AccountHistoryResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetPasswordByUsernameAndGUIDV2 retrieves the current LAPS password for a specific username and guid on a target device.
// Use when multiple accounts share the same username; guid disambiguates them.
// Note: Once viewed, the password will be rotated based on rotation time settings.
// URL: GET /api/v2/local-admin-password/{clientManagementId}/account/{username}/{guid}/password
// https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-account-username-guid-password
func (s *LocalAdminPassword) GetPasswordByUsernameAndGUIDV2(ctx context.Context, clientManagementID string, username string, guid string) (*CurrentPasswordResponse, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}
	if username == "" {
		return nil, nil, fmt.Errorf("username is required")
	}
	if guid == "" {
		return nil, nil, fmt.Errorf("guid is required")
	}

	endpoint := fmt.Sprintf("%s/%s/account/%s/%s/password", constants.EndpointJamfProLocalAdminPasswordV2, clientManagementID, username, guid)

	var result CurrentPasswordResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetFullHistoryByClientManagementIDV2 retrieves the complete history of all local admin passwords for all accounts
// on a specific device, including both viewing and rotation history.
// URL: GET /api/v2/local-admin-password/{clientManagementId}/history
// https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-history
func (s *LocalAdminPassword) GetFullHistoryByClientManagementIDV2(ctx context.Context, clientManagementID string) (*FullHistoryResponse, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProLocalAdminPasswordV2, clientManagementID)

	var result FullHistoryResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetCapableAccountsByClientManagementIDV2 retrieves a list of all admin accounts that are LAPS capable for a specific device.
// Capable accounts are returned in the AutoSetupAdminAccounts from QueryResponses.
// URL: GET /api/v2/local-admin-password/{clientManagementId}/accounts
// https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-accounts
func (s *LocalAdminPassword) GetCapableAccountsByClientManagementIDV2(ctx context.Context, clientManagementID string) (*CapableAccountsResponse, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/accounts", constants.EndpointJamfProLocalAdminPasswordV2, clientManagementID)

	var result CapableAccountsResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// SetPasswordByClientManagementIDV2 sets LAPS passwords for all capable accounts on a device.
// The passwords are provided as a list of username/password pairs.
// URL: PUT /api/v2/local-admin-password/{clientManagementId}/set-password
// https://developer.jamf.com/jamf-pro/reference/put_v2-local-admin-password-clientmanagementid-set-password
func (s *LocalAdminPassword) SetPasswordByClientManagementIDV2(ctx context.Context, clientManagementID string, passwordList *SetPasswordRequest) (*SetPasswordResponse, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}
	if passwordList == nil {
		return nil, nil, fmt.Errorf("passwordList is required")
	}

	endpoint := fmt.Sprintf("%s/%s/set-password", constants.EndpointJamfProLocalAdminPasswordV2, clientManagementID)

	var result SetPasswordResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(passwordList).
		SetResult(&result).
		Put(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
