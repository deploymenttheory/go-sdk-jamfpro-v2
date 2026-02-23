package macos_configuration_profile_custom_settings

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/macos_configuration_profile_custom_settings/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testPayloadUUID = "test-uuid-12345"

func TestUnitGetSchemaList_Success(t *testing.T) {
	mock := mocks.NewMacOSConfigProfileCustomSettingsMock()
	mock.RegisterGetSchemaListMock()
	service := NewService(mock)
	ctx := context.Background()

	result, resp, err := service.GetSchemaList(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, *result, 1)
	assert.Equal(t, "com.example.app", (*result)[0].BucketName)
	assert.Equal(t, "Example App Settings", (*result)[0].DisplayName)
}

func TestUnitGetSchemaList_Error(t *testing.T) {
	mock := mocks.NewMacOSConfigProfileCustomSettingsMock()
	mock.RegisterGetSchemaListErrorMock()
	service := NewService(mock)
	ctx := context.Background()

	result, resp, err := service.GetSchemaList(ctx)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request failed")
}

func TestUnitGetSchemaList_NoMockRegistered(t *testing.T) {
	mock := mocks.NewMacOSConfigProfileCustomSettingsMock()
	service := NewService(mock)
	ctx := context.Background()

	result, resp, err := service.GetSchemaList(ctx)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no mock registered")
}

func TestUnitGetByPayloadUUID_Success(t *testing.T) {
	mock := mocks.NewMacOSConfigProfileCustomSettingsMock()
	mock.RegisterGetByPayloadUUIDMock(testPayloadUUID)
	service := NewService(mock)
	ctx := context.Background()

	result, resp, err := service.GetByPayloadUUID(ctx, testPayloadUUID)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, testPayloadUUID, result.PayloadUUID)
	assert.Len(t, result.PayloadContent, 1)
	assert.Equal(t, "com.apple.ManagedClient.preferences", result.PayloadContent[0].PayloadType)
}

func TestUnitGetByPayloadUUID_EmptyID(t *testing.T) {
	mock := mocks.NewMacOSConfigProfileCustomSettingsMock()
	service := NewService(mock)
	ctx := context.Background()

	result, resp, err := service.GetByPayloadUUID(ctx, "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "payload UUID is required")
}

func TestUnitGetByPayloadUUID_Error(t *testing.T) {
	mock := mocks.NewMacOSConfigProfileCustomSettingsMock()
	mock.RegisterGetByPayloadUUIDErrorMock(testPayloadUUID)
	service := NewService(mock)
	ctx := context.Background()

	result, resp, err := service.GetByPayloadUUID(ctx, testPayloadUUID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request failed")
}

func TestUnitGetByPayloadUUID_NoMockRegistered(t *testing.T) {
	mock := mocks.NewMacOSConfigProfileCustomSettingsMock()
	service := NewService(mock)
	ctx := context.Background()

	result, resp, err := service.GetByPayloadUUID(ctx, testPayloadUUID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no mock registered")
}

func TestUnitCreate_Success(t *testing.T) {
	mock := mocks.NewMacOSConfigProfileCustomSettingsMock()
	mock.RegisterCreateMock()
	service := NewService(mock)
	ctx := context.Background()

	profile := &ResourceConfigProfile{
		PayloadUUID: "create-test-uuid",
		PayloadContent: []PayloadContentItem{
			{
				PayloadType:       "com.apple.ManagedClient.preferences",
				PayloadVersion:    1,
				PayloadIdentifier: "com.example.app",
				PayloadUUID:       "payload-uuid-1",
			},
		},
	}

	result, resp, err := service.Create(ctx, profile)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "newly-created-uuid-67890", result.UUID)
}

func TestUnitCreate_NilProfile(t *testing.T) {
	mock := mocks.NewMacOSConfigProfileCustomSettingsMock()
	service := NewService(mock)
	ctx := context.Background()

	result, resp, err := service.Create(ctx, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "profile is required")
}

func TestUnitCreate_Error(t *testing.T) {
	mock := mocks.NewMacOSConfigProfileCustomSettingsMock()
	mock.RegisterCreateErrorMock()
	service := NewService(mock)
	ctx := context.Background()

	profile := &ResourceConfigProfile{
		PayloadUUID:    "create-test-uuid",
		PayloadContent: []PayloadContentItem{},
	}

	result, resp, err := service.Create(ctx, profile)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request failed")
}

func TestUnitCreate_NoMockRegistered(t *testing.T) {
	mock := mocks.NewMacOSConfigProfileCustomSettingsMock()
	service := NewService(mock)
	ctx := context.Background()

	profile := &ResourceConfigProfile{
		PayloadUUID:    "create-test-uuid",
		PayloadContent: []PayloadContentItem{},
	}

	result, resp, err := service.Create(ctx, profile)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no mock registered")
}
