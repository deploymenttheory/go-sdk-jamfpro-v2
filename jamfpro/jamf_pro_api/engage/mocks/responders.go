package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type EngageMock struct {
	*mocks.GenericMock
}

func NewEngageMock() *EngageMock {
	return &EngageMock{
		GenericMock: mocks.NewJSONMock("EngageMock"),
	}
}

func (m *EngageMock) RegisterGetMock() {
	m.Register("GET", "/api/v2/engage", 200, "validate_get.json")
}

func (m *EngageMock) RegisterGetInvalidJSONMock() {
	m.Register("GET", "/api/v2/engage", 200, "validate_get_invalid.json")
}

func (m *EngageMock) RegisterUpdateMock() {
	m.Register("PUT", "/api/v2/engage", 200, "validate_update.json")
}

func (m *EngageMock) RegisterGetHistoryMock() {
	m.Register("GET", "/api/v2/engage/history", 200, "validate_history.json")
}

func (m *EngageMock) RegisterAddHistoryNotesMock() {
	m.Register("POST", "/api/v2/engage/history", 201, "validate_add_history_notes.json")
}
