package file_share_distribution_points_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/file_share_distribution_points"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/file_share_distribution_points/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_FileShareDistributionPoints_List(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	mockClient.RegisterListFileShareDistributionPointsMock()
	svc := file_share_distribution_points.NewService(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 2, resp.Size)
	assert.Len(t, resp.Results, 2)
	assert.Equal(t, "Main File Share DP", resp.Results[0].Name)
	assert.Equal(t, "Secondary File Share DP", resp.Results[1].Name)
}

func TestUnit_FileShareDistributionPoints_GetByID(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	mockClient.RegisterGetFileShareDistributionPointByIDMock()
	svc := file_share_distribution_points.NewService(mockClient)

	resp, _, err := svc.GetByID(context.Background(), 1)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "Main File Share DP", resp.Name)
	assert.True(t, resp.IsMaster)
	assert.Equal(t, "/path/to/share", resp.LocalPath)
	assert.Equal(t, "SMB", resp.ConnectionType)
	assert.Equal(t, "JamfShare", resp.ShareName)
}

func TestUnit_FileShareDistributionPoints_GetByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	svc := file_share_distribution_points.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "distribution point ID must be a positive integer")
}

func TestUnit_FileShareDistributionPoints_GetByName(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	mockClient.RegisterGetFileShareDistributionPointByNameMock()
	svc := file_share_distribution_points.NewService(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "Main File Share DP")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "Main File Share DP", resp.Name)
}

func TestUnit_FileShareDistributionPoints_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	svc := file_share_distribution_points.NewService(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "distribution point name cannot be empty")
}

func TestUnit_FileShareDistributionPoints_Create(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	mockClient.RegisterCreateFileShareDistributionPointMock()
	svc := file_share_distribution_points.NewService(mockClient)

	req := &file_share_distribution_points.RequestFileShareDistributionPoint{
		Name:                  "Test DP",
		IsMaster:              true,
		LocalPath:             "/path/to/share",
		ConnectionType:        "SMB",
		ShareName:             "JamfShare",
		HTTPDownloadsEnabled:  true,
	}

	resp, _, err := svc.Create(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.ID)
}

func TestUnit_FileShareDistributionPoints_Create_NilRequest(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	svc := file_share_distribution_points.NewService(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_FileShareDistributionPoints_Create_EmptyName(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	svc := file_share_distribution_points.NewService(mockClient)

	req := &file_share_distribution_points.RequestFileShareDistributionPoint{
		Name: "",
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "distribution point name is required")
}

func TestUnit_FileShareDistributionPoints_UpdateByID(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	mockClient.RegisterUpdateFileShareDistributionPointByIDMock()
	svc := file_share_distribution_points.NewService(mockClient)

	req := &file_share_distribution_points.RequestFileShareDistributionPoint{
		Name:     "Updated DP",
		IsMaster: false,
	}

	resp, _, err := svc.UpdateByID(context.Background(), 1, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_FileShareDistributionPoints_UpdateByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	svc := file_share_distribution_points.NewService(mockClient)

	req := &file_share_distribution_points.RequestFileShareDistributionPoint{
		Name: "Test",
	}

	_, _, err := svc.UpdateByID(context.Background(), 0, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "distribution point ID must be a positive integer")
}

func TestUnit_FileShareDistributionPoints_UpdateByName(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	mockClient.RegisterUpdateFileShareDistributionPointByNameMock()
	svc := file_share_distribution_points.NewService(mockClient)

	req := &file_share_distribution_points.RequestFileShareDistributionPoint{
		Name:     "Updated DP",
		IsMaster: false,
	}

	resp, _, err := svc.UpdateByName(context.Background(), "Main File Share DP", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_FileShareDistributionPoints_UpdateByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	svc := file_share_distribution_points.NewService(mockClient)

	req := &file_share_distribution_points.RequestFileShareDistributionPoint{
		Name: "Test",
	}

	_, _, err := svc.UpdateByName(context.Background(), "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "distribution point name cannot be empty")
}

func TestUnit_FileShareDistributionPoints_DeleteByID(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	mockClient.RegisterDeleteFileShareDistributionPointByIDMock()
	svc := file_share_distribution_points.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 1)

	require.NoError(t, err)
}

func TestUnit_FileShareDistributionPoints_DeleteByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	svc := file_share_distribution_points.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "distribution point ID must be a positive integer")
}

func TestUnit_FileShareDistributionPoints_DeleteByName(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	mockClient.RegisterDeleteFileShareDistributionPointByNameMock()
	svc := file_share_distribution_points.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "Main File Share DP")

	require.NoError(t, err)
}

func TestUnit_FileShareDistributionPoints_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	svc := file_share_distribution_points.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "distribution point name cannot be empty")
}

func TestUnit_FileShareDistributionPoints_NotFound(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := file_share_distribution_points.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 999)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}

func TestUnit_FileShareDistributionPoints_Conflict(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	mockClient.RegisterConflictErrorMock()
	svc := file_share_distribution_points.NewService(mockClient)

	req := &file_share_distribution_points.RequestFileShareDistributionPoint{
		Name:     "Duplicate DP",
		IsMaster: false,
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "distribution point with that name already exists")
}

func TestUnit_FileShareDistributionPoints_List_Error(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	svc := file_share_distribution_points.NewService(mockClient)
	_, _, err := svc.List(context.Background())
	require.Error(t, err)
}

func TestUnit_FileShareDistributionPoints_GetByName_Error(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	svc := file_share_distribution_points.NewService(mockClient)
	_, _, err := svc.GetByName(context.Background(), "Main File Share DP")
	require.Error(t, err)
}

func TestUnit_FileShareDistributionPoints_UpdateByID_NilRequest(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	svc := file_share_distribution_points.NewService(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_FileShareDistributionPoints_UpdateByID_EmptyName(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	svc := file_share_distribution_points.NewService(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), 1, &file_share_distribution_points.RequestFileShareDistributionPoint{Name: ""})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "distribution point name is required")
}

func TestUnit_FileShareDistributionPoints_UpdateByID_Error(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	svc := file_share_distribution_points.NewService(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), 1, &file_share_distribution_points.RequestFileShareDistributionPoint{Name: "Test"})
	require.Error(t, err)
}

func TestUnit_FileShareDistributionPoints_UpdateByName_NilRequest(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	svc := file_share_distribution_points.NewService(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "Main File Share DP", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_FileShareDistributionPoints_UpdateByName_EmptyReqName(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	svc := file_share_distribution_points.NewService(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "Main File Share DP", &file_share_distribution_points.RequestFileShareDistributionPoint{Name: ""})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "distribution point name is required in request")
}

func TestUnit_FileShareDistributionPoints_UpdateByName_Error(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	svc := file_share_distribution_points.NewService(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "Main File Share DP", &file_share_distribution_points.RequestFileShareDistributionPoint{Name: "Updated"})
	require.Error(t, err)
}

func TestUnit_FileShareDistributionPoints_DeleteByID_Error(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	svc := file_share_distribution_points.NewService(mockClient)
	_, err := svc.DeleteByID(context.Background(), 1)
	require.Error(t, err)
}

func TestUnit_FileShareDistributionPoints_DeleteByName_Error(t *testing.T) {
	mockClient := mocks.NewFileShareDistributionPointsMock()
	svc := file_share_distribution_points.NewService(mockClient)
	_, err := svc.DeleteByName(context.Background(), "Main File Share DP")
	require.Error(t, err)
}
