package activation_code

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/activation_code/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh ActivationCodeMock.
func setupMockService(t *testing.T) (*ActivationCode, *mocks.ActivationCodeMock) {
	t.Helper()
	mock := mocks.NewActivationCodeMock()
	return NewActivationCode(mock), mock
}

// =============================================================================
// GetActivationCode
// =============================================================================

func TestUnit_ActivationCode_Get_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetActivationCodeMock()

	result, resp, err := svc.GetActivationCode(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "Example Organization", result.OrganizationName)
	assert.Equal(t, "ABCD-1234-EFGH-5678", result.Code)
}

// =============================================================================
// UpdateActivationCode
// =============================================================================

func TestUnit_ActivationCode_Update_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateActivationCodeMock()

	req := &RequestActivationCode{
		OrganizationName: "Updated Organization",
		Code:             "WXYZ-9876-QRST-5432",
	}
	resp, err := svc.UpdateActivationCode(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_ActivationCode_Update_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.UpdateActivationCode(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_ActivationCode_Get_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetActivationCode(context.Background())
	require.Error(t, err)
}

func TestUnit_ActivationCode_Update_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestActivationCode{OrganizationName: "Test", Code: "ABCD-1234"}
	_, err := svc.UpdateActivationCode(context.Background(), req)
	require.Error(t, err)
}
