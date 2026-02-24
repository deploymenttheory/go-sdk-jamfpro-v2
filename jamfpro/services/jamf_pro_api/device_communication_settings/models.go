package device_communication_settings

// ResourceDeviceCommunicationSettings represents the device communication settings resource.
type ResourceDeviceCommunicationSettings struct {
	AutoRenewMobileDeviceMdmProfileWhenCaRenewed                   bool `json:"autoRenewMobileDeviceMdmProfileWhenCaRenewed"`
	AutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring  bool `json:"autoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring"`
	AutoRenewComputerMdmProfileWhenCaRenewed                       bool `json:"autoRenewComputerMdmProfileWhenCaRenewed"`
	AutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring     bool `json:"autoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring"`
	MdmProfileMobileDeviceExpirationLimitInDays                   int  `json:"mdmProfileMobileDeviceExpirationLimitInDays"`
	MdmProfileComputerExpirationLimitInDays                       int  `json:"mdmProfileComputerExpirationLimitInDays"`
}

// HistoryItem represents a single device communication settings history entry.
type HistoryItem struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

// HistoryResponse is the response for GetHistoryV1.
type HistoryResponse struct {
	TotalCount int           `json:"totalCount"`
	Results    []HistoryItem `json:"results"`
}
