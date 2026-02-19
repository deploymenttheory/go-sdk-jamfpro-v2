package mobile_device_groups

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mobile_device_groups/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.MobileDeviceGroupsMock) {
	t.Helper()
	mock := mocks.NewMobileDeviceGroupsMock()
	return NewService(mock), mock
}

func TestUnitListSmartV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListSmartMock()

	result, resp, err := svc.ListSmartV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "iPhones", result.Results[0].Name)
}

func TestUnitGetSmartByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetSmartMock()

	result, resp, err := svc.GetSmartByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "iPhones", result.Name)
	require.Len(t, result.Criteria, 1)
	assert.Equal(t, "Model", result.Criteria[0].Name)
}

func TestUnitCreateSmartV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateSmartMock()

	req := &RequestSmartMobileDeviceGroup{
		Name:        "New Smart",
		Description: "Desc",
		Criteria:    []CriteriaJamfProAPI{{Name: "Model", Priority: 0, AndOr: "and", SearchType: "is", Value: "iPhone"}},
	}
	result, resp, err := svc.CreateSmartV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "2", result.ID)
	assert.NotEmpty(t, result.Href)
}

func TestUnitUpdateSmartByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateSmartMock()

	req := &RequestSmartMobileDeviceGroup{
		Name:        "iPhones Updated",
		Description: "Updated",
		Criteria:    []CriteriaJamfProAPI{{Name: "Model", Priority: 0, AndOr: "and", SearchType: "is", Value: "iPhone"}},
	}
	result, resp, err := svc.UpdateSmartByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "iPhones Updated", result.Name)
}

func TestUnitDeleteSmartByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteSmartMock()

	resp, err := svc.DeleteSmartByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitListStaticV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListStaticMock()

	result, resp, err := svc.ListStaticV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "10", result.Results[0].ID)
	assert.Equal(t, "Static Devices", result.Results[0].Name)
}

func TestUnitGetStaticByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetStaticMock()

	result, resp, err := svc.GetStaticByIDV1(context.Background(), "10")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "10", result.ID)
	assert.Equal(t, "Static Devices", result.Name)
}

func TestUnitCreateStaticV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateStaticMock()

	req := &RequestStaticMobileDeviceGroup{
		Name:        "New Static",
		Description: "Desc",
		SiteId:      "-1",
		Assignments: []StaticMobileDeviceGroupAssignment{},
	}
	result, resp, err := svc.CreateStaticV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "11", result.ID)
	assert.NotEmpty(t, result.Href)
}

func TestUnitUpdateStaticByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateStaticMock()

	req := &RequestStaticMobileDeviceGroup{
		Name:        "Static Devices Updated",
		Description: "Updated",
		SiteId:      "-1",
		Assignments: []StaticMobileDeviceGroupAssignment{},
	}
	result, resp, err := svc.UpdateStaticByIDV1(context.Background(), "10", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "10", result.ID)
	assert.Equal(t, "Static Devices Updated", result.Name)
}

func TestUnitDeleteStaticByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteStaticMock()

	resp, err := svc.DeleteStaticByIDV1(context.Background(), "10")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}
