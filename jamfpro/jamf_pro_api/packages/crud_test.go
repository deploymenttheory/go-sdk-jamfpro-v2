package packages

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/packages/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Packages, *mocks.PackagesMock) {
	t.Helper()
	mock := mocks.NewPackagesMock()
	return NewPackages(mock), mock
}

func TestUnit_Packages_List_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListPackagesMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Firefox", result.Results[0].PackageName)
	assert.Equal(t, "2", result.Results[1].ID)
	assert.Equal(t, "Chrome", result.Results[1].PackageName)
}

func TestUnit_Packages_List_WithrsqlQuery(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListPackagesMock()

	params := map[string]string{"page": "0", "page-size": "50", "sort": "name:asc"}
	result, resp, err := svc.ListV1(context.Background(), params)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_Packages_List_WithRSQLFilter(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListPackagesRSQLMock()

	rsqlQuery := map[string]string{"filter": `name=="Chrome"`}
	result, resp, err := svc.ListV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "2", result.Results[0].ID)
	assert.Equal(t, "Chrome", result.Results[0].PackageName)
	assert.Equal(t, rsqlQuery, mock.LastRSQLQuery)
}

func TestUnit_Packages_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetPackageMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Firefox", result.PackageName)
	assert.Equal(t, "Firefox.pkg", result.FileName)
	assert.Equal(t, "2", result.CategoryID)
}

func TestUnit_Packages_GetByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "package ID is required")
}

func TestUnit_Packages_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_Packages_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetCloudDistributionPointMock("JAMF_CLOUD")
	mock.RegisterCreatePackageMock()

	req := &RequestPackage{
		PackageName:          "Safari Extension",
		FileName:             "SafariExtension.pkg",
		CategoryID:           "2",
		Info:                 "Sample package",
		Notes:                "For testing",
		Priority:             10,
		FillUserTemplate:     BoolPtr(true),
		FillExistingUsers:    BoolPtr(false),
		RebootRequired:       BoolPtr(false),
		OSInstall:            BoolPtr(false),
		SuppressUpdates:      BoolPtr(false),
		SuppressFromDock:     BoolPtr(false),
		SuppressEula:         BoolPtr(false),
		SuppressRegistration: BoolPtr(false),
	}
	result, resp, err := svc.CreateV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "3", result.ID)
	assert.Contains(t, result.Href, "/api/v1/packages/3")
}

func TestUnit_Packages_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Packages_UploadV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUploadPackageMock()

	tmp := t.TempDir()
	pkgPath := filepath.Join(tmp, "test.pkg")
	require.NoError(t, os.WriteFile(pkgPath, []byte("test content"), 0644))

	result, resp, err := svc.UploadV1(context.Background(), "1", pkgPath)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "3", result.ID)
}

func TestUnit_Packages_UploadV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UploadV1(context.Background(), "", "/tmp/test.pkg")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "package ID is required")
}

func TestUnit_Packages_UploadV1_EmptyPath(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UploadV1(context.Background(), "1", "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "file path is required")
}

func TestUnit_Packages_AssignManifestToPackageV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAssignManifestMock()

	tmp := t.TempDir()
	manifestPath := filepath.Join(tmp, "manifest.plist")
	require.NoError(t, os.WriteFile(manifestPath, []byte("<?xml"), 0644))

	result, resp, err := svc.AssignManifestToPackageV1(context.Background(), "1", manifestPath)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
}

func TestUnit_Packages_DeletePackageManifestV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteManifestMock()

	resp, err := svc.DeletePackageManifestV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_Packages_Create_CDPNotEnabled(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetCloudDistributionPointMock("NONE")

	req := &RequestPackage{PackageName: "Test", FileName: "test.pkg", CategoryID: "1", Priority: 1}
	result, resp, err := svc.CreateV1(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "content delivery network")
}

func TestUnit_Packages_Create_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetCloudDistributionPointMock("JAMF_CLOUD")
	mock.RegisterConflictErrorMock()

	req := &RequestPackage{
		PackageName:          "Duplicate",
		FileName:             "Duplicate.pkg",
		CategoryID:           "1",
		Priority:             1,
		FillUserTemplate:     BoolPtr(false),
		RebootRequired:       BoolPtr(false),
		OSInstall:            BoolPtr(false),
		SuppressUpdates:      BoolPtr(false),
		SuppressFromDock:     BoolPtr(false),
		SuppressEula:         BoolPtr(false),
		SuppressRegistration: BoolPtr(false),
	}
	result, resp, err := svc.CreateV1(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode())
}

func TestUnit_Packages_UpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdatePackageMock()

	req := &ResourcePackage{
		PackageName:       "Firefox Updated",
		FileName:          "Firefox.pkg",
		CategoryID:        "2",
		Info:              "Mozilla Firefox ESR",
		Notes:             "Updated to ESR",
		Priority:          15,
		FillUserTemplate:  BoolPtr(true),
		FillExistingUsers: BoolPtr(true),
	}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Firefox Updated", result.PackageName)
}

func TestUnit_Packages_UpdateByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByIDV1(context.Background(), "", &ResourcePackage{PackageName: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_Packages_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.UpdateByIDV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Packages_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeletePackageMock()

	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_Packages_DeleteByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "package ID is required")
}

func TestUnit_Packages_DeleteMultipleByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeletePackagesByIDMock()

	req := &DeletePackagesByIDRequest{IDs: []string{"1", "2"}}
	resp, err := svc.DeletePackagesByIDV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_Packages_DeleteMultipleByID_EmptyIDs(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeletePackagesByIDV1(context.Background(), &DeletePackagesByIDRequest{IDs: []string{}})
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "ids are required")
}

func TestUnit_Packages_DeleteMultipleByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeletePackagesByIDV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "ids are required")
}

func TestUnit_Packages_GetHistoryV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetPackageHistoryMock()

	result, resp, err := svc.GetHistoryV1(context.Background(), "1", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", string(result.Results[0].ID))
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "Package created", result.Results[0].Note)
}

func TestUnit_Packages_GetHistoryV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetHistoryV1(context.Background(), "", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "package ID is required")
}

func TestUnit_Packages_AddHistoryNotesV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddPackageHistoryNotesMock()

	req := &AddHistoryNotesRequest{Note: "Added via SDK"}
	resp, err := svc.AddHistoryNotesV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
}

func TestUnit_Packages_AddHistoryNotesV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.AddHistoryNotesV1(context.Background(), "", &AddHistoryNotesRequest{Note: "x"})
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "package ID is required")
}

func TestUnit_Packages_AddHistoryNotesV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.AddHistoryNotesV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request body is required")
}

// -----------------------------------------------------------------------------
// ExportV1 tests
// -----------------------------------------------------------------------------

func TestUnit_Packages_ExportV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportPackagesMock()

	body, resp, err := svc.ExportV1(context.Background(), nil, nil, constants.ApplicationJSON)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, body)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Contains(t, string(body), "Firefox")
}

func TestUnit_Packages_ExportV1_WithQueryAndBody(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportPackagesMock()

	query := map[string]string{"page": "0", "page-size": "100"}
	page, pageSize := 0, 50
	exportBody := &ExportRequest{Page: &page, PageSize: &pageSize}
	body, resp, err := svc.ExportV1(context.Background(), query, exportBody, constants.ApplicationJSON)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, body)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_Packages_ExportV1_CSV(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportPackagesMock()

	body, resp, err := svc.ExportV1(context.Background(), nil, nil, constants.TextCSV)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, body)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_Packages_ExportV1_EmptyAcceptDefaultsToJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportPackagesMock()

	body, resp, err := svc.ExportV1(context.Background(), nil, nil, "")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, body)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_Packages_ExportV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportPackagesErrorMock()

	body, resp, err := svc.ExportV1(context.Background(), nil, nil, constants.ApplicationJSON)
	assert.Error(t, err)
	assert.Nil(t, body)
	assert.NotNil(t, resp)
}

// -----------------------------------------------------------------------------
// ExportHistoryV1 tests
// -----------------------------------------------------------------------------

func TestUnit_Packages_ExportHistoryV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportHistoryMock()

	body, resp, err := svc.ExportHistoryV1(context.Background(), "1", nil, nil, constants.ApplicationJSON)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, body)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Contains(t, string(body), "admin")
}

func TestUnit_Packages_ExportHistoryV1_WithQueryAndBody(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportHistoryMock()

	query := map[string]string{"page": "0"}
	page := 0
	exportBody := &ExportRequest{Page: &page}
	body, resp, err := svc.ExportHistoryV1(context.Background(), "1", query, exportBody, constants.ApplicationJSON)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, body)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_Packages_ExportHistoryV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	body, resp, err := svc.ExportHistoryV1(context.Background(), "", nil, nil, constants.ApplicationJSON)
	assert.Error(t, err)
	assert.Nil(t, body)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "package ID is required")
}

func TestUnit_Packages_ExportHistoryV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportHistoryErrorMock()

	body, resp, err := svc.ExportHistoryV1(context.Background(), "1", nil, nil, constants.ApplicationJSON)
	assert.Error(t, err)
	assert.Nil(t, body)
	assert.NotNil(t, resp)
}

// -----------------------------------------------------------------------------
// CreateAndUpload tests
// -----------------------------------------------------------------------------

func TestUnit_Packages_CreateAndUpload_Success(t *testing.T) {
	svc, mock := setupMockService(t)

	tmp := t.TempDir()
	pkgPath := filepath.Join(tmp, "test.pkg")
	content := []byte("test content")
	require.NoError(t, os.WriteFile(pkgPath, content, 0644))

	hash, err := crypto.CalculateSHA3_512(pkgPath)
	require.NoError(t, err)

	mock.RegisterGetCloudDistributionPointMock("JAMF_CLOUD")
	mock.RegisterCreatePackageMock()
	mock.RegisterUploadPackageMockForID("3")
	mock.RegisterRefreshCloudDistributionPointMock()
	mock.RegisterGetPackageWithHashMock("3", hash)

	req := &RequestPackage{
		PackageName:          "Test Package",
		FileName:             "test.pkg",
		CategoryID:           "1",
		Priority:             1,
		FillUserTemplate:     BoolPtr(false),
		RebootRequired:       BoolPtr(false),
		OSInstall:            BoolPtr(false),
		SuppressUpdates:      BoolPtr(false),
		SuppressFromDock:     BoolPtr(false),
		SuppressEula:         BoolPtr(false),
		SuppressRegistration: BoolPtr(false),
	}

	result, resp, err := svc.CreateAndUpload(context.Background(), pkgPath, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, "3", result.ID)
}

func TestUnit_Packages_CreateAndUpload_EmptyFilePath(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestPackage{PackageName: "x", FileName: "x.pkg", CategoryID: "1", Priority: 1}
	result, resp, err := svc.CreateAndUpload(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "file path is required")
}

func TestUnit_Packages_CreateAndUpload_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	tmp := t.TempDir()
	pkgPath := filepath.Join(tmp, "test.pkg")
	require.NoError(t, os.WriteFile(pkgPath, []byte("x"), 0644))

	result, resp, err := svc.CreateAndUpload(context.Background(), pkgPath, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Packages_CreateAndUpload_FileNotFound(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestPackage{PackageName: "x", FileName: "x.pkg", CategoryID: "1", Priority: 1}
	result, resp, err := svc.CreateAndUpload(context.Background(), "/nonexistent/path.pkg", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "SHA3_512")
}

func TestUnit_Packages_CreateAndUpload_HashVerificationFailed(t *testing.T) {
	svc, mock := setupMockService(t)

	tmp := t.TempDir()
	pkgPath := filepath.Join(tmp, "test.pkg")
	require.NoError(t, os.WriteFile(pkgPath, []byte("test content"), 0644))

	mock.RegisterGetCloudDistributionPointMock("JAMF_CLOUD")
	mock.RegisterCreatePackageMock()
	mock.RegisterUploadPackageMockForID("3")
	mock.RegisterRefreshCloudDistributionPointMock()
	// Return wrong hash - verification should fail
	mock.RegisterGetPackageWithHashMock("3", "wrong-hash-value")

	req := &RequestPackage{
		PackageName:          "Test",
		FileName:             "test.pkg",
		CategoryID:           "1",
		Priority:             1,
		FillUserTemplate:     BoolPtr(false),
		RebootRequired:       BoolPtr(false),
		OSInstall:            BoolPtr(false),
		SuppressUpdates:      BoolPtr(false),
		SuppressFromDock:     BoolPtr(false),
		SuppressEula:         BoolPtr(false),
		SuppressRegistration: BoolPtr(false),
	}

	result, resp, err := svc.CreateAndUpload(context.Background(), pkgPath, req)
	assert.Error(t, err)
	// result is non-nil: the package was created on the server even though hash verification failed.
	// Callers should use result.ID to clean up the orphaned package.
	assert.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "hash verification failed")
}

func TestUnit_Packages_CreateAndUpload_CreateError(t *testing.T) {
	svc, mock := setupMockService(t)

	tmp := t.TempDir()
	pkgPath := filepath.Join(tmp, "test.pkg")
	require.NoError(t, os.WriteFile(pkgPath, []byte("x"), 0644))

	mock.RegisterGetCloudDistributionPointMock("JAMF_CLOUD")
	mock.RegisterConflictErrorMock() // Create returns 409

	req := &RequestPackage{
		PackageName:          "Test",
		FileName:             "test.pkg",
		CategoryID:           "1",
		Priority:             1,
		FillUserTemplate:     BoolPtr(false),
		RebootRequired:       BoolPtr(false),
		OSInstall:            BoolPtr(false),
		SuppressUpdates:      BoolPtr(false),
		SuppressFromDock:     BoolPtr(false),
		SuppressEula:         BoolPtr(false),
		SuppressRegistration: BoolPtr(false),
	}

	result, resp, err := svc.CreateAndUpload(context.Background(), pkgPath, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "create metadata")
}

func TestUnit_Packages_CreateAndUpload_UploadError(t *testing.T) {
	svc, mock := setupMockService(t)

	tmp := t.TempDir()
	pkgPath := filepath.Join(tmp, "test.pkg")
	require.NoError(t, os.WriteFile(pkgPath, []byte("x"), 0644))

	mock.RegisterGetCloudDistributionPointMock("JAMF_CLOUD")
	mock.RegisterCreatePackageMock()
	mock.RegisterAPIError("POST", "/api/v1/packages/3/upload", 500, "upload failed")

	req := &RequestPackage{
		PackageName:          "Test",
		FileName:             "test.pkg",
		CategoryID:           "1",
		Priority:             1,
		FillUserTemplate:     BoolPtr(false),
		RebootRequired:       BoolPtr(false),
		OSInstall:            BoolPtr(false),
		SuppressUpdates:      BoolPtr(false),
		SuppressFromDock:     BoolPtr(false),
		SuppressEula:         BoolPtr(false),
		SuppressRegistration: BoolPtr(false),
	}

	result, resp, err := svc.CreateAndUpload(context.Background(), pkgPath, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "upload file")
}

func TestUnit_Packages_UploadV1_APIError(t *testing.T) {
	svc, mock := setupMockService(t)

	tmp := t.TempDir()
	pkgPath := filepath.Join(tmp, "test.pkg")
	require.NoError(t, os.WriteFile(pkgPath, []byte("x"), 0644))

	mock.RegisterAPIError("POST", "/api/v1/packages/1/upload", 500, "server error")

	result, resp, err := svc.UploadV1(context.Background(), "1", pkgPath)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnit_Packages_AssignManifestToPackageV1_APIError(t *testing.T) {
	svc, mock := setupMockService(t)

	tmp := t.TempDir()
	manifestPath := filepath.Join(tmp, "manifest.plist")
	require.NoError(t, os.WriteFile(manifestPath, []byte("<?xml"), 0644))

	mock.RegisterAPIError("POST", "/api/v1/packages/1/manifest", 500, "server error")

	result, resp, err := svc.AssignManifestToPackageV1(context.Background(), "1", manifestPath)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnit_Packages_DeletePackageManifestV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAPIError("DELETE", "/api/v1/packages/1/manifest", 500, "server error")

	resp, err := svc.DeletePackageManifestV1(context.Background(), "1")
	assert.Error(t, err)
	require.NotNil(t, resp)
}

func TestUnit_Packages_DeleteByIDV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAPIError("DELETE", "/api/v1/packages/1", 500, "server error")

	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	assert.Error(t, err)
	require.NotNil(t, resp)
}

func TestUnit_Packages_DeletePackagesByIDV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAPIError("POST", "/api/v1/packages/delete-multiple", 500, "server error")

	resp, err := svc.DeletePackagesByIDV1(context.Background(), &DeletePackagesByIDRequest{IDs: []string{"1", "2"}})
	assert.Error(t, err)
	require.NotNil(t, resp)
}

// -----------------------------------------------------------------------------
// Additional validation and error path tests
// -----------------------------------------------------------------------------

func TestUnit_Packages_AssignManifestToPackageV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	tmp := t.TempDir()
	manifestPath := filepath.Join(tmp, "manifest.plist")
	require.NoError(t, os.WriteFile(manifestPath, []byte("<?xml"), 0644))

	result, resp, err := svc.AssignManifestToPackageV1(context.Background(), "", manifestPath)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "package ID is required")
}

func TestUnit_Packages_AssignManifestToPackageV1_EmptyManifestPath(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.AssignManifestToPackageV1(context.Background(), "1", "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "manifest path is required")
}

func TestUnit_Packages_AssignManifestToPackageV1_FileNotFound(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.AssignManifestToPackageV1(context.Background(), "1", "/nonexistent/manifest.plist")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "open manifest file")
}

func TestUnit_Packages_DeletePackageManifestV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeletePackageManifestV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "package ID is required")
}

func TestUnit_Packages_UploadV1_FileNotFound(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UploadV1(context.Background(), "1", "/nonexistent/test.pkg")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "open package file")
}

func TestUnit_Packages_ListV1_InvalidJSONMergePage(t *testing.T) {
	svc, mock := setupMockService(t)
	// Invalid JSON causes mergePage unmarshal to fail
	malformed := []byte(`{invalid json`)
	mock.RegisterRawBody("GET", "/api/v1/packages", 200, malformed)

	result, resp, err := svc.ListV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "mergePage failed")
}

// -----------------------------------------------------------------------------
// UpdateAndUpload tests
// -----------------------------------------------------------------------------

func TestUnit_Packages_UpdateAndUpload_Success(t *testing.T) {
	svc, mock := setupMockService(t)

	tmp := t.TempDir()
	pkgPath := filepath.Join(tmp, "updated.pkg")
	content := []byte("updated content")
	require.NoError(t, os.WriteFile(pkgPath, content, 0644))

	hash, err := crypto.CalculateSHA3_512(pkgPath)
	require.NoError(t, err)

	mock.RegisterUpdatePackageMock()
	mock.RegisterUploadPackageMockForID("1")
	mock.RegisterRefreshCloudDistributionPointMock()
	mock.RegisterGetPackageWithHashMock("1", hash)

	req := &ResourcePackage{
		PackageName:          "Updated Package",
		FileName:             "updated.pkg",
		CategoryID:           "1",
		Priority:             15,
		FillUserTemplate:     BoolPtr(false),
		RebootRequired:       BoolPtr(false),
		OSInstall:            BoolPtr(false),
		SuppressUpdates:      BoolPtr(false),
		SuppressFromDock:     BoolPtr(false),
		SuppressEula:         BoolPtr(false),
		SuppressRegistration: BoolPtr(false),
	}

	result, resp, err := svc.UpdateAndUpload(context.Background(), "1", pkgPath, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Firefox Updated", result.PackageName)
}

func TestUnit_Packages_UpdateAndUpload_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	tmp := t.TempDir()
	pkgPath := filepath.Join(tmp, "test.pkg")
	require.NoError(t, os.WriteFile(pkgPath, []byte("x"), 0644))

	req := &ResourcePackage{PackageName: "x", FileName: "x.pkg", CategoryID: "1", Priority: 1}
	result, resp, err := svc.UpdateAndUpload(context.Background(), "", pkgPath, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "package ID is required")
}

func TestUnit_Packages_UpdateAndUpload_EmptyFilePath(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourcePackage{PackageName: "x", FileName: "x.pkg", CategoryID: "1", Priority: 1}
	result, resp, err := svc.UpdateAndUpload(context.Background(), "1", "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "file path is required")
}

func TestUnit_Packages_UpdateAndUpload_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	tmp := t.TempDir()
	pkgPath := filepath.Join(tmp, "test.pkg")
	require.NoError(t, os.WriteFile(pkgPath, []byte("x"), 0644))

	result, resp, err := svc.UpdateAndUpload(context.Background(), "1", pkgPath, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Packages_UpdateAndUpload_FileNotFound(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourcePackage{PackageName: "x", FileName: "x.pkg", CategoryID: "1", Priority: 1}
	result, resp, err := svc.UpdateAndUpload(context.Background(), "1", "/nonexistent/path.pkg", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "SHA3_512")
}

func TestUnit_Packages_UpdateAndUpload_HashVerificationFailed(t *testing.T) {
	svc, mock := setupMockService(t)

	tmp := t.TempDir()
	pkgPath := filepath.Join(tmp, "test.pkg")
	require.NoError(t, os.WriteFile(pkgPath, []byte("test content"), 0644))

	mock.RegisterUpdatePackageMock()
	mock.RegisterUploadPackageMockForID("1")
	mock.RegisterRefreshCloudDistributionPointMock()
	mock.RegisterGetPackageWithHashMock("1", "wrong-hash-value")

	req := &ResourcePackage{
		PackageName:          "Test",
		FileName:             "test.pkg",
		CategoryID:           "1",
		Priority:             1,
		FillUserTemplate:     BoolPtr(false),
		RebootRequired:       BoolPtr(false),
		OSInstall:            BoolPtr(false),
		SuppressUpdates:      BoolPtr(false),
		SuppressFromDock:     BoolPtr(false),
		SuppressEula:         BoolPtr(false),
		SuppressRegistration: BoolPtr(false),
	}

	result, resp, err := svc.UpdateAndUpload(context.Background(), "1", pkgPath, req)
	assert.Error(t, err)
	// result is non-nil: the package was updated on the server even though hash verification failed.
	// Callers should use result.ID to investigate or roll back the orphaned package.
	assert.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "hash verification failed")
}

func TestUnit_Packages_UpdateAndUpload_UpdateError(t *testing.T) {
	svc, mock := setupMockService(t)

	tmp := t.TempDir()
	pkgPath := filepath.Join(tmp, "test.pkg")
	require.NoError(t, os.WriteFile(pkgPath, []byte("x"), 0644))

	mock.RegisterAPIError("PUT", "/api/v1/packages/1", 500, "update failed")

	req := &ResourcePackage{
		PackageName:          "Test",
		FileName:             "test.pkg",
		CategoryID:           "1",
		Priority:             1,
		FillUserTemplate:     BoolPtr(false),
		RebootRequired:       BoolPtr(false),
		OSInstall:            BoolPtr(false),
		SuppressUpdates:      BoolPtr(false),
		SuppressFromDock:     BoolPtr(false),
		SuppressEula:         BoolPtr(false),
		SuppressRegistration: BoolPtr(false),
	}

	result, resp, err := svc.UpdateAndUpload(context.Background(), "1", pkgPath, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "update metadata")
}

func TestUnit_Packages_UpdateAndUpload_UploadError(t *testing.T) {
	svc, mock := setupMockService(t)

	tmp := t.TempDir()
	pkgPath := filepath.Join(tmp, "test.pkg")
	require.NoError(t, os.WriteFile(pkgPath, []byte("x"), 0644))

	mock.RegisterUpdatePackageMock()
	mock.RegisterAPIError("POST", "/api/v1/packages/1/upload", 500, "upload failed")

	req := &ResourcePackage{
		PackageName:          "Test",
		FileName:             "test.pkg",
		CategoryID:           "1",
		Priority:             1,
		FillUserTemplate:     BoolPtr(false),
		RebootRequired:       BoolPtr(false),
		OSInstall:            BoolPtr(false),
		SuppressUpdates:      BoolPtr(false),
		SuppressFromDock:     BoolPtr(false),
		SuppressEula:         BoolPtr(false),
		SuppressRegistration: BoolPtr(false),
	}

	result, resp, err := svc.UpdateAndUpload(context.Background(), "1", pkgPath, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "upload file")
}
