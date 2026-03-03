package ebooks

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/ebooks/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.EbooksMock) {
	t.Helper()
	mock := mocks.NewEbooksMock()
	return NewService(mock), mock
}

func TestUnit_Ebooks_NewService(t *testing.T) {
	mock := mocks.NewEbooksMock()
	svc := NewService(mock)
	require.NotNil(t, svc)
}

func TestUnit_Ebooks_ListV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListEbooksMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Getting Started with Jamf Pro", result.Results[0].Name)
	assert.Equal(t, "PDF", result.Results[0].Kind)
	assert.Equal(t, "https://example.com/ebook1.pdf", result.Results[0].URL)
	assert.True(t, result.Results[0].Free)
	assert.Equal(t, "1.0", result.Results[0].Version)
	assert.Equal(t, "Jamf", result.Results[0].Author)
	assert.Equal(t, "2", result.Results[1].ID)
	assert.Equal(t, "Advanced Deployment", result.Results[1].Name)
	assert.Equal(t, "EPUB", result.Results[1].Kind)
}

func TestUnit_Ebooks_ListV1_WithRSQLFilter(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListEbooksMock()

	rsqlQuery := map[string]string{
		"filter": `name=="Getting Started with Jamf Pro"`,
		"sort":   "name:asc",
	}

	result, resp, err := svc.ListV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, rsqlQuery, mock.LastRSQLQuery)
}

func TestUnit_Ebooks_ListV1_WithPagination(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListEbooksMock()

	rsqlQuery := map[string]string{
		"page":      "0",
		"page-size": "50",
		"sort":      "name:desc",
	}

	result, resp, err := svc.ListV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "0", mock.LastRSQLQuery["page"])
	assert.Equal(t, "50", mock.LastRSQLQuery["page-size"])
}

func TestUnit_Ebooks_ListV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.ListV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to list ebooks")
}

func TestUnit_Ebooks_ListV1_InvalidJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterInvalidJSONMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "mergePage failed")
}

func TestUnit_Ebooks_GetByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetEbookMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Getting Started with Jamf Pro", result.Name)
	assert.Equal(t, "PDF", result.Kind)
	assert.Equal(t, "https://example.com/ebook1.pdf", result.URL)
	assert.True(t, result.Free)
	assert.Equal(t, "1.0", result.Version)
	assert.Equal(t, "Jamf", result.Author)
	assert.False(t, result.DeployAsManaged)
	assert.False(t, result.InstallAutomatically)
	assert.Equal(t, "1", result.CategoryID)
	assert.Equal(t, "2", result.SiteID)
}

func TestUnit_Ebooks_GetByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "ebook ID is required")
}

func TestUnit_Ebooks_GetByIDV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestUnit_Ebooks_GetByIDV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_Ebooks_GetScopeByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetEbookScopeMock()

	result, resp, err := svc.GetScopeByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.False(t, result.AllComputers)
	assert.True(t, result.AllMobileDevices)
	assert.False(t, result.AllUsers)
	assert.Equal(t, []string{"1", "2"}, result.ComputerIDs)
	assert.Equal(t, []string{"1"}, result.ComputerGroupIDs)
	assert.Equal(t, []string{"10", "11"}, result.MobileDeviceIDs)
	assert.Equal(t, []string{"2"}, result.MobileDeviceGroupIDs)
	assert.Equal(t, []string{"1"}, result.BuildingIDs)
	assert.Equal(t, []string{"1"}, result.DepartmentIDs)
	assert.Equal(t, []string{"1", "2"}, result.UserIDs)
	assert.Equal(t, []string{"1"}, result.UserGroupIDs)
	require.NotNil(t, result.Limitations)
	assert.Equal(t, []string{"1"}, result.Limitations.NetworkSegments)
	require.Len(t, result.Limitations.Users, 1)
	assert.Equal(t, "admin", result.Limitations.Users[0].Name)
	assert.Equal(t, []string{"admins"}, result.Limitations.UserGroups)
	require.NotNil(t, result.Exclusions)
	assert.Equal(t, []string{"99"}, result.Exclusions.ComputerIDs)
	assert.Equal(t, []string{"99"}, result.Exclusions.UserIDs)
}

func TestUnit_Ebooks_GetScopeByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetScopeByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "ebook ID is required")
}

func TestUnit_Ebooks_GetScopeByIDV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetScopeByIDV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestUnit_Ebooks_GetScopeByIDV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetScopeByIDV1(context.Background(), "1")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}
