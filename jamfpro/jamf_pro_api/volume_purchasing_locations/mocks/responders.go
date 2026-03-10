package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type VolumePurchasingLocationsMock struct {
	*mocks.GenericMock
}

func NewVolumePurchasingLocationsMock() *VolumePurchasingLocationsMock {
	return &VolumePurchasingLocationsMock{
		GenericMock: mocks.NewJSONMock("VolumePurchasingLocationsMock"),
	}
}

func (m *VolumePurchasingLocationsMock) RegisterListMock() {
	m.Register("GET", "/api/v1/volume-purchasing-locations", 200, "validate_list.json")
}

func (m *VolumePurchasingLocationsMock) RegisterListInvalidJSONMock() {
	m.Register("GET", "/api/v1/volume-purchasing-locations", 200, "validate_list_invalid.json")
}

func (m *VolumePurchasingLocationsMock) RegisterGetMock() {
	m.Register("GET", "/api/v1/volume-purchasing-locations/1", 200, "validate_get.json")
}

func (m *VolumePurchasingLocationsMock) RegisterCreateMock() {
	m.Register("POST", "/api/v1/volume-purchasing-locations", 201, "validate_create.json")
}

func (m *VolumePurchasingLocationsMock) RegisterUpdateMock() {
	m.Register("PATCH", "/api/v1/volume-purchasing-locations/1", 200, "validate_update.json")
}

func (m *VolumePurchasingLocationsMock) RegisterDeleteMock() {
	m.Register("DELETE", "/api/v1/volume-purchasing-locations/1", 204, "")
}

func (m *VolumePurchasingLocationsMock) RegisterReclaimMock() {
	m.Register("POST", "/api/v1/volume-purchasing-locations/1/reclaim", 202, "")
}

func (m *VolumePurchasingLocationsMock) RegisterGetContentMock() {
	m.Register("GET", "/api/v1/volume-purchasing-locations/1/content", 200, "validate_content.json")
}

func (m *VolumePurchasingLocationsMock) RegisterGetContentInvalidJSONMock() {
	m.Register("GET", "/api/v1/volume-purchasing-locations/99/content", 200, "validate_list_invalid.json")
}

func (m *VolumePurchasingLocationsMock) RegisterGetHistoryInvalidJSONMock() {
	m.Register("GET", "/api/v1/volume-purchasing-locations/99/history", 200, "validate_list_invalid.json")
}

func (m *VolumePurchasingLocationsMock) RegisterGetContentWithResultsMock() {
	m.Register("GET", "/api/v1/volume-purchasing-locations/2/content", 200, "validate_content_with_results.json")
}

func (m *VolumePurchasingLocationsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/volume-purchasing-locations/999", 404, "error_not_found.json", "")
}

func (m *VolumePurchasingLocationsMock) RegisterGetHistoryMock() {
	m.Register("GET", "/api/v1/volume-purchasing-locations/1/history", 200, "validate_history.json")
}

func (m *VolumePurchasingLocationsMock) RegisterAddHistoryNotesMock() {
	m.Register("POST", "/api/v1/volume-purchasing-locations/1/history", 201, "")
}

func (m *VolumePurchasingLocationsMock) RegisterRevokeLicensesMock() {
	m.Register("POST", "/api/v1/volume-purchasing-locations/1/revoke-licenses", 200, "")
}

func (m *VolumePurchasingLocationsMock) RegisterListNoResponseErrorMock() {
	m.RegisterError("GET", "/api/v1/volume-purchasing-locations", 500, "error_internal.json", "")
}

func (m *VolumePurchasingLocationsMock) RegisterGetByIDNoResponseErrorMock() {
	m.RegisterError("GET", "/api/v1/volume-purchasing-locations/1", 500, "error_internal.json", "")
}

func (m *VolumePurchasingLocationsMock) RegisterCreateNoResponseErrorMock() {
	m.RegisterError("POST", "/api/v1/volume-purchasing-locations", 500, "error_internal.json", "")
}

func (m *VolumePurchasingLocationsMock) RegisterUpdateNoResponseErrorMock() {
	m.RegisterError("PATCH", "/api/v1/volume-purchasing-locations/1", 500, "error_internal.json", "")
}

func (m *VolumePurchasingLocationsMock) RegisterDeleteNoResponseErrorMock() {
	m.RegisterError("DELETE", "/api/v1/volume-purchasing-locations/1", 500, "error_internal.json", "")
}

func (m *VolumePurchasingLocationsMock) RegisterReclaimNoResponseErrorMock() {
	m.RegisterError("POST", "/api/v1/volume-purchasing-locations/1/reclaim", 500, "error_internal.json", "")
}

func (m *VolumePurchasingLocationsMock) RegisterGetContentNoResponseErrorMock() {
	m.RegisterError("GET", "/api/v1/volume-purchasing-locations/1/content", 500, "error_internal.json", "")
}

func (m *VolumePurchasingLocationsMock) RegisterGetHistoryNoResponseErrorMock() {
	m.RegisterError("GET", "/api/v1/volume-purchasing-locations/1/history", 500, "error_internal.json", "")
}

func (m *VolumePurchasingLocationsMock) RegisterAddHistoryNotesNoResponseErrorMock() {
	m.RegisterError("POST", "/api/v1/volume-purchasing-locations/1/history", 500, "error_internal.json", "")
}

func (m *VolumePurchasingLocationsMock) RegisterRevokeLicensesNoResponseErrorMock() {
	m.RegisterError("POST", "/api/v1/volume-purchasing-locations/1/revoke-licenses", 500, "error_internal.json", "")
}
