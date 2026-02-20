package ldap

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/ldap/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.LdapMock) {
	t.Helper()
	mock := mocks.NewLdapMock()
	return NewService(mock), mock
}

func TestUnitGetLdapGroupsV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetLdapGroupsMock()

	result, resp, err := svc.GetLdapGroupsV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "Engineering", result.Results[0].Name)
	assert.Equal(t, 1, result.Results[0].LdapServerID)
}

func TestUnitGetLdapGroupsV1_WithQuery(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetLdapGroupsMock()

	params := map[string]string{"q": "eng"}
	result, resp, err := svc.GetLdapGroupsV1(context.Background(), params)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitGetLdapServersV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetLdapServersMock()

	result, resp, err := svc.GetLdapServersV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	require.Len(t, result, 2)
	assert.Equal(t, 1, result[0].ID)
	assert.Equal(t, "Corporate LDAP", result[0].Name)
	assert.Equal(t, 2, result[1].ID)
	assert.Equal(t, "Cloud IdP", result[1].Name)
}
