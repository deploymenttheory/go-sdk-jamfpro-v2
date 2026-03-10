package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type VPPAssignmentsMock struct {
	*mocks.GenericMock
}

func NewVPPAssignmentsMock() *VPPAssignmentsMock {
	return &VPPAssignmentsMock{
		GenericMock: mocks.NewXMLMock("VPPAssignmentsMock"),
	}
}

func (m *VPPAssignmentsMock) RegisterMocks() {
	m.RegisterListVPPAssignmentsMock()
	m.RegisterGetVPPAssignmentByIDMock()
	m.RegisterCreateVPPAssignmentMock()
	m.RegisterUpdateVPPAssignmentByIDMock()
	m.RegisterDeleteVPPAssignmentByIDMock()
}

func (m *VPPAssignmentsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *VPPAssignmentsMock) RegisterListVPPAssignmentsMock() {
	m.Register("GET", "/JSSResource/vppassignments", 200, "validate_list_vpp_assignments.xml")
}

func (m *VPPAssignmentsMock) RegisterGetVPPAssignmentByIDMock() {
	m.Register("GET", "/JSSResource/vppassignments/id/1", 200, "validate_get_vpp_assignment.xml")
}

func (m *VPPAssignmentsMock) RegisterCreateVPPAssignmentMock() {
	m.Register("POST", "/JSSResource/vppassignments/id/0", 201, "validate_create_vpp_assignment.xml")
}

func (m *VPPAssignmentsMock) RegisterUpdateVPPAssignmentByIDMock() {
	m.Register("PUT", "/JSSResource/vppassignments/id/1", 200, "validate_update_vpp_assignment.xml")
}

func (m *VPPAssignmentsMock) RegisterDeleteVPPAssignmentByIDMock() {
	m.Register("DELETE", "/JSSResource/vppassignments/id/1", 200, "")
}

func (m *VPPAssignmentsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/vppassignments/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *VPPAssignmentsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/vppassignments/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A VPP assignment with that name already exists")
}

