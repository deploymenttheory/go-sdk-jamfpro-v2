package webhooks

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/webhooks/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh WebhooksMock.
func setupMockService(t *testing.T) (*Service, *mocks.WebhooksMock) {
	t.Helper()
	mock := mocks.NewWebhooksMock()
	return NewService(mock), mock
}

// =============================================================================
// ListWebhooks
// =============================================================================

func TestUnit_Webhooks_List_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.List(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.Size)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "Computer Enrolled", result.Results[0].Name)
	assert.Equal(t, 2, result.Results[1].ID)
	assert.Equal(t, "Policy Completed", result.Results[1].Name)
}

// =============================================================================
// GetWebhookByID
// =============================================================================

func TestUnit_Webhooks_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDMock()

	result, resp, err := svc.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Computer Enrolled", result.Name)
	assert.True(t, result.Enabled)
	assert.Equal(t, "https://hooks.example.com/enrolled", result.URL)
	assert.Equal(t, "ComputerAdded", result.Event)
}

func TestUnit_Webhooks_GetByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "webhook ID must be a positive integer")
}

func TestUnit_Webhooks_GetByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "webhook ID must be a positive integer")
}

func TestUnit_Webhooks_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

// =============================================================================
// GetWebhookByName
// =============================================================================

func TestUnit_Webhooks_GetByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByNameMock()

	result, resp, err := svc.GetByName(context.Background(), "Computer Enrolled")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Computer Enrolled", result.Name)
}

func TestUnit_Webhooks_GetByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "webhook name is required")
}

// =============================================================================
// CreateWebhook
// =============================================================================

func TestUnit_Webhooks_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	req := &RequestWebhook{
		Name:    "New Webhook",
		Enabled: true,
		URL:     "https://hooks.example.com/new",
		Event:   "ComputerAdded",
	}
	result, resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, 3, result.ID)
	assert.Equal(t, "New Webhook", result.Name)
}

func TestUnit_Webhooks_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.Create(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Webhooks_Create_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestWebhook{Name: "Computer Enrolled", Enabled: true}
	result, resp, err := svc.Create(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode())
}

// =============================================================================
// UpdateWebhookByID
// =============================================================================

func TestUnit_Webhooks_UpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateByIDMock()

	req := &RequestWebhook{Name: "Computer Enrolled Updated", Enabled: true}
	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Computer Enrolled Updated", result.Name)
}

func TestUnit_Webhooks_UpdateByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByID(context.Background(), 0, &RequestWebhook{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "webhook ID must be a positive integer")
}

func TestUnit_Webhooks_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateWebhookByName
// =============================================================================

func TestUnit_Webhooks_UpdateByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateByNameMock()

	req := &RequestWebhook{Name: "Computer Enrolled Updated", Enabled: true}
	result, resp, err := svc.UpdateByName(context.Background(), "Computer Enrolled", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Computer Enrolled Updated", result.Name)
}

func TestUnit_Webhooks_UpdateByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByName(context.Background(), "", &RequestWebhook{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "webhook name is required")
}

func TestUnit_Webhooks_UpdateByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByName(context.Background(), "Computer Enrolled", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteWebhookByID
// =============================================================================

func TestUnit_Webhooks_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteByIDMock()

	resp, err := svc.DeleteByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_Webhooks_DeleteByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "webhook ID must be a positive integer")
}

// =============================================================================
// DeleteWebhookByName
// =============================================================================

func TestUnit_Webhooks_DeleteByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteByNameMock()

	resp, err := svc.DeleteByName(context.Background(), "Computer Enrolled")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_Webhooks_DeleteByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "webhook name is required")
}
