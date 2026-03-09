package vpp_assignments_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/vpp_assignments"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/vpp_assignments/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_VPPAssignments_List(t *testing.T) {
	mockClient := mocks.NewVPPAssignmentsMock()
	mockClient.RegisterListVPPAssignmentsMock()
	svc := vpp_assignments.NewVppAssignments(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Len(t, resp.VPPAssignments, 2)
	assert.Equal(t, 1, resp.VPPAssignments[0].ID)
	assert.Equal(t, "Test VPP Assignment 1", resp.VPPAssignments[0].Name)
	assert.Equal(t, 2, resp.VPPAssignments[1].ID)
	assert.Equal(t, "Test VPP Assignment 2", resp.VPPAssignments[1].Name)
}

func TestUnit_VPPAssignments_GetByID(t *testing.T) {
	mockClient := mocks.NewVPPAssignmentsMock()
	mockClient.RegisterGetVPPAssignmentByIDMock()
	svc := vpp_assignments.NewVppAssignments(mockClient)

	resp, _, err := svc.GetByID(context.Background(), 1)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Test VPP Assignment", resp.General.Name)
	assert.Equal(t, 1, resp.General.VPPAdminAccountID)
	assert.Equal(t, "Test VPP Account", resp.General.VPPAdminAccountName)
	assert.Len(t, resp.IOSApps, 1)
	assert.Equal(t, 123456789, resp.IOSApps[0].AdamID)
	assert.Equal(t, "Test iOS App", resp.IOSApps[0].Name)
	assert.Len(t, resp.MacApps, 1)
	assert.Equal(t, 987654321, resp.MacApps[0].AdamID)
	assert.False(t, resp.Scope.AllJSSUsers)
	assert.Len(t, resp.Scope.JSSUsers, 1)
	assert.Equal(t, 1, resp.Scope.JSSUsers[0].ID)
	assert.Equal(t, "testuser", resp.Scope.JSSUsers[0].Name)
}

func TestUnit_VPPAssignments_GetByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewVPPAssignmentsMock()
	svc := vpp_assignments.NewVppAssignments(mockClient)

	_, _, err := svc.GetByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "VPP assignment ID must be a positive integer")
}

func TestUnit_VPPAssignments_Create(t *testing.T) {
	mockClient := mocks.NewVPPAssignmentsMock()
	mockClient.RegisterCreateVPPAssignmentMock()
	svc := vpp_assignments.NewVppAssignments(mockClient)

	req := &vpp_assignments.Resource{
		General: vpp_assignments.SubsetGeneral{
			Name:                "New VPP Assignment",
			VPPAdminAccountID:   1,
			VPPAdminAccountName: "Test Account",
		},
		Scope: vpp_assignments.SubsetScope{
			AllJSSUsers: false,
		},
	}

	resp, _, err := svc.Create(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.ID)
}

func TestUnit_VPPAssignments_Create_NilRequest(t *testing.T) {
	mockClient := mocks.NewVPPAssignmentsMock()
	svc := vpp_assignments.NewVppAssignments(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_VPPAssignments_Create_EmptyName(t *testing.T) {
	mockClient := mocks.NewVPPAssignmentsMock()
	svc := vpp_assignments.NewVppAssignments(mockClient)

	req := &vpp_assignments.Resource{
		General: vpp_assignments.SubsetGeneral{
			Name: "",
		},
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "VPP assignment name is required")
}

func TestUnit_VPPAssignments_UpdateByID(t *testing.T) {
	mockClient := mocks.NewVPPAssignmentsMock()
	mockClient.RegisterUpdateVPPAssignmentByIDMock()
	svc := vpp_assignments.NewVppAssignments(mockClient)

	req := &vpp_assignments.Resource{
		General: vpp_assignments.SubsetGeneral{
			Name:                "Updated VPP Assignment",
			VPPAdminAccountID:   1,
			VPPAdminAccountName: "Test Account",
		},
	}

	resp, _, err := svc.UpdateByID(context.Background(), 1, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_VPPAssignments_UpdateByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewVPPAssignmentsMock()
	svc := vpp_assignments.NewVppAssignments(mockClient)

	req := &vpp_assignments.Resource{
		General: vpp_assignments.SubsetGeneral{Name: "Test"},
	}

	_, _, err := svc.UpdateByID(context.Background(), 0, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "VPP assignment ID must be a positive integer")
}

func TestUnit_VPPAssignments_UpdateByID_NilRequest(t *testing.T) {
	mockClient := mocks.NewVPPAssignmentsMock()
	svc := vpp_assignments.NewVppAssignments(mockClient)

	_, _, err := svc.UpdateByID(context.Background(), 1, nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_VPPAssignments_DeleteByID(t *testing.T) {
	mockClient := mocks.NewVPPAssignmentsMock()
	mockClient.RegisterDeleteVPPAssignmentByIDMock()
	svc := vpp_assignments.NewVppAssignments(mockClient)

	_, err := svc.DeleteByID(context.Background(), 1)

	require.NoError(t, err)
}

func TestUnit_VPPAssignments_DeleteByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewVPPAssignmentsMock()
	svc := vpp_assignments.NewVppAssignments(mockClient)

	_, err := svc.DeleteByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "VPP assignment ID must be a positive integer")
}

func TestUnit_VPPAssignments_NotFound(t *testing.T) {
	mockClient := mocks.NewVPPAssignmentsMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := vpp_assignments.NewVppAssignments(mockClient)

	_, _, err := svc.GetByID(context.Background(), 999)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}

func TestUnit_VPPAssignments_Conflict(t *testing.T) {
	mockClient := mocks.NewVPPAssignmentsMock()
	mockClient.RegisterConflictErrorMock()
	svc := vpp_assignments.NewVppAssignments(mockClient)

	req := &vpp_assignments.Resource{
		General: vpp_assignments.SubsetGeneral{
			Name: "Duplicate Assignment",
		},
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "VPP assignment with that name already exists")
}
