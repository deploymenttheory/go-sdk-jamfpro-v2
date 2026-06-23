package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type MobileDevicesMock struct {
	*mocks.GenericMock
}

func NewMobileDevicesMock() *MobileDevicesMock {
	return &MobileDevicesMock{
		GenericMock: mocks.NewJSONMock("MobileDevicesMock"),
	}
}

func (m *MobileDevicesMock) RegisterListMock() {
	m.Register("GET", "/api/v2/mobile-devices", 200, "validate_list.json")
}

func (m *MobileDevicesMock) RegisterGetByIDMock(id string) {
	m.Register("GET", "/api/v2/mobile-devices/"+id, 200, "validate_get.json")
}

func (m *MobileDevicesMock) RegisterGetDetailMock() {
	m.Register("GET", "/api/v2/mobile-devices/detail", 200, "validate_detail_list.json")
}

func (m *MobileDevicesMock) RegisterGetDetailByIDMock(id string) {
	m.Register("GET", "/api/v2/mobile-devices/"+id+"/detail", 200, "validate_detail_get.json")
}

func (m *MobileDevicesMock) RegisterGetPairedDevicesByIDMock(id string) {
	m.Register("GET", "/api/v2/mobile-devices/"+id+"/paired-devices", 200, "validate_detail_list.json")
}

func (m *MobileDevicesMock) RegisterGetByIDNotFoundMock(id string) {
	m.RegisterNotFoundError("GET", "/api/v2/mobile-devices/"+id)
}

func (m *MobileDevicesMock) RegisterListErrorMock() {
	m.RegisterError("GET", "/api/v2/mobile-devices", 500, "", "simulated ListV2 API error")
}

func (m *MobileDevicesMock) RegisterGetDetailErrorMock() {
	m.RegisterError("GET", "/api/v2/mobile-devices/detail", 500, "", "simulated GetDetailV2 API error")
}
