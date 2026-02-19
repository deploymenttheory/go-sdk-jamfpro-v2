package computer_extension_attributes

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_extension_attributes/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.ComputerExtensionAttributesMock) {
	t.Helper()
	mock := mocks.NewComputerExtensionAttributesMock()
	return NewService(mock), mock
}

func TestUnitListComputerExtensionAttributes_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.ListComputerExtensionAttributesV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "EA One", result.Results[0].Name)
	assert.Equal(t, "2", result.Results[1].ID)
	assert.Equal(t, "EA Two", result.Results[1].Name)
}

func TestUnitListComputerExtensionAttributes_WithQueryParams(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	params := map[string]string{"page": "0", "page-size": "50"}
	result, resp, err := svc.ListComputerExtensionAttributesV1(context.Background(), params)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitGetComputerExtensionAttributeByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.GetComputerExtensionAttributeByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "EA One", result.Name)
	assert.Equal(t, "String", result.DataType)
	assert.NotNil(t, result.Enabled)
	assert.True(t, *result.Enabled)
}

func TestUnitGetComputerExtensionAttributeByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetComputerExtensionAttributeByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitGetComputerExtensionAttributeByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetComputerExtensionAttributeByIDV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestUnitCreateComputerExtensionAttribute_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	req := &RequestComputerExtensionAttribute{
		Name:                 "New EA",
		DataType:             "String",
		InventoryDisplayType: "General",
		InputType:            "Text Field",
	}
	result, resp, err := svc.CreateComputerExtensionAttributeV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "3", result.ID)
	assert.NotEmpty(t, result.Href)
}

func TestUnitCreateComputerExtensionAttribute_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateComputerExtensionAttributeV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitUpdateComputerExtensionAttributeByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	req := &RequestComputerExtensionAttribute{
		Name:                 "EA One Updated",
		Description:          "Updated",
		DataType:             "String",
		InventoryDisplayType: "General",
		InputType:            "Text Field",
	}
	result, resp, err := svc.UpdateComputerExtensionAttributeByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "EA One Updated", result.Name)
	assert.Equal(t, "Updated", result.Description)
}

func TestUnitUpdateComputerExtensionAttributeByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestComputerExtensionAttribute{Name: "x", DataType: "String", InventoryDisplayType: "General", InputType: "Text Field"}

	result, resp, err := svc.UpdateComputerExtensionAttributeByIDV1(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitUpdateComputerExtensionAttributeByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateComputerExtensionAttributeByIDV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitDeleteComputerExtensionAttributeByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteMock()

	resp, err := svc.DeleteComputerExtensionAttributeByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitDeleteComputerExtensionAttributeByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteComputerExtensionAttributeByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnitDeleteComputerExtensionAttributesByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteMultipleMock()

	req := &DeleteComputerExtensionAttributesByIDRequest{IDs: []string{"1", "2"}}
	resp, err := svc.DeleteComputerExtensionAttributesByIDV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitDeleteComputerExtensionAttributesByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteComputerExtensionAttributesByIDV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnitDeleteComputerExtensionAttributesByID_EmptyIDs(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteComputerExtensionAttributesByIDV1(context.Background(), &DeleteComputerExtensionAttributesByIDRequest{IDs: nil})
	assert.Error(t, err)
	assert.Nil(t, resp)
}
