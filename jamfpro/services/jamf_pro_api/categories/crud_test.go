package categories

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/categories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh CategoriesMock.
func setupMockService(t *testing.T) (*Service, *mocks.CategoriesMock) {
	t.Helper()
	mock := mocks.NewCategoriesMock()
	return NewService(mock), mock
}

// =============================================================================
// ListCategoriesV1
// =============================================================================

func TestUnitListCategories_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListCategoriesMock()

	result, resp, err := svc.ListCategoriesV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "No priority", result.Results[0].Name)
	assert.Equal(t, 9, result.Results[0].Priority)
	assert.Equal(t, "2", result.Results[1].ID)
	assert.Equal(t, "Critical", result.Results[1].Name)
	assert.Equal(t, 1, result.Results[1].Priority)
}

func TestUnitListCategories_WithQueryParams(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListCategoriesMock()

	params := map[string]string{"page": "0", "page-size": "50", "sort": "name:asc"}
	result, resp, err := svc.ListCategoriesV1(context.Background(), params)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitListCategories_WithRSQLFilter(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListCategoriesRSQLMock()

	rsqlQuery := map[string]string{"filter": `name=="Critical"`}
	result, resp, err := svc.ListCategoriesV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount, "filtered result should contain exactly one category")
	require.Len(t, result.Results, 1)
	assert.Equal(t, "2", result.Results[0].ID)
	assert.Equal(t, "Critical", result.Results[0].Name)
	assert.Equal(t, 1, result.Results[0].Priority)

	// Verify the service forwarded the RSQL filter to the HTTP client unchanged.
	assert.Equal(t, rsqlQuery, mock.LastRSQLQuery, "rsqlQuery should be passed through to the HTTP client")
}

// =============================================================================
// GetCategoryByIDV1
// =============================================================================

func TestUnitGetCategoryByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetCategoryMock()

	result, resp, err := svc.GetCategoryByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "No priority", result.Name)
	assert.Equal(t, 9, result.Priority)
}

func TestUnitGetCategoryByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetCategoryByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "category ID is required")
}

func TestUnitGetCategoryByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetCategoryByIDV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

// =============================================================================
// CreateCategoryV1
// =============================================================================

func TestUnitCreateCategory_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateCategoryMock()

	req := &RequestCategory{Name: "Test Category", Priority: 5}
	result, resp, err := svc.CreateCategoryV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "3", result.ID)
	assert.Contains(t, result.Href, "/api/v1/categories/3")
}

func TestUnitCreateCategory_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateCategoryV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateCategory_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestCategory{Name: "Duplicate Category", Priority: 5}
	result, resp, err := svc.CreateCategoryV1(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

// =============================================================================
// UpdateCategoryByIDV1
// =============================================================================

func TestUnitUpdateCategoryByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateCategoryMock()

	req := &RequestCategory{Name: "Updated Category", Priority: 3}
	result, resp, err := svc.UpdateCategoryByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Contains(t, result.Href, "/api/v1/categories/1")
}

func TestUnitUpdateCategoryByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateCategoryByIDV1(context.Background(), "", &RequestCategory{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnitUpdateCategoryByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateCategoryByIDV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteCategoryByIDV1
// =============================================================================

func TestUnitDeleteCategoryByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteCategoryMock()

	resp, err := svc.DeleteCategoryByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitDeleteCategoryByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteCategoryByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "category ID is required")
}

// =============================================================================
// DeleteCategoriesByIDV1 (bulk)
// =============================================================================

func TestUnitDeleteCategoriesByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteCategoriesBulkMock()

	req := &DeleteCategoriesByIDRequest{IDs: []string{"1", "2"}}
	resp, err := svc.DeleteCategoriesByIDV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitDeleteCategoriesByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteCategoriesByIDV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "ids are required")
}

func TestUnitDeleteCategoriesByID_EmptyIDs(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &DeleteCategoriesByIDRequest{IDs: []string{}}
	resp, err := svc.DeleteCategoriesByIDV1(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "ids are required")
}

// =============================================================================
// GetCategoryHistoryV1
// =============================================================================

func TestUnitGetCategoryHistory_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetCategoryHistoryMock()

	result, resp, err := svc.GetCategoryHistoryV1(context.Background(), "1", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "Category created", result.Results[0].Note)
}

func TestUnitGetCategoryHistory_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetCategoryHistoryV1(context.Background(), "", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "category ID is required")
}

// =============================================================================
// AddCategoryHistoryNotesV1
// =============================================================================

func TestUnitAddCategoryHistoryNotes_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddCategoryHistoryNotesMock()

	req := &AddCategoryHistoryNotesRequest{Note: "Test note added"}
	resp, err := svc.AddCategoryHistoryNotesV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
}

func TestUnitAddCategoryHistoryNotes_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &AddCategoryHistoryNotesRequest{Note: "Test note"}
	resp, err := svc.AddCategoryHistoryNotesV1(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "category ID is required")
}

func TestUnitAddCategoryHistoryNotes_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.AddCategoryHistoryNotesV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request body is required")
}
