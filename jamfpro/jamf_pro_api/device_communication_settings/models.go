package device_communication_settings

import "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"

// ResourceDeviceCommunicationSettings represents the device communication settings resource.
type ResourceDeviceCommunicationSettings struct {
	AutoRenewMobileDeviceMdmProfileWhenCaRenewed                   bool `json:"autoRenewMobileDeviceMdmProfileWhenCaRenewed"`
	AutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring  bool `json:"autoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring"`
	AutoRenewComputerMdmProfileWhenCaRenewed                       bool `json:"autoRenewComputerMdmProfileWhenCaRenewed"`
	AutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring     bool `json:"autoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring"`
	MdmProfileMobileDeviceExpirationLimitInDays                   int  `json:"mdmProfileMobileDeviceExpirationLimitInDays"`
	MdmProfileComputerExpirationLimitInDays                       int  `json:"mdmProfileComputerExpirationLimitInDays"`
}

// HistoryItem is an alias to the shared history item struct with string IDs.
type HistoryItem = shared.SharedHistoryItemString

// HistoryResponse is an alias to the shared history response struct with string IDs.
type HistoryResponse = shared.SharedHistoryResponseString

// RequestAddHistoryNotes is an alias to the shared history note request struct.
type RequestAddHistoryNotes = shared.SharedHistoryNoteRequest

// ResponseAddHistoryNotes is the response for AddHistoryNotesV1 (201 Created).
type ResponseAddHistoryNotes struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}
