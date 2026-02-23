package user_extension_attributes_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/user_extension_attributes"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/user_extension_attributes/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_UserExtensionAttributes_List(t *testing.T) {
	mockClient := mocks.NewUserExtensionAttributesMock()
	mockClient.RegisterListUserExtensionAttributesMock()
	svc := user_extension_attributes.NewService(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 2, resp.Size)
	assert.Len(t, resp.UserExtensionAttributes, 2)
	assert.Equal(t, "Department", resp.UserExtensionAttributes[0].Name)
	assert.Equal(t, 1, resp.UserExtensionAttributes[0].ID)
	assert.Equal(t, "Employee ID", resp.UserExtensionAttributes[1].Name)
	assert.Equal(t, 2, resp.UserExtensionAttributes[1].ID)
}

func TestUnit_UserExtensionAttributes_GetByID(t *testing.T) {
	mockClient := mocks.NewUserExtensionAttributesMock()
	mockClient.RegisterGetUserExtensionAttributeByIDMock()
	svc := user_extension_attributes.NewService(mockClient)

	resp, _, err := svc.GetByID(context.Background(), 1)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "Department", resp.Name)
	assert.Equal(t, "User department", resp.Description)
	assert.Equal(t, "String", resp.DataType)
	assert.Equal(t, "Text Field", resp.InputType.Type)
}

func TestUnit_UserExtensionAttributes_GetByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewUserExtensionAttributesMock()
	svc := user_extension_attributes.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user extension attribute ID must be a positive integer")
}

func TestUnit_UserExtensionAttributes_GetByName(t *testing.T) {
	mockClient := mocks.NewUserExtensionAttributesMock()
	mockClient.RegisterGetUserExtensionAttributeByNameMock()
	svc := user_extension_attributes.NewService(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "Department")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "Department", resp.Name)
	assert.Equal(t, "User department", resp.Description)
}

func TestUnit_UserExtensionAttributes_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewUserExtensionAttributesMock()
	svc := user_extension_attributes.NewService(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user extension attribute name cannot be empty")
}

func TestUnit_UserExtensionAttributes_Create(t *testing.T) {
	mockClient := mocks.NewUserExtensionAttributesMock()
	mockClient.RegisterCreateUserExtensionAttributeMock()
	svc := user_extension_attributes.NewService(mockClient)

	req := &user_extension_attributes.RequestUserExtensionAttribute{
		Name:        "Test Extension Attribute",
		Description: "Created for testing",
		DataType:    "String",
		InputType: user_extension_attributes.ResourceUserExtensionAttributeInputType{
			Type: "Text Field",
		},
	}

	resp, _, err := svc.Create(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.ID)
	assert.Equal(t, "Test Extension Attribute", resp.Name)
	assert.Equal(t, "Created for testing", resp.Description)
}

func TestUnit_UserExtensionAttributes_Create_NilRequest(t *testing.T) {
	mockClient := mocks.NewUserExtensionAttributesMock()
	svc := user_extension_attributes.NewService(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_UserExtensionAttributes_Create_EmptyName(t *testing.T) {
	mockClient := mocks.NewUserExtensionAttributesMock()
	svc := user_extension_attributes.NewService(mockClient)

	req := &user_extension_attributes.RequestUserExtensionAttribute{
		Name: "",
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user extension attribute name is required")
}

func TestUnit_UserExtensionAttributes_UpdateByID(t *testing.T) {
	mockClient := mocks.NewUserExtensionAttributesMock()
	mockClient.RegisterUpdateUserExtensionAttributeByIDMock()
	svc := user_extension_attributes.NewService(mockClient)

	req := &user_extension_attributes.RequestUserExtensionAttribute{
		Name:        "Department Updated",
		Description: "Updated user department",
		DataType:    "String",
		InputType: user_extension_attributes.ResourceUserExtensionAttributeInputType{
			Type: "Text Field",
		},
	}

	resp, _, err := svc.UpdateByID(context.Background(), 1, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "Department Updated", resp.Name)
	assert.Equal(t, "Updated user department", resp.Description)
}

func TestUnit_UserExtensionAttributes_UpdateByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewUserExtensionAttributesMock()
	svc := user_extension_attributes.NewService(mockClient)

	req := &user_extension_attributes.RequestUserExtensionAttribute{
		Name: "Test",
	}

	_, _, err := svc.UpdateByID(context.Background(), 0, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user extension attribute ID must be a positive integer")
}

func TestUnit_UserExtensionAttributes_UpdateByName(t *testing.T) {
	mockClient := mocks.NewUserExtensionAttributesMock()
	mockClient.RegisterUpdateUserExtensionAttributeByNameMock()
	svc := user_extension_attributes.NewService(mockClient)

	req := &user_extension_attributes.RequestUserExtensionAttribute{
		Name:        "Department Updated",
		Description: "Updated user department",
		DataType:    "String",
		InputType: user_extension_attributes.ResourceUserExtensionAttributeInputType{
			Type: "Text Field",
		},
	}

	resp, _, err := svc.UpdateByName(context.Background(), "Department", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "Department Updated", resp.Name)
}

func TestUnit_UserExtensionAttributes_UpdateByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewUserExtensionAttributesMock()
	svc := user_extension_attributes.NewService(mockClient)

	req := &user_extension_attributes.RequestUserExtensionAttribute{
		Name: "Test",
	}

	_, _, err := svc.UpdateByName(context.Background(), "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user extension attribute name cannot be empty")
}

func TestUnit_UserExtensionAttributes_DeleteByID(t *testing.T) {
	mockClient := mocks.NewUserExtensionAttributesMock()
	mockClient.RegisterDeleteUserExtensionAttributeByIDMock()
	svc := user_extension_attributes.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 1)

	require.NoError(t, err)
}

func TestUnit_UserExtensionAttributes_DeleteByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewUserExtensionAttributesMock()
	svc := user_extension_attributes.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user extension attribute ID must be a positive integer")
}

func TestUnit_UserExtensionAttributes_DeleteByName(t *testing.T) {
	mockClient := mocks.NewUserExtensionAttributesMock()
	mockClient.RegisterDeleteUserExtensionAttributeByNameMock()
	svc := user_extension_attributes.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "Department")

	require.NoError(t, err)
}

func TestUnit_UserExtensionAttributes_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewUserExtensionAttributesMock()
	svc := user_extension_attributes.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user extension attribute name cannot be empty")
}

func TestUnit_UserExtensionAttributes_NotFound(t *testing.T) {
	mockClient := mocks.NewUserExtensionAttributesMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := user_extension_attributes.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 999)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}

func TestUnit_UserExtensionAttributes_Conflict(t *testing.T) {
	mockClient := mocks.NewUserExtensionAttributesMock()
	mockClient.RegisterConflictErrorMock()
	svc := user_extension_attributes.NewService(mockClient)

	req := &user_extension_attributes.RequestUserExtensionAttribute{
		Name:        "Duplicate Attribute",
		Description: "Test",
		DataType:    "String",
		InputType: user_extension_attributes.ResourceUserExtensionAttributeInputType{
			Type: "Text Field",
		},
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user extension attribute with that name already exists")
}
