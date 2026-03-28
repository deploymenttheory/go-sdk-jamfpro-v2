package mdm

// MdmCommandState* constants represent the possible states of an MDM command.
const (
	MdmCommandStatePending      = "PENDING"
	MdmCommandStateAcknowledged = "ACKNOWLEDGED"
	MdmCommandStateNotNow       = "NOT_NOW"
	MdmCommandStateError        = "ERROR"
)

// MdmClientType* constants represent the types of MDM client devices.
const (
	MdmClientTypeMobileDevice     = "MOBILE_DEVICE"
	MdmClientTypeTV               = "TV"
	MdmClientTypeVisionPro        = "VISION_PRO"
	MdmClientTypeWatch            = "WATCH"
	MdmClientTypeComputer         = "COMPUTER"
	MdmClientTypeComputerUser     = "COMPUTER_USER"
	MdmClientTypeMobileDeviceUser = "MOBILE_DEVICE_USER"
	MdmClientTypeUnknown          = "UNKNOWN"
)
