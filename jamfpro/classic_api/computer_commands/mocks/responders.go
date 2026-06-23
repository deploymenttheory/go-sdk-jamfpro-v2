package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

// ComputerCommandsMock is a test double for the computer commands Classic API.
type ComputerCommandsMock struct {
	*mocks.GenericMock
}

// NewComputerCommandsMock returns a ComputerCommandsMock configured for XML responses.
func NewComputerCommandsMock() *ComputerCommandsMock {
	return &ComputerCommandsMock{GenericMock: mocks.NewXMLMock("ComputerCommandsMock")}
}

// RegisterSendCommandMock registers a successful 201 response for sending the
// given command to the given id list (comma-joined).
func (m *ComputerCommandsMock) RegisterSendCommandMock(command, ids string) {
	m.Register("POST", "/JSSResource/computercommands/command/"+command+"/id/"+ids, 201, "")
}
