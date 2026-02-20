package advanced_user_searches

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/advanced_user_searches/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh AdvancedUserSearchesMock.
func setupMockService(t *testing.T) (*Service, *mocks.AdvancedUserSearchesMock) {
	t.Helper()
	mock := mocks.NewAdvancedUserSearchesMock()
	return NewService(mock), mock
}

// =============================================================================
// ListAdvancedUserSearches
// =============================================================================

func TestUnitListAdvancedUserSearches_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListAdvancedUserSearchesMock()

	result, resp, err := svc.ListAdvancedUserSearches(context.Background())
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
// GetAdvancedUserSearchByID
// =============================================================================

func TestUnitGetAdvancedUserSearchByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAdvancedUserSearchByIDMock()

	result, resp, err := svc.GetAdvancedUserSearchByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Test Search", result.Name)
	assert.Equal(t, 2, result.Criteria.Size)
	require.Len(t, result.Criteria.Criterion, 2)
	assert.Equal(t, "Username", result.Criteria.Criterion[0].Name)
	assert.Equal(t, "test", result.Criteria.Criterion[0].Value)
	require.Len(t, result.DisplayFields, 3)
	assert.Equal(t, "Username", result.DisplayFields[0].Name)
	require.Len(t, result.Users, 2)
	assert.Equal(t, 100, result.Users[0].ID)
	assert.Equal(t, "Test User 1", result.Users[0].Name)
	assert.Equal(t, "testuser1", result.Users[0].Username)
}

func TestUnitGetAdvancedUserSearchByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAdvancedUserSearchByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "advanced user search ID must be a positive integer")
}

func TestUnitGetAdvancedUserSearchByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAdvancedUserSearchByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "advanced user search ID must be a positive integer")
}

func TestUnitGetAdvancedUserSearchByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetAdvancedUserSearchByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

// =============================================================================
// GetAdvancedUserSearchByName
// =============================================================================

func TestUnitGetAdvancedUserSearchByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAdvancedUserSearchByNameMock()

	result, resp, err := svc.GetAdvancedUserSearchByName(context.Background(), "Test Search")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Test Search", result.Name)
}

func TestUnitGetAdvancedUserSearchByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAdvancedUserSearchByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "advanced user search name is required")
}

// =============================================================================
// CreateAdvancedUserSearch
// =============================================================================

func TestUnitCreateAdvancedUserSearch_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateAdvancedUserSearchMock()

	req := &RequestAdvancedUserSearch{
		Name: "New Test Search",
		Criteria: CriteriaContainer{
			Size: 1,
			Criterion: []Criterion{
				{
					Name:       "Username",
					Priority:   0,
					SearchType: "like",
					Value:      "test",
				},
			},
		},
	}
	result, resp, err := svc.CreateAdvancedUserSearch(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 100, result.ID)
}

func TestUnitCreateAdvancedUserSearch_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateAdvancedUserSearch(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateAdvancedUserSearch_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestAdvancedUserSearch{Name: "Duplicate Search"}
	result, resp, err := svc.CreateAdvancedUserSearch(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

// =============================================================================
// UpdateAdvancedUserSearchByID
// =============================================================================

func TestUnitUpdateAdvancedUserSearchByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateAdvancedUserSearchByIDMock()

	req := &RequestAdvancedUserSearch{Name: "Updated Search Name"}
	result, resp, err := svc.UpdateAdvancedUserSearchByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitUpdateAdvancedUserSearchByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestAdvancedUserSearch{Name: "Updated Search Name"}
	result, resp, err := svc.UpdateAdvancedUserSearchByID(context.Background(), 0, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "advanced user search ID must be a positive integer")
}

func TestUnitUpdateAdvancedUserSearchByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateAdvancedUserSearchByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateAdvancedUserSearchByName
// =============================================================================

func TestUnitUpdateAdvancedUserSearchByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateAdvancedUserSearchByNameMock()

	req := &RequestAdvancedUserSearch{Name: "Updated Search Name"}
	result, resp, err := svc.UpdateAdvancedUserSearchByName(context.Background(), "Test Search", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitUpdateAdvancedUserSearchByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestAdvancedUserSearch{Name: "Updated Search Name"}
	result, resp, err := svc.UpdateAdvancedUserSearchByName(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "advanced user search name is required")
}

func TestUnitUpdateAdvancedUserSearchByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateAdvancedUserSearchByName(context.Background(), "Test Search", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteAdvancedUserSearchByID
// =============================================================================

func TestUnitDeleteAdvancedUserSearchByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteAdvancedUserSearchByIDMock()

	resp, err := svc.DeleteAdvancedUserSearchByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteAdvancedUserSearchByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteAdvancedUserSearchByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "advanced user search ID must be a positive integer")
}

func TestUnitDeleteAdvancedUserSearchByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteAdvancedUserSearchByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "advanced user search ID must be a positive integer")
}

// =============================================================================
// DeleteAdvancedUserSearchByName
// =============================================================================

func TestUnitDeleteAdvancedUserSearchByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteAdvancedUserSearchByNameMock()

	resp, err := svc.DeleteAdvancedUserSearchByName(context.Background(), "Test Search")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteAdvancedUserSearchByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteAdvancedUserSearchByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "advanced user search name is required")
}
