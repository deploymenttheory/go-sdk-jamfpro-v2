package licensed_software_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/licensed_software"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/licensed_software/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_LicensedSoftware_List(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	mockClient.RegisterListLicensedSoftwareMock()
	svc := licensed_software.NewService(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 2, resp.Size)
	assert.Len(t, resp.Results, 2)
	assert.Equal(t, "Sample Licensed Software 1", resp.Results[0].Name)
	assert.Equal(t, 1, resp.Results[0].ID)
	assert.Equal(t, "Sample Licensed Software 2", resp.Results[1].Name)
	assert.Equal(t, 2, resp.Results[1].ID)
}

func TestUnit_LicensedSoftware_GetByID(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	mockClient.RegisterGetLicensedSoftwareByIDMock()
	svc := licensed_software.NewService(mockClient)

	resp, _, err := svc.GetByID(context.Background(), 1)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Sample Licensed Software", resp.General.Name)
	assert.Equal(t, "Test Publisher", resp.General.Publisher)
	assert.Equal(t, "Mac", resp.General.Platform)
	assert.False(t, resp.General.SendEmailOnViolation)
	assert.Equal(t, "Test notes", resp.General.Notes)
	assert.Equal(t, -1, resp.General.Site.ID)
	assert.Equal(t, "None", resp.General.Site.Name)
}

func TestUnit_LicensedSoftware_GetByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	svc := licensed_software.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "licensed software ID must be a positive integer")
}

func TestUnit_LicensedSoftware_GetByName(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	mockClient.RegisterGetLicensedSoftwareByNameMock()
	svc := licensed_software.NewService(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "Sample Licensed Software")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Sample Licensed Software", resp.General.Name)
}

func TestUnit_LicensedSoftware_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	svc := licensed_software.NewService(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "licensed software name cannot be empty")
}

func TestUnit_LicensedSoftware_Create(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	mockClient.RegisterCreateLicensedSoftwareMock()
	svc := licensed_software.NewService(mockClient)

	req := &licensed_software.Resource{
		General: licensed_software.SubsetGeneral{
			Name:      "Test Licensed Software",
			Publisher: "Test Publisher",
			Platform:  "Mac",
			Site:      shared.SharedResourceSite{ID: -1, Name: "None"},
		},
	}

	resp, _, err := svc.Create(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.ID)
}

func TestUnit_LicensedSoftware_Create_NilRequest(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	svc := licensed_software.NewService(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_LicensedSoftware_Create_EmptyName(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	svc := licensed_software.NewService(mockClient)

	req := &licensed_software.Resource{
		General: licensed_software.SubsetGeneral{
			Name: "",
		},
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "licensed software name is required")
}

func TestUnit_LicensedSoftware_UpdateByID(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	mockClient.RegisterUpdateLicensedSoftwareByIDMock()
	svc := licensed_software.NewService(mockClient)

	req := &licensed_software.Resource{
		General: licensed_software.SubsetGeneral{
			Name:      "Updated Licensed Software",
			Publisher: "Test Publisher",
			Platform:  "Mac",
		},
	}

	resp, _, err := svc.UpdateByID(context.Background(), 1, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_LicensedSoftware_UpdateByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	svc := licensed_software.NewService(mockClient)

	req := &licensed_software.Resource{
		General: licensed_software.SubsetGeneral{Name: "Test"},
	}

	_, _, err := svc.UpdateByID(context.Background(), 0, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "licensed software ID must be a positive integer")
}

func TestUnit_LicensedSoftware_UpdateByName(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	mockClient.RegisterUpdateLicensedSoftwareByNameMock()
	svc := licensed_software.NewService(mockClient)

	req := &licensed_software.Resource{
		General: licensed_software.SubsetGeneral{
			Name:      "Updated Licensed Software",
			Publisher: "Test Publisher",
			Platform:  "Mac",
		},
	}

	resp, _, err := svc.UpdateByName(context.Background(), "Sample Licensed Software", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_LicensedSoftware_UpdateByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	svc := licensed_software.NewService(mockClient)

	req := &licensed_software.Resource{
		General: licensed_software.SubsetGeneral{Name: "Test"},
	}

	_, _, err := svc.UpdateByName(context.Background(), "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "licensed software name cannot be empty")
}

func TestUnit_LicensedSoftware_DeleteByID(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	mockClient.RegisterDeleteLicensedSoftwareByIDMock()
	svc := licensed_software.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 1)

	require.NoError(t, err)
}

func TestUnit_LicensedSoftware_DeleteByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	svc := licensed_software.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "licensed software ID must be a positive integer")
}

func TestUnit_LicensedSoftware_DeleteByName(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	mockClient.RegisterDeleteLicensedSoftwareByNameMock()
	svc := licensed_software.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "Sample Licensed Software")

	require.NoError(t, err)
}

func TestUnit_LicensedSoftware_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	svc := licensed_software.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "licensed software name cannot be empty")
}

func TestUnit_LicensedSoftware_NotFound(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := licensed_software.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 999)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}

func TestUnit_LicensedSoftware_Conflict(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	mockClient.RegisterConflictErrorMock()
	svc := licensed_software.NewService(mockClient)

	req := &licensed_software.Resource{
		General: licensed_software.SubsetGeneral{
			Name:      "Duplicate Licensed Software",
			Publisher: "Test Publisher",
			Platform:  "Mac",
		},
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Licensed software with that name already exists")
}

func TestUnit_LicensedSoftware_List_Error(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	svc := licensed_software.NewService(mockClient)
	_, _, err := svc.List(context.Background())
	require.Error(t, err)
}

func TestUnit_LicensedSoftware_GetByName_Error(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	svc := licensed_software.NewService(mockClient)
	_, _, err := svc.GetByName(context.Background(), "Sample Licensed Software 1")
	require.Error(t, err)
}

func TestUnit_LicensedSoftware_UpdateByID_NilRequest(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	svc := licensed_software.NewService(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_LicensedSoftware_UpdateByID_EmptyName(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	svc := licensed_software.NewService(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), 1, &licensed_software.Resource{General: licensed_software.SubsetGeneral{Name: ""}})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "licensed software name is required")
}

func TestUnit_LicensedSoftware_UpdateByID_Error(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	svc := licensed_software.NewService(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), 1, &licensed_software.Resource{General: licensed_software.SubsetGeneral{Name: "Test"}})
	require.Error(t, err)
}

func TestUnit_LicensedSoftware_UpdateByName_NilRequest(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	svc := licensed_software.NewService(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "Sample Licensed Software 1", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_LicensedSoftware_UpdateByName_EmptyReqName(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	svc := licensed_software.NewService(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "Sample Licensed Software 1", &licensed_software.Resource{General: licensed_software.SubsetGeneral{Name: ""}})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "licensed software name is required in request")
}

func TestUnit_LicensedSoftware_UpdateByName_Error(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	svc := licensed_software.NewService(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "Sample Licensed Software 1", &licensed_software.Resource{General: licensed_software.SubsetGeneral{Name: "Updated"}})
	require.Error(t, err)
}

func TestUnit_LicensedSoftware_DeleteByID_Error(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	svc := licensed_software.NewService(mockClient)
	_, err := svc.DeleteByID(context.Background(), 1)
	require.Error(t, err)
}

func TestUnit_LicensedSoftware_DeleteByName_Error(t *testing.T) {
	mockClient := mocks.NewLicensedSoftwareMock()
	svc := licensed_software.NewService(mockClient)
	_, err := svc.DeleteByName(context.Background(), "Sample Licensed Software 1")
	require.Error(t, err)
}
