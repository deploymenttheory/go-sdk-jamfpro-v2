package cloud_idp

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_idp/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListV1(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.ListV1(ctx, nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	assert.Len(t, result.Results, 2)
	assert.Equal(t, "Test Azure IDP", result.Results[0].DisplayName)
	assert.Equal(t, "AZURE", result.Results[0].ProviderName)
	assert.True(t, result.Results[0].Enabled)
}

func TestGetByIDV1(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	mock.RegisterGetByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByIDV1(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Test Azure IDP", result.DisplayName)
	assert.Equal(t, "AZURE", result.ProviderName)
}

func TestGetByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByIDV1(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestGetByNameV1(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	mock.RegisterListMock()
	mock.RegisterGetByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByNameV1(ctx, "Test Azure IDP")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Test Azure IDP", result.DisplayName)
}

func TestGetByNameV1_EmptyName(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByNameV1(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "name is required")
}

func TestGetByNameV1_NotFound(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByNameV1(ctx, "Nonexistent IDP")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not found")
}

func TestExportV1(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	mock.RegisterExportMock()

	svc := NewService(mock)
	ctx := context.Background()

	request := &ExportRequest{
		Fields: []ExportField{
			{Name: "id"},
			{Name: "displayName"},
		},
	}

	resp, data, err := svc.ExportV1(ctx, nil, request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, data)
}

func TestExportV1_QueryOnly(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	mock.RegisterExportMock()

	svc := NewService(mock)
	ctx := context.Background()

	query := map[string]string{
		"export-fields": "id,displayName",
	}

	resp, data, err := svc.ExportV1(ctx, query, nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, data)
}

func TestGetHistoryByIDV1(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	mock.RegisterGetHistoryByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetHistoryByIDV1(ctx, "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "Test history note", result.Results[0].Note)
}

func TestGetHistoryByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetHistoryByIDV1(ctx, "", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestAddHistoryNoteByIDV1(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	mock.RegisterAddHistoryNoteMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	request := &HistoryNoteRequest{
		Note: "Test note",
	}

	resp, err := svc.AddHistoryNoteByIDV1(ctx, "1", request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
}

func TestAddHistoryNoteByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	svc := NewService(mock)
	ctx := context.Background()

	request := &HistoryNoteRequest{
		Note: "Test note",
	}

	resp, err := svc.AddHistoryNoteByIDV1(ctx, "", request)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestAddHistoryNoteByIDV1_NilRequest(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.AddHistoryNoteByIDV1(ctx, "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestTestGroupSearchByIDV1(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	mock.RegisterTestGroupSearchMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	request := &TestGroupSearchRequest{
		GroupName: "TestGroup",
	}

	result, resp, err := svc.TestGroupSearchByIDV1(ctx, "1", request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "TestGroup", result.Results[0].Name)
}

func TestTestGroupSearchByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	svc := NewService(mock)
	ctx := context.Background()

	request := &TestGroupSearchRequest{
		GroupName: "TestGroup",
	}

	result, resp, err := svc.TestGroupSearchByIDV1(ctx, "", request)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestTestGroupSearchByIDV1_NilRequest(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.TestGroupSearchByIDV1(ctx, "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestTestUserSearchByIDV1(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	mock.RegisterTestUserSearchMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	request := &TestUserSearchRequest{
		Username: "testuser",
	}

	result, resp, err := svc.TestUserSearchByIDV1(ctx, "1", request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "testuser", result.Results[0].Name)
	assert.Equal(t, "Test User", result.Results[0].Attributes.FullName)
	assert.Equal(t, "testuser@example.com", result.Results[0].Attributes.EmailAddress)
}

func TestTestUserSearchByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	svc := NewService(mock)
	ctx := context.Background()

	request := &TestUserSearchRequest{
		Username: "testuser",
	}

	result, resp, err := svc.TestUserSearchByIDV1(ctx, "", request)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestTestUserSearchByIDV1_NilRequest(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.TestUserSearchByIDV1(ctx, "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestTestUserMembershipByIDV1(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	mock.RegisterTestUserMembershipMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	request := &TestUserMembershipRequest{
		Username:  "testuser",
		GroupName: "TestGroup",
	}

	result, resp, err := svc.TestUserMembershipByIDV1(ctx, "1", request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "TestGroup", result.Results[0].Name)
}

func TestTestUserMembershipByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	svc := NewService(mock)
	ctx := context.Background()

	request := &TestUserMembershipRequest{
		Username:  "testuser",
		GroupName: "TestGroup",
	}

	result, resp, err := svc.TestUserMembershipByIDV1(ctx, "", request)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestTestUserMembershipByIDV1_NilRequest(t *testing.T) {
	mock := mocks.NewCloudIdpMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.TestUserMembershipByIDV1(ctx, "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}
