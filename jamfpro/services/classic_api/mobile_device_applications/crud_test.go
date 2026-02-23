package mobile_device_applications_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_device_applications"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_device_applications/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_MobileDeviceApplications_List(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	mockClient.RegisterListMobileDeviceApplicationsMock()
	svc := mobile_device_applications.NewService(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 2, resp.Size)
	assert.Len(t, resp.Results, 2)
	assert.Equal(t, "Sample iOS App 1", resp.Results[0].Name)
	assert.Equal(t, "com.example.app1", resp.Results[0].BundleID)
	assert.Equal(t, "Sample iOS App 2", resp.Results[1].Name)
	assert.Equal(t, "com.example.app2", resp.Results[1].BundleID)
}

func TestUnit_MobileDeviceApplications_GetByID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	mockClient.RegisterGetMobileDeviceApplicationByIDMock()
	svc := mobile_device_applications.NewService(mockClient)

	resp, _, err := svc.GetByID(context.Background(), 1)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Sample iOS App", resp.General.Name)
	assert.Equal(t, "com.example.app", resp.General.BundleID)
	assert.Equal(t, "iOS", resp.General.OsType)
	assert.NotNil(t, resp.General.Site)
	assert.Equal(t, -1, resp.General.Site.ID)
}

func TestUnit_MobileDeviceApplications_GetByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	svc := mobile_device_applications.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device application ID must be a positive integer")
}

func TestUnit_MobileDeviceApplications_GetByName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	mockClient.RegisterGetMobileDeviceApplicationByNameMock()
	svc := mobile_device_applications.NewService(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "Sample iOS App")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Sample iOS App", resp.General.Name)
}

func TestUnit_MobileDeviceApplications_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	svc := mobile_device_applications.NewService(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device application name cannot be empty")
}

func TestUnit_MobileDeviceApplications_GetByBundleID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	mockClient.RegisterGetMobileDeviceApplicationByBundleIDMock()
	svc := mobile_device_applications.NewService(mockClient)

	resp, _, err := svc.GetByBundleID(context.Background(), "com.example.app")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "com.example.app", resp.General.BundleID)
}

func TestUnit_MobileDeviceApplications_GetByBundleIDAndVersion(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	mockClient.RegisterGetMobileDeviceApplicationByBundleIDAndVersionMock()
	svc := mobile_device_applications.NewService(mockClient)

	resp, _, err := svc.GetByBundleIDAndVersion(context.Background(), "com.example.app", "1.0")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "com.example.app", resp.General.BundleID)
	assert.Equal(t, "1.0", resp.General.Version)
}

func TestUnit_MobileDeviceApplications_GetByIDAndSubset(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	mockClient.RegisterGetMobileDeviceApplicationByIDAndSubsetMock()
	svc := mobile_device_applications.NewService(mockClient)

	resp, _, err := svc.GetByIDAndSubset(context.Background(), 1, "General")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "Sample iOS App", resp.General.Name)
}

func TestUnit_MobileDeviceApplications_GetByNameAndSubset(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	mockClient.RegisterGetMobileDeviceApplicationByNameAndSubsetMock()
	svc := mobile_device_applications.NewService(mockClient)

	resp, _, err := svc.GetByNameAndSubset(context.Background(), "Sample iOS App", "General")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "Sample iOS App", resp.General.Name)
}

func TestUnit_MobileDeviceApplications_Create(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	mockClient.RegisterCreateMobileDeviceApplicationMock()
	svc := mobile_device_applications.NewService(mockClient)

	internalApp := false
	req := &mobile_device_applications.Resource{
		General: mobile_device_applications.SubsetGeneral{
			Name:        "Test iOS App",
			DisplayName: "Test iOS App",
			BundleID:    "com.test.app",
			Version:     "1.0",
			InternalApp: &internalApp,
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: mobile_device_applications.SubsetScope{
			AllMobileDevices: boolPtr(true),
			AllJSSUsers:      boolPtr(false),
		},
		SelfService: mobile_device_applications.SubsetSelfService{
			SelfServiceDescription: "Test app for self service",
		},
	}

	resp, _, err := svc.Create(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.ID)
}

func TestUnit_MobileDeviceApplications_Create_NilRequest(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	svc := mobile_device_applications.NewService(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceApplications_Create_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	svc := mobile_device_applications.NewService(mockClient)

	req := &mobile_device_applications.Resource{
		General: mobile_device_applications.SubsetGeneral{
			Name: "",
		},
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device application name is required")
}

func TestUnit_MobileDeviceApplications_UpdateByID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	mockClient.RegisterUpdateMobileDeviceApplicationByIDMock()
	svc := mobile_device_applications.NewService(mockClient)

	req := &mobile_device_applications.Resource{
		General: mobile_device_applications.SubsetGeneral{
			Name:        "Updated iOS App",
			DisplayName: "Updated iOS App",
			BundleID:    "com.example.app",
			Version:     "1.1",
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: mobile_device_applications.SubsetScope{
			AllMobileDevices: boolPtr(true),
			AllJSSUsers:      boolPtr(false),
		},
		SelfService: mobile_device_applications.SubsetSelfService{
			SelfServiceDescription: "Updated description",
		},
	}

	resp, _, err := svc.UpdateByID(context.Background(), 1, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "Sample iOS App Updated", resp.General.Name)
}

func TestUnit_MobileDeviceApplications_UpdateByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	svc := mobile_device_applications.NewService(mockClient)

	req := &mobile_device_applications.Resource{
		General: mobile_device_applications.SubsetGeneral{
			Name:     "Test",
			BundleID: "com.test.app",
		},
	}

	_, _, err := svc.UpdateByID(context.Background(), 0, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device application ID must be a positive integer")
}

func TestUnit_MobileDeviceApplications_UpdateByName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	mockClient.RegisterUpdateMobileDeviceApplicationByNameMock()
	svc := mobile_device_applications.NewService(mockClient)

	req := &mobile_device_applications.Resource{
		General: mobile_device_applications.SubsetGeneral{
			Name:        "Updated iOS App",
			BundleID:    "com.example.app",
			Version:     "1.1",
			Site:        &shared.SharedResourceSite{ID: -1, Name: "None"},
		},
		Scope:       mobile_device_applications.SubsetScope{AllMobileDevices: boolPtr(true), AllJSSUsers: boolPtr(false)},
		SelfService: mobile_device_applications.SubsetSelfService{SelfServiceDescription: "Updated"},
	}

	resp, _, err := svc.UpdateByName(context.Background(), "Sample iOS App", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "Sample iOS App Updated", resp.General.Name)
}

func TestUnit_MobileDeviceApplications_UpdateByBundleID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	mockClient.RegisterUpdateMobileDeviceApplicationByBundleIDMock()
	svc := mobile_device_applications.NewService(mockClient)

	req := &mobile_device_applications.Resource{
		General: mobile_device_applications.SubsetGeneral{
			Name:     "Updated",
			BundleID: "com.example.app",
			Site:     &shared.SharedResourceSite{ID: -1, Name: "None"},
		},
		Scope:       mobile_device_applications.SubsetScope{AllMobileDevices: boolPtr(true), AllJSSUsers: boolPtr(false)},
		SelfService: mobile_device_applications.SubsetSelfService{},
	}

	resp, _, err := svc.UpdateByBundleID(context.Background(), "com.example.app", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestUnit_MobileDeviceApplications_UpdateByIDAndVersion(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	mockClient.RegisterUpdateMobileDeviceApplicationByIDAndVersionMock()
	svc := mobile_device_applications.NewService(mockClient)

	req := &mobile_device_applications.Resource{
		General: mobile_device_applications.SubsetGeneral{
			Name:     "Updated",
			BundleID: "com.example.app",
			Version:  "1.0",
			Site:     &shared.SharedResourceSite{ID: -1, Name: "None"},
		},
		Scope:       mobile_device_applications.SubsetScope{AllMobileDevices: boolPtr(true), AllJSSUsers: boolPtr(false)},
		SelfService: mobile_device_applications.SubsetSelfService{},
	}

	resp, _, err := svc.UpdateByIDAndVersion(context.Background(), 1, "1.0", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestUnit_MobileDeviceApplications_DeleteByID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	mockClient.RegisterDeleteMobileDeviceApplicationByIDMock()
	svc := mobile_device_applications.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 1)

	require.NoError(t, err)
}

func TestUnit_MobileDeviceApplications_DeleteByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	svc := mobile_device_applications.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device application ID must be a positive integer")
}

func TestUnit_MobileDeviceApplications_DeleteByName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	mockClient.RegisterDeleteMobileDeviceApplicationByNameMock()
	svc := mobile_device_applications.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "Sample iOS App")

	require.NoError(t, err)
}

func TestUnit_MobileDeviceApplications_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	svc := mobile_device_applications.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device application name cannot be empty")
}

func TestUnit_MobileDeviceApplications_DeleteByBundleID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	mockClient.RegisterDeleteMobileDeviceApplicationByBundleIDMock()
	svc := mobile_device_applications.NewService(mockClient)

	_, err := svc.DeleteByBundleID(context.Background(), "com.example.app")

	require.NoError(t, err)
}

func TestUnit_MobileDeviceApplications_DeleteByBundleIDAndVersion(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	mockClient.RegisterDeleteMobileDeviceApplicationByBundleIDAndVersionMock()
	svc := mobile_device_applications.NewService(mockClient)

	_, err := svc.DeleteByBundleIDAndVersion(context.Background(), "com.example.app", "1.0")

	require.NoError(t, err)
}

func TestUnit_MobileDeviceApplications_NotFound(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := mobile_device_applications.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 999)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}

func TestUnit_MobileDeviceApplications_Conflict(t *testing.T) {
	mockClient := mocks.NewMobileDeviceApplicationsMock()
	mockClient.RegisterConflictErrorMock()
	svc := mobile_device_applications.NewService(mockClient)

	req := &mobile_device_applications.Resource{
		General: mobile_device_applications.SubsetGeneral{
			Name:     "Duplicate App",
			BundleID: "com.dup.app",
			Site:     &shared.SharedResourceSite{ID: -1, Name: "None"},
		},
		Scope:       mobile_device_applications.SubsetScope{AllMobileDevices: boolPtr(true), AllJSSUsers: boolPtr(false)},
		SelfService: mobile_device_applications.SubsetSelfService{},
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device application with that name already exists")
}

func boolPtr(b bool) *bool {
	return &b
}
