package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type CommandFlushMock struct {
	*mocks.GenericMock
}

func NewCommandFlushMock() *CommandFlushMock {
	return &CommandFlushMock{
		GenericMock: mocks.NewXMLMock("CommandFlushMock"),
	}
}

func (m *CommandFlushMock) RegisterMocks() {
	// Computers - all status combinations
	m.RegisterFlushComputersPendingMock()
	m.RegisterFlushComputersFailedMock()
	m.RegisterFlushComputersPendingAndFailedMock()

	// Computer groups - all status combinations
	m.RegisterFlushComputerGroupsPendingMock()
	m.RegisterFlushComputerGroupsFailedMock()
	m.RegisterFlushComputerGroupsPendingAndFailedMock()

	// Mobile devices - all status combinations
	m.RegisterFlushMobileDevicesPendingMock()
	m.RegisterFlushMobileDevicesFailedMock()
	m.RegisterFlushMobileDevicesPendingAndFailedMock()

	// Mobile device groups - all status combinations
	m.RegisterFlushMobileDeviceGroupsPendingMock()
	m.RegisterFlushMobileDeviceGroupsFailedMock()
	m.RegisterFlushMobileDeviceGroupsPendingAndFailedMock()

	// XML batch operation
	m.RegisterFlushWithXMLMock()
}

func (m *CommandFlushMock) RegisterFlushComputersPendingMock() {
	m.Register("DELETE", "/JSSResource/commandflush/computers/id/123/status/Pending", 204, "")
}

func (m *CommandFlushMock) RegisterFlushComputersFailedMock() {
	m.Register("DELETE", "/JSSResource/commandflush/computers/id/123/status/Failed", 204, "")
}

func (m *CommandFlushMock) RegisterFlushComputersPendingAndFailedMock() {
	m.Register("DELETE", "/JSSResource/commandflush/computers/id/123/status/Pending%2BFailed", 204, "")
}

func (m *CommandFlushMock) RegisterFlushComputerGroupsPendingMock() {
	m.Register("DELETE", "/JSSResource/commandflush/computergroups/id/456/status/Pending", 204, "")
}

func (m *CommandFlushMock) RegisterFlushComputerGroupsFailedMock() {
	m.Register("DELETE", "/JSSResource/commandflush/computergroups/id/456/status/Failed", 204, "")
}

func (m *CommandFlushMock) RegisterFlushComputerGroupsPendingAndFailedMock() {
	m.Register("DELETE", "/JSSResource/commandflush/computergroups/id/456/status/Pending%2BFailed", 204, "")
}

func (m *CommandFlushMock) RegisterFlushMobileDevicesPendingMock() {
	m.Register("DELETE", "/JSSResource/commandflush/mobiledevices/id/789/status/Pending", 204, "")
}

func (m *CommandFlushMock) RegisterFlushMobileDevicesFailedMock() {
	m.Register("DELETE", "/JSSResource/commandflush/mobiledevices/id/789/status/Failed", 204, "")
}

func (m *CommandFlushMock) RegisterFlushMobileDevicesPendingAndFailedMock() {
	m.Register("DELETE", "/JSSResource/commandflush/mobiledevices/id/789/status/Pending%2BFailed", 204, "")
}

func (m *CommandFlushMock) RegisterFlushMobileDeviceGroupsPendingMock() {
	m.Register("DELETE", "/JSSResource/commandflush/mobiledevicegroups/id/101112/status/Pending", 204, "")
}

func (m *CommandFlushMock) RegisterFlushMobileDeviceGroupsFailedMock() {
	m.Register("DELETE", "/JSSResource/commandflush/mobiledevicegroups/id/101112/status/Failed", 204, "")
}

func (m *CommandFlushMock) RegisterFlushMobileDeviceGroupsPendingAndFailedMock() {
	m.Register("DELETE", "/JSSResource/commandflush/mobiledevicegroups/id/101112/status/Pending%2BFailed", 204, "")
}

func (m *CommandFlushMock) RegisterFlushWithXMLMock() {
	m.Register("DELETE", "/JSSResource/commandflush", 204, "")
}

