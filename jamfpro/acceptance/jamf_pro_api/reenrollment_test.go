package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/reenrollment"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Acceptance tests run in doc order: Get → Update → GetHistory → AddHistoryNotes → ExportHistory.

func TestAcceptance_Reenrollment_get(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Reenrollment
	ctx := context.Background()

	result, resp, err := svc.Get(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestAcceptance_Reenrollment_update(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Reenrollment
	ctx := context.Background()

	current, _, err := svc.Get(ctx)
	require.NoError(t, err)
	require.NotNil(t, current)

	request := *current
	request.FlushPolicyHistory = !request.FlushPolicyHistory
	updated, resp, err := svc.Update(ctx, &request)
	require.NoError(t, err)
	require.NotNil(t, updated)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	request.FlushPolicyHistory = current.FlushPolicyHistory
	_, _, _ = svc.Update(ctx, &request)
}

func TestAcceptance_Reenrollment_get_history(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Reenrollment
	ctx := context.Background()

	// Add history note first
	noteReq := &shared.SharedHistoryNoteRequest{
		Note: "Acceptance test history note for reenrollment",
	}
	addResult, addResp, err := svc.AddHistoryNotes(ctx, noteReq)
	require.NoError(t, err)
	require.NotNil(t, addResult)
	assert.Contains(t, []int{200, 201}, addResp.StatusCode)
	t.Logf("Added history note: %s", addResult.Note)

	result, resp, err := svc.GetHistory(ctx, map[string]string{"page": "0", "page-size": "100", "sort": "date:desc"})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 1, "Should have at least the note we just added")
	assert.NotNil(t, result.Results)
}

func TestAcceptance_Reenrollment_add_history_notes(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Reenrollment
	ctx := context.Background()

	request := &reenrollment.AddReenrollmentHistoryNotesRequest{Note: "Acceptance test note from go-sdk-jamfpro-v2"}
	result, resp, err := svc.AddHistoryNotes(ctx, request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, []int{200, 201}, resp.StatusCode, "API may return 200 or 201")
	assert.Equal(t, request.Note, result.Note)
}

func TestAcceptance_Reenrollment_export_history(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Reenrollment
	ctx := context.Background()

	resp, body, err := svc.ExportHistory(ctx, map[string]string{"page": "0", "page-size": "100"}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, body)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Greater(t, len(body), 0)
}
