package jamf_pro_api

import (
	"context"
	"math/rand"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/icons"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Icons
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetByIDV1(ctx, id) - Retrieves icon metadata by ID
//   • UploadV1(ctx, fileReader, fileSize, fileName) - Uploads an icon image (Create)
//   • UploadV1FromFile(ctx, filePath) - Helper to upload from file path
//   • DownloadV1(ctx, id, res, scale) - Downloads icon image bytes (various resolutions)
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 4: Read-Only with Existing Data
//     -- Reason: Tests use existing icons from the system (IDs 1-100)
//     -- Tests: TestAcceptance_Icons_GetByID
//     -- Flow: Find existing icon → GetByID → Verify metadata
//     -- Note: Upload operation not tested in acceptance tests
//
//   ✗ Download operation (commented out)
//     -- TestAcceptance_Icons_Download exists but is commented out
//     -- Reason: Server-side issues with corrupted/missing icon files causing 500 errors
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Get icon metadata by ID (using existing icons)
//   ✗ Upload icon (not yet tested - should be added)
//   ✗ Download icon image (commented out - server issues)
//   ✗ Delete icon (no delete operation available in API)
//
// Notes
// -----------------------------------------------------------------------------
//   • Icons are used for Self Service applications and policies
//   • No List operation available - tests search IDs 1-100 to find valid icons
//   • No Update or Delete operations available in the API
//   • Upload creates new icons but no way to clean them up (no delete)
//   • Download test commented out due to server-side corrupted files (500 errors)
//   • Resolution options for download: "original", "300", "512"
//   • Scale options: "0" (original), non-zero (scaled to 300)
//   • Upload uses multipart/form-data with "file" field
//   • TODO: Add upload test (but note cleanup issues - no delete available)
//   • TODO: Uncomment/fix download test when server-side issues resolved
//
// =============================================================================

// findValidIconID tries IDs in random order (1–100) until GetByIDV1 returns 200.
// Logs success to test notifications. Returns (id, result, true) when found, (0, nil, false) when none exist.
func findValidIconID(t *testing.T) (int, *icons.ResourceIcon, bool) {
	t.Helper()
	acc.RequireClient(t)
	svc := acc.Client.Icons
	ctx := context.Background()

	acc.LogTestStage(t, "Icons", "Finding icon: trying IDs 1–100 until GetByID returns 200")
	perm := rand.Perm(100)
	for i := range perm {
		id := perm[i] + 1
		result, resp, err := svc.GetByIDV1(ctx, id)
		if err != nil {
			continue
		}
		if resp != nil && resp.StatusCode == 200 && result != nil {
			acc.LogTestSuccess(t, "GetByID succeeded for icon ID=%d", id)
			return id, result, true
		}
	}
	return 0, nil, false
}

func TestAcceptance_Icons_GetByID(t *testing.T) {
	id, result, ok := findValidIconID(t)
	if !ok {
		t.Skip("no icon with ID in 1–100 found in this environment")
		return
	}
	require.NotNil(t, result)
	assert.GreaterOrEqual(t, id, 1)
	assert.Equal(t, id, result.ID)
}

// func TestAcceptance_Icons_Download(t *testing.T) {
// 	acc.RequireClient(t)
// 	svc := acc.Client.Icons
// 	ctx := context.Background()

// 	// Try multiple icons since some may have corrupted/missing files on server (500 errors)
// 	acc.LogTestStage(t, "Icons", "Finding downloadable icon: trying IDs 1–100")
// 	perm := rand.Perm(100)
// 	var lastErr error

// 	for i := range perm {
// 		id := perm[i] + 1
// 		// First check if icon exists
// 		_, resp, err := svc.GetByIDV1(ctx, id)
// 		if err != nil || resp == nil || resp.StatusCode != 200 {
// 			continue
// 		}

// 		// Try to download it
// 		acc.LogTestStage(t, "Icons", "Attempting download for icon ID=%d", id)
// 		body, resp, err := svc.DownloadV1(ctx, id, "original", "0")
// 		if err != nil {
// 			lastErr = err
// 			acc.LogTestWarning(t, "Download failed for icon ID=%d (may be corrupted on server): %v", id, err)
// 			continue
// 		}

// 		// Success!
// 		require.NotNil(t, resp)
// 		assert.Equal(t, 200, resp.StatusCode)
// 		assert.NotEmpty(t, body)
// 		acc.LogTestSuccess(t, "Successfully downloaded icon ID=%d (%d bytes)", id, len(body))
// 		return
// 	}

// 	// If we get here, no icons were downloadable
// 	if lastErr != nil {
// 		t.Skipf("no downloadable icons found (all returned errors, last: %v)", lastErr)
// 	} else {
// 		t.Skip("no icons with ID in 1–100 found in this environment")
// 	}
// }
