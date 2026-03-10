package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type VPPAccountsMock struct {
	*mocks.GenericMock
}

func NewVPPAccountsMock() *VPPAccountsMock {
	return &VPPAccountsMock{
		GenericMock: mocks.NewXMLMock("VPPAccountsMock"),
	}
}

func (m *VPPAccountsMock) RegisterMocks() {
	m.RegisterListVPPAccountsMock()
	m.RegisterGetVPPAccountByIDMock()
	m.RegisterCreateVPPAccountMock()
	m.RegisterUpdateVPPAccountByIDMock()
	m.RegisterDeleteVPPAccountByIDMock()
}

func (m *VPPAccountsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *VPPAccountsMock) RegisterListVPPAccountsMock() {
	m.Register("GET", "/JSSResource/vppaccounts", 200, "validate_list_vpp_accounts.xml")
}

func (m *VPPAccountsMock) RegisterGetVPPAccountByIDMock() {
	m.Register("GET", "/JSSResource/vppaccounts/id/1", 200, "validate_get_vpp_account.xml")
}

func (m *VPPAccountsMock) RegisterCreateVPPAccountMock() {
	m.Register("POST", "/JSSResource/vppaccounts/id/0", 201, "validate_create_vpp_account.xml")
}

func (m *VPPAccountsMock) RegisterUpdateVPPAccountByIDMock() {
	m.Register("PUT", "/JSSResource/vppaccounts/id/1", 200, "validate_update_vpp_account.xml")
}

func (m *VPPAccountsMock) RegisterDeleteVPPAccountByIDMock() {
	m.Register("DELETE", "/JSSResource/vppaccounts/id/1", 200, "")
}

func (m *VPPAccountsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/vppaccounts/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *VPPAccountsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/vppaccounts/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A VPP account with that name already exists")
}

