package mobile_device_extension_attributes

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mobile_device_extension_attributes/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.MobileDeviceExtensionAttributesMock) {
	t.Helper()
	mock := mocks.NewMobileDeviceExtensionAttributesMock()
	return NewService(mock), mock
}

func TestUnitListMobileDeviceExtensionAttributes_Success(t *testing.T) {
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
	assert.Equal(t, "MDEA One", result.Results[0].Name)
}

func TestUnitGetMobileDeviceExtensionAttributeByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "MDEA One", result.Name)
	assert.Equal(t, "String", result.DataType)
}

func TestUnitGetMobileDeviceExtensionAttributeByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitGetMobileDeviceExtensionAttributeByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestUnitCreateMobileDeviceExtensionAttribute_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	req := &RequestMobileDeviceExtensionAttribute{
		Name:                 "New MDEA",
		DataType:             "String",
		InventoryDisplayType: "General",
		InputType:            "Text Field",
	}
	result, resp, err := svc.CreateV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "3", result.ID)
	assert.NotEmpty(t, result.Href)
}

func TestUnitCreateMobileDeviceExtensionAttribute_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitUpdateMobileDeviceExtensionAttributeByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	req := &RequestMobileDeviceExtensionAttribute{
		Name:                 "MDEA One Updated",
		Description:          "Updated",
		DataType:             "String",
		InventoryDisplayType: "General",
		InputType:            "Text Field",
	}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "MDEA One Updated", result.Name)
}

func TestUnitDeleteMobileDeviceExtensionAttributeByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteMock()

	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitDeleteMobileDeviceExtensionAttributesByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteMultipleMock()

	req := &DeleteMobileDeviceExtensionAttributesByIDRequest{IDs: []string{"1", "2"}}
	resp, err := svc.DeleteMobileDeviceExtensionAttributesByIDV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitDeleteMobileDeviceExtensionAttributesByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteMobileDeviceExtensionAttributesByIDV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
}
