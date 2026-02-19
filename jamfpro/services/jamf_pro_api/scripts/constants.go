package scripts

// Endpoints for the scripts API (Jamf Pro API v1).
const (
	EndpointScriptsV1 = "/api/v1/scripts"
)

// Script priority values for the Priority field.
const (
	ScriptPriorityBefore   = "BEFORE"
	ScriptPriorityAfter    = "AFTER"
	ScriptPriorityAtReboot = "AT_REBOOT"
)
