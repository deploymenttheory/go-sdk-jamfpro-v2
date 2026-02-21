package dss_declarations

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/dss_declarations/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetByUUIDV1(t *testing.T) {
	mock := mocks.NewDSSDeclarationsMock()
	mocks.RegisterGetByUUIDMock(mock)

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByUUIDV1(ctx, "550e8400-e29b-41d4-a716-446655440000")

	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, result.Declarations, 1)
	assert.Equal(t, "550e8400-e29b-41d4-a716-446655440000", result.Declarations[0].UUID)
	assert.Equal(t, "com.apple.configuration.management.status-subscriptions", result.Declarations[0].Type)
}

func TestGetByUUIDV1_EmptyUUID(t *testing.T) {
	mock := mocks.NewDSSDeclarationsMock()
	svc := NewService(mock)
	ctx := context.Background()

	_, _, err := svc.GetByUUIDV1(ctx, "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "uuid is required")
}
