package classes

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/classes/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh ClassesMock.
func setupMockService(t *testing.T) (*Service, *mocks.ClassesMock) {
	t.Helper()
	mock := mocks.NewClassesMock()
	return NewService(mock), mock
}

// =============================================================================
// ListClasses
// =============================================================================

func TestUnitListClasses_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListClassesMock()

	result, resp, err := svc.ListClasses(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.Size)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "Test Class 1", result.Results[0].Name)
	assert.Equal(t, 2, result.Results[1].ID)
	assert.Equal(t, "Test Class 2", result.Results[1].Name)
}

// =============================================================================
// GetClassByID
// =============================================================================

func TestUnitGetClassByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetClassByIDMock()

	result, resp, err := svc.GetClassByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Test Class 1", result.Name)
	assert.Equal(t, "Test class description", result.Description)
}

func TestUnitGetClassByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetClassByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "class ID must be a positive integer")
}

func TestUnitGetClassByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetClassByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "class ID must be a positive integer")
}

func TestUnitGetClassByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetClassByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

// =============================================================================
// GetClassByName
// =============================================================================

func TestUnitGetClassByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetClassByNameMock()

	result, resp, err := svc.GetClassByName(context.Background(), "Test Class 1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Test Class 1", result.Name)
}

func TestUnitGetClassByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetClassByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "class name is required")
}

// =============================================================================
// CreateClass
// =============================================================================

func TestUnitCreateClass_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateClassMock()

	req := &RequestClass{Name: "New Test Class"}
	result, resp, err := svc.CreateClass(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 100, result.ID)
}

func TestUnitCreateClass_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateClass(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateClass_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestClass{Name: "Duplicate Class"}
	result, resp, err := svc.CreateClass(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

// =============================================================================
// UpdateClassByID
// =============================================================================

func TestUnitUpdateClassByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateClassByIDMock()

	req := &RequestClass{Name: "Updated Class Name"}
	result, resp, err := svc.UpdateClassByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitUpdateClassByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestClass{Name: "Updated Class Name"}
	result, resp, err := svc.UpdateClassByID(context.Background(), 0, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "class ID must be a positive integer")
}

func TestUnitUpdateClassByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateClassByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateClassByName
// =============================================================================

func TestUnitUpdateClassByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateClassByNameMock()

	req := &RequestClass{Name: "Updated Class Name"}
	result, resp, err := svc.UpdateClassByName(context.Background(), "Test Class 1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitUpdateClassByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestClass{Name: "Updated Class Name"}
	result, resp, err := svc.UpdateClassByName(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "class name is required")
}

func TestUnitUpdateClassByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateClassByName(context.Background(), "Test Class 1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteClassByID
// =============================================================================

func TestUnitDeleteClassByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteClassByIDMock()

	resp, err := svc.DeleteClassByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteClassByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteClassByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "class ID must be a positive integer")
}

func TestUnitDeleteClassByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteClassByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "class ID must be a positive integer")
}

// =============================================================================
// DeleteClassByName
// =============================================================================

func TestUnitDeleteClassByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteClassByNameMock()

	resp, err := svc.DeleteClassByName(context.Background(), "Test Class 1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteClassByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteClassByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "class name is required")
}
