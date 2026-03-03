package jamf_pro_api

import (
	"context"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/log_flushing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Log Flushing
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetSettingsV1(ctx) - Gets log flushing settings and retention policies
//   • ListTasksV1(ctx) - Lists all log flushing tasks
//   • GetTaskByIDV1(ctx, id) - Gets a specific log flushing task by ID
//   • QueueTaskV1(ctx, request) - Creates/queues a new log flushing task
//   • DeleteTaskByIDV1(ctx, id) - Deletes a log flushing task by ID
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 2: Settings/Configuration (read-only)
//     -- Tests: TestAcceptance_LogFlushing_settings_v1
//
//   ✓ Pattern 1: Task Lifecycle (Queue → Get → ListTasks → Delete)
//     -- Requires at least one retention policy in settings
//     -- Tests: TestAcceptance_LogFlushing_task_lifecycle
//     -- Uses a very large retention period to avoid flushing real logs
//
//   ✓ Pattern 7: Validation Errors
//     -- Tests: TestAcceptance_LogFlushing_validation_errors
//
// Notes
// -----------------------------------------------------------------------------
//   • A large RetentionPeriod (9999) is used in task lifecycle to prevent
//     actual log deletion in test environments
//   • The qualifier is sourced from GetSettingsV1 retention policies
//
// =============================================================================

// TestAcceptance_LogFlushing_settings_v1 verifies the log flushing settings endpoint.
func TestAcceptance_LogFlushing_settings_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.LogFlushing
	ctx := context.Background()

	acc.LogTestStage(t, "GetSettingsV1", "Getting log flushing settings")

	settings, resp, err := svc.GetSettingsV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, settings)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, settings.HourOfDay, 0)
	assert.LessOrEqual(t, settings.HourOfDay, 23)

	acc.LogTestSuccess(t, "GetSettingsV1: hourOfDay=%d retentionPolicies=%d",
		settings.HourOfDay, len(settings.RetentionPolicies))
}

// TestAcceptance_LogFlushing_task_lifecycle exercises the task queue/get/list/delete lifecycle.
func TestAcceptance_LogFlushing_task_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.LogFlushing
	ctx := context.Background()

	// Get settings to find a valid qualifier
	settings, _, err := svc.GetSettingsV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, settings)

	if len(settings.RetentionPolicies) == 0 {
		t.Skip("No retention policies configured; skipping log flushing task lifecycle test")
	}

	firstPolicy := settings.RetentionPolicies[0]
	acc.LogTestSuccess(t, "Using retention policy: qualifier=%s unit=%s",
		firstPolicy.Qualifier, firstPolicy.RetentionPeriodUnit)

	// 1. QueueTask — use a very large period to avoid deleting real logs
	acc.LogTestStage(t, "QueueTaskV1", "Queuing log flushing task qualifier=%s", firstPolicy.Qualifier)

	taskReq := &log_flushing.RequestLogFlushingTask{
		Qualifier:           firstPolicy.Qualifier,
		RetentionPeriod:     9999,
		RetentionPeriodUnit: firstPolicy.RetentionPeriodUnit,
	}

	created, createResp, err := svc.QueueTaskV1(ctx, taskReq)
	require.NoError(t, err, "QueueTaskV1 should not return an error")
	require.NotNil(t, created)
	require.Contains(t, []int{200, 201, 202}, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	taskID := created.ID
	acc.LogTestSuccess(t, "Log flushing task queued ID=%s", taskID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteTaskByIDV1(cleanupCtx, taskID)
		acc.LogCleanupDeleteError(t, "log flushing task", taskID, delErr)
	})

	// 2. GetTaskByIDV1 (with retry for eventual consistency)
	acc.LogTestStage(t, "GetTaskByIDV1", "Getting log flushing task ID=%s", taskID)

	var task *log_flushing.ResourceLogFlushingTask
	var getResp *interfaces.Response
	err = acc.RetryOnNotFound(t, 3, 500*time.Millisecond, func() error {
		var getErr error
		task, getResp, getErr = svc.GetTaskByIDV1(ctx, taskID)
		return getErr
	})
	require.NoError(t, err)
	require.NotNil(t, task)
	assert.Equal(t, 200, getResp.StatusCode)
	assert.Equal(t, taskID, task.ID)
	assert.Equal(t, firstPolicy.Qualifier, task.Qualifier)
	acc.LogTestSuccess(t, "GetTaskByIDV1: ID=%s state=%s", task.ID, task.State)

	// 3. ListTasksV1
	acc.LogTestStage(t, "ListTasksV1", "Listing all log flushing tasks")

	tasks, listResp, err := svc.ListTasksV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, listResp)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, tk := range tasks {
		if tk.ID == taskID {
			found = true
			break
		}
	}
	assert.True(t, found, "queued task should appear in ListTasksV1")
	acc.LogTestSuccess(t, "ListTasksV1: %d task(s) found=%v", len(tasks), found)

	// 4. DeleteTaskByIDV1
	acc.LogTestStage(t, "DeleteTaskByIDV1", "Deleting log flushing task ID=%s", taskID)

	deleteResp, err := svc.DeleteTaskByIDV1(ctx, taskID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	require.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "DeleteTaskByIDV1: ID=%s deleted", taskID)
}

// TestAcceptance_LogFlushing_validation_errors verifies input validation.
func TestAcceptance_LogFlushing_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.LogFlushing

	t.Run("QueueTaskV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.QueueTaskV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "log flushing task request cannot be nil")
	})

	t.Run("GetTaskByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetTaskByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "task ID is required")
	})

	t.Run("DeleteTaskByIDV1_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteTaskByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "task ID is required")
	})
}
