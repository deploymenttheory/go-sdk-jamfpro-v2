package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/enrollment"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Enrollment
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   V4 Endpoints:
//   • GetV4(ctx) - Retrieves enrollment settings
//   • UpdateV4(ctx, request) - Updates enrollment settings
//
//   V3 Endpoints:
//   • ListAccessGroupsV3(ctx, rsqlQuery) - Lists ADUE access groups
//   • GetAccessGroupByIDV3(ctx, id) - Retrieves access group by ID
//   • CreateAccessGroupV3(ctx, request) - Creates ADUE access group
//   • UpdateAccessGroupByIDV3(ctx, id, request) - Updates access group
//   • DeleteAccessGroupByIDV3(ctx, id) - Deletes access group
//   • ListLanguageMessagesV3(ctx) - Lists all language messages
//   • GetLanguageMessageV3(ctx, languageCode) - Gets language message
//   • UpdateLanguageMessageV3(ctx, languageCode, request) - Updates language message
//   • DeleteLanguageMessageV3(ctx, languageCode) - Deletes language message
//   • DeleteMultipleLanguageMessagesV3(ctx, request) - Bulk deletes language messages
//   • ListLanguageCodesV3(ctx) - Lists available language codes
//
//   V2 Endpoints:
//   • GetHistoryV2(ctx, rsqlQuery) - Retrieves enrollment history
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 2: Settings/Configuration Testing
//     -- Reason: V4 settings endpoint returns singleton configuration
//     -- Tests: TestAcceptance_Enrollment_get_and_update_v4
//     -- Flow: Get settings → Update settings → Verify changes
//
//   ✓ Pattern 4: Full Lifecycle Testing
//     -- Reason: Access groups support full CRUD operations
//     -- Tests: TestAcceptance_Enrollment_access_group_lifecycle_v3
//     -- Flow: Create → Get → Update → Delete
//
//   ✓ Pattern 3: Read-Only List Testing
//     -- Reason: Language messages are pre-configured in Jamf Pro
//     -- Tests: TestAcceptance_Enrollment_language_messages_v3
//     -- Flow: List all → Get by code → List available codes
//
//   ✓ Pattern 3: Read-Only List Testing
//     -- Reason: Enrollment history is read-only audit data
//     -- Tests: TestAcceptance_Enrollment_history_v2
//     -- Flow: Get history → Verify structure
//
//   ✓ Pattern 7: Validation Errors
//     -- Reason: Client-side validation prevents invalid API calls
//     -- Tests: TestAcceptance_Enrollment_validation_errors
//     -- Cases: Empty IDs, invalid language codes, nil requests
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ V4 operations (Get, Update)
//   ✓ V3 ADUE access group lifecycle (Create, Get, Update, Delete)
//   ✓ V3 language message operations (List, Get, validation)
//   ✓ V3 language code listing
//   ✓ V2 history retrieval
//   ✓ Input validation and error handling
//
// Notes
// -----------------------------------------------------------------------------
//   • Enrollment settings are global configuration (v4)
//   • Access groups link LDAP groups to enrollment options (v3)
//   • Language messages customize enrollment UI per locale (v3)
//   • Language code validation enforces valid locales (v3)
//   • Enrollment history provides audit trail (v2)
//   • Tests handle multi-version API gracefully
//
// =============================================================================

func TestAcceptance_Enrollment_get_and_update_v4(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Enrollment
	ctx := context.Background()

	// Get current enrollment settings
	result, resp, err := svc.GetV4(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify structure
	assert.NotNil(t, result.InstallSingleProfile)
	assert.NotNil(t, result.SigningMdmProfileEnabled)
	assert.NotNil(t, result.MacOsEnterpriseEnrollmentEnabled)

	// Update settings (toggle a setting and toggle back)
	originalRestrictReenrollment := result.RestrictReenrollment
	updateRequest := &enrollment.ResourceEnrollment{
		InstallSingleProfile:             result.InstallSingleProfile,
		SigningMdmProfileEnabled:         result.SigningMdmProfileEnabled,
		MdmSigningCertificate:            result.MdmSigningCertificate,
		RestrictReenrollment:             !originalRestrictReenrollment,
		FlushLocationInformation:         result.FlushLocationInformation,
		FlushLocationHistoryInformation:  result.FlushLocationHistoryInformation,
		FlushPolicyHistory:               result.FlushPolicyHistory,
		FlushExtensionAttributes:         result.FlushExtensionAttributes,
		FlushSoftwareUpdatePlans:         result.FlushSoftwareUpdatePlans,
		MacOsEnterpriseEnrollmentEnabled: result.MacOsEnterpriseEnrollmentEnabled,
		ManagementUsername:               result.ManagementUsername,
		CreateManagementAccount:          result.CreateManagementAccount,
		HideManagementAccount:            result.HideManagementAccount,
		AllowSshOnlyManagementAccount:    result.AllowSshOnlyManagementAccount,
		EnsureSshRunning:                 result.EnsureSshRunning,
		LaunchSelfService:                result.LaunchSelfService,
		SignQuickAdd:                     result.SignQuickAdd,
		DeveloperCertificateIdentity:     result.DeveloperCertificateIdentity,
		DeveloperCertificateIdentityDetails: result.DeveloperCertificateIdentityDetails,
		MdmSigningCertificateDetails:     result.MdmSigningCertificateDetails,
		IosEnterpriseEnrollmentEnabled:   result.IosEnterpriseEnrollmentEnabled,
		IosPersonalEnrollmentEnabled:     result.IosPersonalEnrollmentEnabled,
		PersonalDeviceEnrollmentType:     result.PersonalDeviceEnrollmentType,
		AccountDrivenUserEnrollmentEnabled: result.AccountDrivenUserEnrollmentEnabled,
	}

	updated, resp, err := svc.UpdateV4(ctx, updateRequest)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, !originalRestrictReenrollment, updated.RestrictReenrollment)

	// Restore original setting
	restoreRequest := &enrollment.ResourceEnrollment{
		InstallSingleProfile:             result.InstallSingleProfile,
		SigningMdmProfileEnabled:         result.SigningMdmProfileEnabled,
		MdmSigningCertificate:            result.MdmSigningCertificate,
		RestrictReenrollment:             originalRestrictReenrollment,
		FlushLocationInformation:         result.FlushLocationInformation,
		FlushLocationHistoryInformation:  result.FlushLocationHistoryInformation,
		FlushPolicyHistory:               result.FlushPolicyHistory,
		FlushExtensionAttributes:         result.FlushExtensionAttributes,
		FlushSoftwareUpdatePlans:         result.FlushSoftwareUpdatePlans,
		MacOsEnterpriseEnrollmentEnabled: result.MacOsEnterpriseEnrollmentEnabled,
		ManagementUsername:               result.ManagementUsername,
		CreateManagementAccount:          result.CreateManagementAccount,
		HideManagementAccount:            result.HideManagementAccount,
		AllowSshOnlyManagementAccount:    result.AllowSshOnlyManagementAccount,
		EnsureSshRunning:                 result.EnsureSshRunning,
		LaunchSelfService:                result.LaunchSelfService,
		SignQuickAdd:                     result.SignQuickAdd,
		DeveloperCertificateIdentity:     result.DeveloperCertificateIdentity,
		DeveloperCertificateIdentityDetails: result.DeveloperCertificateIdentityDetails,
		MdmSigningCertificateDetails:     result.MdmSigningCertificateDetails,
		IosEnterpriseEnrollmentEnabled:   result.IosEnterpriseEnrollmentEnabled,
		IosPersonalEnrollmentEnabled:     result.IosPersonalEnrollmentEnabled,
		PersonalDeviceEnrollmentType:     result.PersonalDeviceEnrollmentType,
		AccountDrivenUserEnrollmentEnabled: result.AccountDrivenUserEnrollmentEnabled,
	}

	restored, resp, err := svc.UpdateV4(ctx, restoreRequest)
	require.NoError(t, err)
	require.NotNil(t, restored)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, originalRestrictReenrollment, restored.RestrictReenrollment)
}

func TestAcceptance_Enrollment_access_group_lifecycle_v3(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Enrollment
	ctx := context.Background()

	// Create access group
	createRequest := &enrollment.ResourceAccountDrivenUserEnrollmentAccessGroup{
		GroupID:                            "test-group-id",
		LdapServerID:                       "1",
		Name:                               "sdkv2_acc_Test ADUE Access Group",
		SiteID:                             "-1",
		EnterpriseEnrollmentEnabled:        true,
		PersonalEnrollmentEnabled:          false,
		AccountDrivenUserEnrollmentEnabled: true,
		RequireEula:                        false,
	}

	created, resp, err := svc.CreateAccessGroupV3(ctx, createRequest)
	require.NoError(t, err)
	require.NotNil(t, created)
	assert.Equal(t, 201, resp.StatusCode)
	assert.NotEmpty(t, created.ID)
	createdID := created.ID

	// Ensure cleanup
	defer func() {
		_, _ = svc.DeleteAccessGroupByIDV3(ctx, createdID)
	}()

	// Get by ID
	fetched, resp, err := svc.GetAccessGroupByIDV3(ctx, createdID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, createdID, fetched.ID)
	assert.Equal(t, "Test ADUE Access Group", fetched.Name)

	// Update
	updateRequest := &enrollment.ResourceAccountDrivenUserEnrollmentAccessGroup{
		GroupID:                            fetched.GroupID,
		LdapServerID:                       fetched.LdapServerID,
		Name:                               "sdkv2_acc_Test ADUE Access Group Updated",
		SiteID:                             fetched.SiteID,
		EnterpriseEnrollmentEnabled:        false,
		PersonalEnrollmentEnabled:          false,
		AccountDrivenUserEnrollmentEnabled: true,
		RequireEula:                        true,
	}

	updated, resp, err := svc.UpdateAccessGroupByIDV3(ctx, createdID, updateRequest)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "Test ADUE Access Group Updated", updated.Name)
	assert.Equal(t, false, updated.EnterpriseEnrollmentEnabled)
	assert.Equal(t, true, updated.RequireEula)

	// Delete
	resp, err = svc.DeleteAccessGroupByIDV3(ctx, createdID)
	require.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode)

	// Verify deletion
	_, resp, err = svc.GetAccessGroupByIDV3(ctx, createdID)
	assert.Error(t, err)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestAcceptance_Enrollment_list_access_groups_v3(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Enrollment
	ctx := context.Background()

	// List all access groups
	listResult, resp, err := svc.ListAccessGroupsV3(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, listResult)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, listResult.TotalCount, 0)
}

func TestAcceptance_Enrollment_language_messages_v3(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Enrollment
	ctx := context.Background()

	// List available language codes
	codes, resp, err := svc.ListLanguageCodesV3(ctx)
	require.NoError(t, err)
	require.NotNil(t, codes)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, codes)

	// List all language messages
	messages, resp, err := svc.ListLanguageMessagesV3(ctx)
	require.NoError(t, err)
	require.NotNil(t, messages)
	assert.Equal(t, 200, resp.StatusCode)

	// If English exists, test get by code
	if len(messages) > 0 {
		firstMessage := messages[0]

		message, resp, err := svc.GetLanguageMessageV3(ctx, firstMessage.LanguageCode)
		require.NoError(t, err)
		require.NotNil(t, message)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, firstMessage.LanguageCode, message.LanguageCode)
	}
}

func TestAcceptance_Enrollment_history_v2(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Enrollment
	ctx := context.Background()

	// Get enrollment history
	history, resp, err := svc.GetHistoryV2(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, history)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, history.TotalCount, 0)
}

func TestAcceptance_Enrollment_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Enrollment
	ctx := context.Background()

	// Test nil request validation (UpdateV4)
	result, resp, err := svc.UpdateV4(ctx, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "enrollment settings request cannot be nil")

	// Test empty ID validation (GetAccessGroupByIDV3)
	accessGroup, resp, err := svc.GetAccessGroupByIDV3(ctx, "")
	assert.Error(t, err)
	assert.Nil(t, accessGroup)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "access group ID is required")

	// Test nil request validation (CreateAccessGroupV3)
	created, resp, err := svc.CreateAccessGroupV3(ctx, nil)
	assert.Error(t, err)
	assert.Nil(t, created)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "access group request cannot be nil")

	// Test empty language code validation (GetLanguageMessageV3)
	message, resp, err := svc.GetLanguageMessageV3(ctx, "")
	assert.Error(t, err)
	assert.Nil(t, message)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "language code is required")

	// Test invalid language code validation (GetLanguageMessageV3 with invalid code)
	message, resp, err = svc.GetLanguageMessageV3(ctx, "invalid-code")
	assert.Error(t, err)
	assert.Nil(t, message)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "invalid language code")

	// Test nil request validation (DeleteMultipleLanguageMessagesV3)
	resp, err = svc.DeleteMultipleLanguageMessagesV3(ctx, nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "delete request cannot be nil")
}
