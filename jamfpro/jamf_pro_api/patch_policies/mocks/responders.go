package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type PatchPoliciesMock struct {
	*mocks.GenericMock
}

func NewPatchPoliciesMock() *PatchPoliciesMock {
	return &PatchPoliciesMock{
		GenericMock: mocks.NewJSONMock("PatchPoliciesMock"),
	}
}

func (m *PatchPoliciesMock) RegisterListMock() {
	m.Register("GET", "/api/v2/patch-policies/policy-details", 200, "validate_list.json")
}

func (m *PatchPoliciesMock) RegisterGetDashboardStatusMock(id string, onDashboard bool) {
	fixture := "validate_dashboard_status_false.json"
	if onDashboard {
		fixture = "validate_dashboard_status.json"
	}
	path := "/api/v2/patch-policies/" + id + "/dashboard"
	m.Register("GET", path, 200, fixture)
}

func (m *PatchPoliciesMock) RegisterAddToDashboardMock(id string) {
	path := "/api/v2/patch-policies/" + id + "/dashboard"
	m.Register("POST", path, 200, "")
}

func (m *PatchPoliciesMock) RegisterRemoveFromDashboardMock(id string) {
	path := "/api/v2/patch-policies/" + id + "/dashboard"
	m.Register("DELETE", path, 200, "")
}

func (m *PatchPoliciesMock) RegisterEmptyListMock() {
	m.Register("GET", "/api/v2/patch-policies/policy-details", 200, "validate_empty_list.json")
}

func (m *PatchPoliciesMock) RegisterListSummaryMock() {
	m.Register("GET", "/api/v2/patch-policies", 200, "validate_list_summary.json")
}

func (m *PatchPoliciesMock) RegisterListSummaryEmptyMock() {
	m.Register("GET", "/api/v2/patch-policies", 200, "validate_list_summary_empty.json")
}

func (m *PatchPoliciesMock) RegisterListErrorMock() {
	m.RegisterError("GET", "/api/v2/patch-policies/policy-details", 500, "error_not_found.json", "Jamf Pro API error")
}

func (m *PatchPoliciesMock) RegisterListInvalidMock() {
	m.Register("GET", "/api/v2/patch-policies/policy-details", 200, "validate_list_invalid.json")
}

func (m *PatchPoliciesMock) RegisterListSummaryErrorMock() {
	m.RegisterError("GET", "/api/v2/patch-policies", 500, "error_not_found.json", "Jamf Pro API error")
}

func (m *PatchPoliciesMock) RegisterGetDashboardStatusErrorMock(id string) {
	path := "/api/v2/patch-policies/" + id + "/dashboard"
	m.RegisterError("GET", path, 404, "error_not_found.json", "Jamf Pro API error")
}

func (m *PatchPoliciesMock) RegisterAddToDashboardErrorMock(id string) {
	path := "/api/v2/patch-policies/" + id + "/dashboard"
	m.RegisterError("POST", path, 500, "error_not_found.json", "Jamf Pro API error")
}

func (m *PatchPoliciesMock) RegisterRemoveFromDashboardErrorMock(id string) {
	path := "/api/v2/patch-policies/" + id + "/dashboard"
	m.RegisterError("DELETE", path, 500, "error_not_found.json", "Jamf Pro API error")
}
