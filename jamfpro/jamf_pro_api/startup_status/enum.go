package startup_status

// ErrorCode* constants represent the possible error codes in the startup status response.
const (
	ErrorCodeCacheConfigurationError      = "CACHE_CONFIGURATION_ERROR"
	ErrorCodeSecondaryNodeStartupError    = "SECONDARY_NODE_STARTUP_ERROR"
	ErrorCodeMoreThanOneClusterSettings   = "MORE_THAN_ONE_CLUSTER_SETTINGS_ERROR"
	ErrorCodePrimaryNodeNotSetError       = "PRIMARY_NODE_NOT_SET_ERROR"
	ErrorCodeDatabaseError                = "DATABASE_ERROR"
	ErrorCodeDatabasePasswordMissing      = "DATABASE_PASSWORD_MISSING"
	ErrorCodeEhcacheError                 = "EHCACHE_ERROR"
	ErrorCodeFlagInitializationFailed     = "FLAG_INITIALIZATION_FAILED"
	ErrorCodeMemcachedError               = "MEMCACHED_ERROR"
	ErrorCodeDatabaseMyisamError          = "DATABASE_MYISAM_ERROR"
	ErrorCodeOldVersionError              = "OLD_VERSION_ERROR"
)

// StepCode* constants represent the possible step codes in the startup status response.
const (
	StepCodeServerInitStart                              = "SERVER_INIT_START"
	StepCodeServerInitAnalyzingWebapp                    = "SERVER_INIT_ANALYZING_WEBAPP"
	StepCodeServerInitPopulatingNavigation               = "SERVER_INIT_POPULATING_NAVIGATION"
	StepCodeServerInitPopulatingObjects                  = "SERVER_INIT_POPULATING_OBJECTS"
	StepCodeServerInitInitializingObj                    = "SERVER_INIT_INITIALIZING_OBJ"
	StepCodeServerInitVerifyingCache                     = "SERVER_INIT_VERIFYING_CACHE"
	StepCodeServerInitInitializingChangeManagement       = "SERVER_INIT_INITIALIZING_CHANGE_MANAGEMENT"
	StepCodeServerInitInitializingCommunicationSystem    = "SERVER_INIT_INITIALIZING_COMMUNICATION_SYSTEM"
	StepCodeServerInitInitializingMdmQueueMonitor        = "SERVER_INIT_INITIALIZING_MDM_QUEUE_MONITOR"
	StepCodeServerInitCalculatingSmartGroups             = "SERVER_INIT_CALCULATING_SMART_GROUPS"
	StepCodeServerInitDbSchemaCompare                    = "SERVER_INIT_DB_SCHEMA_COMPARE"
	StepCodeServerInitDbTableCheckForRename              = "SERVER_INIT_DB_TABLE_CHECK_FOR_RENAME"
	StepCodeServerInitDbTableAlter                       = "SERVER_INIT_DB_TABLE_ALTER"
	StepCodeServerInitDbTableAnalyzing                   = "SERVER_INIT_DB_TABLE_ANALYZING"
	StepCodeServerInitDbTableCreate                      = "SERVER_INIT_DB_TABLE_CREATE"
	StepCodeServerInitDbTableDrop                        = "SERVER_INIT_DB_TABLE_DROP"
	StepCodeServerInitDbTableRename                      = "SERVER_INIT_DB_TABLE_RENAME"
	StepCodeServerInitDbColumnRename                     = "SERVER_INIT_DB_COLUMN_RENAME"
	StepCodeServerInitDbColumnEncodingChangeStep1        = "SERVER_INIT_DB_COLUMN_ENCODING_CHANGE_STEP_1"
	StepCodeServerInitDbColumnEncodingChangeStep2        = "SERVER_INIT_DB_COLUMN_ENCODING_CHANGE_STEP_2"
	StepCodeServerInitDbColumnEncodingChangeStep3        = "SERVER_INIT_DB_COLUMN_ENCODING_CHANGE_STEP_3"
	StepCodeServerInitDbUpgradeCheck                     = "SERVER_INIT_DB_UPGRADE_CHECK"
	StepCodeServerInitDbUpgradeComplete                  = "SERVER_INIT_DB_UPGRADE_COMPLETE"
	StepCodeServerInitSsGenerateNotifications            = "SERVER_INIT_SS_GENERATE_NOTIFICATIONS"
	StepCodeServerInitSsGenerateNotificationsStatus      = "SERVER_INIT_SS_GENERATE_NOTIFICATIONS_STATUS"
	StepCodeServerInitSsGenerateNotificationsFinalize    = "SERVER_INIT_SS_GENERATE_NOTIFICATIONS_FINALIZE"
	StepCodeServerInitPkiMigrationDone                   = "SERVER_INIT_PKI_MIGRATION_DONE"
	StepCodeServerInitPkiMigrationStatus                 = "SERVER_INIT_PKI_MIGRATION_STATUS"
	StepCodeServerInitMemcachedEndpointsCheck            = "SERVER_INIT_MEMCACHED_ENDPOINTS_CHECK"
	StepCodeServerInitCacheFlushing                      = "SERVER_INIT_CACHE_FLUSHING"
	StepCodeServerInitComplete                           = "SERVER_INIT_COMPLETE"
)

// WarningCode* constants represent the possible warning codes in the startup status response.
const (
	WarningCodeServerInitWarningDbTableEncoding = "SERVER_INIT_WARNING_DB_TABLE_ENCODING"
)
