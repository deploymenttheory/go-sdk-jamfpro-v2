package mobile_device_provisioning_profiles_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_device_provisioning_profiles"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_device_provisioning_profiles/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_MobileDeviceProvisioningProfiles_List(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	mockClient.RegisterListMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 2, resp.Size)
	assert.Len(t, resp.Profiles, 2)
	assert.Equal(t, "Test Provisioning Profile", resp.Profiles[0].Name)
	assert.Equal(t, 1, resp.Profiles[0].ID)
	assert.Equal(t, "Another Provisioning Profile", resp.Profiles[1].Name)
	assert.Equal(t, 2, resp.Profiles[1].ID)
}

func TestUnit_MobileDeviceProvisioningProfiles_GetByID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	mockClient.RegisterGetByIDMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	resp, _, err := svc.GetByID(context.Background(), 1)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Test Provisioning Profile", resp.General.Name)
	assert.Equal(t, "550e8400-e29b-41d4-a716-446655440000", resp.General.UUID)
}

func TestUnit_MobileDeviceProvisioningProfiles_GetByID_NegativeID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), -1)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device provisioning profile ID must be a non-negative integer")
}

func TestUnit_MobileDeviceProvisioningProfiles_GetByName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	mockClient.RegisterGetByNameMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "Test Provisioning Profile")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Test Provisioning Profile", resp.General.Name)
}

func TestUnit_MobileDeviceProvisioningProfiles_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device provisioning profile name cannot be empty")
}

func TestUnit_MobileDeviceProvisioningProfiles_GetByUUID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	mockClient.RegisterGetByUUIDMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	resp, _, err := svc.GetByUUID(context.Background(), "550e8400-e29b-41d4-a716-446655440000")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "550e8400-e29b-41d4-a716-446655440000", resp.General.UUID)
}

func TestUnit_MobileDeviceProvisioningProfiles_GetByUUID_EmptyUUID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	_, _, err := svc.GetByUUID(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device provisioning profile UUID cannot be empty")
}

func TestUnit_MobileDeviceProvisioningProfiles_CreateByID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	mockClient.RegisterCreateByIDMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	req := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{
			Name:        "Created Provisioning Profile",
			DisplayName: "Created Provisioning Profile",
			UUID:        "550e8400-e29b-41d4-a716-446655440001",
		},
	}

	resp, _, err := svc.CreateByID(context.Background(), 0, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.ID)
}

func TestUnit_MobileDeviceProvisioningProfiles_CreateByID_NilRequest(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	_, _, err := svc.CreateByID(context.Background(), 0, nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceProvisioningProfiles_CreateByID_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	req := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{
			Name: "",
		},
	}

	_, _, err := svc.CreateByID(context.Background(), 0, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device provisioning profile name is required")
}

func TestUnit_MobileDeviceProvisioningProfiles_CreateByName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	mockClient.RegisterCreateByNameMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	req := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{
			Name: "New Profile",
		},
	}

	resp, _, err := svc.CreateByName(context.Background(), "New Profile", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.ID)
}

func TestUnit_MobileDeviceProvisioningProfiles_CreateByUUID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	mockClient.RegisterCreateByUUIDMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	req := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{
			Name: "Created Provisioning Profile",
			UUID: "550e8400-e29b-41d4-a716-446655440001",
		},
	}

	resp, _, err := svc.CreateByUUID(context.Background(), "550e8400-e29b-41d4-a716-446655440001", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.ID)
}

func TestUnit_MobileDeviceProvisioningProfiles_UpdateByID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	mockClient.RegisterUpdateByIDMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	req := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{
			Name: "Updated Provisioning Profile",
		},
	}

	resp, _, err := svc.UpdateByID(context.Background(), 1, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_MobileDeviceProvisioningProfiles_UpdateByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	req := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{Name: "Test"},
	}

	_, _, err := svc.UpdateByID(context.Background(), 0, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device provisioning profile ID must be a positive integer")
}

func TestUnit_MobileDeviceProvisioningProfiles_UpdateByName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	mockClient.RegisterUpdateByNameMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	req := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{
			Name: "Updated Provisioning Profile",
		},
	}

	resp, _, err := svc.UpdateByName(context.Background(), "Test Provisioning Profile", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_MobileDeviceProvisioningProfiles_UpdateByUUID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	mockClient.RegisterUpdateByUUIDMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	req := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{
			Name: "Updated Provisioning Profile",
		},
	}

	resp, _, err := svc.UpdateByUUID(context.Background(), "550e8400-e29b-41d4-a716-446655440000", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_MobileDeviceProvisioningProfiles_DeleteByID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	mockClient.RegisterDeleteByIDMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 1)

	require.NoError(t, err)
}

func TestUnit_MobileDeviceProvisioningProfiles_DeleteByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device provisioning profile ID must be a positive integer")
}

func TestUnit_MobileDeviceProvisioningProfiles_DeleteByName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	mockClient.RegisterDeleteByNameMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "Test Provisioning Profile")

	require.NoError(t, err)
}

func TestUnit_MobileDeviceProvisioningProfiles_DeleteByUUID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	mockClient.RegisterDeleteByUUIDMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	_, err := svc.DeleteByUUID(context.Background(), "550e8400-e29b-41d4-a716-446655440000")

	require.NoError(t, err)
}

func TestUnit_MobileDeviceProvisioningProfiles_NotFound(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 999)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}

func TestUnit_MobileDeviceProvisioningProfiles_CreateByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	req := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{Name: "Test"},
	}
	_, _, err := svc.CreateByName(context.Background(), "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "name cannot be empty")
}

func TestUnit_MobileDeviceProvisioningProfiles_CreateByName_NilRequest(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	_, _, err := svc.CreateByName(context.Background(), "New Profile", nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceProvisioningProfiles_CreateByUUID_EmptyUUID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	req := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{Name: "Test"},
	}
	_, _, err := svc.CreateByUUID(context.Background(), "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "UUID cannot be empty")
}

func TestUnit_MobileDeviceProvisioningProfiles_CreateByUUID_NilRequest(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	_, _, err := svc.CreateByUUID(context.Background(), "550e8400-e29b-41d4-a716-446655440001", nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceProvisioningProfiles_UpdateByID_NilRequest(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	_, _, err := svc.UpdateByID(context.Background(), 1, nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceProvisioningProfiles_UpdateByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	req := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{Name: "Test"},
	}
	_, _, err := svc.UpdateByName(context.Background(), "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "name cannot be empty")
}

func TestUnit_MobileDeviceProvisioningProfiles_UpdateByName_NilRequest(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	_, _, err := svc.UpdateByName(context.Background(), "Test Provisioning Profile", nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceProvisioningProfiles_UpdateByUUID_EmptyUUID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	req := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{Name: "Test"},
	}
	_, _, err := svc.UpdateByUUID(context.Background(), "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "UUID cannot be empty")
}

func TestUnit_MobileDeviceProvisioningProfiles_UpdateByUUID_NilRequest(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	_, _, err := svc.UpdateByUUID(context.Background(), "550e8400-e29b-41d4-a716-446655440000", nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceProvisioningProfiles_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "name cannot be empty")
}

func TestUnit_MobileDeviceProvisioningProfiles_DeleteByUUID_EmptyUUID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	_, err := svc.DeleteByUUID(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "UUID cannot be empty")
}

func TestUnit_MobileDeviceProvisioningProfiles_CreateByName_EmptyReqName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	req := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{Name: ""},
	}
	_, _, err := svc.CreateByName(context.Background(), "New Profile", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "name is required")
}

func TestUnit_MobileDeviceProvisioningProfiles_CreateByUUID_EmptyReqName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	req := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{Name: ""},
	}
	_, _, err := svc.CreateByUUID(context.Background(), "550e8400-e29b-41d4-a716-446655440001", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "name is required")
}

func TestUnit_MobileDeviceProvisioningProfiles_UpdateByID_EmptyReqName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	req := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{Name: ""},
	}
	_, _, err := svc.UpdateByID(context.Background(), 1, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "name is required")
}

func TestUnit_MobileDeviceProvisioningProfiles_UpdateByName_EmptyReqName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	req := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{Name: ""},
	}
	_, _, err := svc.UpdateByName(context.Background(), "Test Provisioning Profile", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "name is required")
}

func TestUnit_MobileDeviceProvisioningProfiles_UpdateByUUID_EmptyReqName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	req := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{Name: ""},
	}
	_, _, err := svc.UpdateByUUID(context.Background(), "550e8400-e29b-41d4-a716-446655440000", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "name is required")
}

func TestUnit_MobileDeviceProvisioningProfiles_DeleteByName_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "nonexistent")

	require.Error(t, err)
}

func TestUnit_MobileDeviceProvisioningProfiles_DeleteByUUID_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	_, err := svc.DeleteByUUID(context.Background(), "00000000-0000-0000-0000-000000000000")

	require.Error(t, err)
}

func TestUnit_MobileDeviceProvisioningProfiles_List_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	_, _, err := svc.List(context.Background())

	require.Error(t, err)
}

func TestUnit_MobileDeviceProvisioningProfiles_Conflict(t *testing.T) {
	mockClient := mocks.NewMobileDeviceProvisioningProfilesMock()
	mockClient.RegisterConflictErrorMock()
	svc := mobile_device_provisioning_profiles.NewService(mockClient)

	req := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{
			Name: "Duplicate Profile",
		},
	}

	_, _, err := svc.CreateByID(context.Background(), 0, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device provisioning profile with that name already exists")
}
