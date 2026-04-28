package account_groups

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/account_groups/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*AccountGroups, *mocks.AccountGroupsMock) {
	t.Helper()
	mock := mocks.NewAccountGroupsMock()
	return NewAccountGroups(mock), mock
}

// Test ListV1 success
func TestUnit_AccountGroups_ListV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListAccountGroupsMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	assert.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Administrators", result.Results[0].Name)
	assert.Equal(t, "FullAccess", result.Results[0].AccessLevel)
	assert.Equal(t, "ADMINISTRATOR", result.Results[0].PrivilegeLevel)
}

// Test ListV1 with RSQL filter
func TestUnit_AccountGroups_ListV1_WithRSQLFilter(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListAccountGroupsMock()

	rsqlQuery := map[string]string{
		"filter": `name=="Administrators"`,
	}

	result, resp, err := svc.ListV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, `name=="Administrators"`, mock.LastRSQLQuery["filter"])
}

// Test GetByIDV1 success
func TestUnit_AccountGroups_GetByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAccountGroupMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Administrators", result.Name)
	assert.Equal(t, "FullAccess", result.AccessLevel)
	assert.Equal(t, "ADMINISTRATOR", result.PrivilegeLevel)
	require.NotNil(t, result.Site)
	assert.Equal(t, -1, result.Site.ID)
}

// Test GetByIDV1 with empty ID
func TestUnit_AccountGroups_GetByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account group ID is required")
}

// Error tests (no mock registered)
func TestUnit_AccountGroups_ListV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.ListV1(context.Background(), nil)
	require.Error(t, err)
}

func TestUnit_AccountGroups_GetByIDV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetByIDV1(context.Background(), "1")
	require.Error(t, err)
}
