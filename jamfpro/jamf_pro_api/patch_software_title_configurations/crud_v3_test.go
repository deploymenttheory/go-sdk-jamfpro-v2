package patch_software_title_configurations

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/patch_software_title_configurations/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestListV3_APIError tests ListV3 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_ListV3_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListNoResponseErrorMockV3()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.ListV3(context.Background())

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestListV3 tests listing all patch software title configurations.
func TestUnit_PatchSoftwareTitleConfigurations_ListV3_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListMockV3()

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.ListV3(context.Background())

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

// TestGetByIDV3_APIError tests GetByIDV3 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_GetByIDV3_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetByIDNoResponseErrorMockV3("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetByIDV3(context.Background(), "1")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestGetByIDV3 tests retrieving a patch software title configuration by ID.
func TestUnit_PatchSoftwareTitleConfigurations_GetByIDV3_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetByIDMockV3("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetByIDV3(context.Background(), "1")

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

// TestGetByIDV3_EmptyID tests retrieving a patch software title configuration with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_GetByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetByIDV3(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetByNameV3_ListError tests GetByNameV3 when ListV3 returns error.
func TestUnit_PatchSoftwareTitleConfigurations_GetByNameV3_ListError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListNoResponseErrorMockV3()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetByNameV3(context.Background(), "Google Chrome")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestGetByNameV3 tests retrieving a patch software title configuration by display name.
func TestUnit_PatchSoftwareTitleConfigurations_GetByNameV3_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListMockV3()

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetByNameV3(context.Background(), "Google Chrome")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Google Chrome", result.DisplayName)
	assert.Equal(t, "101", result.SoftwareTitleID)
}

// TestGetByNameV3_NotFound tests retrieving a patch software title configuration by name when not found.
func TestUnit_PatchSoftwareTitleConfigurations_GetByNameV3_NotFound(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListMockV3()

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetByNameV3(context.Background(), "Nonexistent Config")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not found")
}

// TestGetByNameV3_EmptyName tests retrieving a patch software title configuration with empty name.
func TestUnit_PatchSoftwareTitleConfigurations_GetByNameV3_EmptyName(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetByNameV3(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "name is required")
}

// TestCreateV3_APIError tests CreateV3 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_CreateV3_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterCreateNoResponseErrorMockV3()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{
		DisplayName:     "New Config",
		SoftwareTitleID: "103",
	}
	result, resp, err := svc.CreateV3(context.Background(), config)

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestCreateV3 tests creating a new patch software title configuration.
func TestUnit_PatchSoftwareTitleConfigurations_CreateV3_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterCreateMockV3()

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

	result, resp, err := svc.CreateV3(context.Background(), config)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "3", result.ID)
	assert.Contains(t, result.Href, "/api/v2/patch-software-title-configurations/3")
}

// TestCreateV3_NilConfig tests creating a patch software title configuration with nil config.
func TestUnit_PatchSoftwareTitleConfigurations_CreateV3_NilConfig(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.CreateV3(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "config is required")
}

// TestCreateV3_EmptyDisplayName tests creating a patch software title configuration with empty display name.
func TestUnit_PatchSoftwareTitleConfigurations_CreateV3_EmptyDisplayName(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{
		SoftwareTitleID: "103",
	}

	result, resp, err := svc.CreateV3(context.Background(), config)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "display name is required")
}

// TestCreateV3_EmptySoftwareTitleID tests creating a patch software title configuration with empty software title ID.
func TestUnit_PatchSoftwareTitleConfigurations_CreateV3_EmptySoftwareTitleID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{
		DisplayName: "New Config",
	}

	result, resp, err := svc.CreateV3(context.Background(), config)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "software title id is required")
}

// TestUpdateByIDV3_APIError tests UpdateByIDV3 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_UpdateByIDV3_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterUpdateByIDNoResponseErrorMockV3("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{DisplayName: "Updated", SoftwareTitleID: "101"}
	result, resp, err := svc.UpdateByIDV3(context.Background(), "1", config)

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestUpdateByIDV3 tests updating a patch software title configuration by ID.
func TestUnit_PatchSoftwareTitleConfigurations_UpdateByIDV3_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterUpdateByIDMockV3("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{
		DisplayName:        "Updated Google Chrome",
		SoftwareTitleID:    "101",
		UINotifications:    false,
		EmailNotifications: true,
	}

	result, resp, err := svc.UpdateByIDV3(context.Background(), "1", config)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Google Chrome", result.DisplayName)
}

// TestUpdateByIDV3_EmptyID tests updating a patch software title configuration with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_UpdateByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{
		DisplayName: "Updated Config",
	}

	result, resp, err := svc.UpdateByIDV3(context.Background(), "", config)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestUpdateByIDV3_NilConfig tests updating a patch software title configuration with nil config.
func TestUnit_PatchSoftwareTitleConfigurations_UpdateByIDV3_NilConfig(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.UpdateByIDV3(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "config is required")
}

// TestUpdateByNameV3_GetByNameError tests UpdateByNameV3 when GetByNameV3 fails.
func TestUnit_PatchSoftwareTitleConfigurations_UpdateByNameV3_GetByNameError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListNoResponseErrorMockV3()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{DisplayName: "Updated", SoftwareTitleID: "101"}
	result, resp, err := svc.UpdateByNameV3(context.Background(), "Google Chrome", config)

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestUpdateByNameV3 tests updating a patch software title configuration by name.
func TestUnit_PatchSoftwareTitleConfigurations_UpdateByNameV3_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListMockV3()
	mock.RegisterUpdateByIDMockV3("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{
		DisplayName:     "Updated Config",
		SoftwareTitleID: "101",
	}

	result, resp, err := svc.UpdateByNameV3(context.Background(), "Google Chrome", config)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
}

// TestUpdateByNameV3_EmptyName tests updating a patch software title configuration with empty name.
func TestUnit_PatchSoftwareTitleConfigurations_UpdateByNameV3_EmptyName(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{
		DisplayName: "Updated Config",
	}

	result, resp, err := svc.UpdateByNameV3(context.Background(), "", config)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "name is required")
}

// TestDeleteByIDV3_APIError tests DeleteByIDV3 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_DeleteByIDV3_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterDeleteByIDNoResponseErrorMockV3("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	resp, err := svc.DeleteByIDV3(context.Background(), "1")

	assert.Error(t, err)
	assert.NotNil(t, resp)
}

// TestDeleteByIDV3 tests deleting a patch software title configuration by ID.
func TestUnit_PatchSoftwareTitleConfigurations_DeleteByIDV3_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterDeleteByIDMockV3("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	resp, err := svc.DeleteByIDV3(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

// TestDeleteByIDV3_EmptyID tests deleting a patch software title configuration with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_DeleteByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	resp, err := svc.DeleteByIDV3(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

// TestDeleteByNameV3_GetByNameError tests DeleteByNameV3 when GetByNameV3 fails.
func TestUnit_PatchSoftwareTitleConfigurations_DeleteByNameV3_GetByNameError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListNoResponseErrorMockV3()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	resp, err := svc.DeleteByNameV3(context.Background(), "Google Chrome")

	assert.Error(t, err)
	assert.NotNil(t, resp)
}

// TestDeleteByNameV3 tests deleting a patch software title configuration by name.
func TestUnit_PatchSoftwareTitleConfigurations_DeleteByNameV3_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListMockV3()
	mock.RegisterDeleteByIDMockV3("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	resp, err := svc.DeleteByNameV3(context.Background(), "Google Chrome")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

// TestDeleteByNameV3_EmptyName tests deleting a patch software title configuration with empty name.
func TestUnit_PatchSoftwareTitleConfigurations_DeleteByNameV3_EmptyName(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	resp, err := svc.DeleteByNameV3(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "name is required")
}

// TestGetDashboardStatusByIDV3_APIError tests GetDashboardStatusByIDV3 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_GetDashboardStatusByIDV3_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetDashboardStatusNoResponseErrorMockV3("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetDashboardStatusByIDV3(context.Background(), "1")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestGetDashboardStatusByIDV3 tests getting dashboard status.
func TestUnit_PatchSoftwareTitleConfigurations_GetDashboardStatusByIDV3_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetDashboardStatusMockV3("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetDashboardStatusByIDV3(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.True(t, result.OnDashboard)
}

// TestGetDashboardStatusByIDV3_EmptyID tests getting dashboard status with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_GetDashboardStatusByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetDashboardStatusByIDV3(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestAddToDashboardByIDV3_EmptyID tests adding to dashboard with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_AddToDashboardByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	resp, err := svc.AddToDashboardByIDV3(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

// TestAddToDashboardByIDV3_APIError tests AddToDashboardByIDV3 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_AddToDashboardByIDV3_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterAddToDashboardNoResponseErrorMockV3("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	resp, err := svc.AddToDashboardByIDV3(context.Background(), "1")

	assert.Error(t, err)
	assert.NotNil(t, resp)
}

// TestAddToDashboardByIDV3 tests adding to dashboard.
func TestUnit_PatchSoftwareTitleConfigurations_AddToDashboardByIDV3_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterAddToDashboardMockV3("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	resp, err := svc.AddToDashboardByIDV3(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

// TestRemoveFromDashboardByIDV3_EmptyID tests removing from dashboard with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_RemoveFromDashboardByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	resp, err := svc.RemoveFromDashboardByIDV3(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

// TestRemoveFromDashboardByIDV3_APIError tests RemoveFromDashboardByIDV3 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_RemoveFromDashboardByIDV3_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterRemoveFromDashboardNoResponseErrorMockV3("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	resp, err := svc.RemoveFromDashboardByIDV3(context.Background(), "1")

	assert.Error(t, err)
	assert.NotNil(t, resp)
}

// TestRemoveFromDashboardByIDV3 tests removing from dashboard.
func TestUnit_PatchSoftwareTitleConfigurations_RemoveFromDashboardByIDV3_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterRemoveFromDashboardMockV3("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	resp, err := svc.RemoveFromDashboardByIDV3(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

// TestGetDefinitionsByIDV3_EmptyID tests getting definitions with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_GetDefinitionsByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetDefinitionsByIDV3(context.Background(), "", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetDefinitionsByIDV3_APIError tests GetDefinitionsByIDV3 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_GetDefinitionsByIDV3_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetDefinitionsNoResponseErrorMockV3("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetDefinitionsByIDV3(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestGetDefinitionsByIDV3 tests getting definitions.
func TestUnit_PatchSoftwareTitleConfigurations_GetDefinitionsByIDV3_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetDefinitionsMockV3("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetDefinitionsByIDV3(context.Background(), "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "10.37.0", result.Results[0].Version)
	assert.Equal(t, "1", result.Results[0].AbsoluteOrderID)
}

// TestGetDependenciesByIDV3_EmptyID tests getting dependencies with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_GetDependenciesByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetDependenciesByIDV3(context.Background(), "", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetDependenciesByIDV3_APIError tests GetDependenciesByIDV3 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_GetDependenciesByIDV3_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetDependenciesNoResponseErrorMockV3("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetDependenciesByIDV3(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestGetDependenciesByIDV3 tests getting dependencies.
func TestUnit_PatchSoftwareTitleConfigurations_GetDependenciesByIDV3_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetDependenciesMockV3("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetDependenciesByIDV3(context.Background(), "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].SmartGroupID)
	assert.Equal(t, "Chrome Out of Date", result.Results[0].SmartGroupName)
}

// TestExportReportByIDV3_EmptyID tests exporting report with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_ExportReportByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	body, resp, err := svc.ExportReportByIDV3(context.Background(), "", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, body)
	assert.Contains(t, err.Error(), "id is required")
}

// TestExportReportByIDV3_APIError tests ExportReportByIDV3 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_ExportReportByIDV3_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterExportReportNoResponseErrorMockV3("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	body, resp, err := svc.ExportReportByIDV3(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, body)
}

// TestExportReportByIDV3 tests exporting report.
func TestUnit_PatchSoftwareTitleConfigurations_ExportReportByIDV3_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterExportReportMockV3("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	body, resp, err := svc.ExportReportByIDV3(context.Background(), "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, body)
	assert.Contains(t, string(body), "computerName")
}

// TestGetExtensionAttributesByIDV3_EmptyID tests getting extension attributes with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_GetExtensionAttributesByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetExtensionAttributesByIDV3(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetExtensionAttributesByIDV3_APIError tests GetExtensionAttributesByIDV3 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_GetExtensionAttributesByIDV3_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetExtensionAttributesNoResponseErrorMockV3("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetExtensionAttributesByIDV3(context.Background(), "1")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestGetExtensionAttributesByIDV3 tests getting extension attributes.
func TestUnit_PatchSoftwareTitleConfigurations_GetExtensionAttributesByIDV3_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetExtensionAttributesMockV3("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetExtensionAttributesByIDV3(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result, 1)
	assert.Equal(t, "google-chrome-ea", result[0].EAID)
	assert.True(t, result[0].Accepted)
}

// TestGetPatchReportByIDV3_EmptyID tests getting patch report with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_GetPatchReportByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetPatchReportByIDV3(context.Background(), "", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetPatchReportByIDV3_APIError tests GetPatchReportByIDV3 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_GetPatchReportByIDV3_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetPatchReportNoResponseErrorMockV3("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetPatchReportByIDV3(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestGetPatchReportByIDV3 tests getting patch report.
func TestUnit_PatchSoftwareTitleConfigurations_GetPatchReportByIDV3_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetPatchReportMockV3("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetPatchReportByIDV3(context.Background(), "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "MacBook", result.Results[0].ComputerName)
	assert.Equal(t, "10.1", result.Results[0].Version)
	// 11.30 renamed results[].lastContactTime to results[].lastCheckIn.
	assert.Equal(t, "1970-01-01T00:00:00Z", result.Results[0].LastCheckIn)
}

// TestGetPatchSummaryByIDV3_EmptyID tests getting patch summary with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_GetPatchSummaryByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetPatchSummaryByIDV3(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetPatchSummaryByIDV3_APIError tests GetPatchSummaryByIDV3 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_GetPatchSummaryByIDV3_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetPatchSummaryNoResponseErrorMockV3("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetPatchSummaryByIDV3(context.Background(), "1")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestGetPatchSummaryByIDV3 tests getting patch summary.
func TestUnit_PatchSoftwareTitleConfigurations_GetPatchSummaryByIDV3_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetPatchSummaryMockV3("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetPatchSummaryByIDV3(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.SoftwareTitleID)
	assert.Equal(t, "Patch title", result.Title)
	assert.Equal(t, 3, result.UpToDate)
	assert.Equal(t, 6, result.OutOfDate)
}

// TestGetHistoryByIDV3_EmptyID tests getting history with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_GetHistoryByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetHistoryByIDV3(context.Background(), "", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetHistoryByIDV3_APIError tests GetHistoryByIDV3 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_GetHistoryByIDV3_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetHistoryNoResponseErrorMockV3("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetHistoryByIDV3(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestGetHistoryByIDV3 tests getting history.
func TestUnit_PatchSoftwareTitleConfigurations_GetHistoryByIDV3_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetHistoryMockV3("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetHistoryByIDV3(context.Background(), "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "Sso settings update", result.Results[0].Note)
}

// TestAddHistoryNoteByIDV3_APIError tests AddHistoryNoteByIDV3 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_AddHistoryNoteByIDV3_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterAddHistoryNoteNoResponseErrorMockV3("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.AddHistoryNoteByIDV3(context.Background(), "1", &RequestAddHistoryNote{Note: "Test"})

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestAddHistoryNoteByIDV3 tests adding history note.
func TestUnit_PatchSoftwareTitleConfigurations_AddHistoryNoteByIDV3_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterAddHistoryNoteMockV3("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.AddHistoryNoteByIDV3(context.Background(), "1", &RequestAddHistoryNote{Note: "Test note"})

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
	assert.Contains(t, result.Href, "/api/v1/resource/1")
}

// TestAddHistoryNoteByIDV3_EmptyID tests adding history note with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_AddHistoryNoteByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.AddHistoryNoteByIDV3(context.Background(), "", &RequestAddHistoryNote{Note: "Test"})

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestAddHistoryNoteByIDV3_NilRequest tests adding history note with nil request.
func TestUnit_PatchSoftwareTitleConfigurations_AddHistoryNoteByIDV3_NilRequest(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.AddHistoryNoteByIDV3(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

// TestAddHistoryNoteByIDV3_EmptyNote tests adding history note with empty note.
func TestUnit_PatchSoftwareTitleConfigurations_AddHistoryNoteByIDV3_EmptyNote(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.AddHistoryNoteByIDV3(context.Background(), "1", &RequestAddHistoryNote{Note: ""})

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "note is required")
}

// TestGetPatchVersionsByIDV3_EmptyID tests getting patch versions with empty ID.
func TestUnit_PatchSoftwareTitleConfigurations_GetPatchVersionsByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetPatchVersionsByIDV3(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetPatchVersionsByIDV3_APIError tests GetPatchVersionsByIDV3 when API returns error.
func TestUnit_PatchSoftwareTitleConfigurations_GetPatchVersionsByIDV3_APIError(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetPatchVersionsNoResponseErrorMockV3("1")
	svc := NewPatchSoftwareTitleConfigurations(mock)

	result, resp, err := svc.GetPatchVersionsByIDV3(context.Background(), "1")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

// TestGetPatchVersionsByIDV3 tests getting patch versions.
func TestUnit_PatchSoftwareTitleConfigurations_GetPatchVersionsByIDV3_Success(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetPatchVersionsMockV3("1")

	svc := NewPatchSoftwareTitleConfigurations(mock)
	result, resp, err := svc.GetPatchVersionsByIDV3(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result, 1)
	assert.Equal(t, "1", result[0].AbsoluteOrderID)
	assert.Equal(t, "3", result[0].Version)
	assert.Equal(t, 1, result[0].OnVersion)
}
