package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/enrollment_customization_preview"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/enrollment_customizations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"resty.dev/v3"
)

// =============================================================================
// Acceptance Tests: Enrollment Customization Preview
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ParseMarkdown(ctx, request) - Parses markdown to HTML
//   • GetAllPanels(ctx, id) - Gets all panels for a customization
//   • GetPanelByID(ctx, id, panelID) - Gets a specific panel
//   • DeletePanel(ctx, id, panelID) - Deletes any panel
//   • CreateLdapPanel / GetLdapPanel / UpdateLdapPanel / DeleteLdapPanel
//   • CreateSsoPanel / GetSsoPanel / UpdateSsoPanel / DeleteSsoPanel
//   • CreateTextPanel / GetTextPanel / UpdateTextPanel / DeleteTextPanel
//   • GetTextPanelMarkdown(ctx, id, panelID) - Gets text panel markdown
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern: ParseMarkdown (standalone, no dependency)
//     -- Tests: TestAcceptance_EnrollmentCustomizationPreview_parse_markdown
//
//   ✓ Pattern 1: Text Panel Lifecycle (Create → Get → Update → Delete)
//     -- Requires an enrollment customization; creates one for the test
//     -- Tests: TestAcceptance_EnrollmentCustomizationPreview_text_panel_lifecycle
//
//   ✓ Pattern 7: Validation Errors
//     -- Tests: TestAcceptance_EnrollmentCustomizationPreview_validation_errors
//
// =============================================================================

// TestAcceptance_EnrollmentCustomizationPreview_parse_markdown verifies markdown parsing.
func TestAcceptance_EnrollmentCustomizationPreview_parse_markdown(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.EnrollmentCustomizationPreview
	ctx := context.Background()

	acc.LogTestStage(t, "ParseMarkdown", "Parsing markdown to HTML")

	req := &enrollment_customization_preview.RequestParseMarkdown{
		Markdown: "# Hello\n\nThis is **acceptance test** markdown.",
	}

	result, resp, err := svc.ParseMarkdown(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, result.Markdown, "parsed HTML should not be empty")

	acc.LogTestSuccess(t, "ParseMarkdown: output length=%d", len(result.Markdown))
}

// TestAcceptance_EnrollmentCustomizationPreview_text_panel_lifecycle exercises
// text panel CRUD within a freshly created enrollment customization.
func TestAcceptance_EnrollmentCustomizationPreview_text_panel_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	previewSvc := acc.Client.EnrollmentCustomizationPreview
	ecSvc := acc.Client.EnrollmentCustomizations
	ctx := context.Background()

	// Create an enrollment customization to use as the parent
	acc.LogTestStage(t, "Setup", "Creating enrollment customization for panel tests")

	ecName := acc.UniqueName("sdkv2_acc_ec-for-panel")
	ec, ecResp, err := ecSvc.CreateV2(ctx, &enrollment_customizations.ResourceEnrollmentCustomization{
		DisplayName: ecName,
		Description: "Acceptance test EC for panel preview",
		BrandingSettings: enrollment_customizations.SubsetBrandingSettings{
			ButtonColor:     "0066CC",
			ButtonTextColor: "FFFFFF",
		},
	})
	require.NoError(t, err, "failed to create enrollment customization")
	require.NotNil(t, ec)
	assert.Contains(t, []int{200, 201}, ecResp.StatusCode)

	ecID := ec.ID
	acc.LogTestSuccess(t, "Created enrollment customization ID=%s", ecID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := ecSvc.DeleteByIDV2(cleanupCtx, ecID)
		acc.LogCleanupDeleteError(t, "enrollment customization", ecID, delErr)
	})

	// 1. Create text panel
	acc.LogTestStage(t, "CreateTextPanel", "Creating text panel in enrollment customization ID=%s", ecID)

	createPanelReq := &enrollment_customization_preview.ResourceTextPanel{
		DisplayName:        "Acceptance Test Panel",
		Rank:               1,
		Title:              "Welcome",
		Body:               "# Welcome\n\nThis is an acceptance test.",
		BackButtonText:     "Back",
		ContinueButtonText: "Continue",
	}

	createdPanel, createPanelResp, err := previewSvc.CreateTextPanel(ctx, ecID, createPanelReq)
	require.NoError(t, err)
	require.NotNil(t, createdPanel)
	assert.Contains(t, []int{200, 201}, createPanelResp.StatusCode)
	assert.Positive(t, createdPanel.ID)

	panelID := fmt.Sprintf("%d", createdPanel.ID)
	acc.LogTestSuccess(t, "Text panel created with ID=%s", panelID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = previewSvc.DeletePanel(cleanupCtx, ecID, panelID)
	})

	// 2. GetAllPanels
	acc.LogTestStage(t, "GetAllPanels", "Listing all panels for EC ID=%s", ecID)

	panels, panelsResp, err := previewSvc.GetAllPanels(ctx, ecID)
	require.NoError(t, err)
	require.NotNil(t, panels)
	assert.Equal(t, 200, panelsResp.StatusCode)

	found := false
	for _, p := range panels.Panels {
		if fmt.Sprintf("%d", p.ID) == panelID {
			found = true
			break
		}
	}
	assert.True(t, found, "created panel should appear in GetAllPanels")

	// 3. GetTextPanel (with retry for eventual consistency)
	acc.LogTestStage(t, "GetTextPanel", "Getting text panel ID=%s", panelID)

	var fetchedPanel *enrollment_customization_preview.ResourceTextPanel
	var fetchPanelResp *resty.Response
	err = acc.RetryOnNotFound(t, 3, 500*time.Millisecond, func() error {
		var getErr error
		fetchedPanel, fetchPanelResp, getErr = previewSvc.GetTextPanel(ctx, ecID, panelID)
		return getErr
	})
	require.NoError(t, err)
	require.NotNil(t, fetchedPanel)
	assert.Equal(t, 200, fetchPanelResp.StatusCode)
	assert.Equal(t, "Acceptance Test Panel", fetchedPanel.DisplayName)

	// 4. UpdateTextPanel
	acc.LogTestStage(t, "UpdateTextPanel", "Updating text panel ID=%s", panelID)

	updatePanelReq := &enrollment_customization_preview.ResourceTextPanel{
		DisplayName:        "Updated Acceptance Test Panel",
		Rank:               1,
		Title:              "Updated Welcome",
		Body:               "# Updated Welcome\n\nThis panel has been updated.",
		BackButtonText:     "Back",
		ContinueButtonText: "Continue",
	}
	updatedPanel, updatePanelResp, err := previewSvc.UpdateTextPanel(ctx, ecID, panelID, updatePanelReq)
	require.NoError(t, err)
	require.NotNil(t, updatedPanel)
	assert.Equal(t, 200, updatePanelResp.StatusCode)
	assert.Equal(t, "Updated Acceptance Test Panel", updatedPanel.DisplayName)
	acc.LogTestSuccess(t, "Text panel updated: ID=%s", panelID)

	// 5. DeletePanel
	acc.LogTestStage(t, "DeletePanel", "Deleting text panel ID=%s", panelID)

	deleteResp, err := previewSvc.DeletePanel(ctx, ecID, panelID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Text panel ID=%s deleted", panelID)
}

// =============================================================================
// TestAcceptance_EnrollmentCustomizationPreview_validation_errors
// =============================================================================

func TestAcceptance_EnrollmentCustomizationPreview_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.EnrollmentCustomizationPreview

	t.Run("ParseMarkdown_NilRequest", func(t *testing.T) {
		_, _, err := svc.ParseMarkdown(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("GetAllPanels_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetAllPanels(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "enrollment customization ID is required")
	})

	t.Run("GetPanelByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetPanelByID(context.Background(), "", "1")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "enrollment customization ID is required")
	})

	t.Run("GetPanelByID_EmptyPanelID", func(t *testing.T) {
		_, _, err := svc.GetPanelByID(context.Background(), "1", "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "panel ID is required")
	})

	t.Run("CreateTextPanel_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateTextPanel(context.Background(), "1", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("CreateTextPanel_EmptyID", func(t *testing.T) {
		_, _, err := svc.CreateTextPanel(context.Background(), "", &enrollment_customization_preview.ResourceTextPanel{
			DisplayName: "x", Title: "x", Body: "x",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "enrollment customization ID is required")
	})
}
