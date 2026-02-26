package patch_software_title_configurations

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/patch_software_title_configurations/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestListV2 tests listing all patch software title configurations.
func TestListV2(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
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

// TestGetByIDV2 tests retrieving a patch software title configuration by ID.
func TestGetByIDV2(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetByIDMock("1")

	svc := NewService(mock)
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
func TestGetByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewService(mock)

	result, resp, err := svc.GetByIDV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetByNameV2 tests retrieving a patch software title configuration by display name.
func TestGetByNameV2(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	result, resp, err := svc.GetByNameV2(context.Background(), "Google Chrome")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Google Chrome", result.DisplayName)
	assert.Equal(t, "101", result.SoftwareTitleID)
}

// TestGetByNameV2_NotFound tests retrieving a patch software title configuration by name when not found.
func TestGetByNameV2_NotFound(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	result, resp, err := svc.GetByNameV2(context.Background(), "Nonexistent Config")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not found")
}

// TestGetByNameV2_EmptyName tests retrieving a patch software title configuration with empty name.
func TestGetByNameV2_EmptyName(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewService(mock)

	result, resp, err := svc.GetByNameV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "name is required")
}

// TestCreateV2 tests creating a new patch software title configuration.
func TestCreateV2(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterCreateMock()

	svc := NewService(mock)

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
func TestCreateV2_NilConfig(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewService(mock)

	result, resp, err := svc.CreateV2(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "config is required")
}

// TestCreateV2_EmptyDisplayName tests creating a patch software title configuration with empty display name.
func TestCreateV2_EmptyDisplayName(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewService(mock)

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
func TestCreateV2_EmptySoftwareTitleID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewService(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{
		DisplayName: "New Config",
	}

	result, resp, err := svc.CreateV2(context.Background(), config)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "software title id is required")
}

// TestUpdateByIDV2 tests updating a patch software title configuration by ID.
func TestUpdateByIDV2(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterUpdateByIDMock("1")

	svc := NewService(mock)

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
func TestUpdateByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewService(mock)

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
func TestUpdateByIDV2_NilConfig(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewService(mock)

	result, resp, err := svc.UpdateByIDV2(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "config is required")
}

// TestUpdateByNameV2 tests updating a patch software title configuration by name.
func TestUpdateByNameV2(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListMock()
	mock.RegisterUpdateByIDMock("1")

	svc := NewService(mock)

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
func TestUpdateByNameV2_EmptyName(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewService(mock)

	config := &ResourcePatchSoftwareTitleConfiguration{
		DisplayName: "Updated Config",
	}

	result, resp, err := svc.UpdateByNameV2(context.Background(), "", config)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "name is required")
}

// TestDeleteByIDV2 tests deleting a patch software title configuration by ID.
func TestDeleteByIDV2(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterDeleteByIDMock("1")

	svc := NewService(mock)
	resp, err := svc.DeleteByIDV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

// TestDeleteByIDV2_EmptyID tests deleting a patch software title configuration with empty ID.
func TestDeleteByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewService(mock)

	resp, err := svc.DeleteByIDV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

// TestDeleteByNameV2 tests deleting a patch software title configuration by name.
func TestDeleteByNameV2(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterListMock()
	mock.RegisterDeleteByIDMock("1")

	svc := NewService(mock)
	resp, err := svc.DeleteByNameV2(context.Background(), "Google Chrome")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

// TestDeleteByNameV2_EmptyName tests deleting a patch software title configuration with empty name.
func TestDeleteByNameV2_EmptyName(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewService(mock)

	resp, err := svc.DeleteByNameV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "name is required")
}

// TestGetDashboardStatusByIDV2 tests getting dashboard status.
func TestGetDashboardStatusByIDV2(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetDashboardStatusMock("1")

	svc := NewService(mock)
	result, resp, err := svc.GetDashboardStatusByIDV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.True(t, result.OnDashboard)
}

// TestGetDashboardStatusByIDV2_EmptyID tests getting dashboard status with empty ID.
func TestGetDashboardStatusByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewService(mock)

	result, resp, err := svc.GetDashboardStatusByIDV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestAddToDashboardByIDV2 tests adding to dashboard.
func TestAddToDashboardByIDV2(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterAddToDashboardMock("1")

	svc := NewService(mock)
	resp, err := svc.AddToDashboardByIDV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

// TestRemoveFromDashboardByIDV2 tests removing from dashboard.
func TestRemoveFromDashboardByIDV2(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterRemoveFromDashboardMock("1")

	svc := NewService(mock)
	resp, err := svc.RemoveFromDashboardByIDV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

// TestGetDefinitionsByIDV2 tests getting definitions.
func TestGetDefinitionsByIDV2(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetDefinitionsMock("1")

	svc := NewService(mock)
	result, resp, err := svc.GetDefinitionsByIDV2(context.Background(), "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "10.37.0", result.Results[0].Version)
	assert.Equal(t, "1", result.Results[0].AbsoluteOrderID)
}

// TestGetDependenciesByIDV2 tests getting dependencies.
func TestGetDependenciesByIDV2(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetDependenciesMock("1")

	svc := NewService(mock)
	result, resp, err := svc.GetDependenciesByIDV2(context.Background(), "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].SmartGroupID)
	assert.Equal(t, "Chrome Out of Date", result.Results[0].SmartGroupName)
}

// TestExportReportByIDV2 tests exporting report.
func TestExportReportByIDV2(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterExportReportMock("1")

	svc := NewService(mock)
	body, resp, err := svc.ExportReportByIDV2(context.Background(), "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, body)
	assert.Contains(t, string(body), "computerName")
}

// TestGetExtensionAttributesByIDV2 tests getting extension attributes.
func TestGetExtensionAttributesByIDV2(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetExtensionAttributesMock("1")

	svc := NewService(mock)
	result, resp, err := svc.GetExtensionAttributesByIDV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result, 1)
	assert.Equal(t, "google-chrome-ea", result[0].EAID)
	assert.True(t, result[0].Accepted)
}

// TestGetPatchReportByIDV2 tests getting patch report.
func TestGetPatchReportByIDV2(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetPatchReportMock("1")

	svc := NewService(mock)
	result, resp, err := svc.GetPatchReportByIDV2(context.Background(), "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "MacBook", result.Results[0].ComputerName)
	assert.Equal(t, "10.1", result.Results[0].Version)
}

// TestGetPatchSummaryByIDV2 tests getting patch summary.
func TestGetPatchSummaryByIDV2(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetPatchSummaryMock("1")

	svc := NewService(mock)
	result, resp, err := svc.GetPatchSummaryByIDV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.SoftwareTitleID)
	assert.Equal(t, "Patch title", result.Title)
	assert.Equal(t, 3, result.UpToDate)
	assert.Equal(t, 6, result.OutOfDate)
}

// TestGetHistoryByIDV2 tests getting history.
func TestGetHistoryByIDV2(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetHistoryMock("1")

	svc := NewService(mock)
	result, resp, err := svc.GetHistoryByIDV2(context.Background(), "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "Sso settings update", result.Results[0].Note)
}

// TestAddHistoryNoteByIDV2 tests adding history note.
func TestAddHistoryNoteByIDV2(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterAddHistoryNoteMock("1")

	svc := NewService(mock)
	result, resp, err := svc.AddHistoryNoteByIDV2(context.Background(), "1", &RequestAddHistoryNote{Note: "Test note"})

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
	assert.Contains(t, result.Href, "/api/v1/resource/1")
}

// TestAddHistoryNoteByIDV2_EmptyNote tests adding history note with empty note.
func TestAddHistoryNoteByIDV2_EmptyNote(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	svc := NewService(mock)

	result, resp, err := svc.AddHistoryNoteByIDV2(context.Background(), "1", &RequestAddHistoryNote{Note: ""})

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "note is required")
}

// TestGetPatchVersionsByIDV2 tests getting patch versions.
func TestGetPatchVersionsByIDV2(t *testing.T) {
	mock := mocks.NewPatchSoftwareTitleConfigurationsMock()
	mock.RegisterGetPatchVersionsMock("1")

	svc := NewService(mock)
	result, resp, err := svc.GetPatchVersionsByIDV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result, 1)
	assert.Equal(t, "1", result[0].AbsoluteOrderID)
	assert.Equal(t, "3", result[0].Version)
	assert.Equal(t, 1, result[0].OnVersion)
}
