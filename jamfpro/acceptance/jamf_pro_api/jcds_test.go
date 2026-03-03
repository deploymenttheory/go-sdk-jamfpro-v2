package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: JCDS (Jamf Cloud Distribution Service)
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetPackagesV1(ctx) - Lists all files stored in JCDS
//   • GetPackageURIByNameV1(ctx, packageName) - Gets S3 URI for a package
//   • RenewCredentialsV1(ctx) - Obtains fresh AWS credentials for JCDS
//   • CreatePackageV1(ctx, filePath) - Uploads a package to JCDS via S3
//   • DeletePackageV1(ctx, filePath) - Deletes a package from JCDS via S3
//   • RefreshInventoryV1(ctx) - Refreshes Jamf Pro's JCDS package inventory
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 4: Read-Only with Existing Data
//     -- Tests: TestAcceptance_JCDS_get_packages_v1
//     -- Flow: GetPackagesV1 → if packages exist: GetPackageURIByNameV1
//
//   ✓ Pattern 3: Read-Only (credentials renewal)
//     -- Tests: TestAcceptance_JCDS_renew_credentials_v1
//
//   ✓ Pattern 7: Validation Errors
//     -- Tests: TestAcceptance_JCDS_validation_errors
//     -- Cases: GetPackageURIByNameV1("") → "package name is required"
//
// Notes
// -----------------------------------------------------------------------------
//   • Upload/download operations require real package files; tested via validation only
//   • JCDS must be configured and enabled on the Jamf Pro instance
//
// =============================================================================

// TestAcceptance_JCDS_get_packages_v1 lists packages and fetches URI for the first one.
func TestAcceptance_JCDS_get_packages_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JCDS
	ctx := context.Background()

	acc.LogTestStage(t, "GetPackagesV1", "Listing JCDS packages")

	packages, resp, err := svc.GetPackagesV1(ctx)
	if err != nil && resp != nil && resp.StatusCode == 404 {
		t.Skip("JCDS not configured on this tenant (404 NOT_FOUND)")
	}
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, len(packages), 0)
	acc.LogTestSuccess(t, "GetPackagesV1: %d package(s) found", len(packages))

	if len(packages) == 0 {
		acc.LogTestWarning(t, "No JCDS packages found; skipping GetPackageURIByNameV1 test")
		return
	}

	// GetPackageURIByNameV1 using first package filename
	firstPkg := packages[0]
	acc.LogTestStage(t, "GetPackageURIByNameV1", "Getting URI for package fileName=%s", firstPkg.FileName)

	uri, uriResp, err := svc.GetPackageURIByNameV1(ctx, firstPkg.FileName)
	require.NoError(t, err)
	require.NotNil(t, uri)
	assert.Equal(t, 200, uriResp.StatusCode)
	assert.NotEmpty(t, uri.URI, "URI should not be empty")
	acc.LogTestSuccess(t, "GetPackageURIByNameV1: fileName=%s uri=%s", firstPkg.FileName, uri.URI)
}

// TestAcceptance_JCDS_renew_credentials_v1 verifies that JCDS credentials can be renewed.
func TestAcceptance_JCDS_renew_credentials_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JCDS
	ctx := context.Background()

	acc.LogTestStage(t, "RenewCredentialsV1", "Renewing JCDS AWS credentials")

	creds, resp, err := svc.RenewCredentialsV1(ctx)
	if err != nil && resp != nil && resp.StatusCode == 404 {
		t.Skip("JCDS not configured on this tenant (404 NOT_FOUND)")
	}
	require.NoError(t, err)
	require.NotNil(t, creds)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, creds.AccessKeyID, "accessKeyId should not be empty")
	assert.NotEmpty(t, creds.SecretAccessKey, "secretAccessKey should not be empty")
	assert.NotEmpty(t, creds.Region, "region should not be empty")
	acc.LogTestSuccess(t, "RenewCredentialsV1: region=%s bucketName=%s", creds.Region, creds.BucketName)
}

// TestAcceptance_JCDS_validation_errors verifies input validation.
func TestAcceptance_JCDS_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JCDS

	t.Run("GetPackageURIByNameV1_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetPackageURIByNameV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "package name is required")
	})
}
