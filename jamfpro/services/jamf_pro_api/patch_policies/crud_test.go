package patch_policies

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/patch_policies/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestListV2 tests listing all patch policies.
func TestListV2(t *testing.T) {
	mock := mocks.NewPatchPoliciesMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	result, resp, err := svc.ListV2(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 2, result.TotalCount)
	assert.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Adobe Acrobat Reader DC Patch Policy", result.Results[0].Name)
	assert.True(t, result.Results[0].Enabled)
	assert.Equal(t, "24.001.20643", result.Results[0].TargetPatchVersion)
	assert.Equal(t, "PROMPT_USERS_TO_SELF_SERVICE", result.Results[0].DeploymentMethod)
	assert.Equal(t, "138", result.Results[0].SoftwareTitleId)
	assert.Equal(t, 5, result.Results[0].KillAppsDelayMinutes)
	assert.True(t, result.Results[0].SelfServiceEnforceDeadline)
	assert.Equal(t, 7, result.Results[0].SelfServiceDeadline)
	assert.True(t, result.Results[0].ReminderEnabled)

	assert.Equal(t, "2", result.Results[1].ID)
	assert.Equal(t, "Google Chrome Patch Policy", result.Results[1].Name)
	assert.Equal(t, "INSTALL_ASAP", result.Results[1].DeploymentMethod)
	assert.True(t, result.Results[1].PatchUnknownVersion)
	assert.False(t, result.Results[1].ReminderEnabled)
}

// TestGetByIDV2 tests retrieving a patch policy by ID.
func TestGetByIDV2(t *testing.T) {
	mock := mocks.NewPatchPoliciesMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	result, resp, err := svc.GetByIDV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Adobe Acrobat Reader DC Patch Policy", result.Name)
	assert.True(t, result.Enabled)
	assert.Equal(t, "24.001.20643", result.TargetPatchVersion)
	assert.Equal(t, "PROMPT_USERS_TO_SELF_SERVICE", result.DeploymentMethod)
	assert.Equal(t, "138", result.SoftwareTitleId)
	assert.Equal(t, "8e9f10a1-2b3c-4d5e-6f7a-8b9c0d1e2f3a", result.SoftwareTitleConfigurationId)
	assert.Equal(t, 5, result.KillAppsDelayMinutes)
	assert.Equal(t, "Adobe Acrobat Reader will quit in 5 minutes to apply updates. Please save your work.", result.KillAppsMessage)
	assert.False(t, result.Downgrade)
	assert.False(t, result.PatchUnknownVersion)
	assert.Equal(t, "Adobe Acrobat Reader Update Available", result.NotificationHeader)
	assert.True(t, result.SelfServiceEnforceDeadline)
	assert.Equal(t, 7, result.SelfServiceDeadline)
	assert.Equal(t, "Install Update", result.InstallButtonText)
	assert.Equal(t, "This update includes security fixes and performance improvements for Adobe Acrobat Reader DC.", result.SelfServiceDescription)
	assert.Equal(t, "101", result.IconId)
	assert.Equal(t, 1, result.ReminderFrequency)
	assert.True(t, result.ReminderEnabled)
}

// TestGetByIDV2_NotFound tests retrieving a patch policy by ID when not found.
func TestGetByIDV2_NotFound(t *testing.T) {
	mock := mocks.NewPatchPoliciesMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	result, resp, err := svc.GetByIDV2(context.Background(), "999")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not found")
	assert.Contains(t, err.Error(), "999")
}

// TestGetByIDV2_EmptyID tests retrieving a patch policy with an empty ID.
func TestGetByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchPoliciesMock()
	svc := NewService(mock)

	result, resp, err := svc.GetByIDV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetByNameV2 tests retrieving a patch policy by name.
func TestGetByNameV2(t *testing.T) {
	mock := mocks.NewPatchPoliciesMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	result, resp, err := svc.GetByNameV2(context.Background(), "Adobe Acrobat Reader DC Patch Policy")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Adobe Acrobat Reader DC Patch Policy", result.Name)
	assert.True(t, result.Enabled)
}

// TestGetByNameV2_NotFound tests retrieving a patch policy by name when not found.
func TestGetByNameV2_NotFound(t *testing.T) {
	mock := mocks.NewPatchPoliciesMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	result, resp, err := svc.GetByNameV2(context.Background(), "Nonexistent Patch Policy")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not found")
}

// TestGetByNameV2_EmptyName tests retrieving a patch policy with an empty name.
func TestGetByNameV2_EmptyName(t *testing.T) {
	mock := mocks.NewPatchPoliciesMock()
	svc := NewService(mock)

	result, resp, err := svc.GetByNameV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "name is required")
}

// TestGetDashboardStatusV2 tests checking if a patch policy is on the dashboard.
func TestGetDashboardStatusV2(t *testing.T) {
	mock := mocks.NewPatchPoliciesMock()
	mock.RegisterGetDashboardStatusMock("1", true)

	svc := NewService(mock)
	result, resp, err := svc.GetDashboardStatusV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.True(t, result.OnDashboard)
}

// TestGetDashboardStatusV2_NotOnDashboard tests checking dashboard status when not on dashboard.
func TestGetDashboardStatusV2_NotOnDashboard(t *testing.T) {
	mock := mocks.NewPatchPoliciesMock()
	mock.RegisterGetDashboardStatusMock("2", false)

	svc := NewService(mock)
	result, resp, err := svc.GetDashboardStatusV2(context.Background(), "2")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.False(t, result.OnDashboard)
}

// TestGetDashboardStatusV2_EmptyID tests getting dashboard status with an empty ID.
func TestGetDashboardStatusV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchPoliciesMock()
	svc := NewService(mock)

	result, resp, err := svc.GetDashboardStatusV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestAddToDashboardV2 tests adding a patch policy to the dashboard.
func TestAddToDashboardV2(t *testing.T) {
	mock := mocks.NewPatchPoliciesMock()
	mock.RegisterAddToDashboardMock("1")

	svc := NewService(mock)
	resp, err := svc.AddToDashboardV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

// TestAddToDashboardV2_EmptyID tests adding to dashboard with an empty ID.
func TestAddToDashboardV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchPoliciesMock()
	svc := NewService(mock)

	resp, err := svc.AddToDashboardV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

// TestRemoveFromDashboardV2 tests removing a patch policy from the dashboard.
func TestRemoveFromDashboardV2(t *testing.T) {
	mock := mocks.NewPatchPoliciesMock()
	mock.RegisterRemoveFromDashboardMock("1")

	svc := NewService(mock)
	resp, err := svc.RemoveFromDashboardV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

// TestRemoveFromDashboardV2_EmptyID tests removing from dashboard with an empty ID.
func TestRemoveFromDashboardV2_EmptyID(t *testing.T) {
	mock := mocks.NewPatchPoliciesMock()
	svc := NewService(mock)

	resp, err := svc.RemoveFromDashboardV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

// TestListV2_Empty tests listing patch policies when the list is empty.
func TestListV2_Empty(t *testing.T) {
	mock := mocks.NewPatchPoliciesMock()
	mock.RegisterEmptyListMock()

	svc := NewService(mock)
	result, resp, err := svc.ListV2(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 0, result.TotalCount)
	assert.Len(t, result.Results, 0)
}

// TestGetByIDV2_EmptyList tests getting by ID when the list is empty.
func TestGetByIDV2_EmptyList(t *testing.T) {
	mock := mocks.NewPatchPoliciesMock()
	mock.RegisterEmptyListMock()

	svc := NewService(mock)
	result, resp, err := svc.GetByIDV2(context.Background(), "1")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not found")
}

// TestGetByNameV2_EmptyList tests getting by name when the list is empty.
func TestGetByNameV2_EmptyList(t *testing.T) {
	mock := mocks.NewPatchPoliciesMock()
	mock.RegisterEmptyListMock()

	svc := NewService(mock)
	result, resp, err := svc.GetByNameV2(context.Background(), "Test Policy")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not found")
}
