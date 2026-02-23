package slasa

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/slasa/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitGetStatusV1_Accepted(t *testing.T) {
	mock := mocks.NewSLASAMock()
	mock.RegisterGetStatusAcceptedMock()
	service := NewService(mock)

	result, resp, err := service.GetStatusV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "ACCEPTED", result.SLASAAcceptanceStatus)
}

func TestUnitGetStatusV1_NotAccepted(t *testing.T) {
	mock := mocks.NewSLASAMock()
	mock.RegisterGetStatusNotAcceptedMock()
	service := NewService(mock)

	result, resp, err := service.GetStatusV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "NOT_ACCEPTED", result.SLASAAcceptanceStatus)
}

func TestUnitGetStatusV1_Error(t *testing.T) {
	mock := mocks.NewSLASAMock()
	mock.RegisterGetStatusErrorMock()
	service := NewService(mock)

	result, resp, err := service.GetStatusV1(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request failed")
}

func TestUnitGetStatusV1_NoMockRegistered(t *testing.T) {
	mock := mocks.NewSLASAMock()
	service := NewService(mock)

	result, resp, err := service.GetStatusV1(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no mock registered")
}

func TestUnitAcceptV1_Success(t *testing.T) {
	mock := mocks.NewSLASAMock()
	mock.RegisterAcceptMock()
	service := NewService(mock)

	resp, err := service.AcceptV1(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitAcceptV1_NoMockRegistered(t *testing.T) {
	mock := mocks.NewSLASAMock()
	service := NewService(mock)

	resp, err := service.AcceptV1(context.Background())
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no mock registered")
}
