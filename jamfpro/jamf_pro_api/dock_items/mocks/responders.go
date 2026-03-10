package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type DockItemsMock struct {
	*mocks.GenericMock
}

func NewDockItemsMock() *DockItemsMock {
	return &DockItemsMock{
		GenericMock: mocks.NewJSONMock("DockItemsMock"),
	}
}

func (m *DockItemsMock) RegisterMocks() {
	m.RegisterGetDockItemMock()
	m.RegisterCreateDockItemMock()
	m.RegisterUpdateDockItemMock()
	m.RegisterDeleteDockItemMock()
}

func (m *DockItemsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *DockItemsMock) RegisterGetDockItemMock() {
	m.Register("GET", "/api/v1/dock-items/1", 200, "validate_get_dock_item.json")
}

func (m *DockItemsMock) RegisterCreateDockItemMock() {
	m.Register("POST", "/api/v1/dock-items", 201, "validate_create_dock_item.json")
}

func (m *DockItemsMock) RegisterUpdateDockItemMock() {
	m.Register("PUT", "/api/v1/dock-items/1", 200, "validate_update_dock_item.json")
}

func (m *DockItemsMock) RegisterDeleteDockItemMock() {
	m.Register("DELETE", "/api/v1/dock-items/1", 204, "")
}

func (m *DockItemsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/dock-items/999", 404, "error_not_found.json", "")
}

func (m *DockItemsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/api/v1/dock-items", 409, "error_conflict.json", "")
}

