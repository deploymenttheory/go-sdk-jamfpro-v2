package onboarding

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/onboarding/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.OnboardingMock) {
	t.Helper()
	mock := mocks.NewOnboardingMock()
	return NewService(mock), mock
}

func TestUnitGetV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()
	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, result.Enabled)
	require.Len(t, result.OnboardingItems, 1)
	assert.Equal(t, "APP", result.OnboardingItems[0].SelfServiceEntityType)
}

func TestUnitUpdateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UpdateV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}
