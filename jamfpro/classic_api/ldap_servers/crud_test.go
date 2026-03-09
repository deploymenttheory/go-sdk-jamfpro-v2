package ldap_servers

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/ldap_servers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh LDAPServersMock.
func setupMockService(t *testing.T) (*LdapServers, *mocks.LDAPServersMock) {
	t.Helper()
	mock := mocks.NewLDAPServersMock()
	return NewLdapServers(mock), mock
}

// =============================================================================
// ListLDAPServers
// =============================================================================

func TestUnit_LdapServers_List_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLDAPServersMock()

	result, resp, err := svc.List(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
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

func TestUnit_LdapServers_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetLDAPServerByIDMock()

	result, resp, err := svc.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.Connection.ID)
	assert.Equal(t, "Test LDAP Server", result.Connection.Name)
	assert.Equal(t, "ldap.example.com", result.Connection.Hostname)
	assert.Equal(t, 389, result.Connection.Port)
}

func TestUnit_LdapServers_GetByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "LDAP server ID must be a positive integer")
}

func TestUnit_LdapServers_GetByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "LDAP server ID must be a positive integer")
}

func TestUnit_LdapServers_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

// =============================================================================
// GetLDAPServerByName
// =============================================================================

func TestUnit_LdapServers_GetByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetLDAPServerByNameMock()

	result, resp, err := svc.GetByName(context.Background(), "Test LDAP Server")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.Connection.ID)
	assert.Equal(t, "Test LDAP Server", result.Connection.Name)
	assert.Equal(t, "ldap.example.com", result.Connection.Hostname)
}

func TestUnit_LdapServers_GetByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "LDAP server name is required")
}

// =============================================================================
// CreateLDAPServer
// =============================================================================

func TestUnit_LdapServers_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateLDAPServerMock()

	req := &RequestLDAPServer{
		Connection: RequestConnection{
			Name:     "Test LDAP Server",
			Hostname: "ldap.example.com",
			Port:     389,
		},
	}
	result, resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, 100, result.ID)
	assert.Equal(t, "Test LDAP Server", result.Name)
}

func TestUnit_LdapServers_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.Create(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_LdapServers_Create_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestLDAPServer{
		Connection: RequestConnection{
			Name:     "Test LDAP Server",
			Hostname: "ldap.example.com",
			Port:     389,
		},
	}
	result, resp, err := svc.Create(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode())
}

// =============================================================================
// UpdateLDAPServerByID
// =============================================================================

func TestUnit_LdapServers_UpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateLDAPServerByIDMock()

	req := &RequestLDAPServer{
		Connection: RequestConnection{
			Name:     "Updated LDAP Server",
			Hostname: "ldap2.example.com",
			Port:     636,
		},
	}
	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.Connection.ID)
	assert.Equal(t, "Updated LDAP Server", result.Connection.Name)
	assert.Equal(t, "ldap2.example.com", result.Connection.Hostname)
}

func TestUnit_LdapServers_UpdateByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestLDAPServer{
		Connection: RequestConnection{
			Name: "Updated LDAP Server",
		},
	}
	result, resp, err := svc.UpdateByID(context.Background(), 0, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "LDAP server ID must be a positive integer")
}

func TestUnit_LdapServers_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateLDAPServerByName
// =============================================================================

func TestUnit_LdapServers_UpdateByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateLDAPServerByNameMock()

	req := &RequestLDAPServer{
		Connection: RequestConnection{
			Name:     "Updated LDAP Server",
			Hostname: "ldap2.example.com",
			Port:     636,
		},
	}
	result, resp, err := svc.UpdateByName(context.Background(), "Test LDAP Server", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.Connection.ID)
	assert.Equal(t, "Updated LDAP Server", result.Connection.Name)
}

func TestUnit_LdapServers_UpdateByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestLDAPServer{
		Connection: RequestConnection{
			Name: "Updated LDAP Server",
		},
	}
	result, resp, err := svc.UpdateByName(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "LDAP server name is required")
}

func TestUnit_LdapServers_UpdateByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByName(context.Background(), "Test LDAP Server", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteLDAPServerByID
// =============================================================================

func TestUnit_LdapServers_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteLDAPServerByIDMock()

	resp, err := svc.DeleteByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_LdapServers_DeleteByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "LDAP server ID must be a positive integer")
}

func TestUnit_LdapServers_DeleteByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "LDAP server ID must be a positive integer")
}

// =============================================================================
// DeleteLDAPServerByName
// =============================================================================

func TestUnit_LdapServers_DeleteByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteLDAPServerByNameMock()

	resp, err := svc.DeleteByName(context.Background(), "Test LDAP Server")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_LdapServers_DeleteByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "LDAP server name is required")
}
