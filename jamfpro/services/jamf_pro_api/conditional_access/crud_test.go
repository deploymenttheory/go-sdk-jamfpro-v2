package conditional_access

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/conditional_access/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetDeviceComplianceFeatureToggleV1(t *testing.T) {
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
