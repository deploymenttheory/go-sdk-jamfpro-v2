package adcs_settings

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/adcs_settings/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_AdcsSettings_CreateV1_Success(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	mock.RegisterCreateMock()

	svc := NewService(mock)
	ctx := context.Background()

	request := &ResourceAdcsSettings{
		DisplayName: "Test ADCS",
		CAName:      "TestCA",
		FQDN:        "adcs.example.com",
		AdcsURL:     "https://adcs.example.com/certsrv",
	}

	result, resp, err := svc.CreateV1(ctx, request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, "/api/v1/pki/adcs-settings/1", result.Href)
}

func TestUnit_AdcsSettings_CreateV1_NilRequest(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.CreateV1(ctx, nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_AdcsSettings_GetByIDV1_Success(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	mock.RegisterGetByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByIDV1(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Test ADCS", result.DisplayName)
	assert.Equal(t, "TestCA", result.CAName)
	assert.Equal(t, "adcs.example.com", result.FQDN)
	assert.True(t, result.RevocationEnabled)
	assert.False(t, result.Outbound)
}

func TestUnit_AdcsSettings_GetByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByIDV1(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_AdcsSettings_UpdateByIDV1_Success(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	mock.RegisterUpdateByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	request := &ResourceAdcsSettings{
		DisplayName: "Updated ADCS",
	}

	resp, err := svc.UpdateByIDV1(ctx, "1", request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_AdcsSettings_UpdateByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	svc := NewService(mock)
	ctx := context.Background()

	request := &ResourceAdcsSettings{
		DisplayName: "Updated ADCS",
	}

	resp, err := svc.UpdateByIDV1(ctx, "", request)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_AdcsSettings_UpdateByIDV1_NilRequest(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.UpdateByIDV1(ctx, "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_AdcsSettings_DeleteByIDV1_Success(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	mock.RegisterDeleteByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteByIDV1(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_AdcsSettings_DeleteByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteByIDV1(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_AdcsSettings_ValidateServerCertificateV1_Success(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	mock.RegisterValidateServerCertificateMock()

	svc := NewService(mock)
	ctx := context.Background()

	request := &ValidateCertificateRequest{
		Filename: "server.cer",
		Data:     []string{"base64encodeddata"},
	}

	resp, err := svc.ValidateServerCertificateV1(ctx, request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_AdcsSettings_ValidateServerCertificateV1_NilRequest(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.ValidateServerCertificateV1(ctx, nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_AdcsSettings_ValidateClientCertificateV1_Success(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	mock.RegisterValidateClientCertificateMock()

	svc := NewService(mock)
	ctx := context.Background()

	password := "test-password"
	request := &ValidateCertificateRequest{
		Filename: "client.p12",
		Data:     []string{"base64encodeddata"},
		Password: &password,
	}

	resp, err := svc.ValidateClientCertificateV1(ctx, request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_AdcsSettings_ValidateClientCertificateV1_NilRequest(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.ValidateClientCertificateV1(ctx, nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_AdcsSettings_GetDependenciesByIDV1_Success(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	mock.RegisterGetDependenciesByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDependenciesByIDV1(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "Test Profile", result.Results[0].ConfigProfileName)
	assert.Equal(t, "OSX_CONFIGURATION_PROFILE", result.Results[0].ConfigProfileType)
}

func TestUnit_AdcsSettings_GetDependenciesByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDependenciesByIDV1(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_AdcsSettings_GetHistoryByIDV1_Success(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	mock.RegisterGetHistoryByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetHistoryByIDV1(ctx, "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "Test note", result.Results[0].Note)
}

func TestUnit_AdcsSettings_GetHistoryByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetHistoryByIDV1(ctx, "", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_AdcsSettings_AddHistoryNoteByIDV1_Success(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	mock.RegisterAddHistoryNoteMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	request := &HistoryNoteRequest{
		Note: "Test history note",
	}

	resp, err := svc.AddHistoryNoteByIDV1(ctx, "1", request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
}

func TestUnit_AdcsSettings_AddHistoryNoteByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	svc := NewService(mock)
	ctx := context.Background()

	request := &HistoryNoteRequest{
		Note: "Test note",
	}

	resp, err := svc.AddHistoryNoteByIDV1(ctx, "", request)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_AdcsSettings_AddHistoryNoteByIDV1_NilRequest(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.AddHistoryNoteByIDV1(ctx, "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_AdcsSettings_CreateV1_Error(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	svc := NewService(mock)
	_, _, err := svc.CreateV1(context.Background(), &ResourceAdcsSettings{DisplayName: "test"})
	require.Error(t, err)
}

func TestUnit_AdcsSettings_GetByIDV1_Error(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	svc := NewService(mock)
	_, _, err := svc.GetByIDV1(context.Background(), "1")
	require.Error(t, err)
}

func TestUnit_AdcsSettings_UpdateByIDV1_Error(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	svc := NewService(mock)
	_, err := svc.UpdateByIDV1(context.Background(), "1", &ResourceAdcsSettings{DisplayName: "test"})
	require.Error(t, err)
}

func TestUnit_AdcsSettings_DeleteByIDV1_Error(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	svc := NewService(mock)
	_, err := svc.DeleteByIDV1(context.Background(), "1")
	require.Error(t, err)
}

func TestUnit_AdcsSettings_ValidateServerCertificateV1_Error(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	svc := NewService(mock)
	_, err := svc.ValidateServerCertificateV1(context.Background(), &ValidateCertificateRequest{Filename: "test.p12"})
	require.Error(t, err)
}

func TestUnit_AdcsSettings_ValidateClientCertificateV1_Error(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	svc := NewService(mock)
	_, err := svc.ValidateClientCertificateV1(context.Background(), &ValidateCertificateRequest{Filename: "test.p12"})
	require.Error(t, err)
}

func TestUnit_AdcsSettings_GetDependenciesByIDV1_Error(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	svc := NewService(mock)
	_, _, err := svc.GetDependenciesByIDV1(context.Background(), "1")
	require.Error(t, err)
}

func TestUnit_AdcsSettings_GetHistoryByIDV1_Error(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	svc := NewService(mock)
	_, _, err := svc.GetHistoryByIDV1(context.Background(), "1", nil)
	require.Error(t, err)
}

func TestUnit_AdcsSettings_AddHistoryNoteByIDV1_HTTPError(t *testing.T) {
	mock := mocks.NewAdcsSettingsMock()
	svc := NewService(mock)
	_, err := svc.AddHistoryNoteByIDV1(context.Background(), "1", &HistoryNoteRequest{Note: "test"})
	require.Error(t, err)
}
