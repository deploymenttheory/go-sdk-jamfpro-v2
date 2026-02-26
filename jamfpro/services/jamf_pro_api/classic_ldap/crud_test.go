package classic_ldap

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/classic_ldap/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.ClassicLdapMock) {
	t.Helper()
	mock := mocks.NewClassicLdapMock()
	return NewService(mock), mock
}

func TestUnit_ClassicLdap_GetMappingsByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMappingsByIDMock("1")

	result, resp, err := svc.GetMappingsByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "mail", result.UserObjectMapIdTo)
	assert.Equal(t, "uid", result.UserObjectMapUsernameTo)
	assert.Equal(t, "displayName", result.UserObjectMapRealNameTo)
	assert.Equal(t, "mail", result.UserObjectMapEmailTo)
	assert.Equal(t, "departmentNumber", result.UserObjectMapDepartmentTo)
	assert.Equal(t, "title", result.UserObjectMapPositionTo)
	assert.Equal(t, "name", result.UserGroupObjectMapIdTo)
	assert.Equal(t, "name", result.UserGroupObjectMapGroupNameTo)
}
