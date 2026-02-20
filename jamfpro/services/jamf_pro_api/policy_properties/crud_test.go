package policy_properties

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/policy_properties/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.PolicyPropertiesMock) {
	t.Helper()
	mock := mocks.NewPolicyPropertiesMock()
	return NewService(mock), mock
}

func TestUnitGet_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()
	result, resp, err := svc.Get(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, result.AllowNetworkStateChangeTriggers)
}

func TestUnitUpdate_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()
	result, resp, err := svc.Update(context.Background(), &ResourcePolicyProperties{
		PoliciesRequireNetworkStateChange: false,
		AllowNetworkStateChangeTriggers:   true,
	})
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitUpdate_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.Update(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}
