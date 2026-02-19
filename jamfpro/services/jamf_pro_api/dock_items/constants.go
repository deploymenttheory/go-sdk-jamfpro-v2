package dock_items

// Endpoints for the dock items API (Jamf Pro API v1).
const (
	EndpointDockItemsV1 = "/api/v1/dock-items"
)

// Dock item type constants (API expects uppercase: APP, FILE, FOLDER).
const (
	TypeApp    = "APP"
	TypeFile   = "FILE"
	TypeFolder = "FOLDER"
)
