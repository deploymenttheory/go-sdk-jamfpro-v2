package login_customization

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/login_customization/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*LoginCustomization, *mocks.LoginCustomizationMock) {
	t.Helper()
	mock := mocks.NewLoginCustomizationMock()
	return NewLoginCustomization(mock), mock
}

func TestUnit_LoginCustomization_GetV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetLoginCustomizationMock()

	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.True(t, result.IncludeCustomDisclaimer)
	assert.Equal(t, "Accept", result.ActionText)
}

func TestUnit_LoginCustomization_GetV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetV1(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_LoginCustomization_UpdateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UpdateV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_LoginCustomization_UpdateV1_Success(t *testing.T) {
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
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "Updated Disclaimer Header", result.DisclaimerHeading)
}
