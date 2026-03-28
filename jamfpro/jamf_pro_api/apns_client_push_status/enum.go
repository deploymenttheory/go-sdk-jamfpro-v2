package apns_client_push_status

// DeviceType* constants represent the types of devices for APNs client push status.
const (
	DeviceTypeMobileDevice     = "MOBILE_DEVICE"
	DeviceTypeMobileDeviceUser = "MOBILE_DEVICE_USER"
	DeviceTypeComputer         = "COMPUTER"
	DeviceTypeComputerUser     = "COMPUTER_USER"
	DeviceTypeTV               = "TV"
	DeviceTypeWatch            = "WATCH"
	DeviceTypeVisionPro        = "VISION_PRO"
	DeviceTypeUnknown          = "UNKNOWN"
)

// PushEnableStatus* constants represent the status values for APNs push enable requests.
const (
	PushEnableStatusQueued    = "QUEUED"
	PushEnableStatusStarted   = "STARTED"
	PushEnableStatusCompleted = "COMPLETED"
)
