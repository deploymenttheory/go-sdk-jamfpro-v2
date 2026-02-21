package csa

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/csa/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetTokenExchangeDetailsV1(t *testing.T) {
	mock := mocks.NewCSAMock()
	mock.RegisterGetTokenExchangeDetailsMock()

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetTokenExchangeDetailsV1(ctx)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, result.TenantID)
	assert.NotEmpty(t, result.Subject)
	assert.Greater(t, result.RefreshExpiration, 0)
	assert.NotEmpty(t, result.Scopes)
}

func TestGetTenantIDV1(t *testing.T) {
	mock := mocks.NewCSAMock()
	mock.RegisterGetTenantIDMock()

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetTenantIDV1(ctx)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, result.TenantID)
}

func TestDeleteTokenExchangeV1(t *testing.T) {
	mock := mocks.NewCSAMock()
	mock.RegisterDeleteTokenExchangeMock()

	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteTokenExchangeV1(ctx)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}
