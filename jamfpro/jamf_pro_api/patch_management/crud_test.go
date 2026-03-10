package patch_management

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/patch_management/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_PatchManagement_AcceptDisclaimerV2_Success(t *testing.T) {
	mock := mocks.NewPatchManagementMock()
	mock.RegisterAcceptDisclaimerMock()
	service := NewPatchManagement(mock)

	resp, err := service.AcceptDisclaimerV2(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_PatchManagement_AcceptDisclaimerV2_Error(t *testing.T) {
	mock := mocks.NewPatchManagementMock()
	mock.RegisterAcceptDisclaimerErrorMock()
	service := NewPatchManagement(mock)

	resp, err := service.AcceptDisclaimerV2(context.Background())
	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Contains(t, err.Error(), "mock error")
}

func TestUnit_PatchManagement_AcceptDisclaimerV2_NoMockRegistered(t *testing.T) {
	mock := mocks.NewPatchManagementMock()
	mock.RegisterAcceptDisclaimerNoResponseErrorMock()
	service := NewPatchManagement(mock)

	resp, err := service.AcceptDisclaimerV2(context.Background())
	assert.Error(t, err)
	assert.NotNil(t, resp)
}
