package volume_purchasing_locations

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/volume_purchasing_locations/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.VolumePurchasingLocationsMock) {
	t.Helper()
	mock := mocks.NewVolumePurchasingLocationsMock()
	return NewService(mock), mock
}

func TestUnitListVolumePurchasingLocations_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "VPL One", result.Results[0].Name)
}

func TestUnitGetVolumePurchasingLocationByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "VPL One", result.Name)
	assert.True(t, result.AutomaticallyPopulatePurchasedContent)
}

func TestUnitGetVolumePurchasingLocationByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitCreateVolumePurchasingLocation_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	req := &RequestVolumePurchasingLocation{
		Name:                                  "New VPL",
		ServiceToken:                          "token",
		AutomaticallyPopulatePurchasedContent: true,
	}
	result, resp, err := svc.CreateV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "2", result.ID)
	assert.NotEmpty(t, result.Href)
}

func TestUnitUpdateVolumePurchasingLocationByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	req := &RequestVolumePurchasingLocation{
		Name:                                 "VPL One Updated",
		ServiceToken:                         "token",
		AutomaticallyPopulatePurchasedContent: false,
		SendNotificationWhenNoLongerAssigned:  true,
	}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "VPL One Updated", result.Name)
}

func TestUnitDeleteVolumePurchasingLocationByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteMock()

	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitReclaimVolumePurchasingLocationByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterReclaimMock()

	resp, err := svc.ReclaimVolumePurchasingLocationByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 202, resp.StatusCode)
}

func TestUnitGetContentV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetContentMock()

	result, resp, err := svc.GetContentV1(context.Background(), "1", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 0, result.TotalCount)
	assert.NotNil(t, result.Results)
}
