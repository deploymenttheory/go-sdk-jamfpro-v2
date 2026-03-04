package cloud_ldap

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_ldap/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_CloudLdap_GetDefaultMappingsV2_Success(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	mock.RegisterGetDefaultMappingsMock("GOOGLE")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDefaultMappingsV2(ctx, "GOOGLE")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "uid", result.UserMappings.UserID)
	assert.Equal(t, "cn", result.GroupMappings.GroupName)
	assert.Equal(t, "member", result.MembershipMappings.GroupMembershipMapping)
}

func TestUnit_CloudLdap_GetDefaultServerConfigurationV2_Success(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	mock.RegisterGetDefaultServerConfigurationMock("GOOGLE")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDefaultServerConfigurationV2(ctx, "GOOGLE")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 636, result.Port)
	assert.Equal(t, "LDAPS", result.ConnectionType)
	assert.Equal(t, 15, result.ConnectionTimeout)
	assert.Equal(t, 60, result.SearchTimeout)
}

func TestUnit_CloudLdap_GetDefaultServerConfigurationV2_EmptyProviderName(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDefaultServerConfigurationV2(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "providerName is required")
}

func TestUnit_CloudLdap_GetDefaultMappingsV2_EmptyProviderName(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDefaultMappingsV2(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "providerName is required")
}

func TestUnit_CloudLdap_CreateV2_Success(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	mock.RegisterCreateMock()

	svc := NewService(mock)
	ctx := context.Background()

	request := &ResourceCloudLdap{
		CloudIdPCommon: &CloudIdPCommon{
			ProviderName: "GOOGLE",
			DisplayName:  "Test Google LDAP",
		},
		Server: &CloudLdapServer{
			Enabled:        true,
			ConnectionType: "LDAPS",
			ServerUrl:      "ldaps://ldap.google.com",
			Port:           636,
		},
	}

	result, resp, err := svc.CreateV2(ctx, request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, "/api/v2/cloud-ldaps/1", result.Href)
}

func TestUnit_CloudLdap_CreateV2_NilRequest(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.CreateV2(ctx, nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_CloudLdap_CreateV2_APICallFails(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	mock.RegisterCreateErrorMock()

	svc := NewService(mock)
	ctx := context.Background()

	request := &ResourceCloudLdap{
		CloudIdPCommon: &CloudIdPCommon{
			ProviderName: "GOOGLE",
			DisplayName:  "Test Google LDAP",
		},
	}

	result, resp, err := svc.CreateV2(ctx, request)

	assert.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "server error")
}

func TestUnit_CloudLdap_GetByIDV2_Success(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	mock.RegisterGetByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByIDV2(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.CloudIdPCommon.ID)
	assert.Equal(t, "Test Google LDAP", result.CloudIdPCommon.DisplayName)
	assert.Equal(t, "GOOGLE", result.CloudIdPCommon.ProviderName)
	assert.True(t, result.Server.Enabled)
	assert.Equal(t, "LDAPS", result.Server.ConnectionType)
}

func TestUnit_CloudLdap_GetByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByIDV2(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_CloudLdap_GetByIDV2_APICallFails(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	mock.RegisterGetByIDErrorMock("999")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByIDV2(ctx, "999")

	assert.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not found")
}

func TestUnit_CloudLdap_UpdateByIDV2_Success(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	mock.RegisterUpdateByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	request := &ResourceCloudLdap{
		CloudIdPCommon: &CloudIdPCommon{
			DisplayName:  "Updated Google LDAP",
			ProviderName: "GOOGLE",
		},
	}

	result, resp, err := svc.UpdateByIDV2(ctx, "1", request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotNil(t, result)
}

func TestUnit_CloudLdap_UpdateByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	svc := NewService(mock)
	ctx := context.Background()

	request := &ResourceCloudLdap{}

	result, resp, err := svc.UpdateByIDV2(ctx, "", request)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_CloudLdap_UpdateByIDV2_NilRequest(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.UpdateByIDV2(ctx, "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_CloudLdap_DeleteByIDV2_Success(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	mock.RegisterDeleteByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteByIDV2(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_CloudLdap_DeleteByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteByIDV2(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_CloudLdap_GetBindConnectionPoolStatsByIDV2_Success(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	mock.RegisterGetBindConnectionPoolStatsMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetBindConnectionPoolStatsByIDV2(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, int64(95), result.NumSuccessfulCheckouts)
	assert.Equal(t, int64(20), result.MaximumAvailableConnections)
	assert.Equal(t, int64(15), result.NumAvailableConnections)
}

func TestUnit_CloudLdap_GetBindConnectionPoolStatsByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetBindConnectionPoolStatsByIDV2(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_CloudLdap_GetSearchConnectionPoolStatsByIDV2_Success(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	mock.RegisterGetSearchConnectionPoolStatsMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetSearchConnectionPoolStatsByIDV2(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, int64(95), result.NumSuccessfulCheckouts)
}

func TestUnit_CloudLdap_GetSearchConnectionPoolStatsByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetSearchConnectionPoolStatsByIDV2(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_CloudLdap_TestConnectionByIDV2_Success(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	mock.RegisterTestConnectionMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.TestConnectionByIDV2(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "Successfully connected", result.Status)
}

func TestUnit_CloudLdap_TestConnectionByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.TestConnectionByIDV2(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_CloudLdap_GetMappingsByIDV2_Success(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	mock.RegisterGetMappingsByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetMappingsByIDV2(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "uid", result.UserMappings.UserID)
	assert.Equal(t, "cn", result.GroupMappings.GroupName)
}

func TestUnit_CloudLdap_GetMappingsByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetMappingsByIDV2(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_CloudLdap_UpdateMappingsByIDV2_Success(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	mock.RegisterUpdateMappingsByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	request := &CloudLdapMappings{
		UserMappings: CloudLdapUserMappings{
			UserID:   "uid",
			Username: "uid",
			RealName: "cn",
		},
		GroupMappings: CloudLdapGroupMappings{
			GroupID:   "cn",
			GroupName: "cn",
		},
		MembershipMappings: CloudLdapMembershipMappings{
			GroupMembershipMapping: "member",
		},
	}

	result, resp, err := svc.UpdateMappingsByIDV2(ctx, "1", request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotNil(t, result)
}

func TestUnit_CloudLdap_UpdateMappingsByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	svc := NewService(mock)
	ctx := context.Background()

	request := &CloudLdapMappings{}

	result, resp, err := svc.UpdateMappingsByIDV2(ctx, "", request)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_CloudLdap_UpdateMappingsByIDV2_NilRequest(t *testing.T) {
	mock := mocks.NewCloudLdapMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.UpdateMappingsByIDV2(ctx, "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}
