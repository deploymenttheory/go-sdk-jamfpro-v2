package return_to_service

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/return_to_service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.ReturnToServiceMock) {
	t.Helper()
	mock := mocks.NewReturnToServiceMock()
	return NewService(mock), mock
}

func TestUnitListV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()
	result, resp, err := svc.ListV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Default", result.Results[0].DisplayName)
}

func TestUnitGetByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetByIDV1(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitCreateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.CreateV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitDeleteByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteByIDV1(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, resp)
}
