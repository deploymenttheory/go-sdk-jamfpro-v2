package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type LogFlushingMock struct {
	*mocks.GenericMock
}

func NewLogFlushingMock() *LogFlushingMock {
	return &LogFlushingMock{
		GenericMock: mocks.NewJSONMock("LogFlushingMock"),
	}
}

func (m *LogFlushingMock) RegisterGetSettingsMock() {
	m.Register("GET", "/api/v1/log-flushing", 200, "validate_settings.json")
}

func (m *LogFlushingMock) RegisterGetSettingsErrorMock() {
	m.RegisterError("GET", "/api/v1/log-flushing", 500, "", "api error")
}

func (m *LogFlushingMock) RegisterListTasksMock() {
	m.Register("GET", "/api/v1/log-flushing/task", 200, "validate_tasks_list.json")
}

func (m *LogFlushingMock) RegisterListTasksErrorMock() {
	m.RegisterError("GET", "/api/v1/log-flushing/task", 500, "", "api error")
}

func (m *LogFlushingMock) RegisterGetTaskByIDMock() {
	m.Register("GET", "/api/v1/log-flushing/task/1", 200, "validate_task_get.json")
}

func (m *LogFlushingMock) RegisterGetTaskByIDErrorMock(id string) {
	m.RegisterError("GET", "/api/v1/log-flushing/task/"+id, 500, "", "api error")
}

func (m *LogFlushingMock) RegisterQueueTaskMock() {
	m.Register("POST", "/api/v1/log-flushing/task", 201, "validate_queue_task.json")
}

func (m *LogFlushingMock) RegisterQueueTaskErrorMock() {
	m.RegisterError("POST", "/api/v1/log-flushing/task", 500, "", "api error")
}

func (m *LogFlushingMock) RegisterDeleteTaskMock() {
	m.Register("DELETE", "/api/v1/log-flushing/task/1", 204, "")
}

func (m *LogFlushingMock) RegisterDeleteTaskErrorMock(id string) {
	m.RegisterError("DELETE", "/api/v1/log-flushing/task/"+id, 500, "", "api error")
}
