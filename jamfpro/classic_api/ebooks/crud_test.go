package ebooks_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/ebooks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/ebooks/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_Ebooks_List(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	mockClient.RegisterListEbooksMock()
	svc := ebooks.NewEbooks(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 2, resp.Size)
	assert.Len(t, resp.Results, 2)
	assert.Equal(t, "Sample Ebook 1", resp.Results[0].Name)
	assert.Equal(t, 1, resp.Results[0].ID)
	assert.Equal(t, "Sample Ebook 2", resp.Results[1].Name)
	assert.Equal(t, 2, resp.Results[1].ID)
}

func TestUnit_Ebooks_GetByID(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	mockClient.RegisterGetEbookByIDMock()
	svc := ebooks.NewEbooks(mockClient)

	resp, _, err := svc.GetByID(context.Background(), 1)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Sample Ebook", resp.General.Name)
	assert.Equal(t, "Test Author", resp.General.Author)
	assert.Equal(t, "1.0", resp.General.Version)
	assert.True(t, resp.General.Free)
	assert.Equal(t, "https://example.com/ebook.pdf", resp.General.URL)
	assert.NotNil(t, resp.General.Category)
	assert.Equal(t, 1, resp.General.Category.ID)
	assert.Equal(t, "Ebooks", resp.General.Category.Name)
	assert.True(t, resp.Scope.AllComputers)
}

func TestUnit_Ebooks_GetByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)

	_, _, err := svc.GetByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "ebook ID must be a positive integer")
}

func TestUnit_Ebooks_GetByName(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	mockClient.RegisterGetEbookByNameMock()
	svc := ebooks.NewEbooks(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "Sample Ebook")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Sample Ebook", resp.General.Name)
}

func TestUnit_Ebooks_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "ebook name cannot be empty")
}

func TestUnit_Ebooks_GetByNameAndSubset(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	mockClient.RegisterGetEbookByNameAndSubsetMock()
	svc := ebooks.NewEbooks(mockClient)

	resp, _, err := svc.GetByNameAndSubset(context.Background(), "Sample Ebook", "General")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Sample Ebook", resp.General.Name)
}

func TestUnit_Ebooks_GetByNameAndSubset_EmptyName(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)

	_, _, err := svc.GetByNameAndSubset(context.Background(), "", "General")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "ebook name cannot be empty")
}

func TestUnit_Ebooks_GetByNameAndSubset_EmptySubset(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)

	_, _, err := svc.GetByNameAndSubset(context.Background(), "Sample Ebook", "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "ebook subset cannot be empty")
}

func TestUnit_Ebooks_Create(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	mockClient.RegisterCreateEbookMock()
	svc := ebooks.NewEbooks(mockClient)

	req := &ebooks.Resource{
		General: ebooks.SubsetGeneral{
			Name:   "Test Ebook",
			Author: "Author",
			URL:    "https://example.com/ebook.pdf",
			Site:   shared.SharedResourceSite{ID: -1, Name: "None"},
		},
	}

	resp, _, err := svc.Create(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.ID)
}

func TestUnit_Ebooks_Create_NilRequest(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Ebooks_Create_EmptyName(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)

	req := &ebooks.Resource{
		General: ebooks.SubsetGeneral{
			Name: "",
		},
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "ebook name is required")
}

func TestUnit_Ebooks_UpdateByID(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	mockClient.RegisterUpdateEbookByIDMock()
	svc := ebooks.NewEbooks(mockClient)

	req := &ebooks.Resource{
		General: ebooks.SubsetGeneral{
			Name:   "Updated Ebook",
			Author: "Author",
			URL:    "https://example.com/ebook.pdf",
		},
	}

	resp, _, err := svc.UpdateByID(context.Background(), 1, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_Ebooks_UpdateByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)

	req := &ebooks.Resource{
		General: ebooks.SubsetGeneral{Name: "Test"},
	}

	_, _, err := svc.UpdateByID(context.Background(), 0, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "ebook ID must be a positive integer")
}

func TestUnit_Ebooks_UpdateByName(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	mockClient.RegisterUpdateEbookByNameMock()
	svc := ebooks.NewEbooks(mockClient)

	req := &ebooks.Resource{
		General: ebooks.SubsetGeneral{
			Name:   "Updated Ebook",
			Author: "Author",
			URL:    "https://example.com/ebook.pdf",
		},
	}

	resp, _, err := svc.UpdateByName(context.Background(), "Sample Ebook", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_Ebooks_UpdateByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)

	req := &ebooks.Resource{
		General: ebooks.SubsetGeneral{Name: "Test"},
	}

	_, _, err := svc.UpdateByName(context.Background(), "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "ebook name cannot be empty")
}

func TestUnit_Ebooks_DeleteByID(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	mockClient.RegisterDeleteEbookByIDMock()
	svc := ebooks.NewEbooks(mockClient)

	_, err := svc.DeleteByID(context.Background(), 1)

	require.NoError(t, err)
}

func TestUnit_Ebooks_DeleteByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)

	_, err := svc.DeleteByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "ebook ID must be a positive integer")
}

func TestUnit_Ebooks_DeleteByName(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	mockClient.RegisterDeleteEbookByNameMock()
	svc := ebooks.NewEbooks(mockClient)

	_, err := svc.DeleteByName(context.Background(), "Sample Ebook")

	require.NoError(t, err)
}

func TestUnit_Ebooks_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "ebook name cannot be empty")
}

func TestUnit_Ebooks_NotFound(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := ebooks.NewEbooks(mockClient)

	_, _, err := svc.GetByID(context.Background(), 999)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}

func TestUnit_Ebooks_Conflict(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	mockClient.RegisterConflictErrorMock()
	svc := ebooks.NewEbooks(mockClient)

	req := &ebooks.Resource{
		General: ebooks.SubsetGeneral{
			Name: "Duplicate Ebook",
			URL:  "https://example.com/ebook.pdf",
		},
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "ebook with that name already exists")
}

func TestUnit_Ebooks_List_Error(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)
	_, _, err := svc.List(context.Background())
	require.Error(t, err)
}

func TestUnit_Ebooks_GetByName_Error(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)
	_, _, err := svc.GetByName(context.Background(), "Sample Ebook")
	require.Error(t, err)
}

func TestUnit_Ebooks_GetByNameAndSubset_Error(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)
	_, _, err := svc.GetByNameAndSubset(context.Background(), "Sample Ebook", "General")
	require.Error(t, err)
}

func TestUnit_Ebooks_UpdateByID_NilRequest(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Ebooks_UpdateByID_EmptyName(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), 1, &ebooks.Resource{General: ebooks.SubsetGeneral{Name: ""}})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "ebook name is required")
}

func TestUnit_Ebooks_UpdateByID_Error(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), 1, &ebooks.Resource{General: ebooks.SubsetGeneral{Name: "Sample Ebook"}})
	require.Error(t, err)
}

func TestUnit_Ebooks_UpdateByName_NilRequest(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "Sample Ebook", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Ebooks_UpdateByName_EmptyReqName(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "Sample Ebook", &ebooks.Resource{General: ebooks.SubsetGeneral{Name: ""}})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "ebook name is required in request")
}

func TestUnit_Ebooks_UpdateByName_Error(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "Sample Ebook", &ebooks.Resource{General: ebooks.SubsetGeneral{Name: "Updated"}})
	require.Error(t, err)
}

func TestUnit_Ebooks_DeleteByID_Error(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)
	_, err := svc.DeleteByID(context.Background(), 1)
	require.Error(t, err)
}

func TestUnit_Ebooks_DeleteByName_Error(t *testing.T) {
	mockClient := mocks.NewEbooksMock()
	svc := ebooks.NewEbooks(mockClient)
	_, err := svc.DeleteByName(context.Background(), "Sample Ebook")
	require.Error(t, err)
}
