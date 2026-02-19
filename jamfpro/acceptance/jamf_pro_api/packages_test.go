package jamf_pro_api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/packages"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// acceptanceTestPackageURL is a public web source used for package upload acceptance tests.
const acceptanceTestPackageURL = "https://ftp.mozilla.org/pub/firefox/releases/147.0/mac/en-GB/Firefox%20147.0.pkg"

func uniquePackageName(base string) string {
	return fmt.Sprintf("%s-%d", base, time.Now().UnixMilli())
}

// =============================================================================
// TestAcceptance_Packages_Lifecycle exercises the full write/read/delete
// lifecycle: DoPackageUpload (create metadata → upload → verify SHA3_512) → List → GetByID → Update → History → Delete.
// =============================================================================

func TestAcceptance_Packages_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Packages
	ctx := context.Background()

	// 1. Download package from web source and DoPackageUpload (create metadata → upload → verify SHA3_512)
	acc.LogTestStage(t, "DoPackageUpload", "Downloading package from web source, creating metadata, uploading file, verifying SHA3_512")

	tmpDir := t.TempDir()
	pkgPath := filepath.Join(tmpDir, "Firefox_147.0.pkg")
	func() {
		resp, err := http.Get(acceptanceTestPackageURL)
		require.NoError(t, err, "download package from web source")
		defer resp.Body.Close()
		require.Equal(t, http.StatusOK, resp.StatusCode, "web source must return 200")
		out, err := os.Create(pkgPath)
		require.NoError(t, err)
		defer out.Close()
		_, err = io.Copy(out, resp.Body)
		require.NoError(t, err)
	}()

	createReq := &packages.RequestPackage{
		PackageName:           uniquePackageName("acc-test-package"),
		FileName:               "", // set by DoPackageUpload from filepath
		CategoryID:            "-1",
		Info:                  "Acceptance test package",
		Notes:                 "Created by SDK acceptance test",
		Priority:              10,
		FillUserTemplate:      packages.BoolPtr(true),
		FillExistingUsers:     packages.BoolPtr(false),
		RebootRequired:        packages.BoolPtr(false),
		OSInstall:             packages.BoolPtr(false),
		SuppressUpdates:       packages.BoolPtr(false),
		SuppressFromDock:      packages.BoolPtr(false),
		SuppressEula:          packages.BoolPtr(false),
		SuppressRegistration: packages.BoolPtr(false),
	}
	created, createResp, err := svc.DoPackageUpload(ctx, pkgPath, createReq)
	require.NoError(t, err, "DoPackageUpload should not return an error")
	require.NotNil(t, created)
	assert.Equal(t, 201, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	packageID := created.ID
	acc.LogTestSuccess(t, "Package created and uploaded with ID=%s (SHA3_512 verified)", packageID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeletePackageByIDV1(cleanupCtx, packageID)
		acc.LogCleanupDeleteError(t, "package", packageID, delErr)
	})

	// 2. List — verify creation
	acc.LogTestStage(t, "List", "Listing packages to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListPackagesV1(ctx2, map[string]string{"page": "0", "page-size": "200"})
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, p := range list.Results {
		if p.ID == packageID {
			found = true
			assert.Equal(t, createReq.PackageName, p.PackageName)
			break
		}
	}
	assert.True(t, found, "newly created package should appear in list")
	acc.LogTestSuccess(t, "Package ID=%s found in list (%d total)", packageID, list.TotalCount)

	// 3. GetByID
	acc.LogTestStage(t, "GetByID", "Fetching package by ID=%s", packageID)

	fetched, fetchResp, err := svc.GetPackageByIDV1(ctx, packageID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, packageID, fetched.ID)
	assert.Equal(t, createReq.PackageName, fetched.PackageName)
	acc.LogTestSuccess(t, "GetByID: packageName=%q", fetched.PackageName)

	// 4. Update (GET first, modify, then PUT full ResourcePackage)
	acc.LogTestStage(t, "Update", "Updating package ID=%s", packageID)

	updateReq := &packages.ResourcePackage{
		ID:                   fetched.ID,
		PackageName:          uniquePackageName("acc-test-package-updated"),
		FileName:             fetched.FileName,
		CategoryID:           fetched.CategoryID,
		Info:                 "Updated acceptance test package",
		Notes:                "Updated by SDK acceptance test",
		Priority:             15,
		FillUserTemplate:     packages.BoolPtr(true),
		FillExistingUsers:    packages.BoolPtr(true),
		RebootRequired:       fetched.RebootRequired,
		OSInstall:            fetched.OSInstall,
		SuppressUpdates:      fetched.SuppressUpdates,
		SuppressFromDock:     fetched.SuppressFromDock,
		SuppressEula:         fetched.SuppressEula,
		SuppressRegistration: fetched.SuppressRegistration,
	}
	updated, updateResp, err := svc.UpdatePackageByIDV1(ctx, packageID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode)
	assert.Equal(t, updateReq.PackageName, updated.PackageName)
	acc.LogTestSuccess(t, "Package updated: ID=%s", packageID)

	// 5. Re-fetch to verify
	fetched2, _, err := svc.GetPackageByIDV1(ctx, packageID)
	require.NoError(t, err)
	assert.Equal(t, updateReq.PackageName, fetched2.PackageName)
	acc.LogTestSuccess(t, "Update verified: packageName=%q", fetched2.PackageName)

	// 5a. Add history note and fetch history
	acc.LogTestStage(t, "History", "Adding history note and fetching history for ID=%s", packageID)

	noteReq := &packages.AddHistoryNotesRequest{
		Note: fmt.Sprintf("Acceptance test note at %s", time.Now().Format(time.RFC3339)),
	}
	noteResp, err := svc.AddPackageHistoryNotesV1(ctx, packageID, noteReq)
	require.NoError(t, err)
	require.NotNil(t, noteResp)
	assert.Equal(t, 201, noteResp.StatusCode)
	acc.LogTestSuccess(t, "History note added")

	history, histResp, err := svc.GetPackageHistoryV1(ctx, packageID, nil)
	require.NoError(t, err)
	require.NotNil(t, history)
	assert.Equal(t, 200, histResp.StatusCode)
	assert.GreaterOrEqual(t, history.TotalCount, 1)
	acc.LogTestSuccess(t, "History entries: %d", history.TotalCount)

	// 6. Delete
	acc.LogTestStage(t, "Delete", "Deleting package ID=%s", packageID)

	deleteResp, err := svc.DeletePackageByIDV1(ctx, packageID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Package ID=%s deleted", packageID)
}

// =============================================================================
// TestAcceptance_Packages_ListWithRSQLFilter
// =============================================================================

func TestAcceptance_Packages_ListWithRSQLFilter(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Packages
	ctx := context.Background()

	name := uniquePackageName("acc-rsql-package")
	createReq := &packages.RequestPackage{
		PackageName:           name,
		FileName:              "acc-rsql.pkg",
		CategoryID:            "-1",
		Info:                  "RSQL filter test",
		Priority:              5,
		FillUserTemplate:      packages.BoolPtr(false),
		RebootRequired:        packages.BoolPtr(false),
		OSInstall:             packages.BoolPtr(false),
		SuppressUpdates:       packages.BoolPtr(false),
		SuppressFromDock:      packages.BoolPtr(false),
		SuppressEula:          packages.BoolPtr(false),
		SuppressRegistration:  packages.BoolPtr(false),
	}

	created, _, err := svc.CreatePackageV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	packageID := created.ID
	acc.LogTestSuccess(t, "Created package ID=%s name=%q", packageID, name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeletePackageByIDV1(cleanupCtx, packageID)
		acc.LogCleanupDeleteError(t, "package", packageID, delErr)
	})

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`packageName=="%s"`, name),
	}

	list, listResp, err := svc.ListPackagesV1(ctx, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, p := range list.Results {
		if p.ID == packageID {
			found = true
			assert.Equal(t, name, p.PackageName)
			break
		}
	}
	assert.True(t, found, "package should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target package found=%v", list.TotalCount, found)
}

// =============================================================================
// TestAcceptance_Packages_DeleteMultiple
// =============================================================================

func TestAcceptance_Packages_DeleteMultiple(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Packages
	ctx := context.Background()

	ids := make([]string, 0, 2)
	for i := 0; i < 2; i++ {
		req := &packages.RequestPackage{
			PackageName:           uniquePackageName(fmt.Sprintf("acc-bulk-delete-%d", i)),
			FileName:              "acc-bulk.pkg",
			CategoryID:            "-1",
			Info:                  "Bulk delete test",
			Priority:              9,
			FillUserTemplate:      packages.BoolPtr(false),
			RebootRequired:        packages.BoolPtr(false),
			OSInstall:             packages.BoolPtr(false),
			SuppressUpdates:       packages.BoolPtr(false),
			SuppressFromDock:      packages.BoolPtr(false),
			SuppressEula:          packages.BoolPtr(false),
			SuppressRegistration:  packages.BoolPtr(false),
		}
		created, _, err := svc.CreatePackageV1(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, created)
		ids = append(ids, created.ID)
		acc.LogTestSuccess(t, "Created package ID=%s", created.ID)
	}

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		for _, id := range ids {
			_, delErr := svc.DeletePackageByIDV1(cleanupCtx, id)
			acc.LogCleanupDeleteError(t, "package", id, delErr)
		}
	})

	bulkReq := &packages.DeletePackagesByIDRequest{IDs: ids}
	bulkResp, err := svc.DeletePackagesByIDV1(ctx, bulkReq)
	require.NoError(t, err)
	require.NotNil(t, bulkResp)
	assert.Equal(t, 204, bulkResp.StatusCode)
	acc.LogTestSuccess(t, "Bulk delete of %d packages succeeded", len(ids))

	for _, id := range ids {
		_, _, getErr := svc.GetPackageByIDV1(ctx, id)
		assert.Error(t, getErr, "deleted package ID=%s should return error on Get", id)
	}
}

// =============================================================================
// TestAcceptance_Packages_ValidationErrors
// =============================================================================

func TestAcceptance_Packages_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Packages

	t.Run("GetPackageByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetPackageByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "package ID is required")
	})

	t.Run("CreatePackageV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreatePackageV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdatePackageByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdatePackageByIDV1(context.Background(), "", &packages.ResourcePackage{PackageName: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("DeletePackageByIDV1_EmptyID", func(t *testing.T) {
		_, err := svc.DeletePackageByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "package ID is required")
	})

	t.Run("DeletePackagesByIDV1_EmptyIDs", func(t *testing.T) {
		_, err := svc.DeletePackagesByIDV1(context.Background(), &packages.DeletePackagesByIDRequest{IDs: []string{}})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "ids are required")
	})

	t.Run("GetPackageHistoryV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetPackageHistoryV1(context.Background(), "", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "package ID is required")
	})

	t.Run("AddPackageHistoryNotesV1_NilRequest", func(t *testing.T) {
		_, err := svc.AddPackageHistoryNotesV1(context.Background(), "1", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request body is required")
	})
}
