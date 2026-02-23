package oauth2_session_tokens

// SessionTokenResponse represents OAuth2 session token information.
type SessionTokenResponse struct {
	AccessToken string `json:"accessToken"`
	IDToken     string `json:"idToken"`
}
