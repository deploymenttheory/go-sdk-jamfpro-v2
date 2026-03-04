package volume_purchasing_subscriptions

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/volume_purchasing_subscriptions/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.VolumePurchasingSubscriptionsMock) {
	t.Helper()
	mock := mocks.NewVolumePurchasingSubscriptionsMock()
	return NewService(mock), mock
}

func TestUnit_VolumePurchasingSubscriptions_List_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "VPS One", result.Results[0].Name)
}

func TestUnit_VolumePurchasingSubscriptions_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "VPS One", result.Name)
	assert.True(t, result.Enabled)
}

func TestUnit_VolumePurchasingSubscriptions_GetByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_VolumePurchasingSubscriptions_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_VolumePurchasingSubscriptions_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	req := &RequestVolumePurchasingSubscription{
		Name:    "New VPS",
		Enabled: true,
	}
	result, resp, err := svc.CreateV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "2", result.ID)
	assert.NotEmpty(t, result.Href)
}

func TestUnit_VolumePurchasingSubscriptions_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_VolumePurchasingSubscriptions_UpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	req := &RequestVolumePurchasingSubscription{
		Name:    "VPS One Updated",
		Enabled: false,
	}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "VPS One Updated", result.Name)
}

func TestUnit_VolumePurchasingSubscriptions_UpdateByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestVolumePurchasingSubscription{Name: "Test", Enabled: true}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_VolumePurchasingSubscriptions_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_VolumePurchasingSubscriptions_UpdateByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateNotFoundErrorMock()

	req := &RequestVolumePurchasingSubscription{Name: "Test", Enabled: true}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "999", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_VolumePurchasingSubscriptions_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteMock()

	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_VolumePurchasingSubscriptions_DeleteByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_VolumePurchasingSubscriptions_DeleteByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteNotFoundErrorMock()

	resp, err := svc.DeleteByIDV1(context.Background(), "999")
	assert.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_VolumePurchasingSubscriptions_List_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListErrorMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
}
