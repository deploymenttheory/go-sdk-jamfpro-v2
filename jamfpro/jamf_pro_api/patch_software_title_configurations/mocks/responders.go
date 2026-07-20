package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type PatchSoftwareTitleConfigurationsMock struct {
	*mocks.GenericMock
}

func NewPatchSoftwareTitleConfigurationsMock() *PatchSoftwareTitleConfigurationsMock {
	return &PatchSoftwareTitleConfigurationsMock{
		GenericMock: mocks.NewJSONMock("PatchSoftwareTitleConfigurationsMock"),
	}
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterListMock() {
	m.Register("GET", "/api/v2/patch-software-title-configurations", 200, "validate_list.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetByIDMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id
	m.Register("GET", path, 200, "validate_get.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterCreateMock() {
	m.Register("POST", "/api/v2/patch-software-title-configurations", 200, "validate_create.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterUpdateByIDMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id
	m.Register("PATCH", path, 200, "validate_get.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterDeleteByIDMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id
	m.Register("DELETE", path, 200, "")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterEmptyListMock() {
	m.Register("GET", "/api/v2/patch-software-title-configurations", 200, "validate_empty_list.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetDashboardStatusMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/dashboard"
	m.Register("GET", path, 200, "validate_dashboard_status.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterAddToDashboardMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/dashboard"
	m.Register("POST", path, 204, "")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterRemoveFromDashboardMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/dashboard"
	m.Register("DELETE", path, 204, "")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetDefinitionsMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/definitions"
	m.Register("GET", path, 200, "validate_definitions.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetDependenciesMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/dependencies"
	m.Register("GET", path, 200, "validate_dependencies.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterExportReportMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/export-report"
	m.Register("GET", path, 200, "validate_export_report.csv")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetExtensionAttributesMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/extension-attributes"
	m.Register("GET", path, 200, "validate_extension_attributes.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetPatchReportMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/patch-report"
	m.Register("GET", path, 200, "validate_patch_report.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetPatchSummaryMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/patch-summary"
	m.Register("GET", path, 200, "validate_patch_summary.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetHistoryMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/history"
	m.Register("GET", path, 200, "validate_history.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterAddHistoryNoteMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/history"
	m.Register("POST", path, 201, "validate_add_history_note.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetPatchVersionsMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/patch-summary/versions"
	m.Register("GET", path, 200, "validate_patch_versions.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterListNoResponseErrorMock() {
	m.RegisterError("GET", "/api/v2/patch-software-title-configurations", 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetByIDNoResponseErrorMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id
	m.RegisterError("GET", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterCreateNoResponseErrorMock() {
	m.RegisterError("POST", "/api/v2/patch-software-title-configurations", 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterUpdateByIDNoResponseErrorMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id
	m.RegisterError("PATCH", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterDeleteByIDNoResponseErrorMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id
	m.RegisterError("DELETE", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetDashboardStatusNoResponseErrorMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/dashboard"
	m.RegisterError("GET", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterAddToDashboardNoResponseErrorMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/dashboard"
	m.RegisterError("POST", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterRemoveFromDashboardNoResponseErrorMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/dashboard"
	m.RegisterError("DELETE", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetDefinitionsNoResponseErrorMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/definitions"
	m.RegisterError("GET", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetDependenciesNoResponseErrorMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/dependencies"
	m.RegisterError("GET", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterExportReportNoResponseErrorMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/export-report"
	m.RegisterError("GET", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetExtensionAttributesNoResponseErrorMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/extension-attributes"
	m.RegisterError("GET", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetPatchReportNoResponseErrorMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/patch-report"
	m.RegisterError("GET", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetPatchSummaryNoResponseErrorMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/patch-summary"
	m.RegisterError("GET", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetHistoryNoResponseErrorMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/history"
	m.RegisterError("GET", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterAddHistoryNoteNoResponseErrorMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/history"
	m.RegisterError("POST", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetPatchVersionsNoResponseErrorMock(id string) {
	path := "/api/v2/patch-software-title-configurations/" + id + "/patch-summary/versions"
	m.RegisterError("GET", path, 500, "error_internal.json", "no response")
}

// -----------------------------------------------------------------------------
// V3 responders (Jamf Pro 11.30)
// -----------------------------------------------------------------------------

func (m *PatchSoftwareTitleConfigurationsMock) RegisterListMockV3() {
	m.Register("GET", "/api/v3/patch-software-title-configurations", 200, "validate_list.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetByIDMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id
	m.Register("GET", path, 200, "validate_get.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterCreateMockV3() {
	m.Register("POST", "/api/v3/patch-software-title-configurations", 200, "validate_create.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterUpdateByIDMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id
	m.Register("PATCH", path, 200, "validate_get.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterDeleteByIDMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id
	m.Register("DELETE", path, 200, "")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterEmptyListMockV3() {
	m.Register("GET", "/api/v3/patch-software-title-configurations", 200, "validate_empty_list.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetDashboardStatusMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/dashboard"
	m.Register("GET", path, 200, "validate_dashboard_status.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterAddToDashboardMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/dashboard"
	m.Register("POST", path, 204, "")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterRemoveFromDashboardMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/dashboard"
	m.Register("DELETE", path, 204, "")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetDefinitionsMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/definitions"
	m.Register("GET", path, 200, "validate_definitions.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetDependenciesMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/dependencies"
	m.Register("GET", path, 200, "validate_dependencies.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterExportReportMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/export-report"
	m.Register("GET", path, 200, "validate_export_report.csv")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetExtensionAttributesMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/extension-attributes"
	m.Register("GET", path, 200, "validate_extension_attributes.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetPatchReportMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/patch-report"
	m.Register("GET", path, 200, "validate_patch_report_v3.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetPatchSummaryMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/patch-summary"
	m.Register("GET", path, 200, "validate_patch_summary.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetHistoryMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/history"
	m.Register("GET", path, 200, "validate_history.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterAddHistoryNoteMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/history"
	m.Register("POST", path, 201, "validate_add_history_note.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetPatchVersionsMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/patch-summary/versions"
	m.Register("GET", path, 200, "validate_patch_versions.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterListNoResponseErrorMockV3() {
	m.RegisterError("GET", "/api/v3/patch-software-title-configurations", 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetByIDNoResponseErrorMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id
	m.RegisterError("GET", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterCreateNoResponseErrorMockV3() {
	m.RegisterError("POST", "/api/v3/patch-software-title-configurations", 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterUpdateByIDNoResponseErrorMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id
	m.RegisterError("PATCH", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterDeleteByIDNoResponseErrorMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id
	m.RegisterError("DELETE", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetDashboardStatusNoResponseErrorMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/dashboard"
	m.RegisterError("GET", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterAddToDashboardNoResponseErrorMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/dashboard"
	m.RegisterError("POST", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterRemoveFromDashboardNoResponseErrorMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/dashboard"
	m.RegisterError("DELETE", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetDefinitionsNoResponseErrorMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/definitions"
	m.RegisterError("GET", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetDependenciesNoResponseErrorMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/dependencies"
	m.RegisterError("GET", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterExportReportNoResponseErrorMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/export-report"
	m.RegisterError("GET", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetExtensionAttributesNoResponseErrorMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/extension-attributes"
	m.RegisterError("GET", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetPatchReportNoResponseErrorMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/patch-report"
	m.RegisterError("GET", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetPatchSummaryNoResponseErrorMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/patch-summary"
	m.RegisterError("GET", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetHistoryNoResponseErrorMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/history"
	m.RegisterError("GET", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterAddHistoryNoteNoResponseErrorMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/history"
	m.RegisterError("POST", path, 500, "error_internal.json", "no response")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetPatchVersionsNoResponseErrorMockV3(id string) {
	path := "/api/v3/patch-software-title-configurations/" + id + "/patch-summary/versions"
	m.RegisterError("GET", path, 500, "error_internal.json", "no response")
}
