package log_flushing

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/log_flushing/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_LogFlushing_GetSettingsV1_Success(t *testing.T) {
	mock := mocks.NewLogFlushingMock()
	mock.RegisterGetSettingsMock()
	service := NewService(mock)

	result, resp, err := service.GetSettingsV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.HourOfDay)
	assert.Len(t, result.RetentionPolicies, 1)
	assert.Equal(t, "Jamf Pro Server Logs", result.RetentionPolicies[0].DisplayName)
}

func TestUnit_LogFlushing_ListTasksV1_Success(t *testing.T) {
	mock := mocks.NewLogFlushingMock()
	mock.RegisterListTasksMock()
	service := NewService(mock)

	result, resp, err := service.ListTasksV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, result, 1)
	assert.Equal(t, "1", result[0].ID)
	assert.Equal(t, "COMPLETED", result[0].State)
}

func TestUnit_LogFlushing_GetTaskByIDV1_Success(t *testing.T) {
	mock := mocks.NewLogFlushingMock()
	mock.RegisterGetTaskByIDMock()
	service := NewService(mock)

	result, resp, err := service.GetTaskByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "JAMFSoftwareServer", result.Qualifier)
	assert.Equal(t, "COMPLETED", result.State)
}

func TestUnit_LogFlushing_GetTaskByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewLogFlushingMock()
	service := NewService(mock)

	result, resp, err := service.GetTaskByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "task ID is required")
}

func TestUnit_LogFlushing_QueueTaskV1_Success(t *testing.T) {
	mock := mocks.NewLogFlushingMock()
	mock.RegisterQueueTaskMock()
	service := NewService(mock)

	request := &RequestLogFlushingTask{
		Qualifier:           "JAMFSoftwareServer",
		RetentionPeriod:     30,
		RetentionPeriodUnit: "Days",
	}

	result, resp, err := service.QueueTaskV1(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "2", result.ID)
	assert.Equal(t, "/api/v1/log-flushing/task/2", result.Href)
}

func TestUnit_LogFlushing_QueueTaskV1_NilRequest(t *testing.T) {
	mock := mocks.NewLogFlushingMock()
	service := NewService(mock)

	result, resp, err := service.QueueTaskV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "log flushing task request cannot be nil")
}

func TestUnit_LogFlushing_DeleteTaskByIDV1_Success(t *testing.T) {
	mock := mocks.NewLogFlushingMock()
	mock.RegisterDeleteTaskMock()
	service := NewService(mock)

	resp, err := service.DeleteTaskByIDV1(context.Background(), "1")
	require.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnit_LogFlushing_DeleteTaskByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewLogFlushingMock()
	service := NewService(mock)

	resp, err := service.DeleteTaskByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "task ID is required")
}

func TestUnit_LogFlushing_GetSettingsV1_Error(t *testing.T) {
	mock := mocks.NewLogFlushingMock()
	mock.RegisterGetSettingsErrorMock()
	service := NewService(mock)

	result, resp, err := service.GetSettingsV1(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to get log flushing settings")
}

func TestUnit_LogFlushing_ListTasksV1_Error(t *testing.T) {
	mock := mocks.NewLogFlushingMock()
	mock.RegisterListTasksErrorMock()
	service := NewService(mock)

	result, resp, err := service.ListTasksV1(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to list log flushing tasks")
}

func TestUnit_LogFlushing_GetTaskByIDV1_Error(t *testing.T) {
	mock := mocks.NewLogFlushingMock()
	mock.RegisterGetTaskByIDErrorMock("1")
	service := NewService(mock)

	result, resp, err := service.GetTaskByIDV1(context.Background(), "1")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to get log flushing task 1")
}

func TestUnit_LogFlushing_QueueTaskV1_Error(t *testing.T) {
	mock := mocks.NewLogFlushingMock()
	mock.RegisterQueueTaskErrorMock()
	service := NewService(mock)

	request := &RequestLogFlushingTask{
		Qualifier:           "JAMFSoftwareServer",
		RetentionPeriod:     30,
		RetentionPeriodUnit: "Days",
	}

	result, resp, err := service.QueueTaskV1(context.Background(), request)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to queue log flushing task")
}

func TestUnit_LogFlushing_DeleteTaskByIDV1_Error(t *testing.T) {
	mock := mocks.NewLogFlushingMock()
	mock.RegisterDeleteTaskErrorMock("1")
	service := NewService(mock)

	resp, err := service.DeleteTaskByIDV1(context.Background(), "1")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to delete log flushing task 1")
}
