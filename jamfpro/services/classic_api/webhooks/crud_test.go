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

func TestUnitListWebhooks_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListWebhooksMock()

	result, resp, err := svc.ListWebhooks(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
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

func TestUnitGetWebhookByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetWebhookByIDMock()

	result, resp, err := svc.GetWebhookByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Computer Enrolled", result.Name)
	assert.True(t, result.Enabled)
	assert.Equal(t, "https://hooks.example.com/enrolled", result.URL)
	assert.Equal(t, "ComputerAdded", result.Event)
}

func TestUnitGetWebhookByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetWebhookByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "webhook ID must be a positive integer")
}

func TestUnitGetWebhookByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetWebhookByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "webhook ID must be a positive integer")
}

func TestUnitGetWebhookByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetWebhookByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

// =============================================================================
// GetWebhookByName
// =============================================================================

func TestUnitGetWebhookByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetWebhookByNameMock()

	result, resp, err := svc.GetWebhookByName(context.Background(), "Computer Enrolled")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Computer Enrolled", result.Name)
}

func TestUnitGetWebhookByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetWebhookByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "webhook name is required")
}

// =============================================================================
// CreateWebhook
// =============================================================================

func TestUnitCreateWebhook_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateWebhookMock()

	req := &RequestWebhook{
		Name:    "New Webhook",
		Enabled: true,
		URL:     "https://hooks.example.com/new",
		Event:   "ComputerAdded",
	}
	result, resp, err := svc.CreateWebhook(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 3, result.ID)
	assert.Equal(t, "New Webhook", result.Name)
}

func TestUnitCreateWebhook_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateWebhook(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateWebhook_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestWebhook{Name: "Computer Enrolled", Enabled: true}
	result, resp, err := svc.CreateWebhook(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

// =============================================================================
// UpdateWebhookByID
// =============================================================================

func TestUnitUpdateWebhookByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateWebhookByIDMock()

	req := &RequestWebhook{Name: "Computer Enrolled Updated", Enabled: true}
	result, resp, err := svc.UpdateWebhookByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Computer Enrolled Updated", result.Name)
}

func TestUnitUpdateWebhookByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateWebhookByID(context.Background(), 0, &RequestWebhook{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "webhook ID must be a positive integer")
}

func TestUnitUpdateWebhookByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateWebhookByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateWebhookByName
// =============================================================================

func TestUnitUpdateWebhookByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateWebhookByNameMock()

	req := &RequestWebhook{Name: "Computer Enrolled Updated", Enabled: true}
	result, resp, err := svc.UpdateWebhookByName(context.Background(), "Computer Enrolled", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Computer Enrolled Updated", result.Name)
}

func TestUnitUpdateWebhookByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateWebhookByName(context.Background(), "", &RequestWebhook{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "webhook name is required")
}

func TestUnitUpdateWebhookByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateWebhookByName(context.Background(), "Computer Enrolled", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteWebhookByID
// =============================================================================

func TestUnitDeleteWebhookByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteWebhookByIDMock()

	resp, err := svc.DeleteWebhookByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteWebhookByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteWebhookByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "webhook ID must be a positive integer")
}

// =============================================================================
// DeleteWebhookByName
// =============================================================================

func TestUnitDeleteWebhookByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteWebhookByNameMock()

	resp, err := svc.DeleteWebhookByName(context.Background(), "Computer Enrolled")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteWebhookByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteWebhookByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "webhook name is required")
}
