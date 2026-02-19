package scripts

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/scripts/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh ScriptsMock.
func setupMockService(t *testing.T) (*Service, *mocks.ScriptsMock) {
	t.Helper()
	mock := mocks.NewScriptsMock()
	return NewService(mock), mock
}

// =============================================================================
// ListScriptsV1
// =============================================================================

func TestUnitListScripts_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListScriptsMock()

	result, resp, err := svc.ListScriptsV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Install Homebrew", result.Results[0].Name)
	assert.Equal(t, "AFTER", result.Results[0].Priority)
	assert.Equal(t, "2", result.Results[1].ID)
	assert.Equal(t, "Set Hostname", result.Results[1].Name)
	assert.Equal(t, "BEFORE", result.Results[1].Priority)
}

func TestUnitListScripts_WithQueryParams(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListScriptsMock()

	params := map[string]string{"page": "0", "page-size": "50", "sort": "name:asc"}
	result, resp, err := svc.ListScriptsV1(context.Background(), params)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitListScripts_WithRSQLFilter(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListScriptsRSQLMock()

	rsqlQuery := map[string]string{"filter": `name=="Install Homebrew"`}
	result, resp, err := svc.ListScriptsV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount, "filtered result should contain exactly one script")
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Install Homebrew", result.Results[0].Name)

	// Verify the service forwarded the RSQL filter to the HTTP client unchanged.
	assert.Equal(t, rsqlQuery, mock.LastRSQLQuery, "rsqlQuery should be passed through to the HTTP client")
}

// =============================================================================
// GetScriptByIDV1
// =============================================================================

func TestUnitGetScriptByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetScriptMock()

	result, resp, err := svc.GetScriptByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Install Homebrew", result.Name)
	assert.Equal(t, "AFTER", result.Priority)
	assert.Equal(t, "Utils", result.CategoryName)
	assert.Equal(t, ">=10.15", result.OSRequirements)
	assert.NotEmpty(t, result.ScriptContents)
}

func TestUnitGetScriptByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetScriptByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "script ID is required")
}

func TestUnitGetScriptByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetScriptByIDV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

// =============================================================================
// CreateScriptV1
// =============================================================================

func TestUnitCreateScript_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateScriptMock()

	req := &RequestScript{
		Name:           "Test Script",
		Priority:       ScriptPriorityAfter,
		ScriptContents: "#!/bin/bash\necho hello",
	}
	result, resp, err := svc.CreateScriptV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "3", result.ID)
	assert.Contains(t, result.Href, "/api/v1/scripts/3")
}

func TestUnitCreateScript_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateScriptV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateScript_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestScript{Name: "Duplicate Script", Priority: ScriptPriorityAfter}
	result, resp, err := svc.CreateScriptV1(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

// =============================================================================
// UpdateScriptByIDV1
// =============================================================================

func TestUnitUpdateScriptByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateScriptMock()

	req := &RequestScript{
		Name:     "Install Homebrew Updated",
		Priority: ScriptPriorityBefore,
	}
	result, resp, err := svc.UpdateScriptByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Install Homebrew Updated", result.Name)
	assert.Equal(t, "BEFORE", result.Priority)
}

func TestUnitUpdateScriptByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateScriptByIDV1(context.Background(), "", &RequestScript{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "script ID is required")
}

func TestUnitUpdateScriptByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateScriptByIDV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteScriptByIDV1
// =============================================================================

func TestUnitDeleteScriptByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteScriptMock()

	resp, err := svc.DeleteScriptByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitDeleteScriptByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteScriptByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "script ID is required")
}

// =============================================================================
// GetScriptHistoryV1
// =============================================================================

func TestUnitGetScriptHistory_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetScriptHistoryMock()

	result, resp, err := svc.GetScriptHistoryV1(context.Background(), "1", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "Script created", result.Results[0].Note)
}

func TestUnitGetScriptHistory_WithRSQLFilter(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetScriptHistoryMock()

	rsqlQuery := map[string]string{"page": "0", "page-size": "10", "sort": "date:desc"}
	result, resp, err := svc.GetScriptHistoryV1(context.Background(), "1", rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	// Verify the rsqlQuery was forwarded to the HTTP client.
	assert.Equal(t, rsqlQuery, mock.LastRSQLQuery, "rsqlQuery should be passed through to the HTTP client")
}

func TestUnitGetScriptHistory_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetScriptHistoryV1(context.Background(), "", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "script ID is required")
}

// =============================================================================
// AddScriptHistoryNotesV1
// =============================================================================

func TestUnitAddScriptHistoryNotes_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddScriptHistoryNotesMock()

	req := &AddScriptHistoryNotesRequest{Note: "Test note added"}
	resp, err := svc.AddScriptHistoryNotesV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
}

func TestUnitAddScriptHistoryNotes_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &AddScriptHistoryNotesRequest{Note: "Test note"}
	resp, err := svc.AddScriptHistoryNotesV1(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "script ID is required")
}

func TestUnitAddScriptHistoryNotes_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.AddScriptHistoryNotesV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request body is required")
}

// =============================================================================
// DownloadScriptByIDV1
// =============================================================================

func TestUnitDownloadScriptByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDownloadScriptMock()

	data, resp, err := svc.DownloadScriptByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, data)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, data, "downloaded script contents should not be empty")
	assert.Contains(t, string(data), "#!/bin/bash")
}

func TestUnitDownloadScriptByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	data, resp, err := svc.DownloadScriptByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, data)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "script ID is required")
}
