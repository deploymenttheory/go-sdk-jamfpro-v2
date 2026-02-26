package conditional_access

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/conditional_access/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_ConditionalAccess_GetDeviceComplianceFeatureToggleV1_Success(t *testing.T) {
	mock := mocks.NewConditionalAccessMock()
	mock.RegisterGetDeviceComplianceFeatureToggleMock()

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDeviceComplianceFeatureToggleV1(ctx)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, result.SharedDeviceFeatureEnabled)
}

func TestUnit_ConditionalAccess_GetDeviceComplianceFeatureToggleV1_Error(t *testing.T) {
	mock := mocks.NewConditionalAccessMock()
	// No mock response registered — HTTP call will fail.
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDeviceComplianceFeatureToggleV1(ctx)

	require.Error(t, err)
	assert.Nil(t, result)
	// resp may be non-nil with 404 when mock has no registered response
	_ = resp
}

func TestUnit_ConditionalAccess_GetDeviceComplianceInformationComputerV1_Success(t *testing.T) {
	mock := mocks.NewConditionalAccessMock()
	mock.RegisterGetDeviceComplianceInformationComputerMock()

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDeviceComplianceInformationComputerV1(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	require.Len(t, result, 1)
	assert.Equal(t, "1", result[0].DeviceId)
	assert.True(t, result[0].Applicable)
	assert.Equal(t, "COMPLIANT", result[0].ComplianceState)
	assert.Equal(t, "Vendor A", result[0].ComplianceVendor)
}

func TestUnit_ConditionalAccess_GetDeviceComplianceInformationMobileV1_Success(t *testing.T) {
	mock := mocks.NewConditionalAccessMock()
	mock.RegisterGetDeviceComplianceInformationMobileMock()

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDeviceComplianceInformationMobileV1(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	require.Len(t, result, 1)
	assert.Equal(t, "1", result[0].DeviceId)
	assert.True(t, result[0].Applicable)
	assert.Equal(t, "COMPLIANT", result[0].ComplianceState)
}
