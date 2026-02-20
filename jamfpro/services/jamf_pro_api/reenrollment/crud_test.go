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
