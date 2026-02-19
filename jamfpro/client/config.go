package client

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

// authConfigFile is the JSON shape for LoadAuthConfigFromFile.
type authConfigFile struct {
	InstanceDomain            string `json:"instance_domain"`
	AuthMethod                string `json:"auth_method"`
	ClientID                  string `json:"client_id"`
	ClientSecret              string `json:"client_secret"`
	Username                  string `json:"basic_auth_username"`
	Password                  string `json:"basic_auth_password"`
	TokenRefreshBufferSeconds int    `json:"token_refresh_buffer_period_seconds"`
	HideSensitiveData         bool   `json:"hide_sensitive_data"`
}

// LoadAuthConfigFromFile loads AuthConfig from a JSON file.
// Expected keys: instance_domain, auth_method; for oauth2: client_id, client_secret; for basic: basic_auth_username, basic_auth_password.
// Optional: token_refresh_buffer_period_seconds (default 300).
func LoadAuthConfigFromFile(path string) (*AuthConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open config file: %w", err)
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}
	var c authConfigFile
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, fmt.Errorf("parse config file: %w", err)
	}
	buffer := time.Duration(c.TokenRefreshBufferSeconds) * time.Second
	if buffer == 0 {
		buffer = 5 * time.Minute
	}
	return &AuthConfig{
		InstanceDomain:           c.InstanceDomain,
		AuthMethod:               c.AuthMethod,
		ClientID:                 c.ClientID,
		ClientSecret:             c.ClientSecret,
		Username:                 c.Username,
		Password:                 c.Password,
		TokenRefreshBufferPeriod: buffer,
		HideSensitiveData:        c.HideSensitiveData,
	}, nil
}
