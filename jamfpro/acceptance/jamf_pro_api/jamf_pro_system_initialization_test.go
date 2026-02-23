package jamf_pro_api

import (
	"context"
	"os"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_pro_system_initialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Jamf Pro System Initialization
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • Initialize(ctx, request) - Initializes a fresh Jamf Pro Server installation
//   • InitializeDatabaseConnection(ctx, password) - Sets up database password during startup
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern: Destructive Setup Operations
//     -- Reason: These endpoints perform one-time initialization of Jamf Pro
//     -- Tests are SKIPPED by default - require JAMF_RUN_SYSTEM_INIT_TESTS=1
//     -- WARNING: Running these tests against an initialized instance may fail
//        or have unintended effects. Use only against uninitialized test instances.
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Initialize - Verify API call succeeds (when run against uninitialized instance)
//   ✓ InitializeDatabaseConnection - Verify API call succeeds (when run during DB setup)
//
// Notes
// -----------------------------------------------------------------------------
//   • These are one-time setup operations - not idempotent
//   • Acceptance tests skip unless JAMF_RUN_SYSTEM_INIT_TESTS=1
//   • Use only in isolated test environments with uninitialized Jamf Pro
//
// =============================================================================

func TestAcceptance_JamfProSystemInitialization_Initialize(t *testing.T) {
	if os.Getenv("JAMF_RUN_SYSTEM_INIT_TESTS") != "1" {
		t.Skip("Skipping system initialization test (set JAMF_RUN_SYSTEM_INIT_TESTS=1 to run)")
	}
	acc.RequireClient(t)
	svc := acc.Client.JamfProSystemInitialization
	ctx := context.Background()

	request := &jamf_pro_system_initialization.ResourceSystemInitialize{
		ActivationCode:  os.Getenv("JAMF_ACTIVATION_CODE"),
		InstitutionName: "Acceptance Test Institution",
		EulaAccepted:    true,
		Username:        "admin",
		Password:        os.Getenv("JAMF_INIT_PASSWORD"),
		Email:           "admin@example.com",
		JssUrl:          acc.Config.InstanceDomain,
	}
	if request.ActivationCode == "" || request.Password == "" {
		t.Skip("JAMF_ACTIVATION_CODE and JAMF_INIT_PASSWORD required for Initialize test")
	}

	resp, err := svc.Initialize(ctx, request)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.True(t, resp.StatusCode >= 200 && resp.StatusCode < 300, "expected success status")
}

func TestAcceptance_JamfProSystemInitialization_InitializeDatabaseConnection(t *testing.T) {
	if os.Getenv("JAMF_RUN_SYSTEM_INIT_TESTS") != "1" {
		t.Skip("Skipping database initialization test (set JAMF_RUN_SYSTEM_INIT_TESTS=1 to run)")
	}
	acc.RequireClient(t)
	svc := acc.Client.JamfProSystemInitialization
	ctx := context.Background()

	password := os.Getenv("JAMF_DB_INIT_PASSWORD")
	if password == "" {
		t.Skip("JAMF_DB_INIT_PASSWORD required for InitializeDatabaseConnection test")
	}

	resp, err := svc.InitializeDatabaseConnection(ctx, password)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.True(t, resp.StatusCode >= 200 && resp.StatusCode < 300, "expected success status")
}
