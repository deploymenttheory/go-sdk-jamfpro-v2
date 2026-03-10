package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ComputerHistoryMock struct {
	*mocks.GenericMock
}

func NewComputerHistoryMock() *ComputerHistoryMock {
	return &ComputerHistoryMock{
		GenericMock: mocks.NewXMLMock("ComputerHistoryMock"),
	}
}

func (m *ComputerHistoryMock) RegisterMocks() {
	m.RegisterGetByIDMock()
	m.RegisterGetByIDAndSubsetMock()
	m.RegisterGetByNameMock()
	m.RegisterGetByNameAndSubsetMock()
	m.RegisterGetByUDIDMock()
	m.RegisterGetByUDIDAndSubsetMock()
	m.RegisterGetBySerialNumberMock()
	m.RegisterGetBySerialNumberAndSubsetMock()
	m.RegisterGetByMACAddressMock()
	m.RegisterGetByMACAddressAndSubsetMock()
}

func (m *ComputerHistoryMock) RegisterGetByIDMock() {
	m.Register("GET", "/JSSResource/computerhistory/id/1", 200, "validate_get_computer_history.xml")
}

func (m *ComputerHistoryMock) RegisterGetByIDAndSubsetMock() {
	m.Register("GET", "/JSSResource/computerhistory/id/1/subset/General", 200, "validate_get_computer_history.xml")
}

func (m *ComputerHistoryMock) RegisterGetByNameMock() {
	m.Register("GET", "/JSSResource/computerhistory/name/Test-MacBook-Pro", 200, "validate_get_computer_history.xml")
}

func (m *ComputerHistoryMock) RegisterGetByNameAndSubsetMock() {
	m.Register("GET", "/JSSResource/computerhistory/name/Test-MacBook-Pro/subset/General", 200, "validate_get_computer_history.xml")
}

func (m *ComputerHistoryMock) RegisterGetByUDIDMock() {
	m.Register("GET", "/JSSResource/computerhistory/udid/00000000-0000-0000-0000-000000000001", 200, "validate_get_computer_history.xml")
}

func (m *ComputerHistoryMock) RegisterGetByUDIDAndSubsetMock() {
	m.Register("GET", "/JSSResource/computerhistory/udid/00000000-0000-0000-0000-000000000001/subset/General", 200, "validate_get_computer_history.xml")
}

func (m *ComputerHistoryMock) RegisterGetBySerialNumberMock() {
	m.Register("GET", "/JSSResource/computerhistory/serialnumber/C02XYZ123456", 200, "validate_get_computer_history.xml")
}

func (m *ComputerHistoryMock) RegisterGetBySerialNumberAndSubsetMock() {
	m.Register("GET", "/JSSResource/computerhistory/serialnumber/C02XYZ123456/subset/General", 200, "validate_get_computer_history.xml")
}

func (m *ComputerHistoryMock) RegisterGetByMACAddressMock() {
	m.Register("GET", "/JSSResource/computerhistory/macaddress/00:11:22:33:44:55", 200, "validate_get_computer_history.xml")
}

func (m *ComputerHistoryMock) RegisterGetByMACAddressAndSubsetMock() {
	m.Register("GET", "/JSSResource/computerhistory/macaddress/00:11:22:33:44:55/subset/General", 200, "validate_get_computer_history.xml")
}

func (m *ComputerHistoryMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/computerhistory/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

