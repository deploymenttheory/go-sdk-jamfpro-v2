package mdm_renewal

// DeviceCommonDetails represents device common details for MDM profile renewal.
// Returned by GetDeviceCommonDetailsV1.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mdm-renewal-device-common-details-clientmanagementid
type DeviceCommonDetails struct {
	ID                                      string  `json:"id,omitempty"`
	ClientManagementID                      string  `json:"clientManagementId,omitempty"`
	RenewMdmProfileStartDate                *string `json:"renewMdmProfileStartDate,omitempty"`
	MdmProfileNeedsRenewalDueToCaRenewed    bool    `json:"mdmProfileNeedsRenewalDueToCaRenewed"`
	MdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring bool `json:"mdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring"`
	MdmCheckinUrl                           *string `json:"mdmCheckinUrl,omitempty"`
	MdmServerUrl                            *string `json:"mdmServerUrl,omitempty"`
}

// RequestDeviceCommonDetailsUpdate is the body for UpdateDeviceCommonDetailsV1 (PATCH).
// All fields except ClientManagementID are optional; only provided fields are updated.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-mdm-renewal-device-common-details
type RequestDeviceCommonDetailsUpdate struct {
	ClientManagementID                                   string  `json:"clientManagementId"` // required
	RenewMdmProfileStartDate                             *string `json:"renewMdmProfileStartDate,omitempty"`
	MdmProfileNeedsRenewalDueToCaRenewed                  *bool   `json:"mdmProfileNeedsRenewalDueToCaRenewed,omitempty"`
	MdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring *bool   `json:"mdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring,omitempty"`
	MdmCheckinUrl                                        *string `json:"mdmCheckinUrl,omitempty"`
	MdmServerUrl                                         *string `json:"mdmServerUrl,omitempty"`
}

// MDMRenewalErrorType represents the type of MDM renewal error.
type MDMRenewalErrorType string

const (
	MDMRenewalErrorTypeServerError   MDMRenewalErrorType = "SERVER_ERROR"
	MDMRenewalErrorTypeCheckInError  MDMRenewalErrorType = "CHECK_IN_ERROR"
	MDMRenewalErrorTypeOther         MDMRenewalErrorType = "OTHER"
)

// MDMRenewalStrategyType represents the type of MDM renewal strategy.
type MDMRenewalStrategyType string

const (
	MDMRenewalStrategyTypeReturnNoCheckInInvitation                    MDMRenewalStrategyType = "RETURN_NO_CHECK_IN_INVITATION"
	MDMRenewalStrategyTypeReturnCheckInInvitationFromMdmInvitationTable MDMRenewalStrategyType = "RETURN_CHECK_IN_INVITATION_FROM_MDM_INVITATION_TABLE"
	MDMRenewalStrategyTypeReturnCheckInInvitationFromEnrollmentUsage   MDMRenewalStrategyType = "RETURN_CHECK_IN_INVITATION_FROM_ENROLLMENT_USAGE_TABLE"
	MDMRenewalStrategyTypeReturnCheckInInvitationFromMdmProfilePrototype MDMRenewalStrategyType = "RETURN_CHECK_IN_INVITATION_FROM_MDM_PROFILE_PROTOTYPE_TABLE"
	MDMRenewalStrategyTypeJssUrlOverride                               MDMRenewalStrategyType = "JSS_URL_OVERRIDE"
	MDMRenewalStrategyTypePayloadIdentifier                            MDMRenewalStrategyType = "PAYLOAD_IDENTIFIER"
)

// MDMRenewalError represents an MDM renewal error.
type MDMRenewalError struct {
	MdmRenewalErrorId   string               `json:"mdmRenewalErrorId"`
	ClientManagementId  string               `json:"clientManagementId"`
	MdmRenewalErrorType MDMRenewalErrorType  `json:"mdmRenewalErrorType"`
	ErrorTimeStamp      string               `json:"errorTimeStamp,omitempty"`
	FailureCount        int                  `json:"failureCount,omitempty"`
}

// RenewalStrategy represents a renewal strategy associated with an MDM renewal error.
type RenewalStrategy struct {
	ID                   string                 `json:"id"`
	MdmRenewalErrorId    string                 `json:"mdmRenewalErrorId"`
	MdmRenewalStrategyType MDMRenewalStrategyType `json:"mdmRenewalStrategyType"`
	StrategyTimeStamp    string                 `json:"strategyTimeStamp,omitempty"`
	MdmRenewalCheckInUrl string                 `json:"mdmRenewalCheckInUrl,omitempty"`
	MdmRenewalServerUrl  string                 `json:"mdmRenewalServerUrl,omitempty"`
}

// RenewalErrorWithStrategies represents an MDM renewal error with its associated strategies.
// Returned by GetRenewalStrategiesV1.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mdm-renewal-renewal-strategies-clientmanagementid
type RenewalErrorWithStrategies struct {
	Error      MDMRenewalError  `json:"error"`
	Strategies []RenewalStrategy `json:"strategies"`
}
