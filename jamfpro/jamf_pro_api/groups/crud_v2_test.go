package groups

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/groups/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_Groups_ListV2(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterV2Mocks()
	svc := NewGroups(mock)

	result, resp, err := svc.ListV2(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, result.Results)
}

func TestUnit_Groups_GetByIDV2(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterV2Mocks()
	svc := NewGroups(mock)

	result, _, err := svc.GetByIDV2(context.Background(), "1630e1d5-e0e9-449d-ad11-a5324dd16c46")
	require.NoError(t, err)
	assert.NotEmpty(t, result.GroupPlatformId)
}

func TestUnit_Groups_GetByIDV2_EmptyID(t *testing.T) {
	svc := NewGroups(mocks.NewGroupsMock())
	_, _, err := svc.GetByIDV2(context.Background(), "")
	require.Error(t, err)
}

func TestUnit_Groups_UpdateByIDV2(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterV2Mocks()
	svc := NewGroups(mock)

	result, _, err := svc.UpdateByIDV2(context.Background(), "1630e1d5", &RequestUpdateGroupV2{
		GroupName: "renamed",
		Criteria:  []SubsetCriterion{{Name: "Operating System", AndOr: "and", SearchType: "like", Value: "macOS"}},
	})
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestUnit_Groups_UpdateByIDV2_InvalidAndOr(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterV2Mocks()
	svc := NewGroups(mock)

	_, _, err := svc.UpdateByIDV2(context.Background(), "1630e1d5", &RequestUpdateGroupV2{
		Criteria: []SubsetCriterion{{Name: "Operating System", AndOr: "nope", SearchType: "like", Value: "macOS"}},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "andOr")
}

func TestUnit_Groups_DeleteByIDV2(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterV2Mocks()
	svc := NewGroups(mock)

	resp, err := svc.DeleteByIDV2(context.Background(), "1630e1d5")
	require.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode())
}
