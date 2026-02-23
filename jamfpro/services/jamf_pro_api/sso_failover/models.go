package sso_failover

// FailoverSettings represents SSO failover configuration.
type FailoverSettings struct {
	FailoverURL    string `json:"failoverUrl"`
	GenerationTime int64  `json:"generationTime"` // Unix timestamp in milliseconds
}
