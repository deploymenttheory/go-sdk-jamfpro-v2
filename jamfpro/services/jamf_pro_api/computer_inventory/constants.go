package computer_inventory

// API endpoint constants for Computer Inventory operations.
const (
	// EndpointComputerInventoryV1 is the base endpoint for Computer Inventory operations (v1 API).
	EndpointComputerInventoryV1 = "/api/v1/computers-inventory"
)

// ComputerInventorySections contains all available sections for computer inventory API requests.
var ComputerInventorySections = []string{
	"GENERAL", "DISK_ENCRYPTION", "PURCHASING", "APPLICATIONS", "STORAGE",
	"USER_AND_LOCATION", "CONFIGURATION_PROFILES", "PRINTERS", "SERVICES",
	"HARDWARE", "LOCAL_USER_ACCOUNTS", "CERTIFICATES", "ATTACHMENTS",
	"PLUGINS", "PACKAGE_RECEIPTS", "FONTS", "SECURITY", "OPERATING_SYSTEM",
	"LICENSED_SOFTWARE", "IBEACONS", "SOFTWARE_UPDATES", "EXTENSION_ATTRIBUTES",
	"CONTENT_CACHING", "GROUP_MEMBERSHIPS",
}
