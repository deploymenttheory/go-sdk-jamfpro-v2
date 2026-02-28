package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Mobile Device Enrollment Profile
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetDownloadProfileV1(ctx, id) - Downloads the MDM enrollment profile (binary)
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 4: Read-Only with Existing Data
//     -- Tests: TestAcceptance_MobileDeviceEnrollmentProfile_download_v1
//     -- Flow: List classic mobile device enrollment profiles → download first
//     -- Skips if no enrollment profiles are configured
//
//   ✓ Pattern 7: Validation Errors
//     -- Tests: TestAcceptance_MobileDeviceEnrollmentProfile_validation_errors
//     -- Cases: GetDownloadProfileV1("") → "id is required"
//
// Notes
// -----------------------------------------------------------------------------
//   • Response is a binary Apple Aspen config file (.mobileconfig)
//   • Profile IDs are sourced from the Classic API mobile device enrollment profiles
//
// =============================================================================

// TestAcceptance_MobileDeviceEnrollmentProfile_download_v1 downloads an enrollment profile.
func TestAcceptance_MobileDeviceEnrollmentProfile_download_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.MobileDeviceEnrollmentProfile
	ctx := context.Background()

	// Get a list of classic mobile device enrollment profiles to find an ID
	profileList, _, err := acc.Client.ClassicMobileDeviceEnrollmentProfiles.List(ctx)
	require.NoError(t, err, "failed to list classic mobile device enrollment profiles")

	if profileList == nil || len(profileList.Results) == 0 {
		t.Skip("No mobile device enrollment profiles configured; skipping download test")
	}

	profileID := fmt.Sprintf("%d", profileList.Results[0].ID)
	acc.LogTestStage(t, "GetDownloadProfileV1", "Downloading enrollment profile ID=%s", profileID)

	data, resp, err := svc.GetDownloadProfileV1(ctx, profileID)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, data, "profile download should return non-empty data")

	acc.LogTestSuccess(t, "GetDownloadProfileV1: ID=%s size=%d bytes", profileID, len(data))
}

// TestAcceptance_MobileDeviceEnrollmentProfile_validation_errors verifies input validation.
func TestAcceptance_MobileDeviceEnrollmentProfile_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.MobileDeviceEnrollmentProfile

	t.Run("GetDownloadProfileV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetDownloadProfileV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})
}
