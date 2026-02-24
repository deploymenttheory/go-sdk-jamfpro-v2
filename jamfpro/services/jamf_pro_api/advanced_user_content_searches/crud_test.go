package advanced_user_content_searches

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/advanced_user_content_searches/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.AdvancedUserContentSearchesMock) {
	t.Helper()
	mock := mocks.NewAdvancedUserContentSearchesMock()
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
	require.Equal(t, "Andy's Search", result.Results[0].Name)
}

func TestUnitGetByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "1", result.ID)
	require.Equal(t, "Andy's Search", result.Name)
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
	search := &ResourceAdvancedUserContentSearch{
		Name:          "Test Search",
		Criteria:      []CriteriaJamfProAPI{{Name: "Account", Priority: 1, AndOr: "and", SearchType: "like", Value: "test"}},
		DisplayFields: []string{"Content Name"},
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

func TestUnitUpdateByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	search := &ResourceAdvancedUserContentSearch{
		Name:          "Updated Search",
		Criteria:      []CriteriaJamfProAPI{{Name: "Account", Priority: 0, AndOr: "and", SearchType: "is", Value: "updated"}},
		DisplayFields: []string{"Content Name", "Price"},
	}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", search)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "1", result.ID)
}

func TestUnitDeleteByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 204, resp.StatusCode)
}
