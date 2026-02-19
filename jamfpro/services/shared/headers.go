package shared

import "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"

// JSONHeaders returns the standard JSON content headers for Jamf Pro API requests.
func JSONHeaders() map[string]string {
	return map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}
}

// XMLHeaders returns the standard XML content headers for Classic API requests.
func XMLHeaders() map[string]string {
	return map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}
}
