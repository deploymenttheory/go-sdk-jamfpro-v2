package device_enrollments

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/device_enrollments/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListV1(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	mocks.RegisterListMock(mock)

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.ListV1(ctx, nil)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Example Device Enrollment Instance", result.Results[0].Name)
}

func TestListV1_WithPagination(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	mocks.RegisterListMock(mock)

	svc := NewService(mock)
	ctx := context.Background()

	rsqlQuery := map[string]string{
		"page":      "0",
		"page-size": "100",
		"sort":      "id:asc",
	}

	result, resp, err := svc.ListV1(ctx, rsqlQuery)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
}

func TestGetByIDV1(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	mocks.RegisterGetByIDMock(mock)

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByIDV1(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Example Device Enrollment Instance", result.Name)
}

func TestGetByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	svc := NewService(mock)
	ctx := context.Background()

	_, _, err := svc.GetByIDV1(ctx, "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "id is required")
}

func TestGetByNameV1(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	mocks.RegisterListMock(mock)

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByNameV1(ctx, "Example Device Enrollment Instance")

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Example Device Enrollment Instance", result.Name)
}

func TestGetByNameV1_EmptyName(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	svc := NewService(mock)
	ctx := context.Background()

	_, _, err := svc.GetByNameV1(ctx, "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "name is required")
}

func TestGetByNameV1_NotFound(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	mocks.RegisterListMock(mock)

	svc := NewService(mock)
	ctx := context.Background()

	_, _, err := svc.GetByNameV1(ctx, "NonExistent")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestGetHistoryV1(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	mocks.RegisterGetHistoryMock(mock)

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetHistoryV1(ctx, "1", nil)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, 1, result.Results[0].ID)
}

func TestGetHistoryV1_EmptyID(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	svc := NewService(mock)
	ctx := context.Background()

	_, _, err := svc.GetHistoryV1(ctx, "", nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "id is required")
}

func TestGetSyncStatesV1(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	mocks.RegisterGetSyncStatesMock(mock)

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetSyncStatesV1(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, result, 1)
	assert.Equal(t, "CONNECTION_ERROR", result[0].SyncState)
	assert.Equal(t, "1", result[0].InstanceID)
}

func TestGetSyncStatesV1_EmptyID(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	svc := NewService(mock)
	ctx := context.Background()

	_, _, err := svc.GetSyncStatesV1(ctx, "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "id is required")
}

func TestGetLatestSyncStateV1(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	mocks.RegisterGetLatestSyncStateMock(mock)

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetLatestSyncStateV1(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "CONNECTION_ERROR", result.SyncState)
	assert.Equal(t, "1", result.InstanceID)
}

func TestGetLatestSyncStateV1_EmptyID(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	svc := NewService(mock)
	ctx := context.Background()

	_, _, err := svc.GetLatestSyncStateV1(ctx, "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "id is required")
}

func TestGetAllSyncStatesV1(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	mocks.RegisterGetAllSyncStatesMock(mock)

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetAllSyncStatesV1(ctx)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, result, 1)
}

func TestGetPublicKeyV1(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	mocks.RegisterGetPublicKeyMock(mock)

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetPublicKeyV1(ctx)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Contains(t, string(result), "BEGIN PUBLIC KEY")
}

func TestCreateWithTokenV1(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	mocks.RegisterCreateWithTokenMock(mock)

	svc := NewService(mock)
	ctx := context.Background()

	req := &RequestTokenUpload{
		TokenFileName: "test-token.p7m",
		EncodedToken:  "base64encodedtoken==",
	}

	result, resp, err := svc.CreateWithTokenV1(ctx, req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.NotEmpty(t, result.Href)
}

func TestCreateWithTokenV1_NilRequest(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	svc := NewService(mock)
	ctx := context.Background()

	_, _, err := svc.CreateWithTokenV1(ctx, nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestCreateWithTokenV1_EmptyToken(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	svc := NewService(mock)
	ctx := context.Background()

	req := &RequestTokenUpload{
		TokenFileName: "test.p7m",
		EncodedToken:  "",
	}

	_, _, err := svc.CreateWithTokenV1(ctx, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "encodedToken is required")
}

func TestUpdateByIDV1(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	mocks.RegisterUpdateByIDMock(mock)

	svc := NewService(mock)
	ctx := context.Background()

	req := &RequestUpdate{
		Name:                  "Updated Name",
		SupervisionIdentityId: "2",
		SiteId:                "1",
	}

	result, resp, err := svc.UpdateByIDV1(ctx, "1", req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
}

func TestUpdateByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	svc := NewService(mock)
	ctx := context.Background()

	req := &RequestUpdate{Name: "Test"}

	_, _, err := svc.UpdateByIDV1(ctx, "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUpdateByIDV1_NilRequest(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	svc := NewService(mock)
	ctx := context.Background()

	_, _, err := svc.UpdateByIDV1(ctx, "1", nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUpdateByIDV1_EmptyName(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	svc := NewService(mock)
	ctx := context.Background()

	req := &RequestUpdate{Name: ""}

	_, _, err := svc.UpdateByIDV1(ctx, "1", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "name is required")
}

func TestUpdateTokenByIDV1(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	mocks.RegisterUpdateTokenByIDMock(mock)

	svc := NewService(mock)
	ctx := context.Background()

	req := &RequestTokenUpload{
		EncodedToken: "newbase64token==",
	}

	result, resp, err := svc.UpdateTokenByIDV1(ctx, "1", req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
}

func TestUpdateTokenByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	svc := NewService(mock)
	ctx := context.Background()

	req := &RequestTokenUpload{EncodedToken: "token"}

	_, _, err := svc.UpdateTokenByIDV1(ctx, "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUpdateTokenByIDV1_NilRequest(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	svc := NewService(mock)
	ctx := context.Background()

	_, _, err := svc.UpdateTokenByIDV1(ctx, "1", nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUpdateTokenByIDV1_EmptyToken(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	svc := NewService(mock)
	ctx := context.Background()

	req := &RequestTokenUpload{EncodedToken: ""}

	_, _, err := svc.UpdateTokenByIDV1(ctx, "1", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "encodedToken is required")
}

func TestDeleteByIDV1(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	mocks.RegisterDeleteByIDMock(mock)

	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteByIDV1(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestDeleteByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	svc := NewService(mock)
	ctx := context.Background()

	_, err := svc.DeleteByIDV1(ctx, "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "id is required")
}

func TestDisownDevicesByIDV1(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	mocks.RegisterDisownDevicesMock(mock)

	svc := NewService(mock)
	ctx := context.Background()

	req := &RequestDisown{
		Devices: []string{"a2s3d4f5", "0o9i8u7y6t"},
	}

	result, resp, err := svc.DisownDevicesByIDV1(ctx, "1", req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, result.Devices, 2)
	assert.Equal(t, "SUCCESS", result.Devices["a2s3d4f5"])
	assert.Equal(t, "FAILED", result.Devices["0o9i8u7y6t"])
}

func TestDisownDevicesByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	svc := NewService(mock)
	ctx := context.Background()

	req := &RequestDisown{Devices: []string{"serial1"}}

	_, _, err := svc.DisownDevicesByIDV1(ctx, "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "id is required")
}

func TestDisownDevicesByIDV1_NilRequest(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	svc := NewService(mock)
	ctx := context.Background()

	_, _, err := svc.DisownDevicesByIDV1(ctx, "1", nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestDisownDevicesByIDV1_EmptyDevicesList(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	svc := NewService(mock)
	ctx := context.Background()

	req := &RequestDisown{Devices: []string{}}

	_, _, err := svc.DisownDevicesByIDV1(ctx, "1", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "devices list is required")
}

func TestAddHistoryNotesV1(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	mocks.RegisterAddHistoryNotesMock(mock)

	svc := NewService(mock)
	ctx := context.Background()

	req := &RequestAddHistoryNotes{
		Note: "Test history note",
	}

	result, resp, err := svc.AddHistoryNotesV1(ctx, "1", req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.NotEmpty(t, result.Href)
}

func TestAddHistoryNotesV1_EmptyID(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	svc := NewService(mock)
	ctx := context.Background()

	req := &RequestAddHistoryNotes{Note: "Test"}

	_, _, err := svc.AddHistoryNotesV1(ctx, "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "id is required")
}

func TestAddHistoryNotesV1_NilRequest(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	svc := NewService(mock)
	ctx := context.Background()

	_, _, err := svc.AddHistoryNotesV1(ctx, "1", nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestAddHistoryNotesV1_EmptyNote(t *testing.T) {
	mock := mocks.NewDeviceEnrollmentsMock()
	svc := NewService(mock)
	ctx := context.Background()

	req := &RequestAddHistoryNotes{Note: ""}

	_, _, err := svc.AddHistoryNotesV1(ctx, "1", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "note is required")
}
