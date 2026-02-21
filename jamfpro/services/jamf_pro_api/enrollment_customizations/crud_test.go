package enrollment_customizations

import (
	"context"
	"strings"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/enrollment_customizations/mocks"
	"github.com/stretchr/testify/assert"
)

func TestListV2(t *testing.T) {
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

func TestListV2_WithPagination(t *testing.T) {
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

func TestGetByIDV2(t *testing.T) {
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

func TestGetByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, resp, err := svc.GetByIDV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")
}

func TestGetByNameV2(t *testing.T) {
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

func TestGetByNameV2_EmptyName(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, resp, err := svc.GetByNameV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "enrollment customization name is required")
}

func TestGetByNameV2_NotFound(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	result, resp, err := svc.GetByNameV2(context.Background(), "NonExistent")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not found")
}

func TestCreateV2(t *testing.T) {
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

func TestCreateV2_NilRequest(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, resp, err := svc.CreateV2(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestCreateV2_EmptyDisplayName(t *testing.T) {
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

func TestUpdateByIDV2(t *testing.T) {
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

func TestUpdateByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	req := &ResourceEnrollmentCustomization{DisplayName: "Test"}
	result, resp, err := svc.UpdateByIDV2(context.Background(), "", req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")
}

func TestUpdateByIDV2_NilRequest(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, resp, err := svc.UpdateByIDV2(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUpdateByIDV2_EmptyDisplayName(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	req := &ResourceEnrollmentCustomization{DisplayName: ""}
	result, resp, err := svc.UpdateByIDV2(context.Background(), "1", req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "display name is required")
}

func TestDeleteByIDV2(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	mock.RegisterDeleteMock()

	svc := NewService(mock)
	resp, err := svc.DeleteByIDV2(context.Background(), "1")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestDeleteByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	resp, err := svc.DeleteByIDV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")
}

func TestGetHistoryV2(t *testing.T) {
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

func TestGetHistoryV2_EmptyID(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, resp, err := svc.GetHistoryV2(context.Background(), "", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")
}

func TestAddHistoryNotesV2(t *testing.T) {
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

func TestAddHistoryNotesV2_EmptyID(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	req := &RequestAddHistoryNotes{Note: "Test"}
	result, resp, err := svc.AddHistoryNotesV2(context.Background(), "", req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")
}

func TestAddHistoryNotesV2_NilRequest(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, resp, err := svc.AddHistoryNotesV2(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request body is required")
}

func TestAddHistoryNotesV2_EmptyNote(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	req := &RequestAddHistoryNotes{Note: ""}
	result, resp, err := svc.AddHistoryNotesV2(context.Background(), "1", req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "note is required")
}

func TestGetPrestagesV2(t *testing.T) {
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

func TestGetPrestagesV2_EmptyID(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, resp, err := svc.GetPrestagesV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")
}

func TestUploadImageV2(t *testing.T) {
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

func TestUploadImageV2_NilFileReader(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	result, resp, err := svc.UploadImageV2(context.Background(), nil, 0, "test.png")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "file reader is required")
}

func TestUploadImageV2_EmptyFileName(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	fileReader := strings.NewReader("fake image data")
	result, resp, err := svc.UploadImageV2(context.Background(), fileReader, 15, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "file name is required")
}

func TestGetImageByIdV2(t *testing.T) {
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

func TestGetImageByIdV2_EmptyID(t *testing.T) {
	mock := mocks.NewEnrollmentCustomizationsMock()
	svc := NewService(mock)

	data, resp, err := svc.GetImageByIdV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, data)
	assert.Contains(t, err.Error(), "image ID is required")
}
