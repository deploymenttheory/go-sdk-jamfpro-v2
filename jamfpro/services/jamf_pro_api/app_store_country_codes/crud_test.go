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

func TestUnitListV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Len(t, result.CountryCodes, 3)
	require.Equal(t, "US", result.CountryCodes[0].Code)
	require.Equal(t, "United States", result.CountryCodes[0].Name)
	require.Equal(t, "GB", result.CountryCodes[1].Code)
	require.Equal(t, "United Kingdom", result.CountryCodes[1].Name)
}
