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

func TestUnitListV3_Success(t *testing.T) {
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

func TestUnitGetByIDV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDMock("1")

	result, resp, err := svc.GetByIDV3(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Test Prestage", result.DisplayName)
}

func TestUnitGetByIDV3_EmptyID(t *testing.T) {
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

func TestUnitCreateV3_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateV3(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitCreateV3_Success(t *testing.T) {
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

func TestUnitUpdateByIDV3_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByIDV3(context.Background(), "", &ResourceComputerPrestage{})
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitUpdateByIDV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateByIDMock("1")

	request := &ResourceComputerPrestage{DisplayName: "Updated", VersionLock: 1}
	result, resp, err := svc.UpdateByIDV3(context.Background(), "1", request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteByIDV3_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByIDV3(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnitDeleteByIDV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteByIDMock("1")

	resp, err := svc.DeleteByIDV3(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitUpdateByNameV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()
	mock.RegisterUpdateByIDMock("1")

	request := &ResourceComputerPrestage{DisplayName: "Updated", VersionLock: 1}
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
	mock.RegisterReplaceDeviceScopeMock("1")

	request := &ReplaceDeviceScopeRequest{SerialNumbers: []string{"XYZ"}, VersionLock: 1}
	result, resp, err := svc.ReplaceDeviceScopeByIDV2(context.Background(), "1", request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.PrestageId)
	assert.Equal(t, 1, result.VersionLock)
	require.Len(t, result.Assignments, 1)
	assert.Equal(t, "XYZ", result.Assignments[0].SerialNumber)
}
