package network_segments

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/network_segments/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh NetworkSegmentsMock.
func setupMockService(t *testing.T) (*Service, *mocks.NetworkSegmentsMock) {
	t.Helper()
	mock := mocks.NewNetworkSegmentsMock()
	return NewService(mock), mock
}

// =============================================================================
// ListNetworkSegments
// =============================================================================

func TestUnitListNetworkSegments_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListNetworkSegmentsMock()

	result, resp, err := svc.ListNetworkSegments(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.Size)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "HQ Network", result.Results[0].Name)
	assert.Equal(t, 2, result.Results[1].ID)
	assert.Equal(t, "Branch Network", result.Results[1].Name)
}

// =============================================================================
// GetNetworkSegmentByID
// =============================================================================

func TestUnitGetNetworkSegmentByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetNetworkSegmentByIDMock()

	result, resp, err := svc.GetNetworkSegmentByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "HQ Network", result.Name)
	assert.Equal(t, "10.0.0.0", result.StartingAddress)
	assert.Equal(t, "10.0.0.255", result.EndingAddress)
}

func TestUnitGetNetworkSegmentByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetNetworkSegmentByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "network segment ID must be a positive integer")
}

func TestUnitGetNetworkSegmentByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetNetworkSegmentByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "network segment ID must be a positive integer")
}

func TestUnitGetNetworkSegmentByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetNetworkSegmentByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

// =============================================================================
// GetNetworkSegmentByName
// =============================================================================

func TestUnitGetNetworkSegmentByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetNetworkSegmentByNameMock()

	result, resp, err := svc.GetNetworkSegmentByName(context.Background(), "HQ Network")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "HQ Network", result.Name)
}

func TestUnitGetNetworkSegmentByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetNetworkSegmentByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "network segment name is required")
}

// =============================================================================
// CreateNetworkSegment
// =============================================================================

func TestUnitCreateNetworkSegment_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateNetworkSegmentMock()

	req := &RequestNetworkSegment{
		Name:            "New Segment",
		StartingAddress: "172.16.0.0",
		EndingAddress:   "172.16.0.255",
	}
	result, resp, err := svc.CreateNetworkSegment(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 3, result.ID)
}

func TestUnitCreateNetworkSegment_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateNetworkSegment(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateNetworkSegment_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestNetworkSegment{Name: "HQ Network", StartingAddress: "10.0.0.0", EndingAddress: "10.0.0.255"}
	result, resp, err := svc.CreateNetworkSegment(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

// =============================================================================
// UpdateNetworkSegmentByID
// =============================================================================

func TestUnitUpdateNetworkSegmentByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateNetworkSegmentByIDMock()

	req := &RequestNetworkSegment{
		Name:            "HQ Network Updated",
		StartingAddress: "10.0.0.0",
		EndingAddress:   "10.0.0.255",
	}
	result, resp, err := svc.UpdateNetworkSegmentByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitUpdateNetworkSegmentByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateNetworkSegmentByID(context.Background(), 0, &RequestNetworkSegment{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "network segment ID must be a positive integer")
}

func TestUnitUpdateNetworkSegmentByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateNetworkSegmentByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateNetworkSegmentByName
// =============================================================================

func TestUnitUpdateNetworkSegmentByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateNetworkSegmentByNameMock()

	req := &RequestNetworkSegment{
		Name:            "HQ Network Updated",
		StartingAddress: "10.0.0.0",
		EndingAddress:   "10.0.0.255",
	}
	result, resp, err := svc.UpdateNetworkSegmentByName(context.Background(), "HQ Network", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitUpdateNetworkSegmentByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateNetworkSegmentByName(context.Background(), "", &RequestNetworkSegment{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "network segment name is required")
}

func TestUnitUpdateNetworkSegmentByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateNetworkSegmentByName(context.Background(), "HQ Network", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteNetworkSegmentByID
// =============================================================================

func TestUnitDeleteNetworkSegmentByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteNetworkSegmentByIDMock()

	resp, err := svc.DeleteNetworkSegmentByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteNetworkSegmentByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteNetworkSegmentByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "network segment ID must be a positive integer")
}

// =============================================================================
// DeleteNetworkSegmentByName
// =============================================================================

func TestUnitDeleteNetworkSegmentByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteNetworkSegmentByNameMock()

	resp, err := svc.DeleteNetworkSegmentByName(context.Background(), "HQ Network")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteNetworkSegmentByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteNetworkSegmentByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "network segment name is required")
}
