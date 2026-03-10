package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ComputerInvitationsMock struct {
	*mocks.GenericMock
}

func NewComputerInvitationsMock() *ComputerInvitationsMock {
	return &ComputerInvitationsMock{
		GenericMock: mocks.NewXMLMock("ComputerInvitationsMock"),
	}
}

func (m *ComputerInvitationsMock) RegisterMocks() {
	m.RegisterListComputerInvitationsMock()
	m.RegisterGetComputerInvitationByIDMock()
	m.RegisterGetComputerInvitationByInvitationIDMock()
	m.RegisterCreateComputerInvitationMock()
	m.RegisterDeleteComputerInvitationByIDMock()
}

func (m *ComputerInvitationsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
}

func (m *ComputerInvitationsMock) RegisterListComputerInvitationsMock() {
	m.Register("GET", "/JSSResource/computerinvitations", 200, "validate_list_computer_invitations.xml")
}

func (m *ComputerInvitationsMock) RegisterGetComputerInvitationByIDMock() {
	m.Register("GET", "/JSSResource/computerinvitations/id/1", 200, "validate_get_computer_invitation.xml")
}

func (m *ComputerInvitationsMock) RegisterGetComputerInvitationByInvitationIDMock() {
	m.Register("GET", "/JSSResource/computerinvitations/invitation/1234567890", 200, "validate_get_computer_invitation.xml")
}

func (m *ComputerInvitationsMock) RegisterCreateComputerInvitationMock() {
	m.Register("POST", "/JSSResource/computerinvitations/id/0", 201, "validate_create_computer_invitation.xml")
}

func (m *ComputerInvitationsMock) RegisterDeleteComputerInvitationByIDMock() {
	m.Register("DELETE", "/JSSResource/computerinvitations/id/1", 200, "")
}

func (m *ComputerInvitationsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/computerinvitations/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

