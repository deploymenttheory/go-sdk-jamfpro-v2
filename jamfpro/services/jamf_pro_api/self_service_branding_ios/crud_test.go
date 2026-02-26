package self_service_branding_ios

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/self_service_branding_ios/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.SelfServiceBrandingMobileMock) {
	t.Helper()
	mock := mocks.NewSelfServiceBrandingMobileMock()
	return NewService(mock), mock
}

func TestUnit_SelfServiceBrandingMobile_ListV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Corporate Branding", result.Results[0].BrandingName)
	assert.Equal(t, "#FFFFFF", result.Results[0].HeaderBackgroundColorCode)
}

func TestUnit_SelfServiceBrandingMobile_GetByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Corporate Branding", result.BrandingName)
	assert.Equal(t, "#FFFFFF", result.HeaderBackgroundColorCode)
	require.NotNil(t, result.IconId)
	assert.Equal(t, 5, *result.IconId)
}

func TestUnit_SelfServiceBrandingMobile_GetByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "self-service branding mobile ID is required")
}

func TestUnit_SelfServiceBrandingMobile_GetByIDV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestUnit_SelfServiceBrandingMobile_GetByNameV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.GetByNameV1(context.Background(), "Corporate Branding")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Corporate Branding", result.BrandingName)
}

func TestUnit_SelfServiceBrandingMobile_GetByNameV1_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByNameV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "self-service branding mobile name is required")
}

func TestUnit_SelfServiceBrandingMobile_GetByNameV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.GetByNameV1(context.Background(), "NonExistent")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "was not found")
}

func TestUnit_SelfServiceBrandingMobile_CreateV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	req := &ResourceSelfServiceBrandingMobile{
		BrandingName:              "New Branding",
		HeaderBackgroundColorCode: "#FFFFFF",
		MenuIconColorCode:         "#000000",
		BrandingNameColorCode:     "#333333",
		StatusBarTextColor:        "light",
	}
	result, resp, err := svc.CreateV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "3", result.ID)
	assert.Contains(t, result.Href, "/api/v1/self-service/branding/ios/3")
}

func TestUnit_SelfServiceBrandingMobile_CreateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_SelfServiceBrandingMobile_CreateV1_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &ResourceSelfServiceBrandingMobile{
		BrandingName:              "Duplicate",
		HeaderBackgroundColorCode: "#FFFFFF",
		MenuIconColorCode:         "#000000",
		BrandingNameColorCode:     "#333333",
		StatusBarTextColor:        "light",
	}
	result, resp, err := svc.CreateV1(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

func TestUnit_SelfServiceBrandingMobile_UpdateByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	req := &ResourceSelfServiceBrandingMobile{
		BrandingName:              "Corporate Branding Updated",
		IconId:                    intPtr(6),
		HeaderBackgroundColorCode: "#F0F0F0",
		MenuIconColorCode:         "#0066CC",
		BrandingNameColorCode:     "#222222",
		StatusBarTextColor:        "dark",
	}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Corporate Branding Updated", result.BrandingName)
}

func TestUnit_SelfServiceBrandingMobile_UpdateByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceSelfServiceBrandingMobile{
		BrandingName:              "x",
		HeaderBackgroundColorCode: "#FFF",
		MenuIconColorCode:         "#000",
		BrandingNameColorCode:     "#333",
		StatusBarTextColor:        "light",
	}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_SelfServiceBrandingMobile_UpdateByIDV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_SelfServiceBrandingMobile_DeleteByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteMock()

	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnit_SelfServiceBrandingMobile_DeleteByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "self-service branding mobile ID is required")
}

func TestUnit_SelfServiceBrandingMobile_UpdateByNameV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()
	mock.RegisterUpdateMock()

	req := &ResourceSelfServiceBrandingMobile{
		BrandingName:              "Corporate Branding Updated",
		HeaderBackgroundColorCode: "#F0F0F0",
		MenuIconColorCode:         "#0066CC",
		BrandingNameColorCode:     "#222222",
		StatusBarTextColor:        "dark",
	}
	result, resp, err := svc.UpdateByNameV1(context.Background(), "Corporate Branding", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_SelfServiceBrandingMobile_DeleteByNameV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()
	mock.RegisterDeleteMock()

	resp, err := svc.DeleteByNameV1(context.Background(), "Corporate Branding")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func intPtr(i int) *int {
	return &i
}
