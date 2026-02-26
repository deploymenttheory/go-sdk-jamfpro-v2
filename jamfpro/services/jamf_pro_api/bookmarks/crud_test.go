package bookmarks

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/bookmarks/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.BookmarksMock) {
	t.Helper()
	mock := mocks.NewBookmarksMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnit_Bookmarks_ListV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	require.Equal(t, "1", result.Results[0].ID)
	require.Equal(t, "Jamf", result.Results[0].Name)
}

func TestUnit_Bookmarks_GetByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "1", result.ID)
}

func TestUnit_Bookmarks_CreateV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	bm := &ResourceBookmark{Name: "Test", URL: "https://example.com", SiteID: "-1", IconID: "1"}
	result, resp, err := svc.CreateV1(context.Background(), bm)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Contains(t, []int{200, 201}, resp.StatusCode)
	require.Equal(t, "2", result.ID)
}

func TestUnit_Bookmarks_DeleteByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 204, resp.StatusCode)
}

func TestUnit_Bookmarks_UpdateByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	bm := &ResourceBookmark{Name: "Updated", URL: "https://updated.example.com", SiteID: "-1", IconID: "1"}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", bm)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "1", result.ID)
}

func TestUnit_Bookmarks_UpdateByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	bm := &ResourceBookmark{Name: "Updated", URL: "https://updated.example.com", SiteID: "-1", IconID: "1"}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "", bm)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "id is required")
}

func TestUnit_Bookmarks_UpdateByIDV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "request is required")
}

func TestUnit_Bookmarks_GetByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetByIDV1(context.Background(), "")
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "id is required")
}

func TestUnit_Bookmarks_CreateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.CreateV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "request is required")
}

func TestUnit_Bookmarks_DeleteByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteByIDV1(context.Background(), "")
	require.Error(t, err)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "id is required")
}
