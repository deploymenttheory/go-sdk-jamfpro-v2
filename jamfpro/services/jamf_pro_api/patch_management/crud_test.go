package patch_management

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/patch_management/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitAcceptDisclaimerV2_Success(t *testing.T) {
	mock := mocks.NewPatchManagementMock()
	mock.RegisterAcceptDisclaimerMock()
	service := NewService(mock)

	resp, err := service.AcceptDisclaimerV2(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitAcceptDisclaimerV2_Error(t *testing.T) {
	mock := mocks.NewPatchManagementMock()
	mock.RegisterAcceptDisclaimerErrorMock()
	service := NewService(mock)

	resp, err := service.AcceptDisclaimerV2(context.Background())
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request failed")
}

func TestUnitAcceptDisclaimerV2_NoMockRegistered(t *testing.T) {
	mock := mocks.NewPatchManagementMock()
	service := NewService(mock)

	resp, err := service.AcceptDisclaimerV2(context.Background())
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no mock registered")
}
