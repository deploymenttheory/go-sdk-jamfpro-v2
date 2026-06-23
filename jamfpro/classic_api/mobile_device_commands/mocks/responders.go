package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

// MobileDeviceCommandsMock is a test double for the mobile device commands Classic API.
type MobileDeviceCommandsMock struct {
	*mocks.GenericMock
}

// NewMobileDeviceCommandsMock returns a MobileDeviceCommandsMock configured for XML responses.
func NewMobileDeviceCommandsMock() *MobileDeviceCommandsMock {
	return &MobileDeviceCommandsMock{GenericMock: mocks.NewXMLMock("MobileDeviceCommandsMock")}
}

// RegisterSendCommandMock registers a successful 201 response for sending the
// given command to the given id list (comma-joined).
func (m *MobileDeviceCommandsMock) RegisterSendCommandMock(command, ids string) {
	m.Register("POST", "/JSSResource/mobiledevicecommands/command/"+command+"/id/"+ids, 201, "")
}
