package ldap

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/ldap/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Ldap, *mocks.LdapMock) {
	t.Helper()
	mock := mocks.NewLdapMock()
	return NewLdap(mock), mock
}

func TestUnit_Ldap_GetLdapGroupsV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetLdapGroupsMock()

	result, resp, err := svc.GetLdapGroupsV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "Engineering", result.Results[0].Name)
	assert.Equal(t, 1, result.Results[0].LdapServerID)
}

func TestUnit_Ldap_GetLdapGroupsV1_WithQuery(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetLdapGroupsMock()

	params := map[string]string{"q": "eng"}
	result, resp, err := svc.GetLdapGroupsV1(context.Background(), params)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_Ldap_GetLdapGroupsV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetLdapGroupsErrorMock()

	result, resp, err := svc.GetLdapGroupsV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
}

func TestUnit_Ldap_GetLdapGroupsV1_NoMock(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered

	result, resp, err := svc.GetLdapGroupsV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_Ldap_GetLdapServersV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetLdapServersMock()

	result, resp, err := svc.GetLdapServersV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	require.Len(t, result, 2)
	assert.Equal(t, 1, result[0].ID)
	assert.Equal(t, "Corporate LDAP", result[0].Name)
	assert.Equal(t, 2, result[1].ID)
	assert.Equal(t, "Cloud IdP", result[1].Name)
}

func TestUnit_Ldap_GetLdapServersV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetLdapServersErrorMock()

	result, resp, err := svc.GetLdapServersV1(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_Ldap_GetLdapServersV1_NoMock(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered

	result, resp, err := svc.GetLdapServersV1(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_Ldap_GetLdapServersOnlyV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetLdapServersOnlyMock()

	result, resp, err := svc.GetLdapServersOnlyV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	require.Len(t, result, 2)
	assert.Equal(t, 1, result[0].ID)
	assert.Equal(t, "Corporate LDAP", result[0].Name)
	assert.Equal(t, 3, result[1].ID)
	assert.Equal(t, "Legacy LDAP", result[1].Name)
}

func TestUnit_Ldap_GetLdapServersOnlyV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetLdapServersOnlyErrorMock()

	result, resp, err := svc.GetLdapServersOnlyV1(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
}

func TestUnit_Ldap_GetLdapServersOnlyV1_NoMock(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered

	result, resp, err := svc.GetLdapServersOnlyV1(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_Ldap_NewService(t *testing.T) {
	mock := mocks.NewLdapMock()
	svc := NewLdap(mock)
	require.NotNil(t, svc)
	assert.NotNil(t, svc.client)
}
