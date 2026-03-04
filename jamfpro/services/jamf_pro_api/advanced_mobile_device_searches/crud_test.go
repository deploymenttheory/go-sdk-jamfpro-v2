package advanced_mobile_device_searches

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/advanced_mobile_device_searches/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.AdvancedMobileDeviceSearchesMock) {
	t.Helper()
	mock := mocks.NewAdvancedMobileDeviceSearchesMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnit_AdvancedMobileDeviceSearches_ListV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode())
	require.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	require.Equal(t, "1", result.Results[0].ID)
	require.Equal(t, "All iPhones", result.Results[0].Name)
}

func TestUnit_AdvancedMobileDeviceSearches_GetByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode())
	require.Equal(t, "1", result.ID)
	require.Equal(t, "All iPhones", result.Name)
}

func TestUnit_AdvancedMobileDeviceSearches_GetByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetByIDV1(context.Background(), "")
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
}

func TestUnit_AdvancedMobileDeviceSearches_CreateV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	search := &ResourceAdvancedMobileDeviceSearch{
		Name:          "Test Search",
		Criteria:      []CriteriaJamfProAPI{{Name: "Device Name", Priority: 1, AndOr: "and", SearchType: "like", Value: "iPhone"}},
		DisplayFields: []string{"Device Name"},
	}
	result, resp, err := svc.CreateV1(context.Background(), search)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Contains(t, []int{200, 201}, resp.StatusCode())
	require.Equal(t, "2", result.ID)
}

func TestUnit_AdvancedMobileDeviceSearches_CreateV1_NilSearch(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.CreateV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
}

func TestUnit_AdvancedMobileDeviceSearches_DeleteByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 204, resp.StatusCode())
}

func TestUnit_AdvancedMobileDeviceSearches_DeleteMultipleV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &DeleteAdvancedMobileDeviceSearchesByIDRequest{IDs: []string{"1", "2"}}
	resp, err := svc.DeleteMultipleV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 204, resp.StatusCode())
}

func TestUnit_AdvancedMobileDeviceSearches_DeleteMultipleV1_EmptyIDs(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteMultipleV1(context.Background(), &DeleteAdvancedMobileDeviceSearchesByIDRequest{IDs: []string{}})
	require.Error(t, err)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "ids are required")
}

func TestUnit_AdvancedMobileDeviceSearches_DeleteMultipleV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteMultipleV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "ids are required")
}

func TestUnit_AdvancedMobileDeviceSearches_DeleteByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteByIDV1(context.Background(), "")
	require.Error(t, err)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "id is required")
}

func TestUnit_AdvancedMobileDeviceSearches_GetByIDV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()
	result, resp, err := svc.GetByIDV1(context.Background(), "999")
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 404, resp.StatusCode())
}

func TestUnit_AdvancedMobileDeviceSearches_UpdateByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	search := &ResourceAdvancedMobileDeviceSearch{
		Name:          "All iPhones Updated",
		Criteria:      []CriteriaJamfProAPI{{Name: "Device Name", Priority: 1, AndOr: "and", SearchType: "like", Value: "iPhone"}},
		DisplayFields: []string{"Device Name"},
	}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", search)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode())
}

func TestUnit_AdvancedMobileDeviceSearches_UpdateByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	search := &ResourceAdvancedMobileDeviceSearch{Name: "Test"}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "", search)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "id is required")
}

func TestUnit_AdvancedMobileDeviceSearches_UpdateByIDV1_NilSearch(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "search is required")
}

func TestUnit_AdvancedMobileDeviceSearches_GetChoicesV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetChoicesV1(context.Background(), "Device Name", "-1", "")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode())
	require.Len(t, result.Choices, 3)
	require.Equal(t, "iPhone", result.Choices[0])
}
