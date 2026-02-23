package computers_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/computers"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/computers/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_Computers_List(t *testing.T) {
	mockClient := mocks.NewComputersMock()
	mockClient.RegisterListComputersMock()
	svc := computers.NewService(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 2, resp.Size)
	assert.Len(t, resp.Results, 2)
	assert.Equal(t, "MacBook-Pro-01", resp.Results[0].Name)
	assert.Equal(t, 1, resp.Results[0].ID)
	assert.Equal(t, "MacBook-Pro-02", resp.Results[1].Name)
	assert.Equal(t, 2, resp.Results[1].ID)
}

func TestUnit_Computers_GetByID(t *testing.T) {
	mockClient := mocks.NewComputersMock()
	mockClient.RegisterGetComputerByIDMock()
	svc := computers.NewService(mockClient)

	resp, _, err := svc.GetByID(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "MacBook-Pro-01", resp.General.Name)
	assert.Equal(t, "00:11:22:33:44:55", resp.General.MacAddress)
	assert.Equal(t, "C02XYZ123456", resp.General.SerialNumber)
	assert.NotNil(t, resp.General.Site)
	assert.Equal(t, -1, resp.General.Site.ID)
	assert.Equal(t, "None", resp.General.Site.Name)
	assert.Equal(t, "Apple Inc.", resp.Hardware.Make)
	assert.Equal(t, "MacBook Pro", resp.Hardware.Model)
}

func TestUnit_Computers_GetByID_EmptyID(t *testing.T) {
	mockClient := mocks.NewComputersMock()
	svc := computers.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "computer ID cannot be empty")
}

func TestUnit_Computers_GetByName(t *testing.T) {
	mockClient := mocks.NewComputersMock()
	mockClient.RegisterGetComputerByNameMock()
	svc := computers.NewService(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "MacBook-Pro-01")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "MacBook-Pro-01", resp.General.Name)
}

func TestUnit_Computers_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewComputersMock()
	svc := computers.NewService(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "computer name cannot be empty")
}

func TestUnit_Computers_Create(t *testing.T) {
	mockClient := mocks.NewComputersMock()
	mockClient.RegisterCreateComputerMock()
	svc := computers.NewService(mockClient)

	computer := &computers.ResponseComputer{
		General: computers.ComputerSubsetGeneral{
			Name:         "test-computer-01",
			MacAddress:   "00:11:22:33:44:55",
			SerialNumber: "C02XYZ123456",
			Site: shared.SharedResourceSite{
				ID:   -1,
				Name: "none",
			},
		},
	}

	resp, _, err := svc.Create(context.Background(), computer)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.General.ID)
	assert.Equal(t, "test-computer-01", resp.General.Name)
}

func TestUnit_Computers_Create_NilComputer(t *testing.T) {
	mockClient := mocks.NewComputersMock()
	svc := computers.NewService(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "computer is required")
}

func TestUnit_Computers_Create_EmptyName(t *testing.T) {
	mockClient := mocks.NewComputersMock()
	svc := computers.NewService(mockClient)

	computer := &computers.ResponseComputer{
		General: computers.ComputerSubsetGeneral{
			Name: "",
		},
	}

	_, _, err := svc.Create(context.Background(), computer)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "computer name is required")
}

func TestUnit_Computers_UpdateByID(t *testing.T) {
	mockClient := mocks.NewComputersMock()
	mockClient.RegisterUpdateComputerByIDMock()
	svc := computers.NewService(mockClient)

	computer := &computers.ResponseComputer{
		General: computers.ComputerSubsetGeneral{
			ID:           1,
			Name:         "MacBook-Pro-01-Updated",
			MacAddress:   "00:11:22:33:44:55",
			SerialNumber: "C02XYZ123456",
		},
	}

	resp, _, err := svc.UpdateByID(context.Background(), "1", computer)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "MacBook-Pro-01-Updated", resp.General.Name)
}

func TestUnit_Computers_UpdateByID_SiteDefault(t *testing.T) {
	mockClient := mocks.NewComputersMock()
	mockClient.RegisterUpdateComputerByIDMock()
	svc := computers.NewService(mockClient)

	computer := &computers.ResponseComputer{
		General: computers.ComputerSubsetGeneral{
			ID:           1,
			Name:         "MacBook-Pro-01-Updated",
			MacAddress:   "00:11:22:33:44:55",
			SerialNumber: "C02XYZ123456",
			Site:         shared.SharedResourceSite{ID: 0, Name: ""},
		},
	}

	resp, _, err := svc.UpdateByID(context.Background(), "1", computer)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, -1, computer.General.Site.ID)
	assert.Equal(t, "none", computer.General.Site.Name)
}

func TestUnit_Computers_UpdateByID_EmptyID(t *testing.T) {
	mockClient := mocks.NewComputersMock()
	svc := computers.NewService(mockClient)

	computer := &computers.ResponseComputer{
		General: computers.ComputerSubsetGeneral{Name: "test"},
	}

	_, _, err := svc.UpdateByID(context.Background(), "", computer)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "computer ID cannot be empty")
}

func TestUnit_Computers_UpdateByName(t *testing.T) {
	mockClient := mocks.NewComputersMock()
	mockClient.RegisterUpdateComputerByNameMock()
	svc := computers.NewService(mockClient)

	computer := &computers.ResponseComputer{
		General: computers.ComputerSubsetGeneral{
			ID:           1,
			Name:         "MacBook-Pro-01-Updated",
			MacAddress:   "00:11:22:33:44:55",
			SerialNumber: "C02XYZ123456",
		},
	}

	resp, _, err := svc.UpdateByName(context.Background(), "MacBook-Pro-01", computer)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "MacBook-Pro-01-Updated", resp.General.Name)
}

func TestUnit_Computers_UpdateByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewComputersMock()
	svc := computers.NewService(mockClient)

	computer := &computers.ResponseComputer{
		General: computers.ComputerSubsetGeneral{Name: "test"},
	}

	_, _, err := svc.UpdateByName(context.Background(), "", computer)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "computer name cannot be empty")
}

func TestUnit_Computers_DeleteByID(t *testing.T) {
	mockClient := mocks.NewComputersMock()
	mockClient.RegisterDeleteComputerByIDMock()
	svc := computers.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), "1")

	require.NoError(t, err)
}

func TestUnit_Computers_DeleteByID_EmptyID(t *testing.T) {
	mockClient := mocks.NewComputersMock()
	svc := computers.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "computer ID cannot be empty")
}

func TestUnit_Computers_DeleteByName(t *testing.T) {
	mockClient := mocks.NewComputersMock()
	mockClient.RegisterDeleteComputerByNameMock()
	svc := computers.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "MacBook-Pro-01")

	require.NoError(t, err)
}

func TestUnit_Computers_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewComputersMock()
	svc := computers.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "computer name cannot be empty")
}

func TestUnit_Computers_NotFound(t *testing.T) {
	mockClient := mocks.NewComputersMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := computers.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), "999")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}
