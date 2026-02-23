package mac_applications_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mac_applications"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mac_applications/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_MacApplications_List(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	mockClient.RegisterListMacApplicationsMock()
	svc := mac_applications.NewService(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Len(t, resp.Results, 2)
	assert.Equal(t, "Sample Mac App 1", resp.Results[0].Name)
	assert.Equal(t, 1, resp.Results[0].ID)
	assert.Equal(t, "Sample Mac App 2", resp.Results[1].Name)
	assert.Equal(t, 2, resp.Results[1].ID)
}

func TestUnit_MacApplications_GetByID(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	mockClient.RegisterGetMacApplicationByIDMock()
	svc := mac_applications.NewService(mockClient)

	resp, _, err := svc.GetByID(context.Background(), 1)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Sample Mac App", resp.General.Name)
	assert.Equal(t, "1.0", resp.General.Version)
	assert.Equal(t, "com.example.app", resp.General.BundleID)
	assert.NotNil(t, resp.General.Site)
	assert.Equal(t, -1, resp.General.Site.ID)
}

func TestUnit_MacApplications_GetByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	svc := mac_applications.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mac application ID must be a positive integer")
}

func TestUnit_MacApplications_GetByName(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	mockClient.RegisterGetMacApplicationByNameMock()
	svc := mac_applications.NewService(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "Sample Mac App")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Sample Mac App", resp.General.Name)
}

func TestUnit_MacApplications_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	svc := mac_applications.NewService(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mac application name cannot be empty")
}

func TestUnit_MacApplications_GetByIDAndSubset(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	mockClient.RegisterGetMacApplicationByIDAndSubsetMock()
	svc := mac_applications.NewService(mockClient)

	resp, _, err := svc.GetByIDAndSubset(context.Background(), 1, "General")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Sample Mac App", resp.General.Name)
}

func TestUnit_MacApplications_GetByIDAndSubset_ZeroID(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	svc := mac_applications.NewService(mockClient)

	_, _, err := svc.GetByIDAndSubset(context.Background(), 0, "General")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mac application ID must be a positive integer")
}

func TestUnit_MacApplications_GetByIDAndSubset_EmptySubset(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	svc := mac_applications.NewService(mockClient)

	_, _, err := svc.GetByIDAndSubset(context.Background(), 1, "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mac application subset cannot be empty")
}

func TestUnit_MacApplications_GetByNameAndSubset(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	mockClient.RegisterGetMacApplicationByNameAndSubsetMock()
	svc := mac_applications.NewService(mockClient)

	resp, _, err := svc.GetByNameAndSubset(context.Background(), "Sample Mac App", "General")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Sample Mac App", resp.General.Name)
}

func TestUnit_MacApplications_GetByNameAndSubset_EmptyName(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	svc := mac_applications.NewService(mockClient)

	_, _, err := svc.GetByNameAndSubset(context.Background(), "", "General")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mac application name cannot be empty")
}

func TestUnit_MacApplications_Create(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	mockClient.RegisterCreateMacApplicationMock()
	svc := mac_applications.NewService(mockClient)

	req := &mac_applications.Resource{
		General: mac_applications.SubsetGeneral{
			Name:           "Test Mac App",
			Version:        "1.0",
			BundleID:       "com.test.app",
			URL:            "https://example.com",
			DeploymentType: "Install Automatically/Prompt Users to Install",
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: mac_applications.SubsetScope{
			AllComputers: boolPtr(true),
			AllJSSUsers:  boolPtr(false),
		},
		SelfService: mac_applications.SubsetSelfService{
			InstallButtonText:      "Install",
			SelfServiceDescription: "Test app",
		},
	}

	resp, _, err := svc.Create(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.ID)
}

func TestUnit_MacApplications_Create_NilRequest(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	svc := mac_applications.NewService(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MacApplications_Create_EmptyName(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	svc := mac_applications.NewService(mockClient)

	req := &mac_applications.Resource{
		General: mac_applications.SubsetGeneral{
			Name: "",
		},
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mac application name is required")
}

func TestUnit_MacApplications_UpdateByID(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	mockClient.RegisterUpdateMacApplicationByIDMock()
	svc := mac_applications.NewService(mockClient)

	req := &mac_applications.Resource{
		General: mac_applications.SubsetGeneral{
			Name:           "Updated Mac App",
			Version:        "1.1",
			BundleID:       "com.example.app",
			URL:            "https://example.com",
			DeploymentType: "Install Automatically/Prompt Users to Install",
		},
		Scope: mac_applications.SubsetScope{
			AllComputers: boolPtr(true),
			AllJSSUsers:  boolPtr(false),
		},
		SelfService: mac_applications.SubsetSelfService{
			InstallButtonText: "Install",
		},
	}

	resp, _, err := svc.UpdateByID(context.Background(), 1, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Updated Mac App", resp.General.Name)
}

func TestUnit_MacApplications_UpdateByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	svc := mac_applications.NewService(mockClient)

	req := &mac_applications.Resource{
		General: mac_applications.SubsetGeneral{Name: "Test"},
	}

	_, _, err := svc.UpdateByID(context.Background(), 0, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mac application ID must be a positive integer")
}

func TestUnit_MacApplications_UpdateByName(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	mockClient.RegisterUpdateMacApplicationByNameMock()
	svc := mac_applications.NewService(mockClient)

	req := &mac_applications.Resource{
		General: mac_applications.SubsetGeneral{
			Name:           "Updated Mac App",
			Version:        "1.1",
			BundleID:       "com.example.app",
			URL:            "https://example.com",
			DeploymentType: "Install Automatically/Prompt Users to Install",
		},
		Scope: mac_applications.SubsetScope{
			AllComputers: boolPtr(true),
			AllJSSUsers:  boolPtr(false),
		},
		SelfService: mac_applications.SubsetSelfService{
			InstallButtonText: "Install",
		},
	}

	resp, _, err := svc.UpdateByName(context.Background(), "Sample Mac App", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Updated Mac App", resp.General.Name)
}

func TestUnit_MacApplications_UpdateByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	svc := mac_applications.NewService(mockClient)

	req := &mac_applications.Resource{
		General: mac_applications.SubsetGeneral{Name: "Test"},
	}

	_, _, err := svc.UpdateByName(context.Background(), "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mac application name cannot be empty")
}

func TestUnit_MacApplications_DeleteByID(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	mockClient.RegisterDeleteMacApplicationByIDMock()
	svc := mac_applications.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 1)

	require.NoError(t, err)
}

func TestUnit_MacApplications_DeleteByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	svc := mac_applications.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mac application ID must be a positive integer")
}

func TestUnit_MacApplications_DeleteByName(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	mockClient.RegisterDeleteMacApplicationByNameMock()
	svc := mac_applications.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "Sample Mac App")

	require.NoError(t, err)
}

func TestUnit_MacApplications_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	svc := mac_applications.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mac application name cannot be empty")
}

func TestUnit_MacApplications_NotFound(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := mac_applications.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 999)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}

func TestUnit_MacApplications_Conflict(t *testing.T) {
	mockClient := mocks.NewMacApplicationsMock()
	mockClient.RegisterConflictErrorMock()
	svc := mac_applications.NewService(mockClient)

	req := &mac_applications.Resource{
		General: mac_applications.SubsetGeneral{
			Name:           "Duplicate App",
			BundleID:       "com.dup.app",
			DeploymentType: "Install Automatically/Prompt Users to Install",
		},
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Mac application with that name already exists")
}

func boolPtr(b bool) *bool {
	return &b
}
