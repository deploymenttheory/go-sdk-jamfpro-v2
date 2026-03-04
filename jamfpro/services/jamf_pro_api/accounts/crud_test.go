package accounts

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/accounts/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.AccountsMock) {
	t.Helper()
	mock := mocks.NewAccountsMock()
	return NewService(mock), mock
}

// Test ListV1 with success response
func TestUnit_Accounts_ListV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListAccountsMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	assert.Len(t, result.Results, 2)
	assert.Equal(t, "testuser1", result.Results[0].Username)
	assert.Equal(t, "testuser2", result.Results[1].Username)
}

// Test ListV1 with RSQL filter query
func TestUnit_Accounts_ListV1_WithRSQLFilter(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListAccountsMock()

	rsqlQuery := map[string]string{
		"filter": `username=="testuser1"`,
		"sort":   "username:asc",
	}

	result, resp, err := svc.ListV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())

	// Verify RSQL query was passed to the client
	assert.NotNil(t, mock.LastRSQLQuery)
	assert.Equal(t, `username=="testuser1"`, mock.LastRSQLQuery["filter"])
	assert.Equal(t, "username:asc", mock.LastRSQLQuery["sort"])
}

// Test ListV1 with pagination parameters
func TestUnit_Accounts_ListV1_WithPagination(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListAccountsMock()

	rsqlQuery := map[string]string{
		"page":      "0",
		"page-size": "50",
		"sort":      "realname:desc",
	}

	result, resp, err := svc.ListV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())

	// Verify pagination parameters were passed
	assert.Equal(t, "0", mock.LastRSQLQuery["page"])
	assert.Equal(t, "50", mock.LastRSQLQuery["page-size"])
	assert.Equal(t, "realname:desc", mock.LastRSQLQuery["sort"])
}

// Test ListV1 with complex RSQL filter
func TestUnit_Accounts_ListV1_WithComplexRSQLFilter(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListAccountsMock()

	rsqlQuery := map[string]string{
		"filter": `accountStatus==Enabled and privilegeLevel==ADMINISTRATOR and failedLoginAttempts==0`,
	}

	result, resp, err := svc.ListV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())

	// Verify complex RSQL query was passed
	assert.Equal(t, `accountStatus==Enabled and privilegeLevel==ADMINISTRATOR and failedLoginAttempts==0`, mock.LastRSQLQuery["filter"])
}

// Test GetByIDV1 with success response
func TestUnit_Accounts_GetByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAccountMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "testuser1", result.Username)
	assert.Equal(t, "Test User One", result.Realname)
	assert.Equal(t, "testuser1@example.com", result.Email)
	assert.Equal(t, "FullAccess", result.AccessLevel)
	assert.Equal(t, "ADMINISTRATOR", result.PrivilegeLevel)
	assert.Equal(t, "Enabled", result.AccountStatus)
}

// Test GetByIDV1 with empty ID
func TestUnit_Accounts_GetByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account ID is required")
}

// Test CreateV1 with success response
func TestUnit_Accounts_CreateV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateAccountMock()

	req := &RequestAccount{
		Username:       "newuser",
		Realname:       "New User",
		Email:          "newuser@example.com",
		Phone:          "555-0003",
		PlainPassword:  "securepassword123",
		AccessLevel:    "FullAccess",
		PrivilegeLevel: "ADMINISTRATOR",
		AccountStatus:  "Enabled",
		AccountType:    "DEFAULT",
	}

	result, resp, err := svc.CreateV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "3", result.ID)
	assert.Equal(t, "newuser", result.Username)
}

// Test CreateV1 with nil request
func TestUnit_Accounts_CreateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// Test DeleteByIDV1 with success response
func TestUnit_Accounts_DeleteByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteAccountMock()

	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

// Test DeleteByIDV1 with empty ID
func TestUnit_Accounts_DeleteByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account ID is required")
}

func TestUnit_Accounts_ListV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.ListV1(context.Background(), nil)
	require.Error(t, err)
}

func TestUnit_Accounts_GetByIDV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetByIDV1(context.Background(), "1")
	require.Error(t, err)
}

func TestUnit_Accounts_CreateV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.CreateV1(context.Background(), &RequestAccount{Username: "user"})
	require.Error(t, err)
}

func TestUnit_Accounts_DeleteByIDV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.DeleteByIDV1(context.Background(), "1")
	require.Error(t, err)
}
