package volume_purchasing_locations

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/volume_purchasing_locations/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*VolumePurchasingLocations, *mocks.VolumePurchasingLocationsMock) {
	t.Helper()
	mock := mocks.NewVolumePurchasingLocationsMock()
	return NewVolumePurchasingLocations(mock), mock
}

func TestUnit_VolumePurchasingLocations_List_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "VPL One", result.Results[0].Name)
}

func TestUnit_VolumePurchasingLocations_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "VPL One", result.Name)
	assert.True(t, result.AutomaticallyPopulatePurchasedContent)
}

func TestUnit_VolumePurchasingLocations_GetByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_VolumePurchasingLocations_Create_Success(t *testing.T) {
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

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "2", result.ID)
	assert.NotEmpty(t, result.Href)
}

func TestUnit_VolumePurchasingLocations_UpdateByID_Success(t *testing.T) {
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

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "VPL One Updated", result.Name)
}

func TestUnit_VolumePurchasingLocations_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteMock()

	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_VolumePurchasingLocations_ReclaimByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterReclaimMock()

	resp, err := svc.ReclaimVolumePurchasingLocationByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 202, resp.StatusCode())
}

func TestUnit_VolumePurchasingLocations_GetContentV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetContentMock()

	result, resp, err := svc.GetContentV1(context.Background(), "1", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 0, result.TotalCount)
	assert.NotNil(t, result.Results)
}

func TestUnit_VolumePurchasingLocations_GetHistoryV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryMock()

	result, resp, err := svc.GetHistoryV1(context.Background(), "1", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "admin", result.Results[0].Username)
}

func TestUnit_VolumePurchasingLocations_AddHistoryNotesV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddHistoryNotesMock()

	req := &AddHistoryNotesRequest{ObjectHistoryNote: "Test note"}
	resp, err := svc.AddHistoryNotesV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
}

func TestUnit_VolumePurchasingLocations_AddHistoryNotesV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &AddHistoryNotesRequest{ObjectHistoryNote: "Test note"}
	resp, err := svc.AddHistoryNotesV1(context.Background(), "", req)
	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_VolumePurchasingLocations_AddHistoryNotesV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.AddHistoryNotesV1(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_VolumePurchasingLocations_RevokeVolumePurchasingLocationLicensesByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterRevokeLicensesMock()

	resp, err := svc.RevokeVolumePurchasingLocationLicensesByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_VolumePurchasingLocations_CreateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_VolumePurchasingLocations_UpdateByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestVolumePurchasingLocation{Name: "Test"}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "", req)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_VolumePurchasingLocations_UpdateByIDV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_VolumePurchasingLocations_DeleteByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByIDV1(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "ID is required")
}

func TestUnit_VolumePurchasingLocations_ReclaimByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.ReclaimVolumePurchasingLocationByIDV1(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "ID is required")
}

func TestUnit_VolumePurchasingLocations_GetContentV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetContentV1(context.Background(), "", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "ID is required")
}

func TestUnit_VolumePurchasingLocations_GetHistoryV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetHistoryV1(context.Background(), "", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "ID is required")
}

func TestUnit_VolumePurchasingLocations_RevokeByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.RevokeVolumePurchasingLocationLicensesByIDV1(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "ID is required")
}

func TestUnit_VolumePurchasingLocations_GetContentV1_WithResults(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetContentWithResultsMock()

	result, resp, err := svc.GetContentV1(context.Background(), "2", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "Test App", result.Results[0].Name)
}

func TestUnit_VolumePurchasingLocations_ListV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListNoResponseErrorMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

func TestUnit_VolumePurchasingLocations_ListV1_InvalidJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListInvalidJSONMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "mergePage failed")
}

func TestUnit_VolumePurchasingLocations_GetByIDV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "999")
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_VolumePurchasingLocations_GetByIDV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDNoResponseErrorMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

func TestUnit_VolumePurchasingLocations_CreateV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateNoResponseErrorMock()

	req := &RequestVolumePurchasingLocation{Name: "New", ServiceToken: "token"}
	result, resp, err := svc.CreateV1(context.Background(), req)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

func TestUnit_VolumePurchasingLocations_UpdateByIDV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateNoResponseErrorMock()

	req := &RequestVolumePurchasingLocation{Name: "Updated", ServiceToken: "token"}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", req)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

func TestUnit_VolumePurchasingLocations_DeleteByIDV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteNoResponseErrorMock()

	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.Error(t, err)
	assert.NotNil(t, resp)
}

func TestUnit_VolumePurchasingLocations_ReclaimByIDV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterReclaimNoResponseErrorMock()

	resp, err := svc.ReclaimVolumePurchasingLocationByIDV1(context.Background(), "1")
	require.Error(t, err)
	assert.NotNil(t, resp)
}

func TestUnit_VolumePurchasingLocations_GetContentV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetContentNoResponseErrorMock()

	result, resp, err := svc.GetContentV1(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

func TestUnit_VolumePurchasingLocations_GetContentV1_InvalidJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetContentInvalidJSONMock()

	result, resp, err := svc.GetContentV1(context.Background(), "99", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "mergePage failed")
}

func TestUnit_VolumePurchasingLocations_GetHistoryV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryNoResponseErrorMock()

	result, resp, err := svc.GetHistoryV1(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

func TestUnit_VolumePurchasingLocations_GetHistoryV1_InvalidJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryInvalidJSONMock()

	result, resp, err := svc.GetHistoryV1(context.Background(), "99", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "mergePage failed")
}

func TestUnit_VolumePurchasingLocations_AddHistoryNotesV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddHistoryNotesNoResponseErrorMock()

	req := &AddHistoryNotesRequest{ObjectHistoryNote: "Note"}
	resp, err := svc.AddHistoryNotesV1(context.Background(), "1", req)
	require.Error(t, err)
	assert.NotNil(t, resp)
}

func TestUnit_VolumePurchasingLocations_RevokeLicensesV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterRevokeLicensesNoResponseErrorMock()

	resp, err := svc.RevokeVolumePurchasingLocationLicensesByIDV1(context.Background(), "1")
	require.Error(t, err)
	assert.NotNil(t, resp)
}
