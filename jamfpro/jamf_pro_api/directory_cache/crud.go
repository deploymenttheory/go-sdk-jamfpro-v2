package directory_cache

import (
	"context"
	"fmt"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

// DirectoryCache provides directory sync status and targeted refresh operations.
type DirectoryCache struct{ client client.Client }

func NewDirectoryCache(client client.Client) *DirectoryCache { return &DirectoryCache{client: client} }

// GetSyncStatusV1 retrieves the current directory cache synchronization state.
func (s *DirectoryCache) GetSyncStatusV1(ctx context.Context) (*ResourceSyncStatus, *resty.Response, error) {
	var result ResourceSyncStatus
	endpoint := constants.EndpointJamfProDirectoryCacheV1 + "/sync-status"
	resp, err := s.client.NewRequest(ctx).SetHeader("Accept", constants.ApplicationJSON).SetResult(&result).Get(endpoint)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// SyncGroupV1 requests a targeted directory cache refresh.
func (s *DirectoryCache) SyncGroupV1(ctx context.Context, directoryServiceID, groupName string) (*resty.Response, error) {
	if directoryServiceID == "" {
		return nil, fmt.Errorf("directoryServiceID is required")
	}
	if groupName == "" {
		return nil, fmt.Errorf("groupName is required")
	}
	endpoint := constants.EndpointJamfProDirectoryCacheV1 + "/sync/group"
	resp, err := s.client.NewRequest(ctx).SetQueryParams(map[string]string{"directory-service-id": directoryServiceID, "group-name": groupName}).Post(endpoint)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// SyncUserV1 requests a targeted directory cache refresh.
func (s *DirectoryCache) SyncUserV1(ctx context.Context, directoryServiceID, username string) (*resty.Response, error) {
	if directoryServiceID == "" {
		return nil, fmt.Errorf("directoryServiceID is required")
	}
	if username == "" {
		return nil, fmt.Errorf("username is required")
	}
	endpoint := constants.EndpointJamfProDirectoryCacheV1 + "/sync/user"
	resp, err := s.client.NewRequest(ctx).SetQueryParams(map[string]string{"directory-service-id": directoryServiceID, "username": username}).Post(endpoint)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
