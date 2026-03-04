package local_admin_password

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// LocalAdminPasswordServiceInterface defines the interface for LAPS operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-pending-rotations
	LocalAdminPasswordServiceInterface interface {
		// GetPendingRotationsV2 retrieves a list of devices and usernames with pending LAPS rotations.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-pending-rotations
		GetPendingRotationsV2(ctx context.Context) (*PendingRotationsResponse, *resty.Response, error)

		// GetSettingsV2 retrieves current Jamf Pro LAPS settings.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-settings
		GetSettingsV2(ctx context.Context) (*SettingsResource, *resty.Response, error)

		// UpdateSettingsV2 updates the current Jamf Pro LAPS settings.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-local-admin-password-settings
		UpdateSettingsV2(ctx context.Context, settings *SettingsResource) (*resty.Response, error)

		// GetPasswordHistoryByClientManagementIDV2 retrieves the password view history for a specific username on a target device.
		// History will include password, who viewed the password and when it was viewed.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-account-username-audit
		GetPasswordHistoryByClientManagementIDV2(ctx context.Context, clientManagementID string, username string) (*PasswordHistoryResponse, *resty.Response, error)

		// GetHistoryByUsernameV2 retrieves the LAPS history for a specific username on a target device.
		// History includes date created, date last seen, expiration time, and rotational status.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-account-username-history
		GetHistoryByUsernameV2(ctx context.Context, clientManagementID string, username string) (*AccountHistoryResponse, *resty.Response, error)

		// GetCurrentPasswordByClientManagementIDV2 retrieves the current LAPS password for a specific username on a target device.
		// Note: Once viewed, the password will be rotated based on rotation time settings.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-account-username-password
		GetCurrentPasswordByClientManagementIDV2(ctx context.Context, clientManagementID string, username string) (*CurrentPasswordResponse, *resty.Response, error)

		// GetAuditByUsernameAndGUIDV2 retrieves the password view history for a specific username and guid on a target device.
		// Use when multiple accounts share the same username; guid disambiguates them.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-account-username-guid-audit
		GetAuditByUsernameAndGUIDV2(ctx context.Context, clientManagementID string, username string, guid string) (*PasswordHistoryResponse, *resty.Response, error)

		// GetHistoryByUsernameAndGUIDV2 retrieves the LAPS history for a specific username and guid on a target device.
		// Use when multiple accounts share the same username; guid disambiguates them.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-account-username-guid-history
		GetHistoryByUsernameAndGUIDV2(ctx context.Context, clientManagementID string, username string, guid string) (*AccountHistoryResponse, *resty.Response, error)

		// GetPasswordByUsernameAndGUIDV2 retrieves the current LAPS password for a specific username and guid on a target device.
		// Use when multiple accounts share the same username; guid disambiguates them.
		// Note: Once viewed, the password will be rotated based on rotation time settings.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-account-username-guid-password
		GetPasswordByUsernameAndGUIDV2(ctx context.Context, clientManagementID string, username string, guid string) (*CurrentPasswordResponse, *resty.Response, error)

		// GetFullHistoryByClientManagementIDV2 retrieves the complete history of all local admin passwords for all accounts
		// on a specific device, including both viewing and rotation history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-history
		GetFullHistoryByClientManagementIDV2(ctx context.Context, clientManagementID string) (*FullHistoryResponse, *resty.Response, error)

		// GetCapableAccountsByClientManagementIDV2 retrieves a list of all admin accounts that are LAPS capable for a specific device.
		// Capable accounts are returned in the AutoSetupAdminAccounts from QueryResponses.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-accounts
		GetCapableAccountsByClientManagementIDV2(ctx context.Context, clientManagementID string) (*CapableAccountsResponse, *resty.Response, error)

		// SetPasswordByClientManagementIDV2 sets LAPS passwords for all capable accounts on a device.
		// The passwords are provided as a list of username/password pairs.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-local-admin-password-clientmanagementid-set-password
		SetPasswordByClientManagementIDV2(ctx context.Context, clientManagementID string, passwordList *SetPasswordRequest) (*SetPasswordResponse, *resty.Response, error)
	}

	// Service handles communication with the local admin password methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-pending-rotations
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ LocalAdminPasswordServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Local Admin Password (LAPS) Operations
// -----------------------------------------------------------------------------

// GetPendingRotationsV2 retrieves a list of devices and usernames with pending LAPS rotations.
// URL: GET /api/v2/local-admin-password/pending-rotations
// https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-pending-rotations
func (s *Service) GetPendingRotationsV2(ctx context.Context) (*PendingRotationsResponse, *resty.Response, error) {
	var result PendingRotationsResponse

	endpoint := fmt.Sprintf("%s/pending-rotations", EndpointLocalAdminPasswordV2)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetSettingsV2 retrieves current Jamf Pro LAPS settings.
// URL: GET /api/v2/local-admin-password/settings
// https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-settings
func (s *Service) GetSettingsV2(ctx context.Context) (*SettingsResource, *resty.Response, error) {
	var result SettingsResource

	endpoint := fmt.Sprintf("%s/settings", EndpointLocalAdminPasswordV2)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateSettingsV2 updates the current Jamf Pro LAPS settings.
// URL: PUT /api/v2/local-admin-password/settings
// https://developer.jamf.com/jamf-pro/reference/put_v2-local-admin-password-settings
func (s *Service) UpdateSettingsV2(ctx context.Context, settings *SettingsResource) (*resty.Response, error) {
	if settings == nil {
		return nil, fmt.Errorf("settings is required")
	}

	endpoint := fmt.Sprintf("%s/settings", EndpointLocalAdminPasswordV2)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, settings, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetPasswordHistoryByClientManagementIDV2 retrieves the password view history for a specific username on a target device.
// History will include password, who viewed the password and when it was viewed.
// URL: GET /api/v2/local-admin-password/{clientManagementId}/account/{username}/audit
// https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-account-username-audit
func (s *Service) GetPasswordHistoryByClientManagementIDV2(ctx context.Context, clientManagementID string, username string) (*PasswordHistoryResponse, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}
	if username == "" {
		return nil, nil, fmt.Errorf("username is required")
	}

	endpoint := fmt.Sprintf("%s/%s/account/%s/audit", EndpointLocalAdminPasswordV2, clientManagementID, username)

	var result PasswordHistoryResponse

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetCurrentPasswordByClientManagementIDV2 retrieves the current LAPS password for a specific username on a target device.
// Note: Once viewed, the password will be rotated based on rotation time settings.
// URL: GET /api/v2/local-admin-password/{clientManagementId}/account/{username}/password
// https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-account-username-password
func (s *Service) GetCurrentPasswordByClientManagementIDV2(ctx context.Context, clientManagementID string, username string) (*CurrentPasswordResponse, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}
	if username == "" {
		return nil, nil, fmt.Errorf("username is required")
	}

	endpoint := fmt.Sprintf("%s/%s/account/%s/password", EndpointLocalAdminPasswordV2, clientManagementID, username)

	var result CurrentPasswordResponse

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetHistoryByUsernameV2 retrieves the LAPS history for a specific username on a target device.
// History includes date created, date last seen, expiration time, and rotational status.
// URL: GET /api/v2/local-admin-password/{clientManagementId}/account/{username}/history
// https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-account-username-history
func (s *Service) GetHistoryByUsernameV2(ctx context.Context, clientManagementID string, username string) (*AccountHistoryResponse, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}
	if username == "" {
		return nil, nil, fmt.Errorf("username is required")
	}

	endpoint := fmt.Sprintf("%s/%s/account/%s/history", EndpointLocalAdminPasswordV2, clientManagementID, username)

	var result AccountHistoryResponse

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetAuditByUsernameAndGUIDV2 retrieves the password view history for a specific username and guid on a target device.
// Use when multiple accounts share the same username; guid disambiguates them.
// URL: GET /api/v2/local-admin-password/{clientManagementId}/account/{username}/{guid}/audit
// https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-account-username-guid-audit
func (s *Service) GetAuditByUsernameAndGUIDV2(ctx context.Context, clientManagementID string, username string, guid string) (*PasswordHistoryResponse, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}
	if username == "" {
		return nil, nil, fmt.Errorf("username is required")
	}
	if guid == "" {
		return nil, nil, fmt.Errorf("guid is required")
	}

	endpoint := fmt.Sprintf("%s/%s/account/%s/%s/audit", EndpointLocalAdminPasswordV2, clientManagementID, username, guid)

	var result PasswordHistoryResponse

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetHistoryByUsernameAndGUIDV2 retrieves the LAPS history for a specific username and guid on a target device.
// Use when multiple accounts share the same username; guid disambiguates them.
// URL: GET /api/v2/local-admin-password/{clientManagementId}/account/{username}/{guid}/history
// https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-account-username-guid-history
func (s *Service) GetHistoryByUsernameAndGUIDV2(ctx context.Context, clientManagementID string, username string, guid string) (*AccountHistoryResponse, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}
	if username == "" {
		return nil, nil, fmt.Errorf("username is required")
	}
	if guid == "" {
		return nil, nil, fmt.Errorf("guid is required")
	}

	endpoint := fmt.Sprintf("%s/%s/account/%s/%s/history", EndpointLocalAdminPasswordV2, clientManagementID, username, guid)

	var result AccountHistoryResponse

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
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
func (s *Service) GetPasswordByUsernameAndGUIDV2(ctx context.Context, clientManagementID string, username string, guid string) (*CurrentPasswordResponse, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}
	if username == "" {
		return nil, nil, fmt.Errorf("username is required")
	}
	if guid == "" {
		return nil, nil, fmt.Errorf("guid is required")
	}

	endpoint := fmt.Sprintf("%s/%s/account/%s/%s/password", EndpointLocalAdminPasswordV2, clientManagementID, username, guid)

	var result CurrentPasswordResponse

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetFullHistoryByClientManagementIDV2 retrieves the complete history of all local admin passwords for all accounts
// on a specific device, including both viewing and rotation history.
// URL: GET /api/v2/local-admin-password/{clientManagementId}/history
// https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-history
func (s *Service) GetFullHistoryByClientManagementIDV2(ctx context.Context, clientManagementID string) (*FullHistoryResponse, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointLocalAdminPasswordV2, clientManagementID)

	var result FullHistoryResponse

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetCapableAccountsByClientManagementIDV2 retrieves a list of all admin accounts that are LAPS capable for a specific device.
// Capable accounts are returned in the AutoSetupAdminAccounts from QueryResponses.
// URL: GET /api/v2/local-admin-password/{clientManagementId}/accounts
// https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-clientmanagementid-accounts
func (s *Service) GetCapableAccountsByClientManagementIDV2(ctx context.Context, clientManagementID string) (*CapableAccountsResponse, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/accounts", EndpointLocalAdminPasswordV2, clientManagementID)

	var result CapableAccountsResponse

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// SetPasswordByClientManagementIDV2 sets LAPS passwords for all capable accounts on a device.
// The passwords are provided as a list of username/password pairs.
// URL: PUT /api/v2/local-admin-password/{clientManagementId}/set-password
// https://developer.jamf.com/jamf-pro/reference/put_v2-local-admin-password-clientmanagementid-set-password
func (s *Service) SetPasswordByClientManagementIDV2(ctx context.Context, clientManagementID string, passwordList *SetPasswordRequest) (*SetPasswordResponse, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}
	if passwordList == nil {
		return nil, nil, fmt.Errorf("passwordList is required")
	}

	endpoint := fmt.Sprintf("%s/%s/set-password", EndpointLocalAdminPasswordV2, clientManagementID)

	var result SetPasswordResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, passwordList, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
