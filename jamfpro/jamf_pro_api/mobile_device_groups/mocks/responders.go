package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type MobileDeviceGroupsMock struct {
	*mocks.GenericMock
}

func NewMobileDeviceGroupsMock() *MobileDeviceGroupsMock {
	return &MobileDeviceGroupsMock{
		GenericMock: mocks.NewJSONMock("MobileDeviceGroupsMock"),
	}
}

func (m *MobileDeviceGroupsMock) RegisterListSmartMock() {
	m.Register("GET", "/api/v1/mobile-device-groups/smart-groups", 200, "validate_list_smart_groups.json")
}

func (m *MobileDeviceGroupsMock) RegisterGetSmartMock() {
	m.Register("GET", "/api/v1/mobile-device-groups/smart-groups/1", 200, "validate_get_smart_group.json")
}

func (m *MobileDeviceGroupsMock) RegisterCreateSmartMock() {
	m.Register("POST", "/api/v1/mobile-device-groups/smart-groups", 201, "validate_create_smart_group.json")
}

func (m *MobileDeviceGroupsMock) RegisterUpdateSmartMock() {
	m.Register("PUT", "/api/v1/mobile-device-groups/smart-groups/1", 200, "validate_update_smart_group.json")
}

func (m *MobileDeviceGroupsMock) RegisterDeleteSmartMock() {
	m.Register("DELETE", "/api/v1/mobile-device-groups/smart-groups/1", 204, "")
}

func (m *MobileDeviceGroupsMock) RegisterListStaticMock() {
	m.Register("GET", "/api/v1/mobile-device-groups/static-groups", 200, "validate_list_static_groups.json")
}

func (m *MobileDeviceGroupsMock) RegisterGetStaticMock() {
	m.Register("GET", "/api/v1/mobile-device-groups/static-groups/10", 200, "validate_get_static_group.json")
}

func (m *MobileDeviceGroupsMock) RegisterCreateStaticMock() {
	m.Register("POST", "/api/v1/mobile-device-groups/static-groups", 201, "validate_create_static_group.json")
}

func (m *MobileDeviceGroupsMock) RegisterUpdateStaticMock() {
	m.Register("PATCH", "/api/v1/mobile-device-groups/static-groups/10", 200, "validate_update_static_group.json")
}

func (m *MobileDeviceGroupsMock) RegisterDeleteStaticMock() {
	m.Register("DELETE", "/api/v1/mobile-device-groups/static-groups/10", 204, "")
}

func (m *MobileDeviceGroupsMock) RegisterListAllMock() {
	m.Register("GET", "/api/v1/mobile-device-groups", 200, "validate_list_all.json")
}

func (m *MobileDeviceGroupsMock) RegisterGetStaticGroupMembershipMock() {
	m.Register("GET", "/api/v1/mobile-device-groups/static-group-membership/10", 200, "validate_get_static_group_membership.json")
}

func (m *MobileDeviceGroupsMock) RegisterGetSmartGroupMembershipMock() {
	m.Register("GET", "/api/v1/mobile-device-groups/smart-group-membership/1", 200, "validate_get_smart_group_membership.json")
}

func (m *MobileDeviceGroupsMock) RegisterEraseDevicesMock() {
	m.Register("POST", "/api/v1/mobile-device-groups/1/erase", 200, "")
}

func (m *MobileDeviceGroupsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/mobile-device-groups/smart-groups/999", 404, "error_not_found.json", "")
}

func (m *MobileDeviceGroupsMock) RegisterGetStaticNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/mobile-device-groups/static-groups/999", 404, "error_not_found.json", "")
}

func (m *MobileDeviceGroupsMock) RegisterListSmartErrorMock() {
	m.RegisterError("GET", "/api/v1/mobile-device-groups/smart-groups", 500, "error_not_found.json", "")
}

func (m *MobileDeviceGroupsMock) RegisterListStaticErrorMock() {
	m.RegisterError("GET", "/api/v1/mobile-device-groups/static-groups", 500, "error_not_found.json", "")
}

func (m *MobileDeviceGroupsMock) RegisterListAllErrorMock() {
	m.RegisterError("GET", "/api/v1/mobile-device-groups", 500, "error_not_found.json", "")
}

func (m *MobileDeviceGroupsMock) RegisterGetStaticGroupMembershipErrorMock() {
	m.RegisterError("GET", "/api/v1/mobile-device-groups/static-group-membership/10", 500, "error_not_found.json", "")
}

func (m *MobileDeviceGroupsMock) RegisterGetSmartGroupMembershipErrorMock() {
	m.RegisterError("GET", "/api/v1/mobile-device-groups/smart-group-membership/1", 500, "error_not_found.json", "")
}

func (m *MobileDeviceGroupsMock) RegisterEraseDevicesErrorMock() {
	m.RegisterError("POST", "/api/v1/mobile-device-groups/1/erase", 500, "error_not_found.json", "")
}

func (m *MobileDeviceGroupsMock) RegisterListSmartNoResponseErrorMock() {
	m.RegisterError("GET", "/api/v1/mobile-device-groups/smart-groups", 500, "error_internal.json", "no response registered")
}

// -----------------------------------------------------------------------------
// V2 registrations (Jamf Pro 11.28 mobile-device-groups v2 surface)
// -----------------------------------------------------------------------------

func (m *MobileDeviceGroupsMock) RegisterListSmartV2Mock() {
	m.Register("GET", "/api/v2/mobile-device-groups/smart-groups", 200, "validate_list_smart_groups.json")
}

func (m *MobileDeviceGroupsMock) RegisterGetSmartV2Mock() {
	m.Register("GET", "/api/v2/mobile-device-groups/smart-groups/1", 200, "validate_get_smart_group.json")
}

func (m *MobileDeviceGroupsMock) RegisterCreateSmartV2Mock() {
	m.Register("POST", "/api/v2/mobile-device-groups/smart-groups", 201, "validate_create_smart_group.json")
}

func (m *MobileDeviceGroupsMock) RegisterUpdateSmartV2Mock() {
	m.Register("PUT", "/api/v2/mobile-device-groups/smart-groups/1", 200, "validate_update_smart_group.json")
}

func (m *MobileDeviceGroupsMock) RegisterDeleteSmartV2Mock() {
	m.Register("DELETE", "/api/v2/mobile-device-groups/smart-groups/1", 204, "")
}

func (m *MobileDeviceGroupsMock) RegisterListStaticV2Mock() {
	m.Register("GET", "/api/v2/mobile-device-groups/static-groups", 200, "validate_list_static_groups.json")
}

func (m *MobileDeviceGroupsMock) RegisterGetStaticV2Mock() {
	m.Register("GET", "/api/v2/mobile-device-groups/static-groups/10", 200, "validate_get_static_group.json")
}

func (m *MobileDeviceGroupsMock) RegisterCreateStaticV2Mock() {
	m.Register("POST", "/api/v2/mobile-device-groups/static-groups", 201, "validate_create_static_group.json")
}

func (m *MobileDeviceGroupsMock) RegisterUpdateStaticV2Mock() {
	m.Register("PATCH", "/api/v2/mobile-device-groups/static-groups/10", 200, "validate_update_static_group.json")
}

func (m *MobileDeviceGroupsMock) RegisterDeleteStaticV2Mock() {
	m.Register("DELETE", "/api/v2/mobile-device-groups/static-groups/10", 204, "")
}

func (m *MobileDeviceGroupsMock) RegisterListAllV2Mock() {
	m.Register("GET", "/api/v2/mobile-device-groups", 200, "validate_list_all.json")
}

func (m *MobileDeviceGroupsMock) RegisterGetStaticGroupMembershipV2Mock() {
	m.Register("GET", "/api/v2/mobile-device-groups/static-group-membership/10", 200, "validate_get_static_group_membership_v2.json")
}

func (m *MobileDeviceGroupsMock) RegisterGetSmartGroupMembershipV2Mock() {
	m.Register("GET", "/api/v2/mobile-device-groups/smart-group-membership/1", 200, "validate_get_smart_group_membership.json")
}

func (m *MobileDeviceGroupsMock) RegisterEraseDevicesV2Mock() {
	m.Register("POST", "/api/v2/mobile-device-groups/1/erase", 200, "")
}

func (m *MobileDeviceGroupsMock) RegisterGetSmartV2NotFoundErrorMock() {
	m.RegisterError("GET", "/api/v2/mobile-device-groups/smart-groups/999", 404, "error_not_found.json", "")
}

func (m *MobileDeviceGroupsMock) RegisterListStaticV2ErrorMock() {
	m.RegisterError("GET", "/api/v2/mobile-device-groups/static-groups", 500, "error_not_found.json", "")
}
