package enrollment_customizations

import (
	"context"
	"strings"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/enrollment_customizations/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUnit_EnrollmentCustomizations_ListV2(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	result, resp, err := svc.ListV2(context.Background(), nil)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, 2, result.TotalCount)
	assert.Len(t, result.Results, 2)
	assert.Equal(t, "Test Customization 1", result.Results[0].DisplayName)
}

func TestUnit_EnrollmentCustomizations_ListV2_ClientError(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, resp, err := svc.ListV2(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
}

func TestUnit_EnrollmentCustomizations_ListV2_WithPagination(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	rsqlQuery := map[string]string{
		"page":      "0",
		"page-size": "10",
		"sort":      "id:asc",
	}
	result, resp, err := svc.ListV2(context.Background(), rsqlQuery)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
}

func TestUnit_EnrollmentCustomizations_GetByIDV2(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	mock.RegisterGetByIDMock()

	svc := NewService(mock)
	result, resp, err := svc.GetByIDV2(context.Background(), "1")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Test Customization 1", result.DisplayName)
}

func TestUnit_EnrollmentCustomizations_GetByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, resp, err := svc.GetByIDV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")
}

func TestUnit_EnrollmentCustomizations_GetByIDV2_ClientError(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, resp, err := svc.GetByIDV2(context.Background(), "1")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizations_GetByNameV2(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	mock.RegisterListMock()
	mock.RegisterGetByIDMock()

	svc := NewService(mock)
	result, resp, err := svc.GetByNameV2(context.Background(), "Test Customization 1")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Test Customization 1", result.DisplayName)
}

func TestUnit_EnrollmentCustomizations_GetByNameV2_EmptyName(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, resp, err := svc.GetByNameV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "enrollment customization name is required")
}

func TestUnit_EnrollmentCustomizations_GetByNameV2_NotFound(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	result, resp, err := svc.GetByNameV2(context.Background(), "NonExistent")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not found")
}

func TestUnit_EnrollmentCustomizations_GetByNameV2_ListError(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, _, err := svc.GetByNameV2(context.Background(), "Test")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUnit_EnrollmentCustomizations_CreateV2(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	mock.RegisterCreateMock()

	svc := NewService(mock)
	req := &ResourceEnrollmentCustomization{
		DisplayName: "New Customization",
		Description: "Test description",
	}
	result, resp, err := svc.CreateV2(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "3", result.ID)
}

func TestUnit_EnrollmentCustomizations_CreateV2_NilRequest(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, resp, err := svc.CreateV2(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_EnrollmentCustomizations_CreateV2_EmptyDisplayName(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	req := &ResourceEnrollmentCustomization{
		DisplayName: "",
	}
	result, resp, err := svc.CreateV2(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "display name is required")
}

func TestUnit_EnrollmentCustomizations_CreateV2_ClientError(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	req := &ResourceEnrollmentCustomization{DisplayName: "New"}
	result, resp, err := svc.CreateV2(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizations_UpdateByIDV2(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	mock.RegisterUpdateMock()

	svc := NewService(mock)
	req := &ResourceEnrollmentCustomization{
		DisplayName: "Updated Customization",
		Description: "Updated description",
	}
	result, resp, err := svc.UpdateByIDV2(context.Background(), "1", req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
}

func TestUnit_EnrollmentCustomizations_UpdateByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	req := &ResourceEnrollmentCustomization{DisplayName: "Test"}
	result, resp, err := svc.UpdateByIDV2(context.Background(), "", req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")
}

func TestUnit_EnrollmentCustomizations_UpdateByIDV2_NilRequest(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, resp, err := svc.UpdateByIDV2(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_EnrollmentCustomizations_UpdateByIDV2_EmptyDisplayName(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	req := &ResourceEnrollmentCustomization{DisplayName: ""}
	result, resp, err := svc.UpdateByIDV2(context.Background(), "1", req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "display name is required")
}

func TestUnit_EnrollmentCustomizations_UpdateByIDV2_ClientError(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	req := &ResourceEnrollmentCustomization{DisplayName: "Updated"}
	result, resp, err := svc.UpdateByIDV2(context.Background(), "1", req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizations_DeleteByIDV2(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	mock.RegisterDeleteMock()

	svc := NewService(mock)
	resp, err := svc.DeleteByIDV2(context.Background(), "1")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestUnit_EnrollmentCustomizations_DeleteByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	resp, err := svc.DeleteByIDV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")
}

func TestUnit_EnrollmentCustomizations_DeleteByIDV2_ClientError(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	resp, err := svc.DeleteByIDV2(context.Background(), "1")

	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizations_GetHistoryV2(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	mock.RegisterGetHistoryMock()

	svc := NewService(mock)
	result, resp, err := svc.GetHistoryV2(context.Background(), "1", nil)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, 2, result.TotalCount)
	assert.Len(t, result.Results, 2)
}

func TestUnit_EnrollmentCustomizations_GetHistoryV2_EmptyID(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, resp, err := svc.GetHistoryV2(context.Background(), "", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")
}

func TestUnit_EnrollmentCustomizations_GetHistoryV2_ClientError(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, resp, err := svc.GetHistoryV2(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizations_AddHistoryNotesV2(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	mock.RegisterAddHistoryNotesMock()

	svc := NewService(mock)
	req := &RequestAddHistoryNotes{Note: "Test note"}
	result, resp, err := svc.AddHistoryNotesV2(context.Background(), "1", req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "Test note added", result.Note)
}

func TestUnit_EnrollmentCustomizations_AddHistoryNotesV2_EmptyID(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	req := &RequestAddHistoryNotes{Note: "Test"}
	result, resp, err := svc.AddHistoryNotesV2(context.Background(), "", req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")
}

func TestUnit_EnrollmentCustomizations_AddHistoryNotesV2_NilRequest(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, resp, err := svc.AddHistoryNotesV2(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request body is required")
}

func TestUnit_EnrollmentCustomizations_AddHistoryNotesV2_EmptyNote(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	req := &RequestAddHistoryNotes{Note: ""}
	result, resp, err := svc.AddHistoryNotesV2(context.Background(), "1", req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "note is required")
}

func TestUnit_EnrollmentCustomizations_AddHistoryNotesV2_ClientError(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	req := &RequestAddHistoryNotes{Note: "Test note"}
	result, resp, err := svc.AddHistoryNotesV2(context.Background(), "1", req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizations_GetPrestagesV2(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	mock.RegisterGetPrestagesMock()

	svc := NewService(mock)
	result, resp, err := svc.GetPrestagesV2(context.Background(), "1")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Len(t, result.Dependencies, 2)
	assert.Equal(t, "Computer Prestage 1", result.Dependencies[0].HumanReadableName)
}

func TestUnit_EnrollmentCustomizations_GetPrestagesV2_EmptyID(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, resp, err := svc.GetPrestagesV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")
}

func TestUnit_EnrollmentCustomizations_GetPrestagesV2_ClientError(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, resp, err := svc.GetPrestagesV2(context.Background(), "1")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizations_UploadImageV2(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	mock.RegisterUploadImageMock()

	svc := NewService(mock)
	fileReader := strings.NewReader("fake image data")
	result, resp, err := svc.UploadImageV2(context.Background(), fileReader, 15, "test.png")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "123", result.ID)
	assert.Contains(t, result.URL, "enrollment-customizations/images/123")
}

func TestUnit_EnrollmentCustomizations_UploadImageV2_NilFileReader(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, resp, err := svc.UploadImageV2(context.Background(), nil, 0, "test.png")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "file reader is required")
}

func TestUnit_EnrollmentCustomizations_UploadImageV2_EmptyFileName(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	fileReader := strings.NewReader("fake image data")
	result, resp, err := svc.UploadImageV2(context.Background(), fileReader, 15, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "file name is required")
}

func TestUnit_EnrollmentCustomizations_UploadImageV2_ClientError(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	fileReader := strings.NewReader("fake image data")
	result, resp, err := svc.UploadImageV2(context.Background(), fileReader, 15, "test.png")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizations_GetImageByIdV2(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	mock.RegisterGetImageByIdMock("123")

	svc := NewService(mock)
	data, resp, err := svc.GetImageByIdV2(context.Background(), "123")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, data)
	assert.Greater(t, len(data), 0)
	// Check for PNG header
	assert.Equal(t, byte(0x89), data[0])
	assert.Equal(t, byte(0x50), data[1])
}

func TestUnit_EnrollmentCustomizations_GetImageByIdV2_EmptyID(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	data, resp, err := svc.GetImageByIdV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, data)
	assert.Contains(t, err.Error(), "image ID is required")
}

func TestUnit_EnrollmentCustomizations_GetImageByIdV2_ClientError(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	data, resp, err := svc.GetImageByIdV2(context.Background(), "123")

	assert.Error(t, err)
	assert.Nil(t, data)
	assert.Nil(t, resp)
}
