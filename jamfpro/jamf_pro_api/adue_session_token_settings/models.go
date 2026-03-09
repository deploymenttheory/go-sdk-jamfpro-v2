package adue_session_token_settings

// ResourceADUETokenSettings is the account-driven user enrollment session token settings resource (get/update).
type ResourceADUETokenSettings struct {
	Enabled                   bool `json:"enabled"`
	ExpirationIntervalDays    int  `json:"expirationIntervalDays,omitempty"`
	ExpirationIntervalSeconds int  `json:"expirationIntervalSeconds,omitempty"`
}
