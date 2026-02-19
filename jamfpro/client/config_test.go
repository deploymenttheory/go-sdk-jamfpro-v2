package client

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadAuthConfigFromFile(t *testing.T) {
	dir := t.TempDir()

	validPath := filepath.Join(dir, "config.json")
	require.NoError(t, os.WriteFile(validPath, []byte(`{
		"instance_domain": "https://example.jamfcloud.com",
		"auth_method": "oauth2",
		"client_id": "cid",
		"client_secret": "secret",
		"token_refresh_buffer_period_seconds": 120,
		"hide_sensitive_data": true
	}`), 0644))

	cfg, err := LoadAuthConfigFromFile(validPath)
	require.NoError(t, err)
	require.NotNil(t, cfg)
	assert.Equal(t, "https://example.jamfcloud.com", cfg.InstanceDomain)
	assert.Equal(t, "oauth2", cfg.AuthMethod)
	assert.Equal(t, "cid", cfg.ClientID)
	assert.Equal(t, "secret", cfg.ClientSecret)
	assert.Equal(t, 120*time.Second, cfg.TokenRefreshBufferPeriod)
	assert.True(t, cfg.HideSensitiveData)
}

func TestLoadAuthConfigFromFile_DefaultBuffer(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")
	require.NoError(t, os.WriteFile(path, []byte(`{"instance_domain":"https://x.com","auth_method":"basic","basic_auth_username":"u","basic_auth_password":"p"}`), 0644))

	cfg, err := LoadAuthConfigFromFile(path)
	require.NoError(t, err)
	require.NotNil(t, cfg)
	assert.Equal(t, 5*time.Minute, cfg.TokenRefreshBufferPeriod)
}

func TestLoadAuthConfigFromFile_MissingFile(t *testing.T) {
	_, err := LoadAuthConfigFromFile(filepath.Join(t.TempDir(), "nonexistent.json"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "open")
}

func TestLoadAuthConfigFromFile_InvalidJSON(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.json")
	require.NoError(t, os.WriteFile(path, []byte("not json"), 0644))

	_, err := LoadAuthConfigFromFile(path)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "parse")
}
