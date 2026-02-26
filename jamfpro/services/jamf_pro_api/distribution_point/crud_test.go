package distribution_point

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/distribution_point/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.DistributionPointMock) {
	t.Helper()
	// Ensure working directory is the distribution_point package so mocks can load fixtures
	dir, err := os.Getwd()
	require.NoError(t, err)
	if filepath.Base(dir) != "distribution_point" {
		// When running from module root, chdir to package
		pkgDir := filepath.Join(dir, "jamfpro", "services", "jamf_pro_api", "distribution_point")
		if _, err := os.Stat(pkgDir); err == nil {
			_ = os.Chdir(pkgDir)
			t.Cleanup(func() { _ = os.Chdir(dir) })
		}
	}
	mock := mocks.NewDistributionPointMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnit_DistributionPoint_ListV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Test DP", result.Results[0].Name)
}

func TestUnit_DistributionPoint_CreateV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestDistributionPoint{
		Name:                  "Test DP",
		ServerName:            "dp.example.com",
		FileSharingConnectionType: "SMB",
		ShareName:             "JamfShare",
		Port:                  445,
	}
	result, resp, err := svc.CreateV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, []int{200, 201}, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
}

func TestUnit_DistributionPoint_CreateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.CreateV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_DistributionPoint_DeleteMultipleV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteMultipleV1(context.Background(), []string{"1", "2"})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnit_DistributionPoint_DeleteMultipleV1_EmptyIDs(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteMultipleV1(context.Background(), []string{})
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "at least one ID is required")
}

func TestUnit_DistributionPoint_GetByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Test DP", result.Name)
}

func TestUnit_DistributionPoint_GetByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "distribution point ID is required")
}

func TestUnit_DistributionPoint_GetByIDV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()
	result, resp, err := svc.GetByIDV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestUnit_DistributionPoint_UpdateByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestDistributionPoint{
		Name:                      "Test DP Updated",
		ServerName:                "dp.example.com",
		FileSharingConnectionType: "SMB",
		ShareName:                 "JamfShare",
		Port:                      445,
	}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
}

func TestUnit_DistributionPoint_UpdateByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestDistributionPoint{Name: "x", ServerName: "y", FileSharingConnectionType: "NONE"}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "distribution point ID is required")
}

func TestUnit_DistributionPoint_UpdateByIDV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_DistributionPoint_DeleteByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnit_DistributionPoint_DeleteByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "distribution point ID is required")
}

func TestUnit_DistributionPoint_PatchByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestDistributionPoint{
		Name:                      "Test DP Patched",
		ServerName:                "dp.example.com",
		FileSharingConnectionType: "SMB",
		ShareName:                 "JamfShare",
		Port:                      445,
	}
	result, resp, err := svc.PatchByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
}

func TestUnit_DistributionPoint_PatchByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestDistributionPoint{Name: "x", ServerName: "y", FileSharingConnectionType: "NONE"}
	result, resp, err := svc.PatchByIDV1(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "distribution point ID is required")
}

func TestUnit_DistributionPoint_GetHistoryByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetHistoryByIDV1(context.Background(), "1", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "admin", result.Results[0].Username)
}

func TestUnit_DistributionPoint_GetHistoryByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetHistoryByIDV1(context.Background(), "", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "distribution point ID is required")
}

func TestUnit_DistributionPoint_CreateHistoryNoteV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.CreateHistoryNoteV1(context.Background(), "1", "Test note")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, []int{200, 201}, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Test note", result.Note)
}

func TestUnit_DistributionPoint_CreateHistoryNoteV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.CreateHistoryNoteV1(context.Background(), "", "note")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "distribution point ID is required")
}

func TestUnit_DistributionPoint_CreateHistoryNoteV1_EmptyNote(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.CreateHistoryNoteV1(context.Background(), "1", "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "note is required")
}
