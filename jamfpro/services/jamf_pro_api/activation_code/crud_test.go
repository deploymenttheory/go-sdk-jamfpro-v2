package activation_code

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/activation_code/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.ActivationCodeMock) {
	t.Helper()
	mock := mocks.NewActivationCodeMock()
	return NewService(mock), mock
}

// Test GetHistoryV1 with success response
func TestUnit_ActivationCode_GetHistoryV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryMock()

	result, resp, err := svc.GetHistoryV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 3, result.TotalCount)
	assert.Len(t, result.Results, 3)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "2024-01-15T10:30:00Z", result.Results[0].Date)
	assert.Equal(t, "Activation code generated", result.Results[0].Note)
	require.NotNil(t, result.Results[0].Details)
	assert.Equal(t, "enabled", *result.Results[0].Details)
}

// Test GetHistoryV1 with RSQL filter query
func TestUnit_ActivationCode_GetHistoryV1_WithRSQLFilter(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryMock()

	rsqlQuery := map[string]string{
		"filter": `username==admin`,
		"sort":   "date:desc",
	}

	result, resp, err := svc.GetHistoryV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())

	// Verify RSQL query was passed to the client
	assert.NotNil(t, mock.LastRSQLQuery)
	assert.Equal(t, `username==admin`, mock.LastRSQLQuery["filter"])
	assert.Equal(t, "date:desc", mock.LastRSQLQuery["sort"])
}

// Test GetHistoryV1 with pagination parameters
func TestUnit_ActivationCode_GetHistoryV1_WithPagination(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryMock()

	rsqlQuery := map[string]string{
		"page":      "0",
		"page-size": "50",
		"sort":      "date:asc",
	}

	result, resp, err := svc.GetHistoryV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())

	// Verify pagination parameters were passed
	assert.Equal(t, "0", mock.LastRSQLQuery["page"])
	assert.Equal(t, "50", mock.LastRSQLQuery["page-size"])
	assert.Equal(t, "date:asc", mock.LastRSQLQuery["sort"])
}

// Test GetHistoryV1 with complex RSQL filter
func TestUnit_ActivationCode_GetHistoryV1_WithComplexRSQLFilter(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryMock()

	rsqlQuery := map[string]string{
		"filter": `username!=admin and details==enabled and date>2024-01-01`,
	}

	result, resp, err := svc.GetHistoryV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())

	// Verify complex RSQL query was passed
	assert.Equal(t, `username!=admin and details==enabled and date>2024-01-01`, mock.LastRSQLQuery["filter"])
}

// Test GetHistoryV1 with sorting
func TestUnit_ActivationCode_GetHistoryV1_WithSorting(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryMock()

	rsqlQuery := map[string]string{
		"sort": "username:asc,date:desc",
	}

	result, resp, err := svc.GetHistoryV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())

	// Verify sort parameter was passed
	assert.Equal(t, "username:asc,date:desc", mock.LastRSQLQuery["sort"])
}

// Test UpdateV1 with success response
func TestUnit_ActivationCode_UpdateV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateActivationCodeMock()

	req := &ActivationCodeRequest{
		ActivationCode: "12345678-1234-1234-1234-123456789012",
	}

	resp, err := svc.UpdateV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 202, resp.StatusCode())
}

// Test UpdateV1 with nil request
func TestUnit_ActivationCode_UpdateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.UpdateV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// Test UpdateOrganizationNameV1 with success response
func TestUnit_ActivationCode_UpdateOrganizationNameV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateOrganizationNameMock()

	req := &OrganizationNameRequest{
		OrganizationName: "Example Organization",
	}

	resp, err := svc.UpdateOrganizationNameV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 202, resp.StatusCode())
}

// Test UpdateOrganizationNameV1 with nil request
func TestUnit_ActivationCode_UpdateOrganizationNameV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.UpdateOrganizationNameV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// Test AddHistoryNoteV1 with success response
func TestUnit_ActivationCode_AddHistoryNoteV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddHistoryNoteMock()

	req := &HistoryNoteRequest{
		Note: "Test history note",
	}

	result, resp, err := svc.AddHistoryNoteV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "Test history note", result.Note)
}

// Test AddHistoryNoteV1 with nil request
func TestUnit_ActivationCode_AddHistoryNoteV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.AddHistoryNoteV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_ActivationCode_ExportHistoryV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportHistoryMock()

	queryParams := map[string]string{
		"page":      "0",
		"page-size": "100",
		"sort":      "date:desc",
	}

	req := &HistoryExportRequest{
		Sort:   []string{"date:desc"},
		Filter: stringPtr("username==admin"),
	}

	result, resp, err := svc.ExportHistoryV1(context.Background(), queryParams, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func stringPtr(s string) *string {
	return &s
}

func TestUnit_ActivationCode_GetHistoryV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetHistoryV1(context.Background(), nil)
	require.Error(t, err)
}

func TestUnit_ActivationCode_UpdateV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.UpdateV1(context.Background(), &ActivationCodeRequest{ActivationCode: "ABCD-EFGH-IJKL-MNOP-QRST"})
	require.Error(t, err)
}

func TestUnit_ActivationCode_UpdateOrganizationNameV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.UpdateOrganizationNameV1(context.Background(), &OrganizationNameRequest{OrganizationName: "Test"})
	require.Error(t, err)
}

func TestUnit_ActivationCode_AddHistoryNoteV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.AddHistoryNoteV1(context.Background(), &HistoryNoteRequest{Note: "test note"})
	require.Error(t, err)
}

func TestUnit_ActivationCode_ExportHistoryV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.ExportHistoryV1(context.Background(), nil, &HistoryExportRequest{})
	require.Error(t, err)
}
