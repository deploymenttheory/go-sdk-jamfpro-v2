package app_request

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/app_request/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh AppRequestMock.
func setupMockService(t *testing.T) (*AppRequest, *mocks.AppRequestMock) {
	t.Helper()
	mock := mocks.NewAppRequestMock()
	return NewAppRequest(mock), mock
}

// =============================================================================
// ListFormInputFieldsV1
// =============================================================================

func TestUnit_AppRequest_ListFormInputFields_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListFormInputFieldsMock()

	result, resp, err := svc.ListFormInputFieldsV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "Quantity", result.Results[0].Title)
	assert.Equal(t, "How many of these would you like?", *result.Results[0].Description)
	assert.Equal(t, 1, result.Results[0].Priority)
	assert.Equal(t, 2, result.Results[1].ID)
	assert.Equal(t, "Justification", result.Results[1].Title)
	assert.Equal(t, "Why do you need this app?", *result.Results[1].Description)
	assert.Equal(t, 2, result.Results[1].Priority)
}

func TestUnit_AppRequest_ListFormInputFields_WithQueryParams(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListFormInputFieldsMock()

	params := map[string]string{"page": "0", "page-size": "50", "sort": "priority:asc"}
	result, resp, err := svc.ListFormInputFieldsV1(context.Background(), params)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
}

// =============================================================================
// ReplaceFormInputFieldsV1
// =============================================================================

func TestUnit_AppRequest_ReplaceFormInputFields_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterReplaceFormInputFieldsMock()

	desc1 := "How many of these would you like?"
	desc2 := "Why do you need this app?"
	req := []RequestFormInputField{
		{Title: "Quantity", Description: &desc1, Priority: 1},
		{Title: "Justification", Description: &desc2, Priority: 2},
	}
	result, resp, err := svc.ReplaceFormInputFieldsV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	require.Len(t, result, 2)
	assert.Equal(t, 1, result[0].ID)
	assert.Equal(t, "Quantity", result[0].Title)
	assert.Equal(t, 2, result[1].ID)
	assert.Equal(t, "Justification", result[1].Title)
}

func TestUnit_AppRequest_ReplaceFormInputFields_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.ReplaceFormInputFieldsV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// CreateFormInputFieldV1
// =============================================================================

func TestUnit_AppRequest_CreateFormInputField_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateFormInputFieldMock()

	desc := "Which department?"
	req := &RequestFormInputField{Title: "Department", Description: &desc, Priority: 3}
	result, resp, err := svc.CreateFormInputFieldV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, 3, result.ID)
	assert.Equal(t, "Department", result.Title)
	assert.Equal(t, "Which department?", *result.Description)
	assert.Equal(t, 3, result.Priority)
}

func TestUnit_AppRequest_CreateFormInputField_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateFormInputFieldV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// GetFormInputFieldByIDV1
// =============================================================================

func TestUnit_AppRequest_GetFormInputFieldByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetFormInputFieldMock()

	result, resp, err := svc.GetFormInputFieldByIDV1(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Quantity", result.Title)
	assert.Equal(t, "How many of these would you like?", *result.Description)
	assert.Equal(t, 1, result.Priority)
}

func TestUnit_AppRequest_GetFormInputFieldByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetFormInputFieldByIDV1(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

// =============================================================================
// UpdateFormInputFieldByIDV1
// =============================================================================

func TestUnit_AppRequest_UpdateFormInputFieldByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateFormInputFieldMock()

	desc := "How many of these would you like?"
	req := &RequestFormInputField{Title: "Quantity", Description: &desc, Priority: 1}
	result, resp, err := svc.UpdateFormInputFieldByIDV1(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Quantity", result.Title)
}

func TestUnit_AppRequest_UpdateFormInputFieldByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateFormInputFieldByIDV1(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteFormInputFieldByIDV1
// =============================================================================

func TestUnit_AppRequest_DeleteFormInputFieldByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteFormInputFieldMock()

	resp, err := svc.DeleteFormInputFieldByIDV1(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

// =============================================================================
// GetSettingsV1
// =============================================================================

func TestUnit_AppRequest_GetSettings_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetSettingsMock()

	result, resp, err := svc.GetSettingsV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.True(t, result.IsEnabled)
	assert.Equal(t, "deviceLocale", result.AppStoreLocale)
	assert.Equal(t, 1, result.RequesterUserGroupID)
	require.Len(t, result.ApproverEmails, 2)
	assert.Equal(t, "jane.doe@company.com", result.ApproverEmails[0])
	assert.Equal(t, "john.doe@company.com", result.ApproverEmails[1])
}

// =============================================================================
// UpdateSettingsV1
// =============================================================================

func TestUnit_AppRequest_UpdateSettings_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateSettingsMock()

	req := &ResourceAppRequestSettings{
		IsEnabled:            true,
		AppStoreLocale:       "deviceLocale",
		RequesterUserGroupID:  1,
		ApproverEmails:       []string{"jane.doe@company.com", "john.doe@company.com"},
	}
	result, resp, err := svc.UpdateSettingsV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.True(t, result.IsEnabled)
	assert.Equal(t, "deviceLocale", result.AppStoreLocale)
	assert.Equal(t, 1, result.RequesterUserGroupID)
}

func TestUnit_AppRequest_UpdateSettings_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateSettingsV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// Error path tests (no mock registered → HTTP call fails)
// =============================================================================

func TestUnit_AppRequest_ListFormInputFields_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterErrorMocks()
	result, resp, err := svc.ListFormInputFieldsV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnit_AppRequest_ReplaceFormInputFields_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterErrorMocks()
	desc := "Some description"
	req := []RequestFormInputField{{Title: "Field", Description: &desc, Priority: 1}}
	result, resp, err := svc.ReplaceFormInputFieldsV1(context.Background(), req)
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnit_AppRequest_CreateFormInputField_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterErrorMocks()
	desc := "Some description"
	req := &RequestFormInputField{Title: "Field", Description: &desc, Priority: 1}
	result, resp, err := svc.CreateFormInputFieldV1(context.Background(), req)
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnit_AppRequest_GetFormInputFieldByID_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterErrorMocks()
	result, resp, err := svc.GetFormInputFieldByIDV1(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnit_AppRequest_UpdateFormInputFieldByID_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterErrorMocks()
	desc := "Some description"
	req := &RequestFormInputField{Title: "Field", Description: &desc, Priority: 1}
	result, resp, err := svc.UpdateFormInputFieldByIDV1(context.Background(), 1, req)
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnit_AppRequest_DeleteFormInputFieldByID_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterErrorMocks()
	resp, err := svc.DeleteFormInputFieldByIDV1(context.Background(), 1)
	require.Error(t, err)
	require.NotNil(t, resp)
}

func TestUnit_AppRequest_GetSettings_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterErrorMocks()
	result, resp, err := svc.GetSettingsV1(context.Background())
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnit_AppRequest_UpdateSettings_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterErrorMocks()
	req := &ResourceAppRequestSettings{
		IsEnabled:      true,
		AppStoreLocale: "deviceLocale",
	}
	result, resp, err := svc.UpdateSettingsV1(context.Background(), req)
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
}
