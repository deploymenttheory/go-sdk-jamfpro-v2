package scripts

// ScriptPriority enum values for ResourceScript.Priority / RequestScript.Priority.
const (
	ScriptPriorityBefore   = "BEFORE"
	ScriptPriorityAfter    = "AFTER"
	ScriptPriorityAtReboot = "AT_REBOOT"
)

// validScriptPriorities is the set of accepted Priority values.
var validScriptPriorities = map[string]struct{}{
	ScriptPriorityBefore:   {},
	ScriptPriorityAfter:    {},
	ScriptPriorityAtReboot: {},
}
