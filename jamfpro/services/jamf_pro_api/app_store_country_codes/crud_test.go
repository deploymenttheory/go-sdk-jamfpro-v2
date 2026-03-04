package app_store_country_codes

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/app_store_country_codes/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.AppStoreCountryCodesMock) {
	t.Helper()
	mock := mocks.NewAppStoreCountryCodesMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnit_AppStoreCountryCodes_NewService(t *testing.T) {
	mock := mocks.NewAppStoreCountryCodesMock()
	svc := NewService(mock)
	require.NotNil(t, svc)
	require.NotNil(t, svc.client)
}

func TestUnit_AppStoreCountryCodes_ListV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode())
	require.Len(t, result.CountryCodes, 3)
	require.Equal(t, "US", result.CountryCodes[0].Code)
	require.Equal(t, "United States", result.CountryCodes[0].Name)
	require.Equal(t, "GB", result.CountryCodes[1].Code)
	require.Equal(t, "United Kingdom", result.CountryCodes[1].Name)
}

func TestUnit_AppStoreCountryCodes_ListV1_NoMockRegistered(t *testing.T) {
	mock := mocks.NewAppStoreCountryCodesMock()
	// Do NOT call mock.RegisterMocks() - no mock for ListV1
	svc := NewService(mock)
	result, resp, err := svc.ListV1(context.Background())
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "no response for")
}

func TestUnit_AppStoreCountryCodes_ListV1_EmptyList(t *testing.T) {
	mock := mocks.NewAppStoreCountryCodesMock()
	mock.RegisterEmptyListMock()
	svc := NewService(mock)
	result, resp, err := svc.ListV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode())
	require.NotNil(t, result.CountryCodes)
	require.Len(t, result.CountryCodes, 0)
}
