package login_customization

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/login_customization/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.LoginCustomizationMock) {
	t.Helper()
	mock := mocks.NewLoginCustomizationMock()
	return NewService(mock), mock
}

func TestUnitGetV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetLoginCustomizationMock()

	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, result.IncludeCustomDisclaimer)
	assert.Equal(t, "Accept", result.ActionText)
}

func TestUnitUpdateV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateLoginCustomizationMock()

	request := &ResourceLoginCustomizationV1{
		RampInstance:            true,
		IncludeCustomDisclaimer: true,
		DisclaimerHeading:       "Updated Disclaimer Header",
		DisclaimerMainText:      "Updated disclaimer main text",
		ActionText:              "Accept",
	}
	result, resp, err := svc.UpdateV1(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "Updated Disclaimer Header", result.DisclaimerHeading)
}
