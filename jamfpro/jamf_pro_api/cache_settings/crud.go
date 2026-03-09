package cache_settings

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// CacheSettingsServiceInterface defines the interface for cache settings operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cache-settings
	CacheSettingsServiceInterface interface {
		// GetV1 retrieves the current cache settings (Get Cache Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cache-settings
		GetV1(ctx context.Context) (*ResourceCacheSettings, *resty.Response, error)

		// UpdateV1 updates the cache settings (Update Cache Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-cache-settings
		UpdateV1(ctx context.Context, request *ResourceCacheSettings) (*ResourceCacheSettings, *resty.Response, error)
	}

	// Service handles communication with the cache settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cache-settings
	CacheSettings struct {
		client transport.HTTPClient
	}
)

var _ CacheSettingsServiceInterface = (*CacheSettings)(nil)

func NewCacheSettings(client transport.HTTPClient) *CacheSettings {
	return &CacheSettings{client: client}
}

// GetV1 retrieves the current cache settings.
// URL: GET /api/v1/cache-settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cache-settings
func (s *CacheSettings) GetV1(ctx context.Context) (*ResourceCacheSettings, *resty.Response, error) {
	var result ResourceCacheSettings
	endpoint := EndpointCacheSettingsV1
	headers := map[string]string{"Accept": mime.ApplicationJSON}
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UpdateV1 updates the cache settings.
// URL: PUT /api/v1/cache-settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-cache-settings
func (s *CacheSettings) UpdateV1(ctx context.Context, request *ResourceCacheSettings) (*ResourceCacheSettings, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	var result ResourceCacheSettings
	endpoint := EndpointCacheSettingsV1
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
