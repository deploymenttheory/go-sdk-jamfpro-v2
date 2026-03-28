package reenrollment

// Reenrollment.flushMDMQueue constants.
const (
	FlushMDMQueueDeleteNothing                      = "DELETE_NOTHING"
	FlushMDMQueueDeleteErrors                       = "DELETE_ERRORS"
	FlushMDMQueueDeleteEverythingExceptAcknowledged = "DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED"
	FlushMDMQueueDeleteEverything                   = "DELETE_EVERYTHING"
)

var validFlushMDMQueue = map[string]struct{}{
	FlushMDMQueueDeleteNothing:                      {},
	FlushMDMQueueDeleteErrors:                       {},
	FlushMDMQueueDeleteEverythingExceptAcknowledged: {},
	FlushMDMQueueDeleteEverything:                   {},
}
