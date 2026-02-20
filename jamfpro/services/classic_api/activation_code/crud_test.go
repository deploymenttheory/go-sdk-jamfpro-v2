package activation_code

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/activation_code/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh ActivationCodeMock.
func setupMockService(t *testing.T) (*Service, *mocks.ActivationCodeMock) {
	t.Helper()
	mock := mocks.NewActivationCodeMock()
	return NewService(mock), mock
}

// =============================================================================
// GetActivationCode
// =============================================================================

func TestUnitGetActivationCode_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetActivationCodeMock()

	result, resp, err := svc.GetActivationCode(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "Example Organization", result.OrganizationName)
	assert.Equal(t, "ABCD-1234-EFGH-5678", result.Code)
}

// =============================================================================
// UpdateActivationCode
// =============================================================================

func TestUnitUpdateActivationCode_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateActivationCodeMock()

	req := &RequestActivationCode{
		OrganizationName: "Updated Organization",
		Code:             "WXYZ-9876-QRST-5432",
	}
	resp, err := svc.UpdateActivationCode(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitUpdateActivationCode_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.UpdateActivationCode(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}
