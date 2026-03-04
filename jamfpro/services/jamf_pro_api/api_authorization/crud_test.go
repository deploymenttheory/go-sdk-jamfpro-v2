package api_authorization

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_authorization/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.ApiAuthorizationMock) {
	t.Helper()
	mock := mocks.NewApiAuthorizationMock()
	return NewService(mock), mock
}

func TestUnit_ApiAuthorization_GetV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetV1Mock()

	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.Account.ID)
	assert.Equal(t, "admin", result.Account.Username)
	assert.Equal(t, "IT Bob", result.Account.RealName)
	assert.Equal(t, "ITBob@Jamf.com", result.Account.Email)
	assert.Equal(t, "FullAccess", result.Account.AccessLevel)
	assert.Equal(t, "CUSTOM", result.Account.PrivilegeSet)
	assert.True(t, result.Account.MultiSiteAdmin)
	assert.Equal(t, "JSS", result.AuthenticationType)
	assert.Len(t, result.AccountGroups, 1)
	assert.Len(t, result.Sites, 1)
	assert.Equal(t, "1", result.Sites[0].ID)
	assert.Equal(t, "Eau Claire", result.Sites[0].Name)
}

func TestUnit_ApiAuthorization_GetV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetV1(context.Background())
	require.Error(t, err)
}
