package directory_cache

import (
	"context"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/directory_cache/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnit_DirectoryCache_GetSyncStatusV1(t *testing.T) {
	mock := mocks.NewDirectoryCacheMock()
	mock.Register("GET", "/api/v1/directory-cache/sync-status", 200, "validate_status.json")
	svc := NewDirectoryCache(mock)
	result, resp, err := svc.GetSyncStatusV1(context.Background())
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode())
	require.Equal(t, "NEVER_RAN", result.State)
	require.Nil(t, result.DirectoryServiceID)
	require.Nil(t, result.StartTime)
	mock.RegisterError("GET", "/api/v1/directory-cache/sync-status", 403, "", "Forbidden")
	result, resp, err = svc.GetSyncStatusV1(context.Background())
	require.Error(t, err)
	require.Nil(t, result)
	require.Equal(t, 403, resp.StatusCode())
}
func TestUnit_DirectoryCache_SyncV1(t *testing.T) {
	for _, kind := range []string{"group", "user"} {
		t.Run(kind, func(t *testing.T) {
			mock := mocks.NewDirectoryCacheMock()
			mock.Register("POST", "/api/v1/directory-cache/sync/"+kind, 202, "")
			svc := NewDirectoryCache(mock)
			call := svc.SyncGroupV1
			key := "group-name"
			if kind == "user" {
				call = svc.SyncUserV1
				key = "username"
			}
			resp, err := call(context.Background(), "1", "a & b")
			require.NoError(t, err)
			require.Equal(t, 202, resp.StatusCode())
			require.Equal(t, "1", mock.LastRSQLQuery["directory-service-id"])
			require.Equal(t, "a & b", mock.LastRSQLQuery[key])
			for _, args := range [][2]string{{"", "name"}, {"1", ""}} {
				resp, err := call(context.Background(), args[0], args[1])
				require.Error(t, err)
				require.Nil(t, resp)
			}
			mock.RegisterError("POST", "/api/v1/directory-cache/sync/"+kind, 409, "", "Sync in progress")
			resp, err = call(context.Background(), "1", "name")
			require.Error(t, err)
			require.Equal(t, 409, resp.StatusCode())
		})
	}
}
