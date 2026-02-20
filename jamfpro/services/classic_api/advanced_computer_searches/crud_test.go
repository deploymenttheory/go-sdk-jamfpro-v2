package advanced_computer_searches

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/advanced_computer_searches/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh AdvancedComputerSearchesMock.
func setupMockService(t *testing.T) (*Service, *mocks.AdvancedComputerSearchesMock) {
	t.Helper()
	mock := mocks.NewAdvancedComputerSearchesMock()
	return NewService(mock), mock
}

// =============================================================================
// ListAdvancedComputerSearches
// =============================================================================

func TestUnitListAdvancedComputerSearches_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListAdvancedComputerSearchesMock()

	result, resp, err := svc.ListAdvancedComputerSearches(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.Size)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "Test Search", result.Results[0].Name)
	assert.Equal(t, 2, result.Results[1].ID)
	assert.Equal(t, "Another Search", result.Results[1].Name)
}

// =============================================================================
// GetAdvancedComputerSearchByID
// =============================================================================

func TestUnitGetAdvancedComputerSearchByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAdvancedComputerSearchByIDMock()

	result, resp, err := svc.GetAdvancedComputerSearchByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Test Search", result.Name)
	assert.Equal(t, "Standard Web Page", result.ViewAs)
	assert.Equal(t, 2, result.Criteria.Size)
	require.Len(t, result.Criteria.Criterion, 2)
	assert.Equal(t, "Operating System", result.Criteria.Criterion[0].Name)
	assert.Equal(t, "macOS 13", result.Criteria.Criterion[0].Value)
	require.Len(t, result.DisplayFields, 3)
	assert.Equal(t, "Computer Name", result.DisplayFields[0].Name)
	require.Len(t, result.Computers, 2)
	assert.Equal(t, 100, result.Computers[0].ID)
	assert.Equal(t, "TestMac001", result.Computers[0].Name)
}

func TestUnitGetAdvancedComputerSearchByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAdvancedComputerSearchByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "advanced computer search ID must be a positive integer")
}

func TestUnitGetAdvancedComputerSearchByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAdvancedComputerSearchByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "advanced computer search ID must be a positive integer")
}

func TestUnitGetAdvancedComputerSearchByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetAdvancedComputerSearchByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

// =============================================================================
// GetAdvancedComputerSearchByName
// =============================================================================

func TestUnitGetAdvancedComputerSearchByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAdvancedComputerSearchByNameMock()

	result, resp, err := svc.GetAdvancedComputerSearchByName(context.Background(), "Test Search")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Test Search", result.Name)
}

func TestUnitGetAdvancedComputerSearchByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAdvancedComputerSearchByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "advanced computer search name is required")
}

// =============================================================================
// CreateAdvancedComputerSearch
// =============================================================================

func TestUnitCreateAdvancedComputerSearch_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateAdvancedComputerSearchMock()

	req := &RequestAdvancedComputerSearch{
		Name: "New Test Search",
		Criteria: CriteriaContainer{
			Size: 1,
			Criterion: []Criterion{
				{
					Name:       "Operating System",
					Priority:   0,
					SearchType: "is",
					Value:      "macOS 14",
				},
			},
		},
	}
	result, resp, err := svc.CreateAdvancedComputerSearch(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 100, result.ID)
}

func TestUnitCreateAdvancedComputerSearch_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateAdvancedComputerSearch(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateAdvancedComputerSearch_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestAdvancedComputerSearch{Name: "Duplicate Search"}
	result, resp, err := svc.CreateAdvancedComputerSearch(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

// =============================================================================
// UpdateAdvancedComputerSearchByID
// =============================================================================

func TestUnitUpdateAdvancedComputerSearchByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateAdvancedComputerSearchByIDMock()

	req := &RequestAdvancedComputerSearch{Name: "Updated Search Name"}
	result, resp, err := svc.UpdateAdvancedComputerSearchByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitUpdateAdvancedComputerSearchByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestAdvancedComputerSearch{Name: "Updated Search Name"}
	result, resp, err := svc.UpdateAdvancedComputerSearchByID(context.Background(), 0, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "advanced computer search ID must be a positive integer")
}

func TestUnitUpdateAdvancedComputerSearchByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateAdvancedComputerSearchByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateAdvancedComputerSearchByName
// =============================================================================

func TestUnitUpdateAdvancedComputerSearchByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateAdvancedComputerSearchByNameMock()

	req := &RequestAdvancedComputerSearch{Name: "Updated Search Name"}
	result, resp, err := svc.UpdateAdvancedComputerSearchByName(context.Background(), "Test Search", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitUpdateAdvancedComputerSearchByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestAdvancedComputerSearch{Name: "Updated Search Name"}
	result, resp, err := svc.UpdateAdvancedComputerSearchByName(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "advanced computer search name is required")
}

func TestUnitUpdateAdvancedComputerSearchByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateAdvancedComputerSearchByName(context.Background(), "Test Search", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteAdvancedComputerSearchByID
// =============================================================================

func TestUnitDeleteAdvancedComputerSearchByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteAdvancedComputerSearchByIDMock()

	resp, err := svc.DeleteAdvancedComputerSearchByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteAdvancedComputerSearchByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteAdvancedComputerSearchByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "advanced computer search ID must be a positive integer")
}

func TestUnitDeleteAdvancedComputerSearchByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteAdvancedComputerSearchByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "advanced computer search ID must be a positive integer")
}

// =============================================================================
// DeleteAdvancedComputerSearchByName
// =============================================================================

func TestUnitDeleteAdvancedComputerSearchByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteAdvancedComputerSearchByNameMock()

	resp, err := svc.DeleteAdvancedComputerSearchByName(context.Background(), "Test Search")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteAdvancedComputerSearchByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteAdvancedComputerSearchByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "advanced computer search name is required")
}
