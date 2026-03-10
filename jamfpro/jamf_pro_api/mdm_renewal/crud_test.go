package mdm_renewal

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/mdm_renewal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*MdmRenewal, *mocks.MDMRenewalMock) {
	t.Helper()
	mock := mocks.NewMDMRenewalMock()
	return NewMdmRenewal(mock), mock
}

func TestUnit_MDMRenewal_NewService(t *testing.T) {
	mock := mocks.NewMDMRenewalMock()
	svc := NewMdmRenewal(mock)
	require.NotNil(t, svc)
}

func TestUnit_MDMRenewal_UpdateDeviceCommonDetailsV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateDeviceCommonDetailsMock()

	req := &RequestDeviceCommonDetailsUpdate{
		ClientManagementID: "abc-123-client-mgmt-id",
	}

	resp, err := svc.UpdateDeviceCommonDetailsV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_MDMRenewal_UpdateDeviceCommonDetailsV1_WithOptionalFields(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateDeviceCommonDetailsMock()

	renewDate := "2024-02-01T00:00:00Z"
	checkinURL := "https://checkin.example.com"
	serverURL := "https://mdm.example.com"
	needsRenewal := true

	req := &RequestDeviceCommonDetailsUpdate{
		ClientManagementID:                                   "abc-123",
		RenewMdmProfileStartDate:                             &renewDate,
		MdmProfileNeedsRenewalDueToCaRenewed:                  &needsRenewal,
		MdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring: &needsRenewal,
		MdmCheckinUrl:                                        &checkinURL,
		MdmServerUrl:                                         &serverURL,
	}

	resp, err := svc.UpdateDeviceCommonDetailsV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_MDMRenewal_UpdateDeviceCommonDetailsV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.UpdateDeviceCommonDetailsV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MDMRenewal_UpdateDeviceCommonDetailsV1_EmptyClientManagementID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestDeviceCommonDetailsUpdate{
		ClientManagementID: "",
	}

	resp, err := svc.UpdateDeviceCommonDetailsV1(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "clientManagementId is required")
}

func TestUnit_MDMRenewal_UpdateDeviceCommonDetailsV1_NoMockRegistered(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateDeviceCommonDetailsErrorMock()

	req := &RequestDeviceCommonDetailsUpdate{
		ClientManagementID: "abc-123",
	}

	resp, err := svc.UpdateDeviceCommonDetailsV1(context.Background(), req)
	assert.Error(t, err)
	assert.NotNil(t, resp)
}

func TestUnit_MDMRenewal_GetDeviceCommonDetailsV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDeviceCommonDetailsMock("abc-123-client-mgmt-id")

	result, resp, err := svc.GetDeviceCommonDetailsV1(context.Background(), "abc-123-client-mgmt-id")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "abc-123-client-mgmt-id", result.ClientManagementID)
	assert.NotNil(t, result.RenewMdmProfileStartDate)
	assert.Equal(t, "2024-01-15T00:00:00Z", *result.RenewMdmProfileStartDate)
	assert.False(t, result.MdmProfileNeedsRenewalDueToCaRenewed)
	assert.True(t, result.MdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring)
	assert.NotNil(t, result.MdmCheckinUrl)
	assert.Equal(t, "https://checkin.example.com", *result.MdmCheckinUrl)
	assert.NotNil(t, result.MdmServerUrl)
	assert.Equal(t, "https://mdm.example.com", *result.MdmServerUrl)
}

func TestUnit_MDMRenewal_GetDeviceCommonDetailsV1_EmptyClientManagementID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetDeviceCommonDetailsV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "clientManagementId is required")
}

func TestUnit_MDMRenewal_GetDeviceCommonDetailsV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock("999")

	result, resp, err := svc.GetDeviceCommonDetailsV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_MDMRenewal_GetDeviceCommonDetailsV1_NoMockRegistered(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDeviceCommonDetailsErrorMock("abc-123")

	result, resp, err := svc.GetDeviceCommonDetailsV1(context.Background(), "abc-123")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

func TestUnit_MDMRenewal_GetRenewalStrategiesV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetRenewalStrategiesMock("abc-123-client-mgmt-id")

	result, resp, err := svc.GetRenewalStrategiesV1(context.Background(), "abc-123-client-mgmt-id")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	require.Len(t, result, 1)
	assert.Equal(t, "err-1", result[0].Error.MdmRenewalErrorId)
	assert.Equal(t, "abc-123-client-mgmt-id", result[0].Error.ClientManagementId)
	assert.Equal(t, MDMRenewalErrorTypeServerError, result[0].Error.MdmRenewalErrorType)
	assert.Equal(t, 2, result[0].Error.FailureCount)
	require.Len(t, result[0].Strategies, 1)
	assert.Equal(t, "strat-1", result[0].Strategies[0].ID)
	assert.Equal(t, MDMRenewalStrategyTypeReturnNoCheckInInvitation, result[0].Strategies[0].MdmRenewalStrategyType)
	assert.Equal(t, "https://renewal.example.com", result[0].Strategies[0].MdmRenewalCheckInUrl)
}

func TestUnit_MDMRenewal_GetRenewalStrategiesV1_EmptyClientManagementID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetRenewalStrategiesV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "clientManagementId is required")
}

func TestUnit_MDMRenewal_GetRenewalStrategiesV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock("999")

	result, resp, err := svc.GetRenewalStrategiesV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_MDMRenewal_GetRenewalStrategiesV1_NoMockRegistered(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetRenewalStrategiesErrorMock("abc-123")

	result, resp, err := svc.GetRenewalStrategiesV1(context.Background(), "abc-123")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

func TestUnit_MDMRenewal_DeleteRenewalStrategiesV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteRenewalStrategiesMock("abc-123-client-mgmt-id")

	resp, err := svc.DeleteRenewalStrategiesV1(context.Background(), "abc-123-client-mgmt-id")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_MDMRenewal_DeleteRenewalStrategiesV1_EmptyClientManagementID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteRenewalStrategiesV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "clientManagementId is required")
}

func TestUnit_MDMRenewal_DeleteRenewalStrategiesV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock("999")

	resp, err := svc.DeleteRenewalStrategiesV1(context.Background(), "999")
	assert.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_MDMRenewal_DeleteRenewalStrategiesV1_NoMockRegistered(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteRenewalStrategiesErrorMock("abc-123")

	resp, err := svc.DeleteRenewalStrategiesV1(context.Background(), "abc-123")
	assert.Error(t, err)
	assert.NotNil(t, resp)
}
