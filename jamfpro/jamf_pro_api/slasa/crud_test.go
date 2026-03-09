package slasa

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/slasa/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_Slasa_GetStatusV1_Accepted(t *testing.T) {
	mock := mocks.NewSLASAMock()
	mock.RegisterGetStatusAcceptedMock()
	service := NewSlasa(mock)

	result, resp, err := service.GetStatusV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "ACCEPTED", result.SLASAAcceptanceStatus)
}

func TestUnit_Slasa_GetStatusV1_NotAccepted(t *testing.T) {
	mock := mocks.NewSLASAMock()
	mock.RegisterGetStatusNotAcceptedMock()
	service := NewSlasa(mock)

	result, resp, err := service.GetStatusV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "NOT_ACCEPTED", result.SLASAAcceptanceStatus)
}

func TestUnit_Slasa_GetStatusV1_Error(t *testing.T) {
	mock := mocks.NewSLASAMock()
	mock.RegisterGetStatusErrorMock()
	service := NewSlasa(mock)

	result, resp, err := service.GetStatusV1(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request failed")
}

func TestUnit_Slasa_GetStatusV1_NoMockRegistered(t *testing.T) {
	mock := mocks.NewSLASAMock()
	service := NewSlasa(mock)

	result, resp, err := service.GetStatusV1(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no mock registered")
}

func TestUnit_Slasa_AcceptV1_Success(t *testing.T) {
	mock := mocks.NewSLASAMock()
	mock.RegisterAcceptMock()
	service := NewSlasa(mock)

	resp, err := service.AcceptV1(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_Slasa_AcceptV1_NoMockRegistered(t *testing.T) {
	mock := mocks.NewSLASAMock()
	service := NewSlasa(mock)

	resp, err := service.AcceptV1(context.Background())
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no mock registered")
}
