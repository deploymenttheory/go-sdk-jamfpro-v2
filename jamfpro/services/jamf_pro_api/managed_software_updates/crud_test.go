package managed_software_updates

import (
	"context"
	"net/url"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/managed_software_updates/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetAvailableUpdates tests retrieving available software updates.
func TestGetAvailableUpdates(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterGetAvailableUpdatesMock()

	svc := NewService(mock)
	result, resp, err := svc.GetAvailableUpdates(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, result.AvailableUpdates)
	assert.Len(t, result.AvailableUpdates.MacOS, 6)
	assert.Len(t, result.AvailableUpdates.IOS, 6)
	assert.Contains(t, result.AvailableUpdates.MacOS, "14.2.1")
	assert.Contains(t, result.AvailableUpdates.IOS, "17.2.1")
}

// TestGetPlans tests listing all managed software update plans.
func TestGetPlans(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterGetPlansMock()

	svc := NewService(mock)
	params := url.Values{}
	result, resp, err := svc.GetPlans(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	assert.Len(t, result.Results, 2)

	// Verify first plan
	plan1 := result.Results[0]
	assert.Equal(t, "a1b2c3d4-e5f6-7890-abcd-ef1234567890", plan1.PlanUuid)
	assert.Equal(t, "12345", plan1.Device.DeviceId)
	assert.Equal(t, "COMPUTER", plan1.Device.ObjectType)
	assert.Equal(t, "DOWNLOAD_INSTALL_ALLOW_DEFERRAL", plan1.UpdateAction)
	assert.Equal(t, "LATEST_MAJOR", plan1.VersionType)
	assert.Equal(t, 3, plan1.MaxDeferrals)
	assert.Equal(t, "PENDING", plan1.Status.State)

	// Verify second plan
	plan2 := result.Results[1]
	assert.Equal(t, "f9e8d7c6-b5a4-3210-fedc-ba0987654321", plan2.PlanUuid)
	assert.Equal(t, "67890", plan2.Device.DeviceId)
	assert.Equal(t, "MOBILE_DEVICE", plan2.Device.ObjectType)
	assert.Equal(t, "DOWNLOAD_INSTALL", plan2.UpdateAction)
	assert.Equal(t, "SPECIFIC_VERSION", plan2.VersionType)
	assert.Equal(t, "17.2.1", plan2.SpecificVersion)
	assert.Equal(t, "COMPLETED", plan2.Status.State)
}

// TestGetPlans_Empty tests listing plans when empty.
func TestGetPlans_Empty(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterEmptyPlansMock()

	svc := NewService(mock)
	params := url.Values{}
	result, resp, err := svc.GetPlans(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 0, result.TotalCount)
	assert.Len(t, result.Results, 0)
}

// TestGetPlanByUUID tests retrieving a plan by UUID.
func TestGetPlanByUUID(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	uuid := "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
	mock.RegisterGetPlanByUUIDMock(uuid)

	svc := NewService(mock)
	result, resp, err := svc.GetPlanByUUID(context.Background(), uuid)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, uuid, result.PlanUuid)
	assert.Equal(t, "12345", result.Device.DeviceId)
	assert.Equal(t, "COMPUTER", result.Device.ObjectType)
	assert.Equal(t, "DOWNLOAD_INSTALL_ALLOW_DEFERRAL", result.UpdateAction)
	assert.Equal(t, "LATEST_MAJOR", result.VersionType)
	assert.Equal(t, 3, result.MaxDeferrals)
	assert.Equal(t, "2024-03-15T14:00:00", result.ForceInstallLocalDateTime)
}

// TestGetPlanByUUID_EmptyUUID tests getting plan with an empty UUID.
func TestGetPlanByUUID_EmptyUUID(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	svc := NewService(mock)

	result, resp, err := svc.GetPlanByUUID(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "uuid is required")
}

// TestGetDeclarationsByPlanUUID tests retrieving declarations by plan UUID.
func TestGetDeclarationsByPlanUUID(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	uuid := "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
	mock.RegisterGetDeclarationsByPlanUUIDMock(uuid)

	svc := NewService(mock)
	result, resp, err := svc.GetDeclarationsByPlanUUID(context.Background(), uuid)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, result.Declarations, 2)

	// Verify first declaration
	decl1 := result.Declarations[0]
	assert.Equal(t, "decl-uuid-1234", decl1.UUID)
	assert.Equal(t, "com.apple.configuration.softwareupdate.enforcement.specific", decl1.Type)
	assert.Equal(t, "updates", decl1.Group)
	assert.NotEmpty(t, decl1.PayloadJson)

	// Verify second declaration
	decl2 := result.Declarations[1]
	assert.Equal(t, "decl-uuid-5678", decl2.UUID)
	assert.Equal(t, "com.apple.configuration.softwareupdate.settings", decl2.Type)
}

// TestGetDeclarationsByPlanUUID_EmptyUUID tests getting declarations with an empty UUID.
func TestGetDeclarationsByPlanUUID_EmptyUUID(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	svc := NewService(mock)

	result, resp, err := svc.GetDeclarationsByPlanUUID(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "uuid is required")
}

// TestCreatePlanByDeviceID tests creating a plan by device ID.
func TestCreatePlanByDeviceID(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterCreatePlanByDeviceIDMock()

	svc := NewService(mock)

	plan := &RequestPlanCreate{
		Devices: []PlanObject{
			{
				ObjectType: "COMPUTER",
				DeviceId:   "12345",
			},
		},
		Config: PlanConfig{
			UpdateAction: "DOWNLOAD_INSTALL_ALLOW_DEFERRAL",
			VersionType:  "LATEST_MAJOR",
			MaxDeferrals: 3,
		},
	}

	result, resp, err := svc.CreatePlanByDeviceID(context.Background(), plan)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 201, resp.StatusCode)
	assert.Len(t, result.Plans, 1)
	assert.Equal(t, "12345", result.Plans[0].Device.DeviceID)
	assert.Equal(t, "COMPUTER", result.Plans[0].Device.ObjectType)
	assert.Equal(t, "a1b2c3d4-e5f6-7890-abcd-ef1234567890", result.Plans[0].PlanID)
}

// TestCreatePlanByDeviceID_NilPlan tests creating a plan with nil input.
func TestCreatePlanByDeviceID_NilPlan(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	svc := NewService(mock)

	result, resp, err := svc.CreatePlanByDeviceID(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "plan is required")
}

// TestCreatePlanByGroupID tests creating a plan by group ID.
func TestCreatePlanByGroupID(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterCreatePlanByGroupIDMock()

	svc := NewService(mock)

	plan := &RequestPlanCreate{
		Group: PlanObject{
			ObjectType: "COMPUTER",
			GroupId:    "100",
		},
		Config: PlanConfig{
			UpdateAction: "DOWNLOAD_INSTALL",
			VersionType:  "SPECIFIC_VERSION",
			SpecificVersion: "14.2.1",
		},
	}

	result, resp, err := svc.CreatePlanByGroupID(context.Background(), plan)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 201, resp.StatusCode)
	assert.Len(t, result.Plans, 1)
}

// TestCreatePlanByGroupID_NilPlan tests creating a group plan with nil input.
func TestCreatePlanByGroupID_NilPlan(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	svc := NewService(mock)

	result, resp, err := svc.CreatePlanByGroupID(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "plan is required")
}

// TestGetPlansByGroupID tests retrieving plans by group ID.
func TestGetPlansByGroupID(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	groupID := "100"
	mock.RegisterGetPlansByGroupIDMock(groupID)

	svc := NewService(mock)
	result, resp, err := svc.GetPlansByGroupID(context.Background(), groupID, "COMPUTER")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	assert.Len(t, result.Results, 2)
}

// TestGetPlansByGroupID_EmptyGroupID tests getting plans with empty group ID.
func TestGetPlansByGroupID_EmptyGroupID(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	svc := NewService(mock)

	result, resp, err := svc.GetPlansByGroupID(context.Background(), "", "COMPUTER")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "groupID is required")
}

// TestGetPlansByGroupID_EmptyGroupType tests getting plans with empty group type.
func TestGetPlansByGroupID_EmptyGroupType(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	svc := NewService(mock)

	result, resp, err := svc.GetPlansByGroupID(context.Background(), "100", "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "groupType is required")
}

// TestGetFeatureToggle tests retrieving the feature toggle.
func TestGetFeatureToggle(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterGetFeatureToggleMock()

	svc := NewService(mock)
	result, resp, err := svc.GetFeatureToggle(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, result.Toggle)
}

// TestUpdateFeatureToggle tests updating the feature toggle.
func TestUpdateFeatureToggle(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterUpdateFeatureToggleMock()

	svc := NewService(mock)

	toggle := &RequestFeatureToggle{
		Toggle: true,
	}

	result, resp, err := svc.UpdateFeatureToggle(context.Background(), toggle)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, result.Toggle)
	assert.True(t, result.ForceInstallLocalDateEnabled)
	assert.True(t, result.DssEnabled)
	assert.False(t, result.RecipeEnabled)
}

// TestUpdateFeatureToggle_NilToggle tests updating feature toggle with nil input.
func TestUpdateFeatureToggle_NilToggle(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	svc := NewService(mock)

	result, resp, err := svc.UpdateFeatureToggle(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "toggle is required")
}

// TestGetFeatureToggleStatus tests retrieving the feature toggle status.
func TestGetFeatureToggleStatus(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterGetFeatureToggleStatusMock()

	svc := NewService(mock)
	result, resp, err := svc.GetFeatureToggleStatus(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, result.ToggleOn)
	assert.Nil(t, result.ToggleOff)
	assert.Equal(t, "COMPLETED", result.ToggleOn.State)
	assert.Equal(t, int64(1000), result.ToggleOn.TotalRecords)
	assert.Equal(t, int64(1000), result.ToggleOn.ProcessedRecords)
	assert.Equal(t, 100.0, result.ToggleOn.PercentComplete)
	assert.Equal(t, "SUCCESS", result.ToggleOn.ExitState)
}

// TestForceStopFeatureToggleProcess tests force stopping the feature toggle process.
func TestForceStopFeatureToggleProcess(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterForceStopFeatureToggleProcessMock()

	svc := NewService(mock)
	result, resp, err := svc.ForceStopFeatureToggleProcess(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 200, result.HTTPStatus)
	assert.Len(t, result.Errors, 1)
	assert.Equal(t, "PROCESS_STOPPED", result.Errors[0].Code)
}
