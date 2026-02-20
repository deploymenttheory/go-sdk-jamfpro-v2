package cloud_azure

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_azure/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetDefaultServerConfigurationV1(t *testing.T) {
	mock := mocks.NewCloudAzureMock()
	mock.RegisterGetDefaultServerConfigurationMock()

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDefaultServerConfigurationV1(ctx)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, result)
	assert.Equal(t, "objectGUID", result.Mappings.UserId)
	assert.Equal(t, "userPrincipalName", result.Mappings.UserName)
}

func TestGetByIDV1(t *testing.T) {
	mock := mocks.NewCloudAzureMock()
	mock.RegisterGetByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByIDV1(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.Server.ID)
	assert.Equal(t, "Test Azure IDP", result.CloudIdPCommon.DisplayName)
	assert.Equal(t, "AZURE", result.CloudIdPCommon.ProviderName)
	assert.True(t, result.Server.Enabled)
}

func TestGetByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewCloudAzureMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByIDV1(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestGetByNameV1(t *testing.T) {
	mock := mocks.NewCloudAzureMock()
	mock.RegisterListMock()
	mock.RegisterGetByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByNameV1(ctx, "Test Azure IDP")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "1", result.Server.ID)
	assert.Equal(t, "Test Azure IDP", result.CloudIdPCommon.DisplayName)
}

func TestGetByNameV1_EmptyName(t *testing.T) {
	mock := mocks.NewCloudAzureMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByNameV1(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "name is required")
}

func TestCreateV1(t *testing.T) {
	mock := mocks.NewCloudAzureMock()
	mock.RegisterCreateMock()

	svc := NewService(mock)
	ctx := context.Background()

	request := &ResourceCloudAzure{
		CloudIdPCommon: CloudIdPCommon{
			DisplayName:  "Test Azure IDP",
			ProviderName: "AZURE",
		},
		Server: CloudAzureServer{
			TenantId: "12345678-1234-1234-1234-123456789012",
			Enabled:  true,
		},
	}

	result, resp, err := svc.CreateV1(ctx, request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, "/api/v1/cloud-azure/1", result.Href)
}

func TestCreateV1_NilRequest(t *testing.T) {
	mock := mocks.NewCloudAzureMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.CreateV1(ctx, nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUpdateByIDV1(t *testing.T) {
	mock := mocks.NewCloudAzureMock()
	mock.RegisterUpdateByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	request := &ResourceCloudAzure{
		CloudIdPCommon: CloudIdPCommon{
			DisplayName:  "Updated Azure IDP",
			ProviderName: "AZURE",
		},
	}

	result, resp, err := svc.UpdateByIDV1(ctx, "1", request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, result)
}

func TestUpdateByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewCloudAzureMock()
	svc := NewService(mock)
	ctx := context.Background()

	request := &ResourceCloudAzure{}

	result, resp, err := svc.UpdateByIDV1(ctx, "", request)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUpdateByIDV1_NilRequest(t *testing.T) {
	mock := mocks.NewCloudAzureMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.UpdateByIDV1(ctx, "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUpdateByNameV1(t *testing.T) {
	mock := mocks.NewCloudAzureMock()
	mock.RegisterListMock()
	mock.RegisterGetByIDMock("1")
	mock.RegisterUpdateByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	request := &ResourceCloudAzure{
		CloudIdPCommon: CloudIdPCommon{
			DisplayName:  "Updated Azure IDP",
			ProviderName: "AZURE",
		},
	}

	result, resp, err := svc.UpdateByNameV1(ctx, "Test Azure IDP", request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.NotNil(t, result)
}

func TestUpdateByNameV1_EmptyName(t *testing.T) {
	mock := mocks.NewCloudAzureMock()
	svc := NewService(mock)
	ctx := context.Background()

	request := &ResourceCloudAzure{}

	result, resp, err := svc.UpdateByNameV1(ctx, "", request)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "name is required")
}

func TestDeleteByIDV1(t *testing.T) {
	mock := mocks.NewCloudAzureMock()
	mock.RegisterDeleteByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteByIDV1(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestDeleteByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewCloudAzureMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteByIDV1(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestDeleteByNameV1(t *testing.T) {
	mock := mocks.NewCloudAzureMock()
	mock.RegisterListMock()
	mock.RegisterGetByIDMock("1")
	mock.RegisterDeleteByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteByNameV1(ctx, "Test Azure IDP")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestDeleteByNameV1_EmptyName(t *testing.T) {
	mock := mocks.NewCloudAzureMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteByNameV1(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "name is required")
}

func TestGetDefaultMappingsV1(t *testing.T) {
	mock := mocks.NewCloudAzureMock()
	mock.RegisterGetDefaultMappingsMock()

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDefaultMappingsV1(ctx)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "objectGUID", result.UserId)
	assert.Equal(t, "userPrincipalName", result.UserName)
	assert.Equal(t, "displayName", result.RealName)
}
