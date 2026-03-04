package computer_extension_attributes

import (
	"bytes"
	"context"
	"strings"
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

func TestUnit_ComputerExtensionAttributes_List_Success(t *testing.T) {
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
	assert.Equal(t, "EA One", result.Results[0].Name)
	assert.Equal(t, "2", result.Results[1].ID)
	assert.Equal(t, "EA Two", result.Results[1].Name)
}

func TestUnit_ComputerExtensionAttributes_List_WithrsqlQuery(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	params := map[string]string{"page": "0", "page-size": "50"}
	result, resp, err := svc.ListV1(context.Background(), params)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_ComputerExtensionAttributes_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "EA One", result.Name)
	assert.Equal(t, "STRING", result.DataType)
	assert.NotNil(t, result.Enabled)
	assert.True(t, *result.Enabled)
}

func TestUnit_ComputerExtensionAttributes_GetByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ComputerExtensionAttributes_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_ComputerExtensionAttributes_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	req := &RequestComputerExtensionAttribute{
		Name:                 "New EA",
		DataType:             "STRING",
		InventoryDisplayType: "GENERAL",
		InputType:            "TEXT",
	}
	result, resp, err := svc.CreateV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "3", result.ID)
	assert.NotEmpty(t, result.Href)
}

func TestUnit_ComputerExtensionAttributes_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ComputerExtensionAttributes_UpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	req := &RequestComputerExtensionAttribute{
		Name:                 "EA One Updated",
		Description:          "Updated",
		DataType:             "STRING",
		InventoryDisplayType: "GENERAL",
		InputType:            "TEXT",
	}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "EA One Updated", result.Name)
	assert.Equal(t, "Updated", result.Description)
}

func TestUnit_ComputerExtensionAttributes_UpdateByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestComputerExtensionAttribute{Name: "x", DataType: "STRING", InventoryDisplayType: "GENERAL", InputType: "TEXT"}

	result, resp, err := svc.UpdateByIDV1(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ComputerExtensionAttributes_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ComputerExtensionAttributes_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteMock()

	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_ComputerExtensionAttributes_DeleteByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_ComputerExtensionAttributes_DeleteMultipleByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteMultipleMock()

	req := &DeleteComputerExtensionAttributesByIDRequest{IDs: []string{"1", "2"}}
	resp, err := svc.DeleteComputerExtensionAttributesByIDV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_ComputerExtensionAttributes_DeleteMultipleByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteComputerExtensionAttributesByIDV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_ComputerExtensionAttributes_DeleteMultipleByID_EmptyIDs(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteComputerExtensionAttributesByIDV1(context.Background(), &DeleteComputerExtensionAttributesByIDRequest{IDs: nil})
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_ComputerExtensionAttributes_GetHistoryByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterHistoryMock()

	result, resp, err := svc.GetHistoryByIDV1(context.Background(), "1", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, 2, result.Results[1].ID)
	assert.Equal(t, "Updated script", result.Results[1].Note)
}

func TestUnit_ComputerExtensionAttributes_GetHistoryByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetHistoryByIDV1(context.Background(), "", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_ComputerExtensionAttributes_AddHistoryNoteByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddHistoryNoteMock()

	req := &AddHistoryNoteRequest{Note: "Manual note added"}
	resp, err := svc.AddHistoryNoteByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
}

func TestUnit_ComputerExtensionAttributes_AddHistoryNoteByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &AddHistoryNoteRequest{Note: "Test note"}
	resp, err := svc.AddHistoryNoteByIDV1(context.Background(), "", req)
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_ComputerExtensionAttributes_AddHistoryNoteByIDV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.AddHistoryNoteByIDV1(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_ComputerExtensionAttributes_ListV1_InvalidJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListInvalidMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "failed to list computer extension attributes")
}

func TestUnit_ComputerExtensionAttributes_GetHistoryByIDV1_InvalidJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterHistoryInvalidMock()

	result, resp, err := svc.GetHistoryByIDV1(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "failed to get computer extension attribute history")
}

func TestUnit_ComputerExtensionAttributes_ListTemplatesV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListTemplatesMock()

	result, resp, err := svc.ListTemplatesV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Template One", result.Results[0].Name)
	assert.Equal(t, "2", result.Results[1].ID)
	assert.Equal(t, "Template Two", result.Results[1].Name)
}

func TestUnit_ComputerExtensionAttributes_ListTemplatesV1_WithRsqlQuery(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListTemplatesMock()

	params := map[string]string{"page": "0", "page-size": "50"}
	result, resp, err := svc.ListTemplatesV1(context.Background(), params)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_ComputerExtensionAttributes_ListTemplatesV1_InvalidJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListTemplatesInvalidMock()

	result, resp, err := svc.ListTemplatesV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "failed to list computer extension attribute templates")
}

func TestUnit_ComputerExtensionAttributes_GetTemplateByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetTemplateMock()

	result, resp, err := svc.GetTemplateByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Template One", result.Name)
	assert.Equal(t, "STRING", result.DataType)
	assert.Equal(t, "General", result.TemplateCategory)
}

func TestUnit_ComputerExtensionAttributes_GetTemplateByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetTemplateByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "template ID is required")
}

func TestUnit_ComputerExtensionAttributes_UploadV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUploadMock()

	fileContent := []byte("#!/bin/bash\necho test")
	reader := bytes.NewReader(fileContent)

	result, resp, err := svc.UploadV1(context.Background(), reader, int64(len(fileContent)), "test_script.sh")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "EA One", result.Name)
}

func TestUnit_ComputerExtensionAttributes_UploadV1_NilFileReader(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UploadV1(context.Background(), nil, 100, "test.sh")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "file reader is required")
}

func TestUnit_ComputerExtensionAttributes_UploadV1_EmptyFilename(t *testing.T) {
	svc, _ := setupMockService(t)

	reader := strings.NewReader("content")
	result, resp, err := svc.UploadV1(context.Background(), reader, 7, "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "filename is required")
}

func TestUnit_ComputerExtensionAttributes_GetDataDependencyByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDataDependencyMock()

	result, resp, err := svc.GetDataDependencyByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "Smart Group A", result.Results[0].IdentifiableName)
	assert.Equal(t, 2, result.Results[1].ID)
	assert.Equal(t, "Advanced Search B", result.Results[1].IdentifiableName)
}

func TestUnit_ComputerExtensionAttributes_GetDataDependencyByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetDataDependencyByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "computer extension attribute ID is required")
}

func TestUnit_ComputerExtensionAttributes_DownloadByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDownloadMock()

	result, resp, err := svc.DownloadByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, result)
	assert.Contains(t, string(result), "<computer_extension_attribute>")
	assert.Contains(t, string(result), "<id>1</id>")
}

func TestUnit_ComputerExtensionAttributes_DownloadByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.DownloadByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "computer extension attribute ID is required")
}
