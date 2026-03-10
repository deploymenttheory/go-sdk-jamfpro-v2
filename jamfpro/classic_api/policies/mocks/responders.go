package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type PoliciesMock struct {
	*mocks.GenericMock
}

func NewPoliciesMock() *PoliciesMock {
	return &PoliciesMock{
		GenericMock: mocks.NewXMLMock("PoliciesMock"),
	}
}

func (m *PoliciesMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterGetByNameMock()
	m.RegisterCreateMock()
	m.RegisterUpdateByIDMock()
	m.RegisterUpdateByNameMock()
	m.RegisterDeleteByIDMock()
	m.RegisterDeleteByNameMock()
	m.RegisterGetByCreatedByMock()
	m.RegisterGetByCategoryMock()
	m.RegisterGetByIDWithSubsetMock()
	m.RegisterGetByNameWithSubsetMock()
}

func (m *PoliciesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *PoliciesMock) RegisterListMock() {
	m.Register("GET", "/JSSResource/policies", 200, "validate_list_policies.xml")
}

func (m *PoliciesMock) RegisterGetByIDMock() {
	m.Register("GET", "/JSSResource/policies/id/1", 200, "validate_get_policy.xml")
}

func (m *PoliciesMock) RegisterGetByNameMock() {
	m.Register("GET", "/JSSResource/policies/name/Test Policy", 200, "validate_get_policy.xml")
}

func (m *PoliciesMock) RegisterCreateMock() {
	m.Register("POST", "/JSSResource/policies/id/0", 201, "validate_create_policy.xml")
}

func (m *PoliciesMock) RegisterUpdateByIDMock() {
	m.Register("PUT", "/JSSResource/policies/id/1", 200, "validate_update_policy.xml")
}

func (m *PoliciesMock) RegisterUpdateByNameMock() {
	m.Register("PUT", "/JSSResource/policies/name/Test Policy", 200, "validate_update_policy.xml")
}

func (m *PoliciesMock) RegisterDeleteByIDMock() {
	m.Register("DELETE", "/JSSResource/policies/id/1", 200, "")
}

func (m *PoliciesMock) RegisterDeleteByNameMock() {
	m.Register("DELETE", "/JSSResource/policies/name/Test Policy", 200, "")
}

func (m *PoliciesMock) RegisterGetByCreatedByMock() {
	m.Register("GET", "/JSSResource/policies/createdBy/jss", 200, "validate_list_policies.xml")
}

func (m *PoliciesMock) RegisterGetByCategoryMock() {
	m.Register("GET", "/JSSResource/policies/category/TestCategory", 200, "validate_list_policies.xml")
}

func (m *PoliciesMock) RegisterGetByIDWithSubsetMock() {
	m.Register("GET", "/JSSResource/policies/id/1/subset/General", 200, "validate_get_policy_subset.xml")
}

func (m *PoliciesMock) RegisterGetByNameWithSubsetMock() {
	m.Register("GET", "/JSSResource/policies/name/Test Policy/subset/General", 200, "validate_get_policy_subset.xml")
}

func (m *PoliciesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/policies/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *PoliciesMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/policies/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A policy with that name already exists")
}
