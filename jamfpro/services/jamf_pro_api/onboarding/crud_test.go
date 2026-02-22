package onboarding

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/onboarding/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.OnboardingMock) {
	t.Helper()
	mock := mocks.NewOnboardingMock()
	return NewService(mock), mock
}

func TestUnitGetV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()
	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, result.Enabled)
	require.Len(t, result.OnboardingItems, 1)
	assert.Equal(t, "APP", result.OnboardingItems[0].SelfServiceEntityType)
}

func TestUnitUpdateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UpdateV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitUpdateV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()
	req := &ResourceUpdateOnboardingSettings{
		Enabled: true,
		OnboardingItems: []SubsetOnboardingItemRequest{
			{
				EntityID:              "123",
				SelfServiceEntityType: "OS_X_POLICY",
				Priority:              1,
			},
		},
	}
	result, resp, err := svc.UpdateV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, result.Enabled)
}

func TestUnitGetEligibleAppsV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetEligibleAppsMock()
	result, resp, err := svc.GetEligibleAppsV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "101", result.Results[0].ID)
	assert.Equal(t, "Microsoft Office", result.Results[0].Name)
}

func TestUnitGetEligibleConfigurationProfilesV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetEligibleConfigurationProfilesMock()
	result, resp, err := svc.GetEligibleConfigurationProfilesV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
}

func TestUnitGetEligiblePoliciesV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetEligiblePoliciesMock()
	result, resp, err := svc.GetEligiblePoliciesV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
}

func TestUnitGetHistoryV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryMock()
	result, resp, err := svc.GetHistoryV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.NotNil(t, result.Results[0].Details)
}

func TestUnitAddHistoryNotesV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.AddHistoryNotesV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitAddHistoryNotesV1_EmptyNote(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestAddHistoryNotes{Note: ""}
	result, resp, err := svc.AddHistoryNotesV1(context.Background(), req)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitAddHistoryNotesV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddHistoryNotesMock()
	req := &RequestAddHistoryNotes{Note: "Test note"}
	result, resp, err := svc.AddHistoryNotesV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "3", result.ID)
	assert.Equal(t, "/api/v1/onboarding/history/3", result.Href)
}

func TestUnitExportHistoryV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportHistoryMock()
	data, resp, err := svc.ExportHistoryV1(context.Background(), "text/csv", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, data)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Contains(t, string(data), "id,username,date,note,details")
}
