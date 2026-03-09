package inventory_information

// ResourceInventoryInformation represents statistics about managed/unmanaged devices and computers in the inventory.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-inventory-information
type ResourceInventoryInformation struct {
	ManagedComputers   int `json:"managedComputers"`
	UnmanagedComputers int `json:"unmanagedComputers"`
	ManagedDevices     int `json:"managedDevices"`
	UnmanagedDevices   int `json:"unmanagedDevices"`
}
