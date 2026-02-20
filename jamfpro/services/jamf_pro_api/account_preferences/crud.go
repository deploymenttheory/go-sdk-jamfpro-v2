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
	// Jamf Pro API docs (undocumented): /api/v2/account-preferences
	AccountPreferencesServiceInterface interface {
		// GetAccountPreferencesV2 returns the current account preferences (Get Account Preferences).
		//
		// Jamf Pro API docs (undocumented): /api/v2/account-preferences
		GetAccountPreferencesV2(ctx context.Context) (*ResourceAccountPreferencesV2, *interfaces.Response, error)

		// UpdateAccountPreferencesV2 updates account preferences (Update Account Preferences / PATCH).
		//
		// Jamf Pro API docs (undocumented): /api/v2/account-preferences
		UpdateAccountPreferencesV2(ctx context.Context, request *ResourceAccountPreferencesV2) (*ResourceAccountPreferencesV2, *interfaces.Response, error)
	}

	// Service handles communication with the account preferences-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs (undocumented): /api/v2/account-preferences
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

// GetAccountPreferencesV2 returns the current account preferences.
// URL: GET /api/v2/account-preferences
func (s *Service) GetAccountPreferencesV2(ctx context.Context) (*ResourceAccountPreferencesV2, *interfaces.Response, error) {
	var result ResourceAccountPreferencesV2

	endpoint := EndpointAccountPreferencesV2

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

// UpdateAccountPreferencesV2 updates account preferences.
// URL: PATCH /api/v2/account-preferences
func (s *Service) UpdateAccountPreferencesV2(ctx context.Context, request *ResourceAccountPreferencesV2) (*ResourceAccountPreferencesV2, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceAccountPreferencesV2

	endpoint := EndpointAccountPreferencesV2

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
