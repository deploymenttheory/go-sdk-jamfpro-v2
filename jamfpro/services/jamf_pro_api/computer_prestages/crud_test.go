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

func TestUnitGetByNameV3_Success(t *testing.T) {
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

func TestUnitUpdateByNameV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()       // GetByNameV3 (also provides version lock)
	mock.RegisterUpdateByIDMock("1") // PUT

	request := &ResourceComputerPrestage{DisplayName: "Updated"}
	result, resp, err := svc.UpdateByNameV3(context.Background(), "Test Prestage", request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitUpdateByNameV3_NilRequest(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.UpdateByNameV3(context.Background(), "Test Prestage", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitDeleteByNameV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()
	mock.RegisterDeleteByIDMock("1")

	resp, err := svc.DeleteByNameV3(context.Background(), "Test Prestage")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteByNameV3_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	resp, err := svc.DeleteByNameV3(context.Background(), "Nonexistent")
	require.Error(t, err)
	require.Contains(t, err.Error(), "not found")
	_ = resp
}

func TestUnitGetDeviceScopeByIDV2_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetDeviceScopeByIDV2(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitGetDeviceScopeByIDV2_Success(t *testing.T) {
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

func TestUnitReplaceDeviceScopeByIDV2_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.ReplaceDeviceScopeByIDV2(context.Background(), "", &ReplaceDeviceScopeRequest{SerialNumbers: []string{"ABC"}, VersionLock: 1})
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitReplaceDeviceScopeByIDV2_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.ReplaceDeviceScopeByIDV2(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitReplaceDeviceScopeByIDV2_Success(t *testing.T) {
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
