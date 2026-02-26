package oauth2_session_tokens

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/oauth2_session_tokens/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.OAuth2SessionTokensMock) {
	t.Helper()
	mock := mocks.NewOAuth2SessionTokensMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnit_OAuth2SessionTokens_GetV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "token123", result.AccessToken)
	require.Equal(t, "id-token-456", result.IDToken)
}

func TestUnit_OAuth2SessionTokens_GetV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetErrorMock()
	result, resp, err := svc.GetV1(context.Background())
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 401, resp.StatusCode)
}
