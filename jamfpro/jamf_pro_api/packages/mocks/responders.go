package mocks

import (
	"encoding/json"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type PackagesMock struct {
	*mocks.GenericMock
}

func NewPackagesMock() *PackagesMock {
	return &PackagesMock{
		GenericMock: mocks.NewJSONMock("PackagesMock"),
	}
}

func (m *PackagesMock) RegisterMocks() {
	m.RegisterListPackagesMock()
	m.RegisterGetPackageMock()
	m.RegisterCreatePackageMock()
	m.RegisterUpdatePackageMock()
	m.RegisterDeletePackageMock()
	m.RegisterDeletePackagesByIDMock()
	m.RegisterGetPackageHistoryMock()
	m.RegisterAddPackageHistoryNotesMock()
}

func (m *PackagesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *PackagesMock) RegisterListPackagesMock() {
	m.Register("GET", "/api/v1/packages", 200, "validate_list_packages.json")
}

func (m *PackagesMock) RegisterListPackagesRSQLMock() {
	m.Register("GET", "/api/v1/packages", 200, "validate_list_packages_rsql.json")
}

func (m *PackagesMock) RegisterGetPackageMock() {
	m.Register("GET", "/api/v1/packages/1", 200, "validate_get_package.json")
}

func (m *PackagesMock) RegisterCreatePackageMock() {
	m.Register("POST", "/api/v1/packages", 201, "validate_create_package.json")
}

func (m *PackagesMock) RegisterUpdatePackageMock() {
	m.Register("PUT", "/api/v1/packages/1", 200, "validate_update_package.json")
}

func (m *PackagesMock) RegisterUploadPackageMock() {
	m.Register("POST", "/api/v1/packages/1/upload", 201, "validate_create_package.json")
}

func (m *PackagesMock) RegisterAssignManifestMock() {
	m.Register("POST", "/api/v1/packages/1/manifest", 201, "validate_create_package.json")
}

func (m *PackagesMock) RegisterDeleteManifestMock() {
	m.Register("DELETE", "/api/v1/packages/1/manifest", 204, "")
}

func (m *PackagesMock) RegisterDeletePackageMock() {
	m.Register("DELETE", "/api/v1/packages/1", 204, "")
}

func (m *PackagesMock) RegisterDeletePackagesByIDMock() {
	m.Register("POST", "/api/v1/packages/delete-multiple", 204, "")
}

func (m *PackagesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/packages/999", 404, "error_not_found.json", "")
}

func (m *PackagesMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/api/v1/packages", 409, "error_conflict.json", "")
}

func (m *PackagesMock) RegisterGetPackageHistoryMock() {
	m.Register("GET", "/api/v1/packages/1/history", 200, "validate_get_history.json")
}

func (m *PackagesMock) RegisterAddPackageHistoryNotesMock() {
	m.Register("POST", "/api/v1/packages/1/history", 201, "")
}

func (m *PackagesMock) RegisterExportPackagesMock() {
	m.Register("POST", "/api/v1/packages/export", 200, "validate_export.json")
}

func (m *PackagesMock) RegisterExportHistoryMock() {
	m.Register("POST", "/api/v1/packages/1/history/export", 200, "validate_export_history.json")
}

func (m *PackagesMock) RegisterRefreshCloudDistributionPointMock() {
	m.Register("POST", "/api/v1/cloud-distribution-point/refresh-inventory", 204, "")
}

func (m *PackagesMock) RegisterGetPackageWithHashMock(id string, hashValue string) {
	path := "/api/v1/packages/" + id
	body, _ := json.Marshal(map[string]any{
		"id": id, "packageName": "Test", "fileName": "test.pkg", "categoryId": "1",
		"hashType": "SHA3_512", "hashValue": hashValue,
	})
	m.RegisterRawBody("GET", path, 200, body)
}

func (m *PackagesMock) RegisterUploadPackageMockForID(id string) {
	path := "/api/v1/packages/" + id + "/upload"
	m.Register("POST", path, 201, "validate_create_package.json")
}

func (m *PackagesMock) RegisterRawBody(method, path string, statusCode int, body []byte) {
	m.GenericMock.RegisterRawBody(method, path, statusCode, body)
}

func (m *PackagesMock) RegisterAPIError(method, path string, statusCode int, errMsg string) {
	m.RegisterError(method, path, statusCode, "", errMsg)
}

func (m *PackagesMock) RegisterExportPackagesErrorMock() {
	m.RegisterError("POST", "/api/v1/packages/export", 500, "error_internal.json", "")
}

func (m *PackagesMock) RegisterExportHistoryErrorMock() {
	m.RegisterError("POST", "/api/v1/packages/1/history/export", 500, "error_internal.json", "")
}
