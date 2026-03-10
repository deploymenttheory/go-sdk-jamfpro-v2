package oauth2_session_tokens

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the OAuth2 session token-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-oauth2-session-tokens
	Oauth2SessionTokens struct {
		client client.Client
	}
)

func NewOauth2SessionTokens(client client.Client) *Oauth2SessionTokens {
	return &Oauth2SessionTokens{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - OAuth2 Session Token Operations
// -----------------------------------------------------------------------------

// GetV1 retrieves OAuth2 session tokens.
// URL: GET /api/v1/oauth2/session-tokens
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-oauth2-session-tokens
func (s *Oauth2SessionTokens) GetV1(ctx context.Context) (*SessionTokenResponse, *resty.Response, error) {
	var result SessionTokenResponse

	endpoint := constants.EndpointJamfProOAuth2SessionTokensV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
