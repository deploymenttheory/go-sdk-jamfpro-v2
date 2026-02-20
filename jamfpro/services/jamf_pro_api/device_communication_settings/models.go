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
