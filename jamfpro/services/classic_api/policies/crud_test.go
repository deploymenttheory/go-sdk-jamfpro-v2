package policies

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/policies/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh PoliciesMock.
func setupMockService(t *testing.T) (*Service, *mocks.PoliciesMock) {
	t.Helper()
	mock := mocks.NewPoliciesMock()
	return NewService(mock), mock
}

// =============================================================================
// List
// =============================================================================

func TestUnitList_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.List(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.Size)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "Test Policy", result.Results[0].Name)
	assert.Equal(t, 2, result.Results[1].ID)
	assert.Equal(t, "Another Policy", result.Results[1].Name)
}

// =============================================================================
// GetByID
// =============================================================================

func TestUnitGetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDMock()

	result, resp, err := svc.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.General.ID)
	assert.Equal(t, "Test Policy", result.General.Name)
	assert.True(t, result.General.Enabled)
	assert.Equal(t, "Once per computer", result.General.Frequency)
	assert.True(t, result.Scope.AllComputers)
}

func TestUnitGetByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "policy ID must be a positive integer")
}

func TestUnitGetByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "policy ID must be a positive integer")
}

func TestUnitGetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

// =============================================================================
// GetByName
// =============================================================================

func TestUnitGetByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByNameMock()

	result, resp, err := svc.GetByName(context.Background(), "Test Policy")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.General.ID)
	assert.Equal(t, "Test Policy", result.General.Name)
	assert.True(t, result.General.Enabled)
}

func TestUnitGetByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "policy name is required")
}

// =============================================================================
// Create
// =============================================================================

func TestUnitCreate_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	req := &ResourcePolicy{
		General: PolicySubsetGeneral{
			Name:      "New Policy",
			Enabled:   true,
			Frequency: "Once per computer",
		},
		Scope: PolicySubsetScope{
			AllComputers: true,
		},
	}

	result, resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 3, result.ID)
}

func TestUnitCreate_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.Create(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "policy is required")
}

func TestUnitCreate_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &ResourcePolicy{
		General: PolicySubsetGeneral{
			Name:      "Duplicate Policy",
			Enabled:   true,
			Frequency: "Once per computer",
		},
		Scope: PolicySubsetScope{
			AllComputers: true,
		},
	}

	result, resp, err := svc.Create(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
	assert.Contains(t, err.Error(), "409")
}

// =============================================================================
// UpdateByID
// =============================================================================

func TestUnitUpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateByIDMock()

	req := &ResourcePolicy{
		General: PolicySubsetGeneral{
			Name:      "Updated Policy",
			Enabled:   false,
			Frequency: "Ongoing",
		},
		Scope: PolicySubsetScope{
			AllComputers: true,
		},
	}

	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitUpdateByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourcePolicy{
		General: PolicySubsetGeneral{
			Name:      "Test",
			Enabled:   true,
			Frequency: "Once per computer",
		},
	}

	result, resp, err := svc.UpdateByID(context.Background(), 0, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "policy ID must be a positive integer")
}

func TestUnitUpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "policy is required")
}

// =============================================================================
// UpdateByName
// =============================================================================

func TestUnitUpdateByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateByNameMock()

	req := &ResourcePolicy{
		General: PolicySubsetGeneral{
			Name:      "Test Policy",
			Enabled:   true,
			Frequency: "Once per computer",
		},
		Scope: PolicySubsetScope{
			AllComputers: true,
		},
	}

	result, resp, err := svc.UpdateByName(context.Background(), "Test Policy", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitUpdateByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourcePolicy{
		General: PolicySubsetGeneral{
			Name:      "Test",
			Enabled:   true,
			Frequency: "Once per computer",
		},
	}

	result, resp, err := svc.UpdateByName(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "policy name is required")
}

func TestUnitUpdateByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByName(context.Background(), "Test Policy", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "policy is required")
}

// =============================================================================
// DeleteByID
// =============================================================================

func TestUnitDeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteByIDMock()

	resp, err := svc.DeleteByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "policy ID must be a positive integer")
}

func TestUnitDeleteByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "policy ID must be a positive integer")
}

// =============================================================================
// DeleteByName
// =============================================================================

func TestUnitDeleteByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteByNameMock()

	resp, err := svc.DeleteByName(context.Background(), "Test Policy")
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "policy name is required")
}

// =============================================================================
// GetByCreatedBy
// =============================================================================

func TestUnitGetByCreatedBy_Success_JSS(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByCreatedByMock()

	result, resp, err := svc.GetByCreatedBy(context.Background(), "jss")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.Size)
	require.Len(t, result.Results, 2)
}

func TestUnitGetByCreatedBy_EmptyCreatedBy(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByCreatedBy(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "createdBy is required")
}

func TestUnitGetByCreatedBy_InvalidCreatedBy(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByCreatedBy(context.Background(), "invalid")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "must be 'jss' or 'casper'")
}

// =============================================================================
// GetByCategory
// =============================================================================

func TestUnitGetByCategory_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByCategoryMock()

	result, resp, err := svc.GetByCategory(context.Background(), "TestCategory")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.Size)
	require.Len(t, result.Results, 2)
}

func TestUnitGetByCategory_EmptyCategory(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByCategory(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "category is required")
}

// =============================================================================
// GetByIDWithSubset
// =============================================================================

func TestUnitGetByIDWithSubset_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDWithSubsetMock()

	result, resp, err := svc.GetByIDWithSubset(context.Background(), 1, "General")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.General.ID)
	assert.Equal(t, "Test Policy", result.General.Name)
}

func TestUnitGetByIDWithSubset_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDWithSubset(context.Background(), 0, "General")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "policy ID must be a positive integer")
}

func TestUnitGetByIDWithSubset_EmptySubset(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDWithSubset(context.Background(), 1, "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "subset is required")
}

func TestUnitGetByIDWithSubset_MultipleSubsets(t *testing.T) {
	svc, mock := setupMockService(t)
	// Register a mock for combined subsets
	mock.Register("GET", "/JSSResource/policies/id/1/subset/General&Scope", 200, "validate_get_policy.xml")

	result, resp, err := svc.GetByIDWithSubset(context.Background(), 1, "General&Scope")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.General.ID)
}

// =============================================================================
// GetByNameWithSubset
// =============================================================================

func TestUnitGetByNameWithSubset_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByNameWithSubsetMock()

	result, resp, err := svc.GetByNameWithSubset(context.Background(), "Test Policy", "General")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.General.ID)
	assert.Equal(t, "Test Policy", result.General.Name)
}

func TestUnitGetByNameWithSubset_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByNameWithSubset(context.Background(), "", "General")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "policy name is required")
}

func TestUnitGetByNameWithSubset_EmptySubset(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByNameWithSubset(context.Background(), "Test Policy", "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "subset is required")
}

func TestUnitGetByNameWithSubset_MultipleSubsets(t *testing.T) {
	svc, mock := setupMockService(t)
	// Register a mock for combined subsets
	mock.Register("GET", "/JSSResource/policies/name/Test Policy/subset/General&Scope", 200, "validate_get_policy.xml")

	result, resp, err := svc.GetByNameWithSubset(context.Background(), "Test Policy", "General&Scope")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.General.ID)
}
