package locales

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/locales/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Locales, *mocks.LocalesMock) {
	t.Helper()
	mock := mocks.NewLocalesMock()
	mock.RegisterMocks()
	return NewLocales(mock), mock
}

func TestUnit_Locales_ListV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode())
	require.Len(t, result, 2)
	require.Equal(t, "en", result[0].Identifier)
	require.Equal(t, "English", result[0].Description)
}

func TestUnit_Locales_ListV1_Error(t *testing.T) {
	mock := mocks.NewLocalesMock()
	mock.RegisterListV1ErrorMock()
	svc := NewLocales(mock)

	result, resp, err := svc.ListV1(context.Background())
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 500, resp.StatusCode())
	require.Contains(t, err.Error(), "mock client error")
}

func TestUnit_Locales_ListV1_NoMockRegistered(t *testing.T) {
	mock := mocks.NewLocalesMock()
	mock.RegisterListV1NoResponseErrorMock()
	svc := NewLocales(mock)

	result, resp, err := svc.ListV1(context.Background())
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnit_Locales_ListV1_InvalidJSON(t *testing.T) {
	mock := mocks.NewLocalesMock()
	mock.RegisterListV1InvalidJSONMock()
	svc := NewLocales(mock)

	result, resp, err := svc.ListV1(context.Background())
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode())
}
