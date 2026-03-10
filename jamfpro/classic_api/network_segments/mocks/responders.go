package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type NetworkSegmentsMock struct {
	*mocks.GenericMock
}

func NewNetworkSegmentsMock() *NetworkSegmentsMock {
	return &NetworkSegmentsMock{
		GenericMock: mocks.NewXMLMock("NetworkSegmentsMock"),
	}
}

func (m *NetworkSegmentsMock) RegisterMocks() {
	m.RegisterListNetworkSegmentsMock()
	m.RegisterGetNetworkSegmentByIDMock()
	m.RegisterGetNetworkSegmentByNameMock()
	m.RegisterCreateNetworkSegmentMock()
	m.RegisterUpdateNetworkSegmentByIDMock()
	m.RegisterUpdateNetworkSegmentByNameMock()
	m.RegisterDeleteNetworkSegmentByIDMock()
	m.RegisterDeleteNetworkSegmentByNameMock()
}

func (m *NetworkSegmentsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *NetworkSegmentsMock) RegisterListNetworkSegmentsMock() {
	m.Register("GET", "/JSSResource/networksegments", 200, "validate_list_network_segments.xml")
}

func (m *NetworkSegmentsMock) RegisterGetNetworkSegmentByIDMock() {
	m.Register("GET", "/JSSResource/networksegments/id/1", 200, "validate_get_network_segment.xml")
}

func (m *NetworkSegmentsMock) RegisterGetNetworkSegmentByNameMock() {
	m.Register("GET", "/JSSResource/networksegments/name/HQ Network", 200, "validate_get_network_segment.xml")
}

func (m *NetworkSegmentsMock) RegisterCreateNetworkSegmentMock() {
	m.Register("POST", "/JSSResource/networksegments/id/0", 201, "validate_create_network_segment.xml")
}

func (m *NetworkSegmentsMock) RegisterUpdateNetworkSegmentByIDMock() {
	m.Register("PUT", "/JSSResource/networksegments/id/1", 200, "validate_update_network_segment.xml")
}

func (m *NetworkSegmentsMock) RegisterUpdateNetworkSegmentByNameMock() {
	m.Register("PUT", "/JSSResource/networksegments/name/HQ Network", 200, "validate_update_network_segment.xml")
}

func (m *NetworkSegmentsMock) RegisterDeleteNetworkSegmentByIDMock() {
	m.Register("DELETE", "/JSSResource/networksegments/id/1", 200, "")
}

func (m *NetworkSegmentsMock) RegisterDeleteNetworkSegmentByNameMock() {
	m.Register("DELETE", "/JSSResource/networksegments/name/HQ Network", 200, "")
}

func (m *NetworkSegmentsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/networksegments/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *NetworkSegmentsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/networksegments/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A network segment with that name already exists")
}

