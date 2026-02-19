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

func TestUnitListV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	require.Equal(t, "1", result.Results[0].ID)
	require.Equal(t, "All iPhones", result.Results[0].Name)
}

func TestUnitGetByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "1", result.ID)
	require.Equal(t, "All iPhones", result.Name)
}

func TestUnitGetByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetByIDV1(context.Background(), "")
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
}

func TestUnitCreateV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	search := &ResourceAdvancedMobileDeviceSearch{
		Name:          "Test Search",
		Criteria:      []CriteriaJamfProAPI{{Name: "Device Name", Priority: 1, AndOr: "and", SearchType: "like", Value: "iPhone"}},
		DisplayFields: []string{"Device Name"},
	}
	result, resp, err := svc.CreateV1(context.Background(), search)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Contains(t, []int{200, 201}, resp.StatusCode)
	require.Equal(t, "2", result.ID)
}

func TestUnitCreateV1_NilSearch(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.CreateV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
}

func TestUnitDeleteByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 204, resp.StatusCode)
}

func TestUnitGetChoicesV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetChoicesV1(context.Background(), "Device Name", "-1", "")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Len(t, result.Choices, 3)
	require.Equal(t, "iPhone", result.Choices[0])
}
