package packages

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/packages/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.PackagesMock) {
	t.Helper()
	mock := mocks.NewPackagesMock()
	return NewService(mock), mock
}

func TestUnitListPackages_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListPackagesMock()

	result, resp, err := svc.ListPackagesV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Firefox", result.Results[0].PackageName)
	assert.Equal(t, "2", result.Results[1].ID)
	assert.Equal(t, "Chrome", result.Results[1].PackageName)
}

func TestUnitListPackages_WithrsqlQuery(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListPackagesMock()

	params := map[string]string{"page": "0", "page-size": "50", "sort": "name:asc"}
	result, resp, err := svc.ListPackagesV1(context.Background(), params)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitListPackages_WithRSQLFilter(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListPackagesRSQLMock()

	rsqlQuery := map[string]string{"filter": `name=="Chrome"`}
	result, resp, err := svc.ListPackagesV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "2", result.Results[0].ID)
	assert.Equal(t, "Chrome", result.Results[0].PackageName)
	assert.Equal(t, rsqlQuery, mock.LastRSQLQuery)
}

func TestUnitGetPackageByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetPackageMock()

	result, resp, err := svc.GetPackageByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Firefox", result.PackageName)
	assert.Equal(t, "Firefox.pkg", result.FileName)
	assert.Equal(t, "2", result.CategoryID)
}

func TestUnitGetPackageByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetPackageByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "package ID is required")
}

func TestUnitGetPackageByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetPackageByIDV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestUnitCreatePackage_Success(t *testing.T) {
	svc, mock := setupMockService(t)
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
	result, resp, err := svc.CreatePackageV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "3", result.ID)
	assert.Contains(t, result.Href, "/api/v1/packages/3")
}

func TestUnitCreatePackage_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreatePackageV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitUploadPackageV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUploadPackageMock()

	tmp := t.TempDir()
	pkgPath := filepath.Join(tmp, "test.pkg")
	require.NoError(t, os.WriteFile(pkgPath, []byte("test content"), 0644))

	result, resp, err := svc.UploadPackageV1(context.Background(), "1", pkgPath)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "3", result.ID)
}

func TestUnitUploadPackageV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UploadPackageV1(context.Background(), "", "/tmp/test.pkg")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "package ID is required")
}

func TestUnitUploadPackageV1_EmptyPath(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UploadPackageV1(context.Background(), "1", "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "file path is required")
}

func TestUnitAssignManifestToPackageV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAssignManifestMock()

	tmp := t.TempDir()
	manifestPath := filepath.Join(tmp, "manifest.plist")
	require.NoError(t, os.WriteFile(manifestPath, []byte("<?xml"), 0644))

	result, resp, err := svc.AssignManifestToPackageV1(context.Background(), "1", manifestPath)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
}

func TestUnitDeletePackageManifestV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteManifestMock()

	resp, err := svc.DeletePackageManifestV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitCreatePackage_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
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
	result, resp, err := svc.CreatePackageV1(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

func TestUnitUpdatePackageByID_Success(t *testing.T) {
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
	result, resp, err := svc.UpdatePackageByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Firefox Updated", result.PackageName)
}

func TestUnitUpdatePackageByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdatePackageByIDV1(context.Background(), "", &ResourcePackage{PackageName: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnitUpdatePackageByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.UpdatePackageByIDV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitDeletePackageByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeletePackageMock()

	resp, err := svc.DeletePackageByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitDeletePackageByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeletePackageByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "package ID is required")
}

func TestUnitDeletePackagesByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeletePackagesByIDMock()

	req := &DeletePackagesByIDRequest{IDs: []string{"1", "2"}}
	resp, err := svc.DeletePackagesByIDV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitDeletePackagesByID_EmptyIDs(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeletePackagesByIDV1(context.Background(), &DeletePackagesByIDRequest{IDs: []string{}})
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "ids are required")
}

func TestUnitDeletePackagesByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeletePackagesByIDV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "ids are required")
}

func TestUnitGetPackageHistoryV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetPackageHistoryMock()

	result, resp, err := svc.GetPackageHistoryV1(context.Background(), "1", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", string(result.Results[0].ID))
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "Package created", result.Results[0].Note)
}

func TestUnitGetPackageHistoryV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetPackageHistoryV1(context.Background(), "", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "package ID is required")
}

func TestUnitAddPackageHistoryNotesV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddPackageHistoryNotesMock()

	req := &AddHistoryNotesRequest{Note: "Added via SDK"}
	resp, err := svc.AddPackageHistoryNotesV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
}

func TestUnitAddPackageHistoryNotesV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.AddPackageHistoryNotesV1(context.Background(), "", &AddHistoryNotesRequest{Note: "x"})
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "package ID is required")
}

func TestUnitAddPackageHistoryNotesV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.AddPackageHistoryNotesV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request body is required")
}
