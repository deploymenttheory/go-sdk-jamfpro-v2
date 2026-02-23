package slasa

// ResourceSLASAStatus represents the SLASA (Software License Agreement Service Acceptance) status response.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-slasa
type ResourceSLASAStatus struct {
	SLASAAcceptanceStatus string `json:"slasaAcceptanceStatus"` // Whether SLASA has been accepted (e.g., "ACCEPTED", "NOT_ACCEPTED")
}
