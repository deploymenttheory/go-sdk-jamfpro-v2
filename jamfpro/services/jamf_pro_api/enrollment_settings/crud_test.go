package enrollment_settings

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/enrollment_settings/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.EnrollmentSettingsMock) {
	t.Helper()
	mock := mocks.NewEnrollmentSettingsMock()
	return NewService(mock), mock
}

func TestUnitGetV4_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.GetV4(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, result.MacOsEnterpriseEnrollmentEnabled)
	assert.True(t, result.IosEnterpriseEnrollmentEnabled)
}
