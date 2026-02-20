package ldap_servers

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/ldap_servers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh LDAPServersMock.
func setupMockService(t *testing.T) (*Service, *mocks.LDAPServersMock) {
	t.Helper()
	mock := mocks.NewLDAPServersMock()
	return NewService(mock), mock
}

// =============================================================================
// ListLDAPServers
// =============================================================================

func TestUnitListLDAPServers_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLDAPServersMock()

	result, resp, err := svc.ListLDAPServers(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.Size)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "Test LDAP Server 1", result.Results[0].Name)
	assert.Equal(t, 2, result.Results[1].ID)
	assert.Equal(t, "Test LDAP Server 2", result.Results[1].Name)
}

// =============================================================================
// GetLDAPServerByID
// =============================================================================

func TestUnitGetLDAPServerByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetLDAPServerByIDMock()

	result, resp, err := svc.GetLDAPServerByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.Connection.ID)
	assert.Equal(t, "Test LDAP Server", result.Connection.Name)
	assert.Equal(t, "ldap.example.com", result.Connection.Hostname)
	assert.Equal(t, 389, result.Connection.Port)
}

func TestUnitGetLDAPServerByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetLDAPServerByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "LDAP server ID must be a positive integer")
}

func TestUnitGetLDAPServerByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetLDAPServerByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "LDAP server ID must be a positive integer")
}

func TestUnitGetLDAPServerByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetLDAPServerByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

// =============================================================================
// GetLDAPServerByName
// =============================================================================

func TestUnitGetLDAPServerByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetLDAPServerByNameMock()

	result, resp, err := svc.GetLDAPServerByName(context.Background(), "Test LDAP Server")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.Connection.ID)
	assert.Equal(t, "Test LDAP Server", result.Connection.Name)
	assert.Equal(t, "ldap.example.com", result.Connection.Hostname)
}

func TestUnitGetLDAPServerByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetLDAPServerByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "LDAP server name is required")
}

// =============================================================================
// CreateLDAPServer
// =============================================================================

func TestUnitCreateLDAPServer_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateLDAPServerMock()

	req := &RequestLDAPServer{
		Connection: RequestConnection{
			Name:     "Test LDAP Server",
			Hostname: "ldap.example.com",
			Port:     389,
		},
	}
	result, resp, err := svc.CreateLDAPServer(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 100, result.ID)
	assert.Equal(t, "Test LDAP Server", result.Name)
}

func TestUnitCreateLDAPServer_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateLDAPServer(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateLDAPServer_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestLDAPServer{
		Connection: RequestConnection{
			Name:     "Test LDAP Server",
			Hostname: "ldap.example.com",
			Port:     389,
		},
	}
	result, resp, err := svc.CreateLDAPServer(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

// =============================================================================
// UpdateLDAPServerByID
// =============================================================================

func TestUnitUpdateLDAPServerByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateLDAPServerByIDMock()

	req := &RequestLDAPServer{
		Connection: RequestConnection{
			Name:     "Updated LDAP Server",
			Hostname: "ldap2.example.com",
			Port:     636,
		},
	}
	result, resp, err := svc.UpdateLDAPServerByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.Connection.ID)
	assert.Equal(t, "Updated LDAP Server", result.Connection.Name)
	assert.Equal(t, "ldap2.example.com", result.Connection.Hostname)
}

func TestUnitUpdateLDAPServerByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestLDAPServer{
		Connection: RequestConnection{
			Name: "Updated LDAP Server",
		},
	}
	result, resp, err := svc.UpdateLDAPServerByID(context.Background(), 0, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "LDAP server ID must be a positive integer")
}

func TestUnitUpdateLDAPServerByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateLDAPServerByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateLDAPServerByName
// =============================================================================

func TestUnitUpdateLDAPServerByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateLDAPServerByNameMock()

	req := &RequestLDAPServer{
		Connection: RequestConnection{
			Name:     "Updated LDAP Server",
			Hostname: "ldap2.example.com",
			Port:     636,
		},
	}
	result, resp, err := svc.UpdateLDAPServerByName(context.Background(), "Test LDAP Server", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.Connection.ID)
	assert.Equal(t, "Updated LDAP Server", result.Connection.Name)
}

func TestUnitUpdateLDAPServerByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestLDAPServer{
		Connection: RequestConnection{
			Name: "Updated LDAP Server",
		},
	}
	result, resp, err := svc.UpdateLDAPServerByName(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "LDAP server name is required")
}

func TestUnitUpdateLDAPServerByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateLDAPServerByName(context.Background(), "Test LDAP Server", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteLDAPServerByID
// =============================================================================

func TestUnitDeleteLDAPServerByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteLDAPServerByIDMock()

	resp, err := svc.DeleteLDAPServerByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteLDAPServerByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteLDAPServerByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "LDAP server ID must be a positive integer")
}

func TestUnitDeleteLDAPServerByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteLDAPServerByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "LDAP server ID must be a positive integer")
}

// =============================================================================
// DeleteLDAPServerByName
// =============================================================================

func TestUnitDeleteLDAPServerByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteLDAPServerByNameMock()

	resp, err := svc.DeleteLDAPServerByName(context.Background(), "Test LDAP Server")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteLDAPServerByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteLDAPServerByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "LDAP server name is required")
}
