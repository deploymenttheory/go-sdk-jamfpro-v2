package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type RemoveableMacAddressesMock struct {
	*mocks.GenericMock
}

func NewRemoveableMacAddressesMock() *RemoveableMacAddressesMock {
	return &RemoveableMacAddressesMock{
		GenericMock: mocks.NewXMLMock("RemoveableMacAddressesMock"),
	}
}

func (m *RemoveableMacAddressesMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterGetByNameMock()
	m.RegisterCreateMock()
	m.RegisterUpdateByIDMock()
	m.RegisterUpdateByNameMock()
	m.RegisterDeleteByIDMock()
	m.RegisterDeleteByNameMock()
}

func (m *RemoveableMacAddressesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *RemoveableMacAddressesMock) RegisterListMock() {
	m.Register("GET", "/JSSResource/removablemacaddresses", 200, "validate_list_removeable_mac_addresses.xml")
}

func (m *RemoveableMacAddressesMock) RegisterGetByIDMock() {
	m.Register("GET", "/JSSResource/removablemacaddresses/id/1", 200, "validate_get_removeable_mac_address.xml")
}

func (m *RemoveableMacAddressesMock) RegisterGetByNameMock() {
	m.Register("GET", "/JSSResource/removablemacaddresses/name/AA:BB:CC:DD:EE:FF", 200, "validate_get_removeable_mac_address.xml")
}

func (m *RemoveableMacAddressesMock) RegisterCreateMock() {
	m.Register("POST", "/JSSResource/removablemacaddresses/id/0", 201, "validate_create_removeable_mac_address.xml")
}

func (m *RemoveableMacAddressesMock) RegisterUpdateByIDMock() {
	m.Register("PUT", "/JSSResource/removablemacaddresses/id/1", 200, "validate_update_removeable_mac_address.xml")
}

func (m *RemoveableMacAddressesMock) RegisterUpdateByNameMock() {
	m.Register("PUT", "/JSSResource/removablemacaddresses/name/AA:BB:CC:DD:EE:FF", 200, "validate_update_removeable_mac_address.xml")
}

func (m *RemoveableMacAddressesMock) RegisterDeleteByIDMock() {
	m.Register("DELETE", "/JSSResource/removablemacaddresses/id/1", 200, "")
}

func (m *RemoveableMacAddressesMock) RegisterDeleteByNameMock() {
	m.Register("DELETE", "/JSSResource/removablemacaddresses/name/AA:BB:CC:DD:EE:FF", 200, "")
}

func (m *RemoveableMacAddressesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/removablemacaddresses/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *RemoveableMacAddressesMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/removablemacaddresses/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A removeable mac address with that name already exists")
}
