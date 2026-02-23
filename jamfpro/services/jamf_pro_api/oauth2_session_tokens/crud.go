package oauth2_session_tokens

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ServiceInterface defines the interface for OAuth2 session token operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-oauth2-session-tokens
	ServiceInterface interface {
		// GetV1 retrieves OAuth2 session tokens (Get OAuth2 Session Tokens).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-oauth2-session-tokens
		GetV1(ctx context.Context) (*SessionTokenResponse, *interfaces.Response, error)
	}

	// Service handles communication with the OAuth2 session token-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-oauth2-session-tokens
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - OAuth2 Session Token Operations
// -----------------------------------------------------------------------------

// GetV1 retrieves OAuth2 session tokens.
// URL: GET /api/v1/oauth2/session-tokens
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-oauth2-session-tokens
func (s *Service) GetV1(ctx context.Context) (*SessionTokenResponse, *interfaces.Response, error) {
	var result SessionTokenResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, EndpointOAuth2SessionTokensV1, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
