package mobile_device_extension_attributes

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/mobile_device_extension_attributes/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*MobileDeviceExtensionAttributes, *mocks.MobileDeviceExtensionAttributesMock) {
	t.Helper()
	mock := mocks.NewMobileDeviceExtensionAttributesMock()
	return NewMobileDeviceExtensionAttributes(mock), mock
}

func TestUnit_MobileDeviceExtensionAttributes_List_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "MDEA One", result.Results[0].Name)
}

func TestUnit_MobileDeviceExtensionAttributes_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "MDEA One", result.Name)
	assert.Equal(t, "String", result.DataType)
}

func TestUnit_MobileDeviceExtensionAttributes_GetByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_MobileDeviceExtensionAttributes_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_MobileDeviceExtensionAttributes_Create_Success(t *testing.T) {
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

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "3", result.ID)
	assert.NotEmpty(t, result.Href)
}

func TestUnit_MobileDeviceExtensionAttributes_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_MobileDeviceExtensionAttributes_UpdateByID_Success(t *testing.T) {
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

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "MDEA One Updated", result.Name)
}

func TestUnit_MobileDeviceExtensionAttributes_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteMock()

	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_MobileDeviceExtensionAttributes_DeleteMultipleByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteMultipleMock()

	req := &DeleteMobileDeviceExtensionAttributesByIDRequest{IDs: []string{"1", "2"}}
	resp, err := svc.DeleteMobileDeviceExtensionAttributesByIDV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_MobileDeviceExtensionAttributes_DeleteMultipleByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteMobileDeviceExtensionAttributesByIDV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_MobileDeviceExtensionAttributes_GetHistoryByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryMock()

	result, resp, err := svc.GetHistoryByIDV1(context.Background(), "1", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "admin", result.Results[0].Username)
}

func TestUnit_MobileDeviceExtensionAttributes_GetHistoryByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetHistoryByIDV1(context.Background(), "", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_MobileDeviceExtensionAttributes_AddHistoryNoteByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddHistoryNoteMock()

	req := &AddHistoryNoteRequest{Note: "Test note"}
	resp, err := svc.AddHistoryNoteByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
}

func TestUnit_MobileDeviceExtensionAttributes_AddHistoryNoteByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &AddHistoryNoteRequest{Note: "Test note"}
	resp, err := svc.AddHistoryNoteByIDV1(context.Background(), "", req)
	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_MobileDeviceExtensionAttributes_AddHistoryNoteByIDV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.AddHistoryNoteByIDV1(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_MobileDeviceExtensionAttributes_ListV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.ListV1(context.Background(), nil)
	require.Error(t, err)
}

func TestUnit_MobileDeviceExtensionAttributes_CreateV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestMobileDeviceExtensionAttribute{
		Name:                 "New MDEA",
		DataType:             "String",
		InventoryDisplayType: "General",
		InputType:            "Text Field",
	}
	_, _, err := svc.CreateV1(context.Background(), req)
	require.Error(t, err)
}

func TestUnit_MobileDeviceExtensionAttributes_UpdateByIDV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestMobileDeviceExtensionAttribute{
		Name:                 "MDEA One Updated",
		DataType:             "String",
		InventoryDisplayType: "General",
		InputType:            "Text Field",
	}
	_, _, err := svc.UpdateByIDV1(context.Background(), "1", req)
	require.Error(t, err)
}

func TestUnit_MobileDeviceExtensionAttributes_DeleteByIDV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.DeleteByIDV1(context.Background(), "1")
	require.Error(t, err)
}

func TestUnit_MobileDeviceExtensionAttributes_DeleteMobileDeviceExtensionAttributesByIDV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &DeleteMobileDeviceExtensionAttributesByIDRequest{IDs: []string{"1", "2"}}
	_, err := svc.DeleteMobileDeviceExtensionAttributesByIDV1(context.Background(), req)
	require.Error(t, err)
}

func TestUnit_MobileDeviceExtensionAttributes_GetHistoryByIDV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetHistoryByIDV1(context.Background(), "1", nil)
	require.Error(t, err)
}

func TestUnit_MobileDeviceExtensionAttributes_AddHistoryNoteByIDV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &AddHistoryNoteRequest{Note: "Test note"}
	_, err := svc.AddHistoryNoteByIDV1(context.Background(), "1", req)
	require.Error(t, err)
}

func TestUnit_MobileDeviceExtensionAttributes_GetDataDependencyByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetDataDependencyByIDV1(context.Background(), "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device extension attribute ID is required")
}

func TestUnit_MobileDeviceExtensionAttributes_GetDataDependencyByIDV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetDataDependencyByIDV1(context.Background(), "1")
	require.Error(t, err)
}
