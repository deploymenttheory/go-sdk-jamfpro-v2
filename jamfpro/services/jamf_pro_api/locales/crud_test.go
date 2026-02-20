package locales

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/locales/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.LocalesMock) {
	t.Helper()
	mock := mocks.NewLocalesMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnitListV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Len(t, result, 2)
	require.Equal(t, "en", result[0].Identifier)
	require.Equal(t, "English", result[0].Description)
}
