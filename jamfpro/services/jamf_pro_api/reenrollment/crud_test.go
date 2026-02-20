package reenrollment

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/reenrollment/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.ReenrollmentMock) {
	t.Helper()
	mock := mocks.NewReenrollmentMock()
	return NewService(mock), mock
}

func TestUnitGet_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.Get(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "DELETE_EVERYTHING", result.FlushMdmQueue)
	assert.False(t, result.FlushPolicyHistory)
}

func TestUnitUpdate_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	request := &ResourceReenrollmentSettings{
		FlushPolicyHistory: false, FlushLocationInformation: false,
		FlushLocationInformationHistory: false, FlushExtensionAttributes: false,
		FlushSoftwareUpdatePlans: false, FlushMdmQueue: "DELETE_EVERYTHING",
	}
	result, resp, err := svc.Update(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitUpdate_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.Update(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitGetHistory_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryMock()

	result, resp, err := svc.GetHistory(context.Background(), map[string]string{"page": "0", "page-size": "100", "sort": "date:desc"})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "Test note", result.Results[0].Note)
}

func TestUnitAddHistoryNotes_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddHistoryNotesMock()

	request := &AddReenrollmentHistoryNotesRequest{Note: "Acceptance test note"}
	result, resp, err := svc.AddHistoryNotes(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "Acceptance test note", result.Note)
}

func TestUnitAddHistoryNotes_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.AddHistoryNotes(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitExportHistory_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportHistoryMock()

	resp, body, err := svc.ExportHistory(context.Background(), map[string]string{"page": "0", "page-size": "100"}, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, body)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Greater(t, len(body), 0)
}
