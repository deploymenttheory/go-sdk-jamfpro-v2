package computer_invitations

import (
	"encoding/xml"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"
)

// ListResponse is the response for List (GET /JSSResource/computerinvitations).
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computerinvitations
type ListResponse struct {
	XMLName            xml.Name                 `xml:"computer_invitations"`
	Size               int                      `xml:"size"`
	ComputerInvitation []ComputerInvitationItem `xml:"computer_invitation"`
}

// ComputerInvitationItem represents a single computer invitation item in the list.
type ComputerInvitationItem struct {
	ID                  int    `xml:"id,omitempty"`
	Invitation          int64  `xml:"invitation,omitempty"`
	InvitationType      string `xml:"invitation_type,omitempty"`
	ExpirationDate      string `xml:"expiration_date,omitempty"`
	ExpirationDateUTC   string `xml:"expiration_date_utc,omitempty"`
	ExpirationDateEpoch int64  `xml:"expiration_date_epoch,omitempty"`
}

// ResourceComputerInvitation represents a Jamf Pro Classic API computer invitation resource.
type ResourceComputerInvitation struct {
	XMLName                     xml.Name                            `xml:"computer_invitation"`
	ID                          int                                 `xml:"id,omitempty"`
	Invitation                  string                              `xml:"invitation,omitempty"`
	InvitationStatus            string                              `xml:"invitation_status,omitempty"`
	InvitationType              string                              `xml:"invitation_type,omitempty"`
	ExpirationDate              string                              `xml:"expiration_date,omitempty"`
	ExpirationDateUTC           string                              `xml:"expiration_date_utc,omitempty"`
	ExpirationDateEpoch         int64                               `xml:"expiration_date_epoch,omitempty"`
	SSHUsername                 string                              `xml:"ssh_username,omitempty"`
	SSHPassword                 string                              `xml:"ssh_password,omitempty"`
	MultipleUsersAllowed        bool                                `xml:"multiple_users_allowed,omitempty"`
	TimesUsed                   int                                 `xml:"times_used,omitempty"`
	CreateAccountIfDoesNotExist bool                                `xml:"create_account_if_does_not_exist,omitempty"`
	HideAccount                 bool                                `xml:"hide_account,omitempty"`
	LockDownSSH                 bool                                `xml:"lock_down_ssh,omitempty"`
	InvitedUserUUID             string                              `xml:"invited_user_uuid,omitempty"`
	EnrollIntoSite              *ComputerInvitationSubsetEnrollIntoState `xml:"enroll_into_site,omitempty"`
	KeepExistingSiteMembership  bool                                `xml:"keep_existing_site_membership,omitempty"`
	Site                        *models.SharedResourceSite           `xml:"site,omitempty"`
}

// ComputerInvitationSubsetEnrollIntoState represents the enroll-into-site subset.
type ComputerInvitationSubsetEnrollIntoState struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}
