package mobile_device_enrollment_profiles_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_device_enrollment_profiles"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_device_enrollment_profiles/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_MobileDeviceEnrollmentProfiles_List(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	mockClient.RegisterListMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 2, resp.Size)
	assert.Len(t, resp.Results, 2)
	assert.Equal(t, "Test Enrollment Profile", resp.Results[0].Name)
	assert.Equal(t, float64(1234567890.123456), resp.Results[0].Invitation)
	assert.Equal(t, "Another Profile", resp.Results[1].Name)
}

func TestUnit_MobileDeviceEnrollmentProfiles_GetByID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	mockClient.RegisterGetByIDMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	resp, _, err := svc.GetByID(context.Background(), 1)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Test Enrollment Profile", resp.General.Name)
	assert.Equal(t, "Test profile for mobile device enrollment", resp.General.Description)
	assert.NotNil(t, resp.Location)
	assert.Equal(t, "jdoe", resp.Location.Username)
	assert.Equal(t, "John Doe", resp.Location.Realname)
	assert.NotNil(t, resp.Purchasing)
	assert.True(t, resp.Purchasing.IsPurchased)
	assert.Len(t, resp.Attachments, 1)
	assert.Equal(t, "profile.mobileconfig", resp.Attachments[0].Filename)
}

func TestUnit_MobileDeviceEnrollmentProfiles_GetByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device enrollment profile ID must be a positive integer")
}

func TestUnit_MobileDeviceEnrollmentProfiles_GetByName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	mockClient.RegisterGetByNameMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "Test Enrollment Profile")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Test Enrollment Profile", resp.General.Name)
}

func TestUnit_MobileDeviceEnrollmentProfiles_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device enrollment profile name cannot be empty")
}

func TestUnit_MobileDeviceEnrollmentProfiles_GetByInvitation(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	mockClient.RegisterGetByInvitationMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	resp, _, err := svc.GetByInvitation(context.Background(), "1234567890.123456")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Test Enrollment Profile", resp.General.Name)
}

func TestUnit_MobileDeviceEnrollmentProfiles_GetByInvitation_EmptyInvitation(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, _, err := svc.GetByInvitation(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device enrollment profile invitation cannot be empty")
}

func TestUnit_MobileDeviceEnrollmentProfiles_GetByIDWithSubset(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	mockClient.RegisterGetByIDWithSubsetMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	resp, _, err := svc.GetByIDWithSubset(context.Background(), 1, "General")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Test Enrollment Profile", resp.General.Name)
}

func TestUnit_MobileDeviceEnrollmentProfiles_GetByIDWithSubset_EmptySubset(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, _, err := svc.GetByIDWithSubset(context.Background(), 1, "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "subset cannot be empty")
}

func TestUnit_MobileDeviceEnrollmentProfiles_GetByNameWithSubset(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	mockClient.RegisterGetByNameWithSubsetMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	resp, _, err := svc.GetByNameWithSubset(context.Background(), "Test Enrollment Profile", "General")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Test Enrollment Profile", resp.General.Name)
}

func TestUnit_MobileDeviceEnrollmentProfiles_Create(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	mockClient.RegisterCreateMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	req := &mobile_device_enrollment_profiles.Resource{
		General: mobile_device_enrollment_profiles.SubsetGeneral{
			Name:        "New Enrollment Profile",
			Description: "Newly created profile",
		},
	}

	resp, _, err := svc.Create(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.General.ID)
	assert.Equal(t, "New Enrollment Profile", resp.General.Name)
}

func TestUnit_MobileDeviceEnrollmentProfiles_Create_NilRequest(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceEnrollmentProfiles_Create_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	req := &mobile_device_enrollment_profiles.Resource{
		General: mobile_device_enrollment_profiles.SubsetGeneral{
			Name: "",
		},
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device enrollment profile name is required")
}

func TestUnit_MobileDeviceEnrollmentProfiles_UpdateByID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	mockClient.RegisterUpdateByIDMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	req := &mobile_device_enrollment_profiles.Resource{
		General: mobile_device_enrollment_profiles.SubsetGeneral{
			Name:        "Updated Enrollment Profile",
			Description: "Updated profile description",
		},
	}

	resp, _, err := svc.UpdateByID(context.Background(), 1, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Updated Enrollment Profile", resp.General.Name)
}

func TestUnit_MobileDeviceEnrollmentProfiles_UpdateByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	req := &mobile_device_enrollment_profiles.Resource{
		General: mobile_device_enrollment_profiles.SubsetGeneral{
			Name: "Test",
		},
	}

	_, _, err := svc.UpdateByID(context.Background(), 0, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device enrollment profile ID must be a positive integer")
}

func TestUnit_MobileDeviceEnrollmentProfiles_UpdateByName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	mockClient.RegisterUpdateByNameMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	req := &mobile_device_enrollment_profiles.Resource{
		General: mobile_device_enrollment_profiles.SubsetGeneral{
			Name:        "Updated Enrollment Profile",
			Description: "Updated profile description",
		},
	}

	resp, _, err := svc.UpdateByName(context.Background(), "Test Enrollment Profile", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
}

func TestUnit_MobileDeviceEnrollmentProfiles_UpdateByInvitation(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	mockClient.RegisterUpdateByInvitationMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	req := &mobile_device_enrollment_profiles.Resource{
		General: mobile_device_enrollment_profiles.SubsetGeneral{
			Name:        "Updated Enrollment Profile",
			Description: "Updated profile description",
		},
	}

	resp, _, err := svc.UpdateByInvitation(context.Background(), "1234567890.123456", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
}

func TestUnit_MobileDeviceEnrollmentProfiles_DeleteByID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	mockClient.RegisterDeleteByIDMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 1)

	require.NoError(t, err)
}

func TestUnit_MobileDeviceEnrollmentProfiles_DeleteByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device enrollment profile ID must be a positive integer")
}

func TestUnit_MobileDeviceEnrollmentProfiles_DeleteByName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	mockClient.RegisterDeleteByNameMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "Test Enrollment Profile")

	require.NoError(t, err)
}

func TestUnit_MobileDeviceEnrollmentProfiles_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device enrollment profile name cannot be empty")
}

func TestUnit_MobileDeviceEnrollmentProfiles_DeleteByInvitation(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	mockClient.RegisterDeleteByInvitationMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, err := svc.DeleteByInvitation(context.Background(), "1234567890.123456")

	require.NoError(t, err)
}

func TestUnit_MobileDeviceEnrollmentProfiles_DeleteByInvitation_EmptyInvitation(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, err := svc.DeleteByInvitation(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device enrollment profile invitation cannot be empty")
}

func TestUnit_MobileDeviceEnrollmentProfiles_NotFound(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 999)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}

func TestUnit_MobileDeviceEnrollmentProfiles_List_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, _, err := svc.List(context.Background())
	require.Error(t, err)
}

func TestUnit_MobileDeviceEnrollmentProfiles_GetByName_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, _, err := svc.GetByName(context.Background(), "Unknown")
	require.Error(t, err)
}

func TestUnit_MobileDeviceEnrollmentProfiles_GetByInvitation_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, _, err := svc.GetByInvitation(context.Background(), "unknown")
	require.Error(t, err)
}

func TestUnit_MobileDeviceEnrollmentProfiles_GetByIDWithSubset_ZeroID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, _, err := svc.GetByIDWithSubset(context.Background(), 0, "General")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device enrollment profile ID must be a positive integer")
}

func TestUnit_MobileDeviceEnrollmentProfiles_GetByIDWithSubset_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, _, err := svc.GetByIDWithSubset(context.Background(), 2, "General")
	require.Error(t, err)
}

func TestUnit_MobileDeviceEnrollmentProfiles_GetByNameWithSubset_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, _, err := svc.GetByNameWithSubset(context.Background(), "", "General")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device enrollment profile name cannot be empty")
}

func TestUnit_MobileDeviceEnrollmentProfiles_GetByNameWithSubset_EmptySubset(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, _, err := svc.GetByNameWithSubset(context.Background(), "Test", "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "subset cannot be empty")
}

func TestUnit_MobileDeviceEnrollmentProfiles_GetByNameWithSubset_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, _, err := svc.GetByNameWithSubset(context.Background(), "Unknown", "General")
	require.Error(t, err)
}

func TestUnit_MobileDeviceEnrollmentProfiles_Create_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	req := &mobile_device_enrollment_profiles.Resource{
		General: mobile_device_enrollment_profiles.SubsetGeneral{Name: "Test"},
	}
	_, _, err := svc.Create(context.Background(), req)
	require.Error(t, err)
}

func TestUnit_MobileDeviceEnrollmentProfiles_UpdateByID_NilRequest(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, _, err := svc.UpdateByID(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceEnrollmentProfiles_UpdateByID_EmptyReqName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	req := &mobile_device_enrollment_profiles.Resource{
		General: mobile_device_enrollment_profiles.SubsetGeneral{Name: ""},
	}
	_, _, err := svc.UpdateByID(context.Background(), 1, req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device enrollment profile name is required in request")
}

func TestUnit_MobileDeviceEnrollmentProfiles_UpdateByID_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	req := &mobile_device_enrollment_profiles.Resource{
		General: mobile_device_enrollment_profiles.SubsetGeneral{Name: "Test"},
	}
	_, _, err := svc.UpdateByID(context.Background(), 2, req)
	require.Error(t, err)
}

func TestUnit_MobileDeviceEnrollmentProfiles_UpdateByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	req := &mobile_device_enrollment_profiles.Resource{
		General: mobile_device_enrollment_profiles.SubsetGeneral{Name: "Test"},
	}
	_, _, err := svc.UpdateByName(context.Background(), "", req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device enrollment profile name cannot be empty")
}

func TestUnit_MobileDeviceEnrollmentProfiles_UpdateByName_NilRequest(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, _, err := svc.UpdateByName(context.Background(), "Test", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceEnrollmentProfiles_UpdateByName_EmptyReqName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	req := &mobile_device_enrollment_profiles.Resource{
		General: mobile_device_enrollment_profiles.SubsetGeneral{Name: ""},
	}
	_, _, err := svc.UpdateByName(context.Background(), "Test", req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device enrollment profile name is required in request")
}

func TestUnit_MobileDeviceEnrollmentProfiles_UpdateByName_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	req := &mobile_device_enrollment_profiles.Resource{
		General: mobile_device_enrollment_profiles.SubsetGeneral{Name: "Test"},
	}
	_, _, err := svc.UpdateByName(context.Background(), "Unknown", req)
	require.Error(t, err)
}

func TestUnit_MobileDeviceEnrollmentProfiles_UpdateByInvitation_EmptyInvitation(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	req := &mobile_device_enrollment_profiles.Resource{
		General: mobile_device_enrollment_profiles.SubsetGeneral{Name: "Test"},
	}
	_, _, err := svc.UpdateByInvitation(context.Background(), "", req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device enrollment profile invitation cannot be empty")
}

func TestUnit_MobileDeviceEnrollmentProfiles_UpdateByInvitation_NilRequest(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, _, err := svc.UpdateByInvitation(context.Background(), "123", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceEnrollmentProfiles_UpdateByInvitation_EmptyReqName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	req := &mobile_device_enrollment_profiles.Resource{
		General: mobile_device_enrollment_profiles.SubsetGeneral{Name: ""},
	}
	_, _, err := svc.UpdateByInvitation(context.Background(), "123", req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device enrollment profile name is required in request")
}

func TestUnit_MobileDeviceEnrollmentProfiles_UpdateByInvitation_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	req := &mobile_device_enrollment_profiles.Resource{
		General: mobile_device_enrollment_profiles.SubsetGeneral{Name: "Test"},
	}
	_, _, err := svc.UpdateByInvitation(context.Background(), "unknown", req)
	require.Error(t, err)
}

func TestUnit_MobileDeviceEnrollmentProfiles_DeleteByID_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 999)
	require.Error(t, err)
}

func TestUnit_MobileDeviceEnrollmentProfiles_DeleteByName_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "Unknown")
	require.Error(t, err)
}

func TestUnit_MobileDeviceEnrollmentProfiles_DeleteByInvitation_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceEnrollmentProfilesMock()
	svc := mobile_device_enrollment_profiles.NewService(mockClient)

	_, err := svc.DeleteByInvitation(context.Background(), "unknown")
	require.Error(t, err)
}
