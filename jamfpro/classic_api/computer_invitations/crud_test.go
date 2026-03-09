package computer_invitations_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/computer_invitations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/computer_invitations/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_ComputerInvitations_List(t *testing.T) {
	mockClient := mocks.NewComputerInvitationsMock()
	mockClient.RegisterListComputerInvitationsMock()
	svc := computer_invitations.NewComputerInvitations(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 2, resp.Size)
	assert.Len(t, resp.ComputerInvitation, 2)
	assert.Equal(t, int64(1234567890), resp.ComputerInvitation[0].Invitation)
	assert.Equal(t, "USER_INITIATED_ENROLLMENT", resp.ComputerInvitation[0].InvitationType)
	assert.Equal(t, int64(9876543210), resp.ComputerInvitation[1].Invitation)
}

func TestUnit_ComputerInvitations_GetByID(t *testing.T) {
	mockClient := mocks.NewComputerInvitationsMock()
	mockClient.RegisterGetComputerInvitationByIDMock()
	svc := computer_invitations.NewComputerInvitations(mockClient)

	resp, _, err := svc.GetByID(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "1234567890", resp.Invitation)
	assert.Equal(t, "PENDING", resp.InvitationStatus)
	assert.Equal(t, "USER_INITIATED_ENROLLMENT", resp.InvitationType)
	assert.NotNil(t, resp.Site)
	assert.Equal(t, -1, resp.Site.ID)
	assert.Equal(t, "None", resp.Site.Name)
	assert.NotNil(t, resp.EnrollIntoSite)
	assert.Equal(t, -1, resp.EnrollIntoSite.ID)
}

func TestUnit_ComputerInvitations_GetByID_EmptyID(t *testing.T) {
	mockClient := mocks.NewComputerInvitationsMock()
	svc := computer_invitations.NewComputerInvitations(mockClient)

	_, _, err := svc.GetByID(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "computer invitation ID cannot be empty")
}

func TestUnit_ComputerInvitations_GetByInvitationID(t *testing.T) {
	mockClient := mocks.NewComputerInvitationsMock()
	mockClient.RegisterGetComputerInvitationByInvitationIDMock()
	svc := computer_invitations.NewComputerInvitations(mockClient)

	resp, _, err := svc.GetByInvitationID(context.Background(), "1234567890")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "1234567890", resp.Invitation)
	assert.Equal(t, "PENDING", resp.InvitationStatus)
}

func TestUnit_ComputerInvitations_GetByInvitationID_EmptyID(t *testing.T) {
	mockClient := mocks.NewComputerInvitationsMock()
	svc := computer_invitations.NewComputerInvitations(mockClient)

	_, _, err := svc.GetByInvitationID(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "computer invitation invitation ID cannot be empty")
}

func TestUnit_ComputerInvitations_Create(t *testing.T) {
	mockClient := mocks.NewComputerInvitationsMock()
	mockClient.RegisterCreateComputerInvitationMock()
	svc := computer_invitations.NewComputerInvitations(mockClient)

	req := &computer_invitations.ResourceComputerInvitation{
		InvitationType:              "USER_INITIATED_ENROLLMENT",
		MultipleUsersAllowed:        false,
		CreateAccountIfDoesNotExist: true,
		KeepExistingSiteMembership:  true,
		Site: &shared.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		EnrollIntoSite: &computer_invitations.ComputerInvitationSubsetEnrollIntoState{
			ID:   -1,
			Name: "None",
		},
	}

	resp, _, err := svc.Create(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 42, resp.ID)
	assert.Equal(t, "5555555555", resp.Invitation)
	assert.Equal(t, "PENDING", resp.InvitationStatus)
}

func TestUnit_ComputerInvitations_Create_NilRequest(t *testing.T) {
	mockClient := mocks.NewComputerInvitationsMock()
	svc := computer_invitations.NewComputerInvitations(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_ComputerInvitations_DeleteByID(t *testing.T) {
	mockClient := mocks.NewComputerInvitationsMock()
	mockClient.RegisterDeleteComputerInvitationByIDMock()
	svc := computer_invitations.NewComputerInvitations(mockClient)

	_, err := svc.DeleteByID(context.Background(), "1")

	require.NoError(t, err)
}

func TestUnit_ComputerInvitations_DeleteByID_EmptyID(t *testing.T) {
	mockClient := mocks.NewComputerInvitationsMock()
	svc := computer_invitations.NewComputerInvitations(mockClient)

	_, err := svc.DeleteByID(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "computer invitation ID cannot be empty")
}

func TestUnit_ComputerInvitations_NotFound(t *testing.T) {
	mockClient := mocks.NewComputerInvitationsMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := computer_invitations.NewComputerInvitations(mockClient)

	_, _, err := svc.GetByID(context.Background(), "999")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}
