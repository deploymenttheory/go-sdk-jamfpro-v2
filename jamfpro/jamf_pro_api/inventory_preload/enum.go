package inventory_preload

// DeviceType enum values for InventoryPreloadRecord.DeviceType.
const (
	DeviceTypeComputer     = "Computer"
	DeviceTypeMobileDevice = "Mobile Device"
	DeviceTypeUnknown      = "Unknown"
)

var validDeviceType = map[string]struct{}{
	DeviceTypeComputer:     {},
	DeviceTypeMobileDevice: {},
	DeviceTypeUnknown:      {},
}
