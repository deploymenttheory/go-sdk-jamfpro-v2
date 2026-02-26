package computer_prestages

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_prestages/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.ComputerPrestagesMock) {
	t.Helper()
	mock := mocks.NewComputerPrestagesMock()
	return NewService(mock), mock
}

func TestUnit_ComputerPrestages_ListV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.ListV3(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "Test Prestage", result.Results[0].DisplayName)
	assert.Equal(t, "1", result.Results[0].ID)
}

func TestUnit_ComputerPrestages_GetByIDV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDMock("1")

	result, resp, err := svc.GetByIDV3(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Test Prestage", result.DisplayName)
}

func TestUnit_ComputerPrestages_GetByIDV3_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDV3(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ComputerPrestages_GetByNameV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.GetByNameV3(context.Background(), "Test Prestage")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "Test Prestage", result.DisplayName)
}

func TestUnit_ComputerPrestages_CreateV3_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateV3(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ComputerPrestages_CreateV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	request := &ResourceComputerPrestage{DisplayName: "New Prestage"}
	result, resp, err := svc.CreateV3(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "/api/v3/computer-prestages/1", result.Href)
}

func TestUnit_ComputerPrestages_UpdateByIDV3_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByIDV3(context.Background(), "", &ResourceComputerPrestage{})
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ComputerPrestages_UpdateByIDV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDMock("1")    // internal GET for version lock
	mock.RegisterUpdateByIDMock("1") // PUT

	request := &ResourceComputerPrestage{DisplayName: "Updated"}
	result, resp, err := svc.UpdateByIDV3(context.Background(), "1", request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_ComputerPrestages_DeleteByIDV3_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByIDV3(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_ComputerPrestages_DeleteByIDV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteByIDMock("1")

	resp, err := svc.DeleteByIDV3(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_ComputerPrestages_UpdateByNameV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()       // GetByNameV3 (also provides version lock)
	mock.RegisterUpdateByIDMock("1") // PUT

	request := &ResourceComputerPrestage{DisplayName: "Updated"}
	result, resp, err := svc.UpdateByNameV3(context.Background(), "Test Prestage", request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_ComputerPrestages_UpdateByNameV3_NilRequest(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.UpdateByNameV3(context.Background(), "Test Prestage", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ComputerPrestages_DeleteByNameV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()
	mock.RegisterDeleteByIDMock("1")

	resp, err := svc.DeleteByNameV3(context.Background(), "Test Prestage")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_ComputerPrestages_DeleteByNameV3_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	resp, err := svc.DeleteByNameV3(context.Background(), "Nonexistent")
	require.Error(t, err)
	require.Contains(t, err.Error(), "not found")
	_ = resp
}

func TestUnit_ComputerPrestages_GetDeviceScopeByIDV2_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetDeviceScopeByIDV2(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ComputerPrestages_GetDeviceScopeByIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDeviceScopeMock("1")

	result, resp, err := svc.GetDeviceScopeByIDV2(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.PrestageId)
	assert.Equal(t, 1, result.VersionLock)
	require.Len(t, result.Assignments, 1)
	assert.Equal(t, "XYZ", result.Assignments[0].SerialNumber)
	assert.Equal(t, "admin", result.Assignments[0].UserAssigned)
}

func TestUnit_ComputerPrestages_ReplaceDeviceScopeByIDV2_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.ReplaceDeviceScopeByIDV2(context.Background(), "", &ReplaceDeviceScopeRequest{SerialNumbers: []string{"ABC"}, VersionLock: 1})
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ComputerPrestages_ReplaceDeviceScopeByIDV2_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.ReplaceDeviceScopeByIDV2(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ComputerPrestages_ReplaceDeviceScopeByIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDeviceScopeMock("1")     // internal GET for version lock
	mock.RegisterReplaceDeviceScopeMock("1") // PUT

	request := &ReplaceDeviceScopeRequest{SerialNumbers: []string{"XYZ"}}
	result, resp, err := svc.ReplaceDeviceScopeByIDV2(context.Background(), "1", request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.PrestageId)
	assert.Equal(t, 1, result.VersionLock)
	require.Len(t, result.Assignments, 1)
	assert.Equal(t, "XYZ", result.Assignments[0].SerialNumber)
}

// Optimistic locking tests

// TestUnit_ComputerPrestages_UpdateByIDV3_VersionLockPropagated verifies that
// UpdateByIDV3 fetches the current resource and injects the versionLock (and
// sub-resource locks) into the caller-supplied request before issuing the PUT.
func TestUnit_ComputerPrestages_UpdateByIDV3_VersionLockPropagated(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDMock("1")    // returns versionLock=1
	mock.RegisterUpdateByIDMock("1") // PUT

	request := &ResourceComputerPrestage{DisplayName: "Updated"}
	_, _, err := svc.UpdateByIDV3(context.Background(), "1", request)
	require.NoError(t, err)
	// EnsureVersionLock must have copied versionLock=1 from the GET response.
	assert.Equal(t, 1, request.VersionLock)
	// Sub-resource locks are propagated from the current resource (0 when absent in fixture).
	assert.Equal(t, 0, request.LocationInformation.VersionLock)
	assert.Equal(t, 0, request.PurchasingInformation.VersionLock)
}

// TestUnit_ComputerPrestages_UpdateByNameV3_VersionLockPropagated verifies that
// UpdateByNameV3 reuses the versionLock obtained via the name-lookup list call.
func TestUnit_ComputerPrestages_UpdateByNameV3_VersionLockPropagated(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()          // "Test Prestage" carries versionLock=1
	mock.RegisterUpdateByIDMock("1") // PUT

	request := &ResourceComputerPrestage{DisplayName: "Updated"}
	_, _, err := svc.UpdateByNameV3(context.Background(), "Test Prestage", request)
	require.NoError(t, err)
	assert.Equal(t, 1, request.VersionLock)
}

// TestUnit_ComputerPrestages_ReplaceDeviceScopeByIDV2_VersionLockPropagated
// verifies that ReplaceDeviceScopeByIDV2 fetches the current scope and injects
// its versionLock into the request before issuing the PUT.
func TestUnit_ComputerPrestages_ReplaceDeviceScopeByIDV2_VersionLockPropagated(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDeviceScopeMock("1")     // returns versionLock=1
	mock.RegisterReplaceDeviceScopeMock("1") // PUT

	request := &ReplaceDeviceScopeRequest{SerialNumbers: []string{"XYZ"}}
	_, _, err := svc.ReplaceDeviceScopeByIDV2(context.Background(), "1", request)
	require.NoError(t, err)
	assert.Equal(t, 1, request.VersionLock)
}

// TestUnit_ComputerPrestages_UpdateByIDV3_FetchVersionLockError verifies that
// UpdateByIDV3 returns an error when the internal GET (for version locking) fails.
func TestUnit_ComputerPrestages_UpdateByIDV3_FetchVersionLockError(t *testing.T) {
	svc, _ := setupMockService(t)
	// No GET mock registered – the internal fetch will fail.

	request := &ResourceComputerPrestage{DisplayName: "Updated"}
	result, resp, err := svc.UpdateByIDV3(context.Background(), "1", request)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to fetch current prestage for version locking")
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

// TestUnit_ComputerPrestages_ReplaceDeviceScopeByIDV2_FetchVersionLockError verifies
// that ReplaceDeviceScopeByIDV2 returns an error when the scope GET (for version
// locking) fails.
func TestUnit_ComputerPrestages_ReplaceDeviceScopeByIDV2_FetchVersionLockError(t *testing.T) {
	svc, _ := setupMockService(t)
	// No scope GET mock registered – the internal fetch will fail.

	request := &ReplaceDeviceScopeRequest{SerialNumbers: []string{"XYZ"}}
	result, resp, err := svc.ReplaceDeviceScopeByIDV2(context.Background(), "1", request)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to fetch current device scope for version locking")
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

// ListV3 error path
func TestUnit_ComputerPrestages_ListV3_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered – dispatch returns nil, err

	result, resp, err := svc.ListV3(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

// GetByNameV3 not found
func TestUnit_ComputerPrestages_GetByNameV3_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.GetByNameV3(context.Background(), "Nonexistent")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

// CreateV3 validation
func TestUnit_ComputerPrestages_CreateV3_ValidationEmptyDisplayName(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	request := &ResourceComputerPrestage{DisplayName: ""}
	result, resp, err := svc.CreateV3(context.Background(), request)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "displayName is required")
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ComputerPrestages_CreateV3_ValidationInvalidRecoveryLockPasswordType(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &ResourceComputerPrestage{DisplayName: "Test", RecoveryLockPasswordType: "INVALID"}
	result, resp, err := svc.CreateV3(context.Background(), request)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid recoveryLockPasswordType")
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ComputerPrestages_CreateV3_ValidationInvalidPrestageMinimumOsTargetVersionType(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &ResourceComputerPrestage{DisplayName: "Test", PrestageMinimumOsTargetVersionType: "INVALID"}
	result, resp, err := svc.CreateV3(context.Background(), request)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid prestageMinimumOsTargetVersionType")
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ComputerPrestages_CreateV3_ValidationInvalidUserAccountType(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &ResourceComputerPrestage{
		DisplayName: "Test",
		AccountSettings: &AccountSettings{UserAccountType: "INVALID"},
	}
	result, resp, err := svc.CreateV3(context.Background(), request)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid userAccountType")
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ComputerPrestages_CreateV3_ValidationInvalidPrefillType(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &ResourceComputerPrestage{
		DisplayName: "Test",
		AccountSettings: &AccountSettings{PrefillType: "INVALID"},
	}
	result, resp, err := svc.CreateV3(context.Background(), request)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid prefillType")
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

// UpdateByIDV3 nil request
func TestUnit_ComputerPrestages_UpdateByIDV3_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByIDV3(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ComputerPrestages_UpdateByIDV3_ValidationFailed(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDMock("1")

	request := &ResourceComputerPrestage{DisplayName: "", RecoveryLockPasswordType: "MANUAL"}
	result, resp, err := svc.UpdateByIDV3(context.Background(), "1", request)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "displayName is required")
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

// UpdateByNameV3 validation and not found
func TestUnit_ComputerPrestages_UpdateByNameV3_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.UpdateByNameV3(context.Background(), "Nonexistent", &ResourceComputerPrestage{DisplayName: "Updated"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

func TestUnit_ComputerPrestages_UpdateByNameV3_ValidationFailed(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	request := &ResourceComputerPrestage{DisplayName: ""}
	result, resp, err := svc.UpdateByNameV3(context.Background(), "Test Prestage", request)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "displayName is required")
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

// GetAllDeviceScopeV2
func TestUnit_ComputerPrestages_GetAllDeviceScopeV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAllDeviceScopeMock()

	result, resp, err := svc.GetAllDeviceScopeV2(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	require.Len(t, result.Prestages, 1)
	assert.Equal(t, "1", result.Prestages[0].PrestageId)
	assert.Equal(t, 1, result.Prestages[0].VersionLock)
	require.Len(t, result.Prestages[0].Assignments, 1)
	assert.Equal(t, "XYZ", result.Prestages[0].Assignments[0].SerialNumber)
}

// AddDeviceScopeByIDV2
func TestUnit_ComputerPrestages_AddDeviceScopeByIDV2_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.AddDeviceScopeByIDV2(context.Background(), "", &AddDeviceScopeRequest{SerialNumbers: []string{"ABC"}})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "id is required")
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ComputerPrestages_AddDeviceScopeByIDV2_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.AddDeviceScopeByIDV2(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ComputerPrestages_AddDeviceScopeByIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDeviceScopeMock("1")
	mock.RegisterAddDeviceScopeMock("1")

	request := &AddDeviceScopeRequest{SerialNumbers: []string{"ABC123"}}
	result, resp, err := svc.AddDeviceScopeByIDV2(context.Background(), "1", request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.PrestageId)
	assert.Equal(t, 1, result.VersionLock)
	require.Len(t, result.Assignments, 1)
	assert.Equal(t, "XYZ", result.Assignments[0].SerialNumber)
}

func TestUnit_ComputerPrestages_AddDeviceScopeByIDV2_VersionLockPropagated(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDeviceScopeMock("1")
	mock.RegisterAddDeviceScopeMock("1")

	request := &AddDeviceScopeRequest{SerialNumbers: []string{"ABC123"}}
	_, _, err := svc.AddDeviceScopeByIDV2(context.Background(), "1", request)
	require.NoError(t, err)
	assert.Equal(t, 1, request.VersionLock)
}

func TestUnit_ComputerPrestages_AddDeviceScopeByIDV2_FetchVersionLockError(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &AddDeviceScopeRequest{SerialNumbers: []string{"ABC"}}
	result, resp, err := svc.AddDeviceScopeByIDV2(context.Background(), "1", request)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to fetch current device scope for version locking")
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

// RemoveDeviceScopeByIDV2
func TestUnit_ComputerPrestages_RemoveDeviceScopeByIDV2_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.RemoveDeviceScopeByIDV2(context.Background(), "", &RemoveDeviceScopeRequest{SerialNumbers: []string{"ABC"}})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "id is required")
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ComputerPrestages_RemoveDeviceScopeByIDV2_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.RemoveDeviceScopeByIDV2(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ComputerPrestages_RemoveDeviceScopeByIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDeviceScopeMock("1")
	mock.RegisterRemoveDeviceScopeMock("1")

	request := &RemoveDeviceScopeRequest{SerialNumbers: []string{"XYZ"}}
	result, resp, err := svc.RemoveDeviceScopeByIDV2(context.Background(), "1", request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.PrestageId)
	assert.Equal(t, 1, result.VersionLock)
}

func TestUnit_ComputerPrestages_RemoveDeviceScopeByIDV2_VersionLockPropagated(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDeviceScopeMock("1")
	mock.RegisterRemoveDeviceScopeMock("1")

	request := &RemoveDeviceScopeRequest{SerialNumbers: []string{"XYZ"}}
	_, _, err := svc.RemoveDeviceScopeByIDV2(context.Background(), "1", request)
	require.NoError(t, err)
	assert.Equal(t, 1, request.VersionLock)
}

func TestUnit_ComputerPrestages_RemoveDeviceScopeByIDV2_FetchVersionLockError(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &RemoveDeviceScopeRequest{SerialNumbers: []string{"XYZ"}}
	result, resp, err := svc.RemoveDeviceScopeByIDV2(context.Background(), "1", request)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to fetch current device scope for version locking")
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

// ListV3 with query params
func TestUnit_ComputerPrestages_ListV3_WithQuery(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	query := map[string]string{"page": "1", "page-size": "10"}
	result, resp, err := svc.ListV3(context.Background(), query)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
}

// CreateV3 with valid enum values (covers validatePrefillType, validateUserAccountType, etc.)
func TestUnit_ComputerPrestages_CreateV3_ValidEnums(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	request := &ResourceComputerPrestage{
		DisplayName:                        "Test",
		RecoveryLockPasswordType:           "MANUAL",
		PrestageMinimumOsTargetVersionType: "NO_ENFORCEMENT",
		AccountSettings: &AccountSettings{
			UserAccountType: "ADMINISTRATOR",
			PrefillType:     "DEVICE_OWNER",
		},
	}
	result, resp, err := svc.CreateV3(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_ComputerPrestages_CreateV3_ValidEnumsCUSTOM(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	request := &ResourceComputerPrestage{
		DisplayName: "Test",
		AccountSettings: &AccountSettings{
			PrefillType: "CUSTOM",
		},
	}
	result, resp, err := svc.CreateV3(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_ComputerPrestages_CreateV3_ValidRecoveryLockRANDOM(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	request := &ResourceComputerPrestage{
		DisplayName:              "Test",
		RecoveryLockPasswordType: "RANDOM",
	}
	result, resp, err := svc.CreateV3(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_ComputerPrestages_CreateV3_ValidPrestageMinimumOsTypes(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	for _, v := range []string{"MINIMUM_OS_LATEST_VERSION", "MINIMUM_OS_LATEST_MAJOR_VERSION", "MINIMUM_OS_LATEST_MINOR_VERSION", "MINIMUM_OS_SPECIFIC_VERSION"} {
		request := &ResourceComputerPrestage{
			DisplayName:                        "Test",
			PrestageMinimumOsTargetVersionType: v,
		}
		result, resp, err := svc.CreateV3(context.Background(), request)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode)
	}
}

func TestUnit_ComputerPrestages_CreateV3_ValidUserAccountTypes(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	for _, v := range []string{"STANDARD", "SKIP"} {
		request := &ResourceComputerPrestage{
			DisplayName: "Test",
			AccountSettings: &AccountSettings{
				UserAccountType: v,
			},
		}
		result, resp, err := svc.CreateV3(context.Background(), request)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode)
	}
}
