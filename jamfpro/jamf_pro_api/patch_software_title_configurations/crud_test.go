package patch_software_title_configurations

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/patch_software_title_configurations/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestListV2_APIError tests ListV2 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_ListV2_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListNoResponseErrorMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.ListV2(context.Background())

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestListV2 tests listing all patch software title configurations.
func TestUnit_PatchSoftwareTitleConfigurations_ListV2_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListMock()

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.ListV2(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, *result, 2)
	assert.Equal(t, "1", (*result)[0].ID)
	assert.Equal(t, "Google Chrome", (*result)[0].DisplayName)
	assert.Equal(t, "101", (*result)[0].SoftwareTitleID)
	assert.True(t, (*result)[0].UINotifications)
	assert.False(t, (*result)[0].EmailNotifications)
	assert.Equal(t, "2", (*result)[1].ID)
	assert.Equal(t, "Mozilla Firefox", (*result)[1].DisplayName)
	assert.Equal(t, "102", (*result)[1].SoftwareTitleID)
	assert.False(t, (*result)[1].UINotifications)
	assert.True(t, (*result)[1].EmailNotifications)
}

// TestGetByIDV2_APIError tests GetByIDV2 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_GetByIDV2_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetByIDNoResponseErrorMock("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetByIDV2(context.Background(), "1")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestGetByIDV2 tests retrieving a patch software title configuration by ID.
func TestUnit_PatchSoftwareTitleConfigurations_GetByIDV2_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetByIDMock("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetByIDV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Google Chrome", result.DisplayName)
	assert.Equal(t, "101", result.SoftwareTitleID)
	assert.Equal(t, "1", result.CategoryID)
	assert.Equal(t, "-1", result.SiteID)
	assert.True(t, result.UINotifications)
	assert.False(t, result.EmailNotifications)
	assert.Equal(t, "Google Chrome", result.SoftwareTitleName)
	assert.Equal(t, "GoogleChrome", result.SoftwareTitleNameID)
	assert.Equal(t, "Google Inc.", result.SoftwareTitlePublisher)
	assert.True(t, result.JamfOfficial)
	assert.Equal(t, "Jamf", result.PatchSourceName)
	assert.True(t, result.PatchSourceEnabled)
	assert.Len(t, result.ExtensionAttributes, 1)
	assert.True(t, result.ExtensionAttributes[0].Accepted)
	assert.Equal(t, "10", result.ExtensionAttributes[0].EAID)
	assert.Len(t, result.Packages, 2)
	assert.Equal(t, "200", result.Packages[0].PackageID)
	assert.Equal(t, "121.0.6167.85", result.Packages[0].Version)
	assert.Equal(t, "GoogleChrome-121.0.6167.85.pkg", result.Packages[0].DisplayName)
}

// TestGetByIDV2_EmptyID tests retrieving a patch software title configuration with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_GetByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetByIDV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetByNameV2_ListError tests GetByNameV2 when ListV2 returns error.
func TestUnit_PatchSoftwareTitleConfigurations_GetByNameV2_ListError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListNoResponseErrorMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetByNameV2(context.Background(), "Google Chrome")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestGetByNameV2 tests retrieving a patch software title configuration by display name.
func TestUnit_PatchSoftwareTitleConfigurations_GetByNameV2_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListMock()

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetByNameV2(context.Background(), "Google Chrome")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Google Chrome", result.DisplayName)
	assert.Equal(t, "101", result.SoftwareTitleID)
}

// TestGetByNameV2_NotFound tests retrieving a patch software title configuration by name when not found.
func TestUnit_PatchSoftwareTitleConfigurations_GetByNameV2_NotFound(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListMock()

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetByNameV2(context.Background(), "Nonexistent Config")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not found")
}

// TestGetByNameV2_EmptyName tests retrieving a patch software title configuration with empty name.
func TestUnit_PatchSoftwareTitleConfigurations_GetByNameV2_EmptyName(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetByNameV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "name is required")
}

// TestCreateV2_APIError tests CreateV2 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_CreateV2_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterCreateNoResponseErrorMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{
		DisplayName:     "New Config",
		SoftwareTitleID: "103",
	}
	result, resp, err := svc.CreateV2(context.Background(), config)

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestCreateV2 tests creating a new patch software title configuration.
func TestUnit_PatchSoftwareTitleConfigurations_CreateV2_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterCreateMock()

	svc := NewPatchSoftwareTitleConfigurations(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{
		DisplayName:     "New Patch Config",
		SoftwareTitleID: "103",
		CategoryID:      "1",
		UINotifications: true,
		ExtensionAttributes: []SubsetExtensionAttribute{
			{
				Accepted: true,
				EAID:     "11",
			},
		},
		Packages: []SubsetPackage{
			{
				PackageID:   "203",
				Version:     "1.0.0",
				DisplayName: "NewApp-1.0.0.pkg",
			},
		},
	}

	result, resp, err := svc.CreateV2(context.Background(), config)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "3", result.ID)
	assert.Contains(t, result.Href, "/api/v2/patch-software-title-configurations/3")
}

// TestCreateV2_NilConfig tests creating a patch software title configuration with nil config.
func TestUnit_PatchSoftwareTitleConfigurations_CreateV2_NilConfig(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.CreateV2(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "config is required")
}

// TestCreateV2_EmptyDisplayName tests creating a patch software title configuration with empty display name.
func TestUnit_PatchSoftwareTitleConfigurations_CreateV2_EmptyDisplayName(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{
		SoftwareTitleID: "103",
	}

	result, resp, err := svc.CreateV2(context.Background(), config)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "display name is required")
}

// TestCreateV2_EmptySoftwareTitleID tests creating a patch software title configuration with empty software title ID.
func TestUnit_PatchSoftwareTitleConfigurations_CreateV2_EmptySoftwareTitleID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{
		DisplayName: "New Config",
	}

	result, resp, err := svc.CreateV2(context.Background(), config)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "software title id is required")
}

// TestUpdateByIDV2_APIError tests UpdateByIDV2 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_UpdateByIDV2_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterUpdateByIDNoResponseErrorMock("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{DisplayName: "Updated", SoftwareTitleID: "101"}
	result, resp, err := svc.UpdateByIDV2(context.Background(), "1", config)

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestUpdateByIDV2 tests updating a patch software title configuration by ID.
func TestUnit_PatchSoftwareTitleConfigurations_UpdateByIDV2_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterUpdateByIDMock("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{
		DisplayName:        "Updated Google Chrome",
		SoftwareTitleID:    "101",
		UINotifications:    false,
		EmailNotifications: true,
	}

	result, resp, err := svc.UpdateByIDV2(context.Background(), "1", config)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Google Chrome", result.DisplayName)
}

// TestUpdateByIDV2_EmptyID tests updating a patch software title configuration with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_UpdateByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{
		DisplayName: "Updated Config",
	}

	result, resp, err := svc.UpdateByIDV2(context.Background(), "", config)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestUpdateByIDV2_NilConfig tests updating a patch software title configuration with nil config.
func TestUnit_PatchSoftwareTitleConfigurations_UpdateByIDV2_NilConfig(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.UpdateByIDV2(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "config is required")
}

// TestUpdateByNameV2_GetByNameError tests UpdateByNameV2 when GetByNameV2 fails.
func TestUnit_PatchSoftwareTitleConfigurations_UpdateByNameV2_GetByNameError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListNoResponseErrorMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{DisplayName: "Updated", SoftwareTitleID: "101"}
	result, resp, err := svc.UpdateByNameV2(context.Background(), "Google Chrome", config)

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestUpdateByNameV2 tests updating a patch software title configuration by name.
func TestUnit_PatchSoftwareTitleConfigurations_UpdateByNameV2_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListMock()
	mock.RegisterUpdateByIDMock("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{
		DisplayName:     "Updated Config",
		SoftwareTitleID: "101",
	}

	result, resp, err := svc.UpdateByNameV2(context.Background(), "Google Chrome", config)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
}

// TestUpdateByNameV2_EmptyName tests updating a patch software title configuration with empty name.
func TestUnit_PatchSoftwareTitleConfigurations_UpdateByNameV2_EmptyName(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{
		DisplayName: "Updated Config",
	}

	result, resp, err := svc.UpdateByNameV2(context.Background(), "", config)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "name is required")
}

// TestDeleteByIDV2_APIError tests DeleteByIDV2 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_DeleteByIDV2_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterDeleteByIDNoResponseErrorMock("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	resp, err := svc.DeleteByIDV2(context.Background(), "1")

	assert.Error(t, err)
	assert.NotNil(t, resp)
}

// TestDeleteByIDV2 tests deleting a patch software title configuration by ID.
func TestUnit_PatchSoftwareTitleConfigurations_DeleteByIDV2_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterDeleteByIDMock("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	resp, err := svc.DeleteByIDV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

// TestDeleteByIDV2_EmptyID tests deleting a patch software title configuration with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_DeleteByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	resp, err := svc.DeleteByIDV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

// TestDeleteByNameV2_GetByNameError tests DeleteByNameV2 when GetByNameV2 fails.
func TestUnit_PatchSoftwareTitleConfigurations_DeleteByNameV2_GetByNameError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListNoResponseErrorMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	resp, err := svc.DeleteByNameV2(context.Background(), "Google Chrome")

	assert.Error(t, err)
	assert.NotNil(t, resp)
}

// TestDeleteByNameV2 tests deleting a patch software title configuration by name.
func TestUnit_PatchSoftwareTitleConfigurations_DeleteByNameV2_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListMock()
	mock.RegisterDeleteByIDMock("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	resp, err := svc.DeleteByNameV2(context.Background(), "Google Chrome")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

// TestDeleteByNameV2_EmptyName tests deleting a patch software title configuration with empty name.
func TestUnit_PatchSoftwareTitleConfigurations_DeleteByNameV2_EmptyName(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	resp, err := svc.DeleteByNameV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "name is required")
}

// TestGetDashboardStatusByIDV2_APIError tests GetDashboardStatusByIDV2 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_GetDashboardStatusByIDV2_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetDashboardStatusByIDV2(context.Background(), "1")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "no response")
}

// TestGetDashboardStatusByIDV2 tests getting dashboard status.
func TestUnit_PatchSoftwareTitleConfigurations_GetDashboardStatusByIDV2_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetDashboardStatusMock("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetDashboardStatusByIDV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.True(t, result.OnDashboard)
}

// TestGetDashboardStatusByIDV2_EmptyID tests getting dashboard status with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_GetDashboardStatusByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetDashboardStatusByIDV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestAddToDashboardByIDV2_EmptyID tests adding to dashboard with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_AddToDashboardByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	resp, err := svc.AddToDashboardByIDV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

// TestAddToDashboardByIDV2_APIError tests AddToDashboardByIDV2 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_AddToDashboardByIDV2_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterAddToDashboardNoResponseErrorMock("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	resp, err := svc.AddToDashboardByIDV2(context.Background(), "1")

	assert.Error(t, err)
	assert.NotNil(t, resp)
}

// TestAddToDashboardByIDV2 tests adding to dashboard.
func TestUnit_PatchSoftwareTitleConfigurations_AddToDashboardByIDV2_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterAddToDashboardMock("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	resp, err := svc.AddToDashboardByIDV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

// TestRemoveFromDashboardByIDV2_EmptyID tests removing from dashboard with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_RemoveFromDashboardByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	resp, err := svc.RemoveFromDashboardByIDV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

// TestRemoveFromDashboardByIDV2_APIError tests RemoveFromDashboardByIDV2 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_RemoveFromDashboardByIDV2_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterRemoveFromDashboardNoResponseErrorMock("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	resp, err := svc.RemoveFromDashboardByIDV2(context.Background(), "1")

	assert.Error(t, err)
	assert.NotNil(t, resp)
}

// TestRemoveFromDashboardByIDV2 tests removing from dashboard.
func TestUnit_PatchSoftwareTitleConfigurations_RemoveFromDashboardByIDV2_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterRemoveFromDashboardMock("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	resp, err := svc.RemoveFromDashboardByIDV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

// TestGetDefinitionsByIDV2_EmptyID tests getting definitions with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_GetDefinitionsByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetDefinitionsByIDV2(context.Background(), "", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetDefinitionsByIDV2_APIError tests GetDefinitionsByIDV2 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_GetDefinitionsByIDV2_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetDefinitionsNoResponseErrorMock("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetDefinitionsByIDV2(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestGetDefinitionsByIDV2 tests getting definitions.
func TestUnit_PatchSoftwareTitleConfigurations_GetDefinitionsByIDV2_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetDefinitionsMock("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetDefinitionsByIDV2(context.Background(), "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "10.37.0", result.Results[0].Version)
	assert.Equal(t, "1", result.Results[0].AbsoluteOrderID)
}

// TestGetDependenciesByIDV2_EmptyID tests getting dependencies with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_GetDependenciesByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetDependenciesByIDV2(context.Background(), "", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetDependenciesByIDV2_APIError tests GetDependenciesByIDV2 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_GetDependenciesByIDV2_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetDependenciesNoResponseErrorMock("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetDependenciesByIDV2(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestGetDependenciesByIDV2 tests getting dependencies.
func TestUnit_PatchSoftwareTitleConfigurations_GetDependenciesByIDV2_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetDependenciesMock("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetDependenciesByIDV2(context.Background(), "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].SmartGroupID)
	assert.Equal(t, "Chrome Out of Date", result.Results[0].SmartGroupName)
}

// TestExportReportByIDV2_EmptyID tests exporting report with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_ExportReportByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	body, resp, err := svc.ExportReportByIDV2(context.Background(), "", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, body)
	assert.Contains(t, err.Error(), "id is required")
}

// TestExportReportByIDV2_APIError tests ExportReportByIDV2 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_ExportReportByIDV2_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterExportReportNoResponseErrorMock("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	body, resp, err := svc.ExportReportByIDV2(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, body)
}

// TestExportReportByIDV2 tests exporting report.
func TestUnit_PatchSoftwareTitleConfigurations_ExportReportByIDV2_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterExportReportMock("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	body, resp, err := svc.ExportReportByIDV2(context.Background(), "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, body)
	assert.Contains(t, string(body), "computerName")
}

// TestGetExtensionAttributesByIDV2_EmptyID tests getting extension attributes with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_GetExtensionAttributesByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetExtensionAttributesByIDV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetExtensionAttributesByIDV2_APIError tests GetExtensionAttributesByIDV2 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_GetExtensionAttributesByIDV2_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetExtensionAttributesNoResponseErrorMock("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetExtensionAttributesByIDV2(context.Background(), "1")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestGetExtensionAttributesByIDV2 tests getting extension attributes.
func TestUnit_PatchSoftwareTitleConfigurations_GetExtensionAttributesByIDV2_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetExtensionAttributesMock("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetExtensionAttributesByIDV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result, 1)
	assert.Equal(t, "google-chrome-ea", result[0].EAID)
	assert.True(t, result[0].Accepted)
}

// TestGetPatchReportByIDV2_EmptyID tests getting patch report with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_GetPatchReportByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetPatchReportByIDV2(context.Background(), "", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetPatchReportByIDV2_APIError tests GetPatchReportByIDV2 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_GetPatchReportByIDV2_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetPatchReportNoResponseErrorMock("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetPatchReportByIDV2(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestGetPatchReportByIDV2 tests getting patch report.
func TestUnit_PatchSoftwareTitleConfigurations_GetPatchReportByIDV2_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetPatchReportMock("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetPatchReportByIDV2(context.Background(), "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "MacBook", result.Results[0].ComputerName)
	assert.Equal(t, "10.1", result.Results[0].Version)
}

// TestGetPatchSummaryByIDV2_EmptyID tests getting patch summary with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_GetPatchSummaryByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetPatchSummaryByIDV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetPatchSummaryByIDV2_APIError tests GetPatchSummaryByIDV2 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_GetPatchSummaryByIDV2_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetPatchSummaryNoResponseErrorMock("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetPatchSummaryByIDV2(context.Background(), "1")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestGetPatchSummaryByIDV2 tests getting patch summary.
func TestUnit_PatchSoftwareTitleConfigurations_GetPatchSummaryByIDV2_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetPatchSummaryMock("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetPatchSummaryByIDV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.SoftwareTitleID)
	assert.Equal(t, "Patch title", result.Title)
	assert.Equal(t, 3, result.UpToDate)
	assert.Equal(t, 6, result.OutOfDate)
}

// TestGetHistoryByIDV2_EmptyID tests getting history with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_GetHistoryByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetHistoryByIDV2(context.Background(), "", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetHistoryByIDV2_APIError tests GetHistoryByIDV2 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_GetHistoryByIDV2_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetHistoryNoResponseErrorMock("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetHistoryByIDV2(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestGetHistoryByIDV2 tests getting history.
func TestUnit_PatchSoftwareTitleConfigurations_GetHistoryByIDV2_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetHistoryMock("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetHistoryByIDV2(context.Background(), "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "Sso settings update", result.Results[0].Note)
}

// TestAddHistoryNoteByIDV2_APIError tests AddHistoryNoteByIDV2 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_AddHistoryNoteByIDV2_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterAddHistoryNoteNoResponseErrorMock("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.AddHistoryNoteByIDV2(context.Background(), "1", &RequestAddHistoryNote{Note: "Test"})

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestAddHistoryNoteByIDV2 tests adding history note.
func TestUnit_PatchSoftwareTitleConfigurations_AddHistoryNoteByIDV2_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterAddHistoryNoteMock("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.AddHistoryNoteByIDV2(context.Background(), "1", &RequestAddHistoryNote{Note: "Test note"})

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
	assert.Contains(t, result.Href, "/api/v1/resource/1")
}

// TestAddHistoryNoteByIDV2_EmptyID tests adding history note with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_AddHistoryNoteByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.AddHistoryNoteByIDV2(context.Background(), "", &RequestAddHistoryNote{Note: "Test"})

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestAddHistoryNoteByIDV2_NilRequest tests adding history note with nil request.
func TestUnit_PatchSoftwareTitleConfigurations_AddHistoryNoteByIDV2_NilRequest(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.AddHistoryNoteByIDV2(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

// TestAddHistoryNoteByIDV2_EmptyNote tests adding history note with empty note.
func TestUnit_PatchSoftwareTitleConfigurations_AddHistoryNoteByIDV2_EmptyNote(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.AddHistoryNoteByIDV2(context.Background(), "1", &RequestAddHistoryNote{Note: ""})

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "note is required")
}

// TestGetPatchVersionsByIDV2_EmptyID tests getting patch versions with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_GetPatchVersionsByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetPatchVersionsByIDV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetPatchVersionsByIDV2_APIError tests GetPatchVersionsByIDV2 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_GetPatchVersionsByIDV2_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetPatchVersionsNoResponseErrorMock("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetPatchVersionsByIDV2(context.Background(), "1")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestGetPatchVersionsByIDV2 tests getting patch versions.
func TestUnit_PatchSoftwareTitleConfigurations_GetPatchVersionsByIDV2_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetPatchVersionsMock("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetPatchVersionsByIDV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result, 1)
	assert.Equal(t, "1", result[0].AbsoluteOrderID)
	assert.Equal(t, "3", result[0].Version)
	assert.Equal(t, 1, result[0].OnVersion)
}
