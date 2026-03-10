package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type DockItemsMock struct {
	*mocks.GenericMock
}

func NewDockItemsMock() *DockItemsMock {
	return &DockItemsMock{
		GenericMock: mocks.NewXMLMock("DockItemsMock"),
	}
}

func (m *DockItemsMock) RegisterMocks() {
	m.RegisterListDockItemsMock()
	m.RegisterGetDockItemByIDMock()
	m.RegisterGetDockItemByNameMock()
	m.RegisterCreateDockItemMock()
	m.RegisterUpdateDockItemByIDMock()
	m.RegisterUpdateDockItemByNameMock()
	m.RegisterDeleteDockItemByIDMock()
	m.RegisterDeleteDockItemByNameMock()
}

func (m *DockItemsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *DockItemsMock) RegisterListDockItemsMock() {
	m.Register("GET", "/JSSResource/dockitems", 200, "validate_list_dock_items.xml")
}

func (m *DockItemsMock) RegisterGetDockItemByIDMock() {
	m.Register("GET", "/JSSResource/dockitems/id/1", 200, "validate_get_dock_item.xml")
}

func (m *DockItemsMock) RegisterGetDockItemByNameMock() {
	m.Register("GET", "/JSSResource/dockitems/name/Safari", 200, "validate_get_dock_item.xml")
}

func (m *DockItemsMock) RegisterCreateDockItemMock() {
	m.Register("POST", "/JSSResource/dockitems/id/0", 201, "validate_create_dock_item.xml")
}

func (m *DockItemsMock) RegisterUpdateDockItemByIDMock() {
	m.Register("PUT", "/JSSResource/dockitems/id/1", 200, "validate_update_dock_item.xml")
}

func (m *DockItemsMock) RegisterUpdateDockItemByNameMock() {
	m.Register("PUT", "/JSSResource/dockitems/name/Safari", 200, "validate_update_dock_item.xml")
}

func (m *DockItemsMock) RegisterDeleteDockItemByIDMock() {
	m.Register("DELETE", "/JSSResource/dockitems/id/1", 200, "")
}

func (m *DockItemsMock) RegisterDeleteDockItemByNameMock() {
	m.Register("DELETE", "/JSSResource/dockitems/name/Safari", 200, "")
}

func (m *DockItemsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/dockitems/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *DockItemsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/dockitems/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A dock item with that name already exists")
}

