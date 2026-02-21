package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_idp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCloudIdp_List(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.CloudIdp
	ctx := context.Background()

	result, _, err := svc.ListV1(ctx, nil)
	if err != nil {
		t.Skipf("Failed to list Cloud Identity Providers (may not be supported on this tenant): %v", err)
		return
	}

	assert.NotNil(t, result)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
}

func TestCloudIdp_GetByID(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.CloudIdp
	ctx := context.Background()

	list, _, err := svc.ListV1(ctx, nil)
	if err != nil || len(list.Results) == 0 {
		t.Skip("No Cloud Identity Providers available for testing")
		return
	}

	firstID := list.Results[0].ID

	result, _, err := svc.GetByIDV1(ctx, firstID)
	require.NoError(t, err)
	assert.Equal(t, firstID, result.ID)
	assert.NotEmpty(t, result.DisplayName)
	assert.NotEmpty(t, result.ProviderName)
}

func TestCloudIdp_GetByName(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.CloudIdp
	ctx := context.Background()

	list, _, err := svc.ListV1(ctx, nil)
	if err != nil || len(list.Results) == 0 {
		t.Skip("No Cloud Identity Providers available for testing")
		return
	}

	firstName := list.Results[0].DisplayName

	result, _, err := svc.GetByNameV1(ctx, firstName)
	require.NoError(t, err)
	assert.Equal(t, firstName, result.DisplayName)
}

func TestCloudIdp_Export(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.CloudIdp
	ctx := context.Background()

	request := &cloud_idp.ExportRequest{
		Fields: []cloud_idp.ExportField{
			{Name: "id"},
			{Name: "displayName"},
		},
	}

	_, data, err := svc.ExportV1(ctx, nil, request)
	if err != nil {
		t.Skipf("Failed to export Cloud Identity Providers (may not be supported on this tenant): %v", err)
		return
	}

	assert.NotEmpty(t, data)
}

func TestCloudIdp_HistoryOperations(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.CloudIdp
	ctx := context.Background()

	list, _, err := svc.ListV1(ctx, nil)
	if err != nil || len(list.Results) == 0 {
		t.Skip("No Cloud Identity Providers available for testing")
		return
	}

	firstID := list.Results[0].ID

	history, _, err := svc.GetHistoryByIDV1(ctx, firstID, nil)
	if err != nil {
		t.Skipf("Failed to get history (may not be supported): %v", err)
		return
	}
	assert.NotNil(t, history)
	assert.GreaterOrEqual(t, history.TotalCount, 0)

	noteReq := &cloud_idp.HistoryNoteRequest{
		Note: "Test history note from SDK",
	}
	_, err = svc.AddHistoryNoteByIDV1(ctx, firstID, noteReq)
	if err != nil {
		t.Skipf("Failed to add history note (may not be supported): %v", err)
		return
	}
}

func TestCloudIdp_TestSearches(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.CloudIdp
	ctx := context.Background()

	list, _, err := svc.ListV1(ctx, nil)
	if err != nil || len(list.Results) == 0 {
		t.Skip("No Cloud Identity Providers available for testing")
		return
	}

	firstID := list.Results[0].ID

	groupReq := &cloud_idp.TestGroupSearchRequest{
		GroupName: "TestGroup",
	}
	_, _, err = svc.TestGroupSearchByIDV1(ctx, firstID, groupReq)
	if err != nil {
		t.Skipf("Failed to test group search (expected - test data may not exist): %v", err)
	}

	userReq := &cloud_idp.TestUserSearchRequest{
		Username: "testuser",
	}
	_, _, err = svc.TestUserSearchByIDV1(ctx, firstID, userReq)
	if err != nil {
		t.Skipf("Failed to test user search (expected - test data may not exist): %v", err)
	}

	membershipReq := &cloud_idp.TestUserMembershipRequest{
		Username:  "testuser",
		GroupName: "TestGroup",
	}
	_, _, err = svc.TestUserMembershipByIDV1(ctx, firstID, membershipReq)
	if err != nil {
		t.Skipf("Failed to test user membership (expected - test data may not exist): %v", err)
	}
}
