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
func TestUnit_ManagedSoftwareUpdates_GetAvailableUpdates_Error(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterErrorMock("GET", "/api/v1/managed-software-updates/available-updates", "api error")

	svc := NewService(mock)
	result, resp, err := svc.GetAvailableUpdates(context.Background())

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "api error")
}

func TestUnit_ManagedSoftwareUpdates_GetAvailableUpdates_Success(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterGetAvailableUpdatesMock()

	svc := NewService(mock)
	result, resp, err := svc.GetAvailableUpdates(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotNil(t, result.AvailableUpdates)
	assert.Len(t, result.AvailableUpdates.MacOS, 6)
	assert.Len(t, result.AvailableUpdates.IOS, 6)
	assert.Contains(t, result.AvailableUpdates.MacOS, "14.2.1")
	assert.Contains(t, result.AvailableUpdates.IOS, "17.2.1")
}

// TestGetPlans tests listing all managed software update plans.
func TestUnit_ManagedSoftwareUpdates_GetPlans_Success(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterGetPlansMock()

	svc := NewService(mock)
	params := url.Values{}
	params.Set("page", "0")
	params.Set("page-size", "100")
	result, resp, err := svc.GetPlans(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
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

// TestGetPlans_Error tests listing plans when API returns error.
func TestUnit_ManagedSoftwareUpdates_GetPlans_Error(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterErrorMock("GET", "/api/v1/managed-software-updates/plans", "api error")

	svc := NewService(mock)
	params := url.Values{}
	result, resp, err := svc.GetPlans(context.Background(), params)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "api error")
}

// TestGetPlans_InvalidJSON tests mergePage error path when response is invalid.
func TestUnit_ManagedSoftwareUpdates_GetPlans_InvalidJSON(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterGetPlansInvalidMock()

	svc := NewService(mock)
	params := url.Values{}
	result, resp, err := svc.GetPlans(context.Background(), params)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "mergePage failed")
}

// TestGetPlans_Empty tests listing plans when empty.
func TestUnit_ManagedSoftwareUpdates_GetPlans_Empty(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterEmptyPlansMock()

	svc := NewService(mock)
	params := url.Values{}
	result, resp, err := svc.GetPlans(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 0, result.TotalCount)
	assert.Len(t, result.Results, 0)
}

// TestGetPlanByUUID tests retrieving a plan by UUID.
func TestUnit_ManagedSoftwareUpdates_GetPlanByUUID_Success(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	uuid := "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
	mock.RegisterGetPlanByUUIDMock(uuid)

	svc := NewService(mock)
	result, resp, err := svc.GetPlanByUUID(context.Background(), uuid)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, uuid, result.PlanUuid)
	assert.Equal(t, "12345", result.Device.DeviceId)
	assert.Equal(t, "COMPUTER", result.Device.ObjectType)
	assert.Equal(t, "DOWNLOAD_INSTALL_ALLOW_DEFERRAL", result.UpdateAction)
	assert.Equal(t, "LATEST_MAJOR", result.VersionType)
	assert.Equal(t, 3, result.MaxDeferrals)
	assert.Equal(t, "2024-03-15T14:00:00", result.ForceInstallLocalDateTime)
}

// TestGetPlanByUUID_Error tests getting plan when API returns error.
func TestUnit_ManagedSoftwareUpdates_GetPlanByUUID_Error(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	uuid := "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
	mock.RegisterErrorMock("GET", "/api/v1/managed-software-updates/plans/"+uuid, "not found")

	svc := NewService(mock)
	result, resp, err := svc.GetPlanByUUID(context.Background(), uuid)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not found")
}

// TestGetPlanByUUID_EmptyUUID tests getting plan with an empty UUID.
func TestUnit_ManagedSoftwareUpdates_GetPlanByUUID_EmptyUUID(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	svc := NewService(mock)

	result, resp, err := svc.GetPlanByUUID(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "uuid is required")
}

// TestGetDeclarationsByPlanUUID tests retrieving declarations by plan UUID.
func TestUnit_ManagedSoftwareUpdates_GetDeclarationsByPlanUUID_Success(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	uuid := "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
	mock.RegisterGetDeclarationsByPlanUUIDMock(uuid)

	svc := NewService(mock)
	result, resp, err := svc.GetDeclarationsByPlanUUID(context.Background(), uuid)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
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

// TestGetDeclarationsByPlanUUID_Error tests getting declarations when API returns error.
func TestUnit_ManagedSoftwareUpdates_GetDeclarationsByPlanUUID_Error(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	uuid := "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
	mock.RegisterErrorMock("GET", "/api/v1/managed-software-updates/plans/"+uuid+"/declarations", "api error")

	svc := NewService(mock)
	result, resp, err := svc.GetDeclarationsByPlanUUID(context.Background(), uuid)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
}

// TestGetDeclarationsByPlanUUID_EmptyUUID tests getting declarations with an empty UUID.
func TestUnit_ManagedSoftwareUpdates_GetDeclarationsByPlanUUID_EmptyUUID(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	svc := NewService(mock)

	result, resp, err := svc.GetDeclarationsByPlanUUID(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "uuid is required")
}

// TestCreatePlanByDeviceID tests creating a plan by device ID.
func TestUnit_ManagedSoftwareUpdates_CreatePlanByDeviceID_Success(t *testing.T) {
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
	assert.Equal(t, 201, resp.StatusCode())
	assert.Len(t, result.Plans, 1)
	assert.Equal(t, "12345", result.Plans[0].Device.DeviceID)
	assert.Equal(t, "COMPUTER", result.Plans[0].Device.ObjectType)
	assert.Equal(t, "a1b2c3d4-e5f6-7890-abcd-ef1234567890", result.Plans[0].PlanID)
}

// TestCreatePlanByDeviceID_Error tests creating a plan when API returns error.
func TestUnit_ManagedSoftwareUpdates_CreatePlanByDeviceID_Error(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterErrorMock("POST", "/api/v1/managed-software-updates/plans", "conflict")

	svc := NewService(mock)
	plan := &RequestPlanCreate{
		Devices: []PlanObject{{ObjectType: "COMPUTER", DeviceId: "12345"}},
		Config:  PlanConfig{UpdateAction: "DOWNLOAD_INSTALL", VersionType: "LATEST_MAJOR"},
	}

	result, resp, err := svc.CreatePlanByDeviceID(context.Background(), plan)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "conflict")
}

// TestCreatePlanByDeviceID_NilPlan tests creating a plan with nil input.
func TestUnit_ManagedSoftwareUpdates_CreatePlanByDeviceID_NilPlan(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	svc := NewService(mock)

	result, resp, err := svc.CreatePlanByDeviceID(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "plan is required")
}

// TestCreatePlanByGroupID tests creating a plan by group ID.
func TestUnit_ManagedSoftwareUpdates_CreatePlanByGroupID_Success(t *testing.T) {
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
	assert.Equal(t, 201, resp.StatusCode())
	assert.Len(t, result.Plans, 1)
}

// TestCreatePlanByGroupID_Error tests creating a group plan when API returns error.
func TestUnit_ManagedSoftwareUpdates_CreatePlanByGroupID_Error(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterErrorMock("POST", "/api/v1/managed-software-updates/plans/group", "api error")

	svc := NewService(mock)
	plan := &RequestPlanCreate{
		Group:  PlanObject{ObjectType: "COMPUTER", GroupId: "100"},
		Config: PlanConfig{UpdateAction: "DOWNLOAD_INSTALL", VersionType: "LATEST_MAJOR"},
	}

	result, resp, err := svc.CreatePlanByGroupID(context.Background(), plan)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
}

// TestCreatePlanByGroupID_NilPlan tests creating a group plan with nil input.
func TestUnit_ManagedSoftwareUpdates_CreatePlanByGroupID_NilPlan(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	svc := NewService(mock)

	result, resp, err := svc.CreatePlanByGroupID(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "plan is required")
}

// TestGetPlansByGroupID tests retrieving plans by group ID.
func TestUnit_ManagedSoftwareUpdates_GetPlansByGroupID_Success(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	groupID := "100"
	mock.RegisterGetPlansByGroupIDMock(groupID)

	svc := NewService(mock)
	result, resp, err := svc.GetPlansByGroupID(context.Background(), groupID, "COMPUTER")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	assert.Len(t, result.Results, 2)
}

// TestGetPlansByGroupID_Error tests getting plans when API returns error.
func TestUnit_ManagedSoftwareUpdates_GetPlansByGroupID_Error(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	groupID := "100"
	mock.RegisterErrorMock("GET", "/api/v1/managed-software-updates/plans/group/"+groupID+"?group-type=COMPUTER", "api error")

	svc := NewService(mock)
	result, resp, err := svc.GetPlansByGroupID(context.Background(), groupID, "COMPUTER")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
}

// TestGetPlansByGroupID_EmptyGroupID tests getting plans with empty group ID.
func TestUnit_ManagedSoftwareUpdates_GetPlansByGroupID_EmptyGroupID(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	svc := NewService(mock)

	result, resp, err := svc.GetPlansByGroupID(context.Background(), "", "COMPUTER")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "groupID is required")
}

// TestGetPlansByGroupID_EmptyGroupType tests getting plans with empty group type.
func TestUnit_ManagedSoftwareUpdates_GetPlansByGroupID_EmptyGroupType(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	svc := NewService(mock)

	result, resp, err := svc.GetPlansByGroupID(context.Background(), "100", "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "groupType is required")
}

// TestGetFeatureToggle_Error tests retrieving feature toggle when API returns error.
func TestUnit_ManagedSoftwareUpdates_GetFeatureToggle_Error(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterErrorMock("GET", "/api/v1/managed-software-updates/plans/feature-toggle", "api error")

	svc := NewService(mock)
	result, resp, err := svc.GetFeatureToggle(context.Background())

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
}

// TestGetFeatureToggle tests retrieving the feature toggle.
func TestUnit_ManagedSoftwareUpdates_GetFeatureToggle_Success(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterGetFeatureToggleMock()

	svc := NewService(mock)
	result, resp, err := svc.GetFeatureToggle(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.True(t, result.Toggle)
}

// TestUpdateFeatureToggle tests updating the feature toggle.
func TestUnit_ManagedSoftwareUpdates_UpdateFeatureToggle_Success(t *testing.T) {
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
	assert.Equal(t, 200, resp.StatusCode())
	assert.True(t, result.Toggle)
	assert.True(t, result.ForceInstallLocalDateEnabled)
	assert.True(t, result.DssEnabled)
	assert.False(t, result.RecipeEnabled)
}

// TestUpdateFeatureToggle_Error tests updating feature toggle when API returns error.
func TestUnit_ManagedSoftwareUpdates_UpdateFeatureToggle_Error(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterErrorMock("PUT", "/api/v1/managed-software-updates/plans/feature-toggle", "api error")

	svc := NewService(mock)
	toggle := &RequestFeatureToggle{Toggle: true}

	result, resp, err := svc.UpdateFeatureToggle(context.Background(), toggle)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
}

// TestUpdateFeatureToggle_NilToggle tests updating feature toggle with nil input.
func TestUnit_ManagedSoftwareUpdates_UpdateFeatureToggle_NilToggle(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	svc := NewService(mock)

	result, resp, err := svc.UpdateFeatureToggle(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "toggle is required")
}

// TestGetFeatureToggleStatus_Error tests retrieving feature toggle status when API returns error.
func TestUnit_ManagedSoftwareUpdates_GetFeatureToggleStatus_Error(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterErrorMock("GET", "/api/v1/managed-software-updates/plans/feature-toggle/status", "api error")

	svc := NewService(mock)
	result, resp, err := svc.GetFeatureToggleStatus(context.Background())

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
}

// TestGetFeatureToggleStatus tests retrieving the feature toggle status.
func TestUnit_ManagedSoftwareUpdates_GetFeatureToggleStatus_Success(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterGetFeatureToggleStatusMock()

	svc := NewService(mock)
	result, resp, err := svc.GetFeatureToggleStatus(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotNil(t, result.ToggleOn)
	assert.Nil(t, result.ToggleOff)
	assert.Equal(t, "COMPLETED", result.ToggleOn.State)
	assert.Equal(t, int64(1000), result.ToggleOn.TotalRecords)
	assert.Equal(t, int64(1000), result.ToggleOn.ProcessedRecords)
	assert.Equal(t, 100.0, result.ToggleOn.PercentComplete)
	assert.Equal(t, "SUCCESS", result.ToggleOn.ExitState)
}

// TestForceStopFeatureToggleProcess_Error tests force stop when API returns error.
func TestUnit_ManagedSoftwareUpdates_ForceStopFeatureToggleProcess_Error(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterErrorMock("POST", "/api/v1/managed-software-updates/plans/feature-toggle/abandon", "api error")

	svc := NewService(mock)
	result, resp, err := svc.ForceStopFeatureToggleProcess(context.Background())

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
}

// TestForceStopFeatureToggleProcess tests force stopping the feature toggle process.
func TestUnit_ManagedSoftwareUpdates_ForceStopFeatureToggleProcess_Success(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterForceStopFeatureToggleProcessMock()

	svc := NewService(mock)
	result, resp, err := svc.ForceStopFeatureToggleProcess(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 200, result.HTTPStatus)
	assert.Len(t, result.Errors, 1)
	assert.Equal(t, "PROCESS_STOPPED", result.Errors[0].Code)
}

// TestGetPlanEventsByUUID tests retrieving plan events by UUID.
func TestUnit_ManagedSoftwareUpdates_GetPlanEventsByUUID_Success(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	uuid := "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
	mock.RegisterGetPlanEventsByUUIDMock(uuid)

	svc := NewService(mock)
	result, resp, err := svc.GetPlanEventsByUUID(context.Background(), uuid)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Len(t, result.Events, 2)
	assert.Equal(t, "PLAN_CREATED", result.Events[0].Type)
	assert.Equal(t, "SCAN_SCHEDULED", result.Events[1].Type)
}

// TestGetPlanEventsByUUID_Error tests getting plan events when API returns error.
func TestUnit_ManagedSoftwareUpdates_GetPlanEventsByUUID_Error(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	uuid := "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
	mock.RegisterErrorMock("GET", "/api/v1/managed-software-updates/plans/"+uuid+"/events", "api error")

	svc := NewService(mock)
	result, resp, err := svc.GetPlanEventsByUUID(context.Background(), uuid)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
}

// TestGetPlanEventsByUUID_EmptyUUID tests getting plan events with an empty UUID.
func TestUnit_ManagedSoftwareUpdates_GetPlanEventsByUUID_EmptyUUID(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	svc := NewService(mock)

	result, resp, err := svc.GetPlanEventsByUUID(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "uuid is required")
}

// TestGetUpdateStatuses_Error tests retrieving update statuses when API returns error.
func TestUnit_ManagedSoftwareUpdates_GetUpdateStatuses_Error(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterErrorMock("GET", "/api/v1/managed-software-updates/update-statuses", "api error")

	svc := NewService(mock)
	params := url.Values{}
	result, resp, err := svc.GetUpdateStatuses(context.Background(), params)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
}

// TestGetUpdateStatuses_ArrayFormat tests mergePage with raw array format (Real API style).
func TestUnit_ManagedSoftwareUpdates_GetUpdateStatuses_ArrayFormat(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterGetUpdateStatusesArrayFormatMock()

	svc := NewService(mock)
	params := url.Values{}
	result, resp, err := svc.GetUpdateStatuses(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "2", result.Results[0].OsUpdatesStatusId)
	assert.Equal(t, "99999", result.Results[0].Device.DeviceId)
	assert.Equal(t, "PENDING", result.Results[0].Status)
}

// TestGetUpdateStatuses_InvalidJSON tests mergePage error when response is invalid.
func TestUnit_ManagedSoftwareUpdates_GetUpdateStatuses_InvalidJSON(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterGetUpdateStatusesInvalidMock()

	svc := NewService(mock)
	params := url.Values{}
	result, resp, err := svc.GetUpdateStatuses(context.Background(), params)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "mergePage failed")
}

// TestGetUpdateStatuses tests retrieving update statuses with pagination.
func TestUnit_ManagedSoftwareUpdates_GetUpdateStatuses_Success(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	mock.RegisterGetUpdateStatusesMock()

	svc := NewService(mock)
	params := url.Values{}
	result, resp, err := svc.GetUpdateStatuses(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].OsUpdatesStatusId)
	assert.Equal(t, "12345", result.Results[0].Device.DeviceId)
	assert.Equal(t, "COMPUTER", result.Results[0].Device.ObjectType)
	assert.Equal(t, "DOWNLOADING", result.Results[0].Status)
	assert.Equal(t, "macOSUpdate19F77", result.Results[0].ProductKey)
}

// TestGetUpdateStatusesByComputerGroup tests retrieving update statuses by computer group.
func TestUnit_ManagedSoftwareUpdates_GetUpdateStatusesByComputerGroup_Success(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	groupID := "100"
	mock.RegisterGetUpdateStatusesByComputerGroupMock(groupID)

	svc := NewService(mock)
	result, resp, err := svc.GetUpdateStatusesByComputerGroup(context.Background(), groupID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].OsUpdatesStatusId)
}

// TestGetUpdateStatusesByComputerGroup_Error tests when API returns error.
func TestUnit_ManagedSoftwareUpdates_GetUpdateStatusesByComputerGroup_Error(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	groupID := "100"
	mock.RegisterErrorMock("GET", "/api/v1/managed-software-updates/update-statuses/computer-groups/"+groupID, "api error")

	svc := NewService(mock)
	result, resp, err := svc.GetUpdateStatusesByComputerGroup(context.Background(), groupID)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
}

// TestGetUpdateStatusesByComputerGroup_EmptyID tests getting update statuses with empty ID.
func TestUnit_ManagedSoftwareUpdates_GetUpdateStatusesByComputerGroup_EmptyID(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	svc := NewService(mock)

	result, resp, err := svc.GetUpdateStatusesByComputerGroup(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetUpdateStatusesByComputer tests retrieving update statuses by computer.
func TestUnit_ManagedSoftwareUpdates_GetUpdateStatusesByComputer_Success(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	computerID := "12345"
	mock.RegisterGetUpdateStatusesByComputerMock(computerID)

	svc := NewService(mock)
	result, resp, err := svc.GetUpdateStatusesByComputer(context.Background(), computerID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Len(t, result.Results, 1)
}

// TestGetUpdateStatusesByComputer_Error tests when API returns error.
func TestUnit_ManagedSoftwareUpdates_GetUpdateStatusesByComputer_Error(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	computerID := "12345"
	mock.RegisterErrorMock("GET", "/api/v1/managed-software-updates/update-statuses/computers/"+computerID, "api error")

	svc := NewService(mock)
	result, resp, err := svc.GetUpdateStatusesByComputer(context.Background(), computerID)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
}

// TestGetUpdateStatusesByComputer_EmptyID tests getting update statuses with empty ID.
func TestUnit_ManagedSoftwareUpdates_GetUpdateStatusesByComputer_EmptyID(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	svc := NewService(mock)

	result, resp, err := svc.GetUpdateStatusesByComputer(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetUpdateStatusesByMobileDeviceGroup tests retrieving update statuses by mobile device group.
func TestUnit_ManagedSoftwareUpdates_GetUpdateStatusesByMobileDeviceGroup_Success(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	groupID := "200"
	mock.RegisterGetUpdateStatusesByMobileDeviceGroupMock(groupID)

	svc := NewService(mock)
	result, resp, err := svc.GetUpdateStatusesByMobileDeviceGroup(context.Background(), groupID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Len(t, result.Results, 1)
}

// TestGetUpdateStatusesByMobileDeviceGroup_Error tests when API returns error.
func TestUnit_ManagedSoftwareUpdates_GetUpdateStatusesByMobileDeviceGroup_Error(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	groupID := "200"
	mock.RegisterErrorMock("GET", "/api/v1/managed-software-updates/update-statuses/mobile-device-groups/"+groupID, "api error")

	svc := NewService(mock)
	result, resp, err := svc.GetUpdateStatusesByMobileDeviceGroup(context.Background(), groupID)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
}

// TestGetUpdateStatusesByMobileDeviceGroup_EmptyID tests getting update statuses with empty ID.
func TestUnit_ManagedSoftwareUpdates_GetUpdateStatusesByMobileDeviceGroup_EmptyID(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	svc := NewService(mock)

	result, resp, err := svc.GetUpdateStatusesByMobileDeviceGroup(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetUpdateStatusesByMobileDevice tests retrieving update statuses by mobile device.
func TestUnit_ManagedSoftwareUpdates_GetUpdateStatusesByMobileDevice_Success(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	deviceID := "67890"
	mock.RegisterGetUpdateStatusesByMobileDeviceMock(deviceID)

	svc := NewService(mock)
	result, resp, err := svc.GetUpdateStatusesByMobileDevice(context.Background(), deviceID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Len(t, result.Results, 1)
}

// TestGetUpdateStatusesByMobileDevice_Error tests when API returns error.
func TestUnit_ManagedSoftwareUpdates_GetUpdateStatusesByMobileDevice_Error(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	deviceID := "67890"
	mock.RegisterErrorMock("GET", "/api/v1/managed-software-updates/update-statuses/mobile-devices/"+deviceID, "api error")

	svc := NewService(mock)
	result, resp, err := svc.GetUpdateStatusesByMobileDevice(context.Background(), deviceID)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
}

// TestGetUpdateStatusesByMobileDevice_EmptyID tests getting update statuses with empty ID.
func TestUnit_ManagedSoftwareUpdates_GetUpdateStatusesByMobileDevice_EmptyID(t *testing.T) {
	mock := mocks.NewManagedSoftwareUpdatesMock()
	svc := NewService(mock)

	result, resp, err := svc.GetUpdateStatusesByMobileDevice(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}
