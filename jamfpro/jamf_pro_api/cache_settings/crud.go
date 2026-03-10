package cache_settings

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the cache settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cache-settings
	CacheSettings struct {
		client client.Client
	}
)

func NewCacheSettings(client client.Client) *CacheSettings {
	return &CacheSettings{client: client}
}

// GetV1 retrieves the current cache settings.
// URL: GET /api/v1/cache-settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cache-settings
func (s *CacheSettings) GetV1(ctx context.Context) (*ResourceCacheSettings, *resty.Response, error) {
	var result ResourceCacheSettings
	endpoint := constants.EndpointJamfProCacheSettingsV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)

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
	endpoint := constants.EndpointJamfProCacheSettingsV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Put(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
