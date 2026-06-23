package computer_groups

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupV3MockService(t *testing.T) *ComputerGroups {
	t.Helper()
	_, mock := setupMockService(t)
	mock.RegisterV3Mocks()
	return NewComputerGroups(mock)
}

func TestUnit_ComputerGroups_ListSmartV3_Success(t *testing.T) {
	svc := setupV3MockService(t)
	result, resp, err := svc.ListSmartV3(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode())
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "All Managed Clients", result.Results[0].Name)
}

func TestUnit_ComputerGroups_GetSmartByIDV3_Success(t *testing.T) {
	svc := setupV3MockService(t)
	result, _, err := svc.GetSmartByIDV3(context.Background(), "1")
	require.NoError(t, err)
	require.Len(t, result.Criteria, 1)
	assert.Equal(t, "and", result.Criteria[0].AndOr)
	assert.False(t, result.Criteria[0].OpeningParen)
}

func TestUnit_ComputerGroups_GetSmartByIDV3_EmptyID(t *testing.T) {
	svc := setupV3MockService(t)
	_, _, err := svc.GetSmartByIDV3(context.Background(), "")
	require.Error(t, err)
}

func TestUnit_ComputerGroups_CreateSmartV3_Success(t *testing.T) {
	svc := setupV3MockService(t)
	result, resp, err := svc.CreateSmartV3(context.Background(), &RequestSmartGroupV3{
		Name:     "sdk-v3-smart",
		SiteID:   "-1",
		Criteria: []CriterionV3{{Name: "Operating System", AndOr: "and", SearchType: "like", Value: "macOS"}},
	})
	require.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "13979", result.ID)
}

func TestUnit_ComputerGroups_CreateSmartV3_InvalidAndOr(t *testing.T) {
	svc := setupV3MockService(t)
	_, _, err := svc.CreateSmartV3(context.Background(), &RequestSmartGroupV3{
		Name:     "bad",
		Criteria: []CriterionV3{{Name: "Operating System", AndOr: "nand", SearchType: "like", Value: "macOS"}},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "andOr")
}

func TestUnit_ComputerGroups_CreateSmartV3_NameTooLong(t *testing.T) {
	svc := setupV3MockService(t)
	_, _, err := svc.CreateSmartV3(context.Background(), &RequestSmartGroupV3{Name: strings.Repeat("a", 256)})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "255")
}

func TestUnit_ComputerGroups_CreateSmartV3_NilRequest(t *testing.T) {
	svc := setupV3MockService(t)
	_, _, err := svc.CreateSmartV3(context.Background(), nil)
	require.Error(t, err)
}

func TestUnit_ComputerGroups_UpdateSmartByIDV3_Success(t *testing.T) {
	svc := setupV3MockService(t)
	_, _, err := svc.UpdateSmartByIDV3(context.Background(), "1", &RequestSmartGroupV3{Name: "renamed"})
	require.NoError(t, err)
}

func TestUnit_ComputerGroups_DeleteSmartByIDV3_Success(t *testing.T) {
	svc := setupV3MockService(t)
	resp, err := svc.DeleteSmartByIDV3(context.Background(), "1")
	require.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_ComputerGroups_GetSmartGroupMembershipByIDV3_Success(t *testing.T) {
	svc := setupV3MockService(t)
	result, _, err := svc.GetSmartGroupMembershipByIDV3(context.Background(), "1")
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestUnit_ComputerGroups_ListStaticV3_Success(t *testing.T) {
	svc := setupV3MockService(t)
	result, _, err := svc.ListStaticV3(context.Background(), nil)
	require.NoError(t, err)
	assert.Equal(t, 0, result.TotalCount)
}

func TestUnit_ComputerGroups_GetStaticByIDV3_Success(t *testing.T) {
	svc := setupV3MockService(t)
	result, _, err := svc.GetStaticByIDV3(context.Background(), "10")
	require.NoError(t, err)
	assert.Equal(t, "10", result.ID)
}

func TestUnit_ComputerGroups_CreateStaticV3_DedupesAssignments(t *testing.T) {
	svc := setupV3MockService(t)
	req := &RequestStaticGroupV3{Name: "sdk-v3-static", Assignments: []string{"1", "2", "2", "1"}}
	_, resp, err := svc.CreateStaticV3(context.Background(), req)
	require.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, []string{"1", "2"}, req.Assignments)
}

func TestUnit_ComputerGroups_UpdateStaticByIDV3_Success(t *testing.T) {
	svc := setupV3MockService(t)
	_, _, err := svc.UpdateStaticByIDV3(context.Background(), "10", &RequestStaticGroupV3{Name: "renamed"})
	require.NoError(t, err)
}

func TestUnit_ComputerGroups_DeleteStaticByIDV3_Success(t *testing.T) {
	svc := setupV3MockService(t)
	resp, err := svc.DeleteStaticByIDV3(context.Background(), "10")
	require.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode())
}
