package cache_settings

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// CacheSettingsServiceInterface defines the interface for cache settings operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cache-settings
	CacheSettingsServiceInterface interface {
		// GetV1 retrieves the current cache settings (Get Cache Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cache-settings
		GetV1(ctx context.Context) (*ResourceCacheSettings, *interfaces.Response, error)

		// UpdateV1 updates the cache settings (Update Cache Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-cache-settings
		UpdateV1(ctx context.Context, request *ResourceCacheSettings) (*ResourceCacheSettings, *interfaces.Response, error)
	}

	// Service handles communication with the cache settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cache-settings
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ CacheSettingsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetV1 retrieves the current cache settings.
// URL: GET /api/v1/cache-settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cache-settings
func (s *Service) GetV1(ctx context.Context) (*ResourceCacheSettings, *interfaces.Response, error) {
	var result ResourceCacheSettings
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Get(ctx, EndpointCacheSettingsV1, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UpdateV1 updates the cache settings.
// URL: PUT /api/v1/cache-settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-cache-settings
func (s *Service) UpdateV1(ctx context.Context, request *ResourceCacheSettings) (*ResourceCacheSettings, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	var result ResourceCacheSettings
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Put(ctx, EndpointCacheSettingsV1, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
