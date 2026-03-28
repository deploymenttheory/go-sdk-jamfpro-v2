package enrollment

// EnrollmentSettings.flushMdmCommandsOnReenroll constants.
const (
	FlushMdmCommandsOnReenrollDeleteNothing                      = "DELETE_NOTHING"
	FlushMdmCommandsOnReenrollDeleteErrors                       = "DELETE_ERRORS"
	FlushMdmCommandsOnReenrollDeleteEverythingExceptAcknowledged = "DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED"
	FlushMdmCommandsOnReenrollDeleteEverything                   = "DELETE_EVERYTHING"
)

var validFlushMdmCommandsOnReenroll = map[string]struct{}{
	FlushMdmCommandsOnReenrollDeleteNothing:                      {},
	FlushMdmCommandsOnReenrollDeleteErrors:                       {},
	FlushMdmCommandsOnReenrollDeleteEverythingExceptAcknowledged: {},
	FlushMdmCommandsOnReenrollDeleteEverything:                   {},
}
