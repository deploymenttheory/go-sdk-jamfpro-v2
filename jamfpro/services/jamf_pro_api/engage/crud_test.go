package engage

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/engage/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUnit_Engage_GetV2_Success(t *testing.T) {
	mock := mocks.NewEngageMock()
	mock.RegisterGetMock()

	svc := NewService(mock)
	result, resp, err := svc.GetV2(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.True(t, result.IsEnabled)
}

func TestUnit_Engage_GetV2_ClientError(t *testing.T) {
	mock := mocks.NewEngageMock()
	// No mock registered - client returns error

	svc := NewService(mock)
	result, resp, err := svc.GetV2(context.Background())

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_Engage_GetV2_InvalidJSON(t *testing.T) {
	mock := mocks.NewEngageMock()
	mock.RegisterGetInvalidJSONMock()

	svc := NewService(mock)
	result, resp, err := svc.GetV2(context.Background())

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

func TestUnit_Engage_UpdateV2_Success(t *testing.T) {
	mock := mocks.NewEngageMock()
	mock.RegisterUpdateMock()

	svc := NewService(mock)
	req := &ResourceEngageSettings{IsEnabled: false}
	result, resp, err := svc.UpdateV2(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.False(t, result.IsEnabled)
}

func TestUnit_Engage_UpdateV2_NilRequest(t *testing.T) {
	mock := mocks.NewEngageMock()
	svc := NewService(mock)

	result, resp, err := svc.UpdateV2(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "settings cannot be nil")
}

func TestUnit_Engage_UpdateV2_ClientError(t *testing.T) {
	mock := mocks.NewEngageMock()
	// No mock registered - client returns error

	svc := NewService(mock)
	req := &ResourceEngageSettings{IsEnabled: false}
	result, resp, err := svc.UpdateV2(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_Engage_GetHistoryV2_Success(t *testing.T) {
	mock := mocks.NewEngageMock()
	mock.RegisterGetHistoryMock()

	svc := NewService(mock)
	result, resp, err := svc.GetHistoryV2(context.Background(), nil)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, 2, result.TotalCount)
	assert.Len(t, result.Results, 2)
	assert.Equal(t, "admin", result.Results[0].Username)
}

func TestUnit_Engage_GetHistoryV2_WithFilter(t *testing.T) {
	mock := mocks.NewEngageMock()
	mock.RegisterGetHistoryMock()

	svc := NewService(mock)
	rsqlQuery := map[string]string{
		"filter": "username==admin",
		"sort":   "date:desc",
	}
	result, resp, err := svc.GetHistoryV2(context.Background(), rsqlQuery)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
}

func TestUnit_Engage_GetHistoryV2_ClientError(t *testing.T) {
	mock := mocks.NewEngageMock()
	// No mock registered - client returns error

	svc := NewService(mock)
	result, resp, err := svc.GetHistoryV2(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_Engage_AddHistoryNotesV2_Success(t *testing.T) {
	mock := mocks.NewEngageMock()
	mock.RegisterAddHistoryNotesMock()

	svc := NewService(mock)
	req := &RequestAddHistoryNotes{Note: "Test note"}
	result, resp, err := svc.AddHistoryNotesV2(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "Test note added", result.Note)
}

func TestUnit_Engage_AddHistoryNotesV2_NilRequest(t *testing.T) {
	mock := mocks.NewEngageMock()
	svc := NewService(mock)

	result, resp, err := svc.AddHistoryNotesV2(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request body is required")
}

func TestUnit_Engage_AddHistoryNotesV2_EmptyNote(t *testing.T) {
	mock := mocks.NewEngageMock()
	svc := NewService(mock)

	req := &RequestAddHistoryNotes{Note: ""}
	result, resp, err := svc.AddHistoryNotesV2(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "note is required")
}

func TestUnit_Engage_AddHistoryNotesV2_ClientError(t *testing.T) {
	mock := mocks.NewEngageMock()
	// No mock registered - client returns error

	svc := NewService(mock)
	req := &RequestAddHistoryNotes{Note: "Test note"}
	result, resp, err := svc.AddHistoryNotesV2(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}
