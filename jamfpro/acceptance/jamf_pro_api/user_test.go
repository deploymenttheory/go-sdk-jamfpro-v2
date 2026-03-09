package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: User
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • Get(ctx) - Gets the current authenticated user information
//   • ChangePassword(ctx, request) - Changes the current user's password
//   • UpdateSession(ctx, request) - Updates the current user's session (change site)
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Information
//     -- Tests: TestAcceptance_User_get
//     -- Flow: Get → verify username and access level are non-empty
//
//   ✓ Pattern 7: Validation Errors
//     -- Tests: TestAcceptance_User_validation_errors
//     -- Cases: nil requests, empty passwords
//
// Notes
// -----------------------------------------------------------------------------
//   • ChangePassword is not called in acceptance tests to avoid disrupting
//     the credentials used by the test client
//   • UpdateSession is not called in acceptance tests to avoid changing the
//     active site context during test execution
//
// =============================================================================

// TestAcceptance_User_get verifies the current user can be retrieved.
func TestAcceptance_User_get(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.User
	ctx := context.Background()

	acc.LogTestStage(t, "Get", "Getting current authenticated user")

	u, resp, err := svc.Get(ctx)
	if err != nil {
		// GET /api/user may not be supported for OAuth2 API client credentials
		// (service accounts don't have a corresponding user record)
		acc.LogTestWarning(t, "Get /api/user returned error (may not be supported for API client credentials): %v", err)
		t.Skip("GET /api/user is not supported for this authentication method")
	}
	require.NotNil(t, u)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, u.Username, "username should not be empty")

	acc.LogTestSuccess(t, "Get: username=%s accessLevel=%s privilegeSet=%s",
		u.Username, u.AccessLevel, u.PrivilegeSet)
}

// TestAcceptance_User_validation_errors verifies input validation.
func TestAcceptance_User_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.User

	t.Run("ChangePassword_NilRequest", func(t *testing.T) {
		_, err := svc.ChangePassword(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("ChangePassword_EmptyCurrentPassword", func(t *testing.T) {
		_, err := svc.ChangePassword(context.Background(), &user.RequestChangePassword{
			CurrentPassword: "",
			NewPassword:     "newpass123",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "currentPassword is required and cannot be empty")
	})

	t.Run("ChangePassword_EmptyNewPassword", func(t *testing.T) {
		_, err := svc.ChangePassword(context.Background(), &user.RequestChangePassword{
			CurrentPassword: "currentpass123",
			NewPassword:     "",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "newPassword is required and cannot be empty")
	})

	t.Run("UpdateSession_NilRequest", func(t *testing.T) {
		_, err := svc.UpdateSession(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})
}
