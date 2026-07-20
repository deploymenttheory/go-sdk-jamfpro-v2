package macos_configuration_profiles

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/macos_configuration_profiles/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testPayloadUUID = "test-uuid-12345"

func TestUnit_MacosConfigurationProfiles_GetSchemaList_Success(t *testing.T) {
	mock := mocks.NewMacOSConfigProfilesMock()
	mock.RegisterGetSchemaListMock()
	service := NewMacosConfigurationProfiles(mock)
	ctx := context.Background()

	result, resp, err := service.GetSchemaList(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Len(t, *result, 1)
	assert.Equal(t, "com.example.app", (*result)[0].BucketName)
	assert.Equal(t, "Example App Settings", (*result)[0].DisplayName)
}

func TestUnit_MacosConfigurationProfiles_GetSchemaList_Error(t *testing.T) {
	mock := mocks.NewMacOSConfigProfilesMock()
	mock.RegisterGetSchemaListErrorMock()
	service := NewMacosConfigurationProfiles(mock)
	ctx := context.Background()

	result, resp, err := service.GetSchemaList(ctx)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Contains(t, err.Error(), "request failed")
}

func TestUnit_MacosConfigurationProfiles_GetSchemaList_NoMockRegistered(t *testing.T) {
	mock := mocks.NewMacOSConfigProfilesMock()
	mock.RegisterGetSchemaListNoResponseErrorMock()
	service := NewMacosConfigurationProfiles(mock)
	ctx := context.Background()

	result, resp, err := service.GetSchemaList(ctx)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

func TestUnit_MacosConfigurationProfiles_GetByPayloadUUID_Success(t *testing.T) {
	mock := mocks.NewMacOSConfigProfilesMock()
	mock.RegisterGetByPayloadUUIDMock(testPayloadUUID)
	service := NewMacosConfigurationProfiles(mock)
	ctx := context.Background()

	result, resp, err := service.GetByPayloadUUID(ctx, testPayloadUUID)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, testPayloadUUID, result.PayloadUUID)
	assert.Len(t, result.PayloadContent, 1)
	assert.Equal(t, "com.apple.ManagedClient.preferences", result.PayloadContent[0].PayloadType)
}

func TestUnit_MacosConfigurationProfiles_GetByPayloadUUID_EmptyID(t *testing.T) {
	mock := mocks.NewMacOSConfigProfilesMock()
	service := NewMacosConfigurationProfiles(mock)
	ctx := context.Background()

	result, resp, err := service.GetByPayloadUUID(ctx, "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "payload UUID is required")
}

func TestUnit_MacosConfigurationProfiles_GetByPayloadUUID_Error(t *testing.T) {
	mock := mocks.NewMacOSConfigProfilesMock()
	mock.RegisterGetByPayloadUUIDErrorMock(testPayloadUUID)
	service := NewMacosConfigurationProfiles(mock)
	ctx := context.Background()

	result, resp, err := service.GetByPayloadUUID(ctx, testPayloadUUID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Contains(t, err.Error(), "request failed")
}

func TestUnit_MacosConfigurationProfiles_GetByPayloadUUID_NoMockRegistered(t *testing.T) {
	mock := mocks.NewMacOSConfigProfilesMock()
	mock.RegisterGetByPayloadUUIDNoResponseErrorMock(testPayloadUUID)
	service := NewMacosConfigurationProfiles(mock)
	ctx := context.Background()

	result, resp, err := service.GetByPayloadUUID(ctx, testPayloadUUID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

func TestUnit_MacosConfigurationProfiles_Create_Success(t *testing.T) {
	mock := mocks.NewMacOSConfigProfilesMock()
	mock.RegisterCreateMock()
	service := NewMacosConfigurationProfiles(mock)
	ctx := context.Background()

	profile := &ResourceConfigProfile{
		PayloadUUID: "create-test-uuid",
		PayloadContent: []PayloadContentItem{
			{
				PayloadType:       PayloadTypeManagedClientPreferences,
				PayloadVersion:    1,
				PayloadIdentifier: "com.example.app",
				PayloadUUID:       "payload-uuid-1",
				PreferenceDomain:  "com.example.app",
			},
		},
	}

	result, resp, err := service.Create(ctx, profile)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "newly-created-uuid-67890", result.UUID)
}

func TestUnit_MacosConfigurationProfiles_Create_NilProfile(t *testing.T) {
	mock := mocks.NewMacOSConfigProfilesMock()
	service := NewMacosConfigurationProfiles(mock)
	ctx := context.Background()

	result, resp, err := service.Create(ctx, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "profile is required")
}

func TestUnit_MacosConfigurationProfiles_Create_Error(t *testing.T) {
	mock := mocks.NewMacOSConfigProfilesMock()
	mock.RegisterCreateErrorMock()
	service := NewMacosConfigurationProfiles(mock)
	ctx := context.Background()

	profile := &ResourceConfigProfile{
		PayloadUUID:    "create-test-uuid",
		PayloadContent: []PayloadContentItem{},
	}

	result, resp, err := service.Create(ctx, profile)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Contains(t, err.Error(), "request failed")
}

func TestUnit_MacosConfigurationProfiles_Create_NoMockRegistered(t *testing.T) {
	mock := mocks.NewMacOSConfigProfilesMock()
	mock.RegisterCreateNoResponseErrorMock()
	service := NewMacosConfigurationProfiles(mock)
	ctx := context.Background()

	profile := &ResourceConfigProfile{
		PayloadUUID:    "create-test-uuid",
		PayloadContent: []PayloadContentItem{},
	}

	result, resp, err := service.Create(ctx, profile)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

// validProfile returns a profile that satisfies every server-side constraint.
func validProfile() *ResourceConfigProfile {
	return &ResourceConfigProfile{
		Level: ConfigProfileLevelSystem,
		PayloadContent: []PayloadContentItem{
			{
				PayloadType:       PayloadTypeManagedClientPreferences,
				PayloadVersion:    1,
				PayloadIdentifier: "com.example.app",
				PreferenceDomain:  "com.example.app",
				Forced: &ForcedSettings{
					Plist: `<?xml version="1.0" encoding="UTF-8"?><plist version="1.0"><dict><key>K</key><string>V</string></dict></plist>`,
				},
			},
		},
	}
}

// invalidProfileCases covers every constraint the API enforces with an opaque
// HTTP 400. Each must be caught before a request is issued.
func invalidProfileCases() []struct {
	name    string
	profile *ResourceConfigProfile
	wantErr string
} {
	return []struct {
		name    string
		profile *ResourceConfigProfile
		wantErr string
	}{
		{
			// The title-case form real .mobileconfig files use in PayloadScope.
			name:    "title case level",
			profile: &ResourceConfigProfile{Level: "System"},
			wantErr: "invalid level",
		},
		{
			name:    "unknown level",
			profile: &ResourceConfigProfile{Level: "COMPUTER"},
			wantErr: "invalid level",
		},
		{
			name: "missing preferenceDomain on custom settings payload",
			profile: &ResourceConfigProfile{
				PayloadContent: []PayloadContentItem{
					{PayloadType: PayloadTypeManagedClientPreferences},
				},
			},
			wantErr: "preferenceDomain is required",
		},
		{
			name: "forced without plist",
			profile: &ResourceConfigProfile{
				PayloadContent: []PayloadContentItem{
					{
						PayloadType:      PayloadTypeManagedClientPreferences,
						PreferenceDomain: "com.example.app",
						Forced:           &ForcedSettings{SchemaDomain: "com.example.app"},
					},
				},
			},
			wantErr: "forced.plist is required",
		},
	}
}

func TestUnit_MacosConfigurationProfiles_Create_ValidationErrors(t *testing.T) {
	for _, tc := range invalidProfileCases() {
		t.Run(tc.name, func(t *testing.T) {
			mock := mocks.NewMacOSConfigProfilesMock()
			service := NewMacosConfigurationProfiles(mock)

			result, resp, err := service.Create(context.Background(), tc.profile)
			assert.Error(t, err)
			assert.Nil(t, result)
			// A nil response proves validation short-circuited before the request.
			assert.Nil(t, resp)
			assert.Contains(t, err.Error(), tc.wantErr)
		})
	}
}

func TestUnit_MacosConfigurationProfiles_UpdateByPayloadUUID_ValidationErrors(t *testing.T) {
	for _, tc := range invalidProfileCases() {
		t.Run(tc.name, func(t *testing.T) {
			mock := mocks.NewMacOSConfigProfilesMock()
			service := NewMacosConfigurationProfiles(mock)

			result, resp, err := service.UpdateByPayloadUUID(context.Background(), testPayloadUUID, tc.profile)
			assert.Error(t, err)
			assert.Nil(t, result)
			assert.Nil(t, resp)
			assert.Contains(t, err.Error(), tc.wantErr)
		})
	}
}

func TestUnit_MacosConfigurationProfiles_ValidLevelsAccepted(t *testing.T) {
	for _, level := range []string{ConfigProfileLevelSystem, ConfigProfileLevelUser, ""} {
		mock := mocks.NewMacOSConfigProfilesMock()
		mock.RegisterCreateMock()
		service := NewMacosConfigurationProfiles(mock)

		profile := validProfile()
		profile.Level = level

		result, _, err := service.Create(context.Background(), profile)
		require.NoError(t, err, "level %q should be accepted", level)
		require.NotNil(t, result)
	}
}

// TestUnit_MacosConfigurationProfiles_NonCustomSettingsPayloadsAccepted guards
// against re-introducing two over-broad client-side rules: a payloadType
// allowlist, and requiring preferenceDomain on every payload. Both were
// verified wrong against Jamf Pro 11.30.0 -- these payloads create successfully
// server-side, so the SDK must not reject them.
func TestUnit_MacosConfigurationProfiles_NonCustomSettingsPayloadsAccepted(t *testing.T) {
	payloadTypes := []string{
		PayloadTypeNotificationSettings,
		PayloadTypePasswordPolicy,
		// A type not in any SDK constant, to prove there is no allowlist.
		"com.apple.dnsSettings.managed",
	}

	for _, payloadType := range payloadTypes {
		t.Run(payloadType, func(t *testing.T) {
			mock := mocks.NewMacOSConfigProfilesMock()
			mock.RegisterCreateMock()
			service := NewMacosConfigurationProfiles(mock)

			profile := &ResourceConfigProfile{
				Level: ConfigProfileLevelSystem,
				PayloadContent: []PayloadContentItem{
					{
						PayloadType:       payloadType,
						PayloadVersion:    1,
						PayloadIdentifier: "com.example.app",
						// No PreferenceDomain: not required for these types.
					},
				},
			}

			result, _, err := service.Create(context.Background(), profile)
			require.NoError(t, err)
			require.NotNil(t, result)
		})
	}
}

func TestUnit_MacosConfigurationProfiles_UpdateByPayloadUUID_Success(t *testing.T) {
	mock := mocks.NewMacOSConfigProfilesMock()
	mock.RegisterUpdateByPayloadUUIDMock(testPayloadUUID)
	service := NewMacosConfigurationProfiles(mock)
	ctx := context.Background()

	result, resp, err := service.UpdateByPayloadUUID(ctx, testPayloadUUID, validProfile())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	// Update returns the create envelope, not the updated resource.
	assert.Equal(t, testPayloadUUID, result.UUID)
}

func TestUnit_MacosConfigurationProfiles_UpdateByPayloadUUID_EmptyID(t *testing.T) {
	mock := mocks.NewMacOSConfigProfilesMock()
	service := NewMacosConfigurationProfiles(mock)
	ctx := context.Background()

	result, resp, err := service.UpdateByPayloadUUID(ctx, "", validProfile())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "payload UUID is required")
}

func TestUnit_MacosConfigurationProfiles_UpdateByPayloadUUID_NilProfile(t *testing.T) {
	mock := mocks.NewMacOSConfigProfilesMock()
	service := NewMacosConfigurationProfiles(mock)
	ctx := context.Background()

	result, resp, err := service.UpdateByPayloadUUID(ctx, testPayloadUUID, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "profile is required")
}

func TestUnit_MacosConfigurationProfiles_UpdateByPayloadUUID_Error(t *testing.T) {
	mock := mocks.NewMacOSConfigProfilesMock()
	mock.RegisterUpdateByPayloadUUIDErrorMock(testPayloadUUID)
	service := NewMacosConfigurationProfiles(mock)
	ctx := context.Background()

	result, resp, err := service.UpdateByPayloadUUID(ctx, testPayloadUUID, validProfile())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Contains(t, err.Error(), "request failed")
}

func TestUnit_MacosConfigurationProfiles_UpdateByPayloadUUID_NoMockRegistered(t *testing.T) {
	mock := mocks.NewMacOSConfigProfilesMock()
	mock.RegisterUpdateByPayloadUUIDNoResponseErrorMock(testPayloadUUID)
	service := NewMacosConfigurationProfiles(mock)
	ctx := context.Background()

	result, resp, err := service.UpdateByPayloadUUID(ctx, testPayloadUUID, validProfile())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

func TestUnit_MacosConfigurationProfiles_DeleteByPayloadUUID_Success(t *testing.T) {
	mock := mocks.NewMacOSConfigProfilesMock()
	mock.RegisterDeleteByPayloadUUIDMock(testPayloadUUID)
	service := NewMacosConfigurationProfiles(mock)
	ctx := context.Background()

	resp, err := service.DeleteByPayloadUUID(ctx, testPayloadUUID)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_MacosConfigurationProfiles_DeleteByPayloadUUID_EmptyID(t *testing.T) {
	mock := mocks.NewMacOSConfigProfilesMock()
	service := NewMacosConfigurationProfiles(mock)
	ctx := context.Background()

	resp, err := service.DeleteByPayloadUUID(ctx, "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "payload UUID is required")
}

func TestUnit_MacosConfigurationProfiles_DeleteByPayloadUUID_Error(t *testing.T) {
	mock := mocks.NewMacOSConfigProfilesMock()
	mock.RegisterDeleteByPayloadUUIDErrorMock(testPayloadUUID)
	service := NewMacosConfigurationProfiles(mock)
	ctx := context.Background()

	resp, err := service.DeleteByPayloadUUID(ctx, testPayloadUUID)
	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Contains(t, err.Error(), "request failed")
}

func TestUnit_MacosConfigurationProfiles_DeleteByPayloadUUID_NoMockRegistered(t *testing.T) {
	mock := mocks.NewMacOSConfigProfilesMock()
	mock.RegisterDeleteByPayloadUUIDNoResponseErrorMock(testPayloadUUID)
	service := NewMacosConfigurationProfiles(mock)
	ctx := context.Background()

	resp, err := service.DeleteByPayloadUUID(ctx, testPayloadUUID)
	assert.Error(t, err)
	assert.NotNil(t, resp)
}
