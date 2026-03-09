package jamf_package

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/jamf_package/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*JamfPackage, *mocks.JamfPackageMock) {
	t.Helper()
	mock := mocks.NewJamfPackageMock()
	return NewJamfPackage(mock), mock
}

func TestUnit_JamfPackage_NewService(t *testing.T) {
	mock := mocks.NewJamfPackageMock()
	svc := NewJamfPackage(mock)
	require.NotNil(t, svc)
}

func TestUnit_JamfPackage_ListV1_Success_Protect(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListV1Mock()

	result, resp, err := svc.ListV1(context.Background(), constants.ApplicationProtect)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	require.Len(t, result, 2)
	assert.Equal(t, "1", result[0].ID)
	assert.Equal(t, "JamfProtect-1.0.pkg", result[0].Filename)
	assert.Equal(t, "1.0", result[0].Version)
	assert.Equal(t, "https://example.com/pkg1.pkg", result[0].URL)
	assert.Equal(t, "2", result[1].ID)
	assert.Equal(t, "JamfProtect-2.0.pkg", result[1].Filename)
	assert.Equal(t, "2.0", result[1].Version)
	assert.Equal(t, "protect", mock.LastRSQLQuery["application"])
}

func TestUnit_JamfPackage_ListV1_Success_Connect(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListV1Mock()

	result, resp, err := svc.ListV1(context.Background(), constants.ApplicationConnect)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	require.Len(t, result, 2)
	assert.Equal(t, "connect", mock.LastRSQLQuery["application"])
}

func TestUnit_JamfPackage_ListV1_Success_ApplicationCaseInsensitive(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListV1Mock()

	result, resp, err := svc.ListV1(context.Background(), "PROTECT")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "protect", mock.LastRSQLQuery["application"])

	result2, resp2, err2 := svc.ListV1(context.Background(), " Connect ")
	require.NoError(t, err2)
	require.NotNil(t, result2)
	require.NotNil(t, resp2)
	assert.Equal(t, 200, resp2.StatusCode())
	assert.Equal(t, "connect", mock.LastRSQLQuery["application"])
}

func TestUnit_JamfPackage_ListV1_InvalidApplication_Empty(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.ListV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "list jamf packages")
	assert.Contains(t, err.Error(), "application must")
}

func TestUnit_JamfPackage_ListV1_InvalidApplication_Unknown(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.ListV1(context.Background(), "invalid")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "list jamf packages")
	assert.Contains(t, err.Error(), "application must")
}

func TestUnit_JamfPackage_ListV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.ListV1(context.Background(), constants.ApplicationProtect)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to list jamf packages")
}

func TestUnit_JamfPackage_ListV1_InvalidJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListV1InvalidJSONMock()

	result, resp, err := svc.ListV1(context.Background(), constants.ApplicationProtect)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Contains(t, err.Error(), "unmarshal")
}

func TestUnit_JamfPackage_ListV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListV1NotFoundErrorMock()

	result, resp, err := svc.ListV1(context.Background(), constants.ApplicationProtect)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_JamfPackage_GetV2_Success_Protect(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetV2Mock()

	result, resp, err := svc.GetV2(context.Background(), constants.ApplicationProtect)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "Jamf Protect", result.DisplayName)
	assert.Equal(t, "https://docs.jamf.com/jamf-protect/release-notes/", result.ReleaseHistoryURL)
	require.Len(t, result.Artifacts, 1)
	assert.Equal(t, "1", result.Artifacts[0].ID)
	assert.Equal(t, "JamfProtect-1.0.pkg", result.Artifacts[0].Filename)
	assert.Equal(t, "1.0", result.Artifacts[0].Version)
	assert.Equal(t, "protect", mock.LastRSQLQuery["application"])
}

func TestUnit_JamfPackage_GetV2_Success_Connect(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetV2Mock()

	result, resp, err := svc.GetV2(context.Background(), constants.ApplicationConnect)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "Jamf Protect", result.DisplayName)
	assert.Equal(t, "connect", mock.LastRSQLQuery["application"])
}

func TestUnit_JamfPackage_GetV2_Success_ApplicationCaseInsensitive(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetV2Mock()

	result, resp, err := svc.GetV2(context.Background(), "PROTECT")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "protect", mock.LastRSQLQuery["application"])
}

func TestUnit_JamfPackage_GetV2_InvalidApplication_Empty(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetV2(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "get jamf package")
	assert.Contains(t, err.Error(), "application must")
}

func TestUnit_JamfPackage_GetV2_InvalidApplication_Unknown(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetV2(context.Background(), "unknown")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "get jamf package")
	assert.Contains(t, err.Error(), "application must")
}

func TestUnit_JamfPackage_GetV2_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetV2(context.Background(), constants.ApplicationProtect)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to get jamf package")
}

func TestUnit_JamfPackage_GetV2_InvalidJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetV2InvalidJSONMock()

	result, resp, err := svc.GetV2(context.Background(), constants.ApplicationProtect)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Contains(t, err.Error(), "unmarshal")
}

func TestUnit_JamfPackage_GetV2_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetV2NotFoundErrorMock()

	result, resp, err := svc.GetV2(context.Background(), constants.ApplicationProtect)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_JamfPackage_ListV1_EmptyResult(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListV1EmptyMock()

	result, resp, err := svc.ListV1(context.Background(), constants.ApplicationProtect)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Empty(t, result)
}
