package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type PatchExternalSourcesMock struct {
	*mocks.GenericMock
}

func NewPatchExternalSourcesMock() *PatchExternalSourcesMock {
	return &PatchExternalSourcesMock{
		GenericMock: mocks.NewXMLMock("PatchExternalSourcesMock"),
	}
}

func (m *PatchExternalSourcesMock) RegisterMocks() {
	m.RegisterListPatchExternalSourcesMock()
	m.RegisterGetPatchExternalSourceByIDMock()
	m.RegisterGetPatchExternalSourceByNameMock()
	m.RegisterCreatePatchExternalSourceMock()
	m.RegisterUpdatePatchExternalSourceByIDMock()
	m.RegisterUpdatePatchExternalSourceByNameMock()
	m.RegisterDeletePatchExternalSourceByIDMock()
}

func (m *PatchExternalSourcesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *PatchExternalSourcesMock) RegisterListPatchExternalSourcesMock() {
	m.Register("GET", "/JSSResource/patchexternalsources", 200, "validate_list_patch_external_sources.xml")
}

func (m *PatchExternalSourcesMock) RegisterGetPatchExternalSourceByIDMock() {
	m.Register("GET", "/JSSResource/patchexternalsources/id/1", 200, "validate_get_patch_external_source.xml")
}

func (m *PatchExternalSourcesMock) RegisterGetPatchExternalSourceByNameMock() {
	m.Register("GET", "/JSSResource/patchexternalsources/name/Primary Patch Source", 200, "validate_get_patch_external_source.xml")
}

func (m *PatchExternalSourcesMock) RegisterCreatePatchExternalSourceMock() {
	m.Register("POST", "/JSSResource/patchexternalsources/id/0", 201, "validate_create_patch_external_source.xml")
}

func (m *PatchExternalSourcesMock) RegisterUpdatePatchExternalSourceByIDMock() {
	m.Register("PUT", "/JSSResource/patchexternalsources/id/1", 200, "validate_update_patch_external_source.xml")
}

func (m *PatchExternalSourcesMock) RegisterUpdatePatchExternalSourceByNameMock() {
	m.Register("PUT", "/JSSResource/patchexternalsources/name/Primary Patch Source", 200, "validate_update_patch_external_source.xml")
}

func (m *PatchExternalSourcesMock) RegisterDeletePatchExternalSourceByIDMock() {
	m.Register("DELETE", "/JSSResource/patchexternalsources/id/1", 200, "")
}

func (m *PatchExternalSourcesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/patchexternalsources/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *PatchExternalSourcesMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/patchexternalsources/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A patch external source with that name already exists")
}

