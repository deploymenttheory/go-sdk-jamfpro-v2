package self_service_branding_macos

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/self_service_branding_macos/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.SelfServiceBrandingMacOSMock) {
	t.Helper()
	mock := mocks.NewSelfServiceBrandingMacOSMock()
	return NewService(mock), mock
}

func TestUnitList_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.List(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	assert.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Corporate Branding", result.Results[0].BrandingName)
	assert.Equal(t, "2", result.Results[1].ID)
	assert.Equal(t, "Test Branding", result.Results[1].BrandingName)
}

func TestUnitGetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDMock()

	result, resp, err := svc.GetByID(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Corporate Branding", result.BrandingName)
	assert.Equal(t, "Self Service", result.ApplicationName)
	assert.Equal(t, "Welcome", result.HomeHeading)
	require.NotNil(t, result.IconId)
	assert.Equal(t, 1, *result.IconId)
}

func TestUnitGetByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "self-service branding configuration ID is required")
}

func TestUnitGetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByID(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestUnitGetByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.GetByName(context.Background(), "Corporate Branding")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Corporate Branding", result.BrandingName)
}

func TestUnitGetByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "self-service branding configuration name is required")
}

func TestUnitGetByName_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.GetByName(context.Background(), "NonExistent")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "not found")
}

func TestUnitCreate_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	req := &ResourceSelfServiceBrandingMacOS{
		ApplicationName:       "Self Service",
		BrandingName:          "New Branding",
		BrandingNameSecondary: "Created via API",
		HomeHeading:           "Welcome",
		HomeSubheading:        "Select an option",
	}
	result, resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "3", result.ID)
	assert.Equal(t, "New Branding", result.BrandingName)
}

func TestUnitCreate_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.Create(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreate_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &ResourceSelfServiceBrandingMacOS{
		BrandingName: "Duplicate",
	}
	result, resp, err := svc.Create(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

func TestUnitUpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	req := &ResourceSelfServiceBrandingMacOS{
		ApplicationName:       "Self Service",
		BrandingName:          "Corporate Branding Updated",
		BrandingNameSecondary: "IT Department",
		IconId:                intPtr(1),
		BrandingHeaderImageId: intPtr(2),
		HomeHeading:           "Welcome Back",
		HomeSubheading:        "Choose an item below",
	}
	result, resp, err := svc.UpdateByID(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Corporate Branding Updated", result.BrandingName)
	assert.Equal(t, "Welcome Back", result.HomeHeading)
}

func TestUnitUpdateByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByID(context.Background(), "", &ResourceSelfServiceBrandingMacOS{
		BrandingName: "x",
	})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnitUpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByID(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitDeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteMock()

	resp, err := svc.DeleteByID(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitDeleteByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "self-service branding configuration ID is required")
}

func TestUnitUpdateByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()
	mock.RegisterUpdateMock()

	req := &ResourceSelfServiceBrandingMacOS{
		BrandingName: "Corporate Branding Updated",
	}
	result, resp, err := svc.UpdateByName(context.Background(), "Corporate Branding", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
}

func TestUnitUpdateByName_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	req := &ResourceSelfServiceBrandingMacOS{BrandingName: "Updated"}
	result, resp, err := svc.UpdateByName(context.Background(), "NonExistent", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnitDeleteByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()
	mock.RegisterDeleteMock()

	resp, err := svc.DeleteByName(context.Background(), "Corporate Branding")
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestUnitDeleteByName_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	resp, err := svc.DeleteByName(context.Background(), "NonExistent")
	assert.Error(t, err)
	_ = resp
}

func TestUnitDeleteByID_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), "999")
	assert.Error(t, err)
	_ = resp
}

func TestUnitList_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.List(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func intPtr(i int) *int {
	return &i
}
