package managed_software_updates

// ManagedSoftwareUpdatePlanUpdateAction represents the update action for a managed software update plan.
const (
	UpdateActionDownloadOnly                  = "DOWNLOAD_ONLY"
	UpdateActionDownloadInstall               = "DOWNLOAD_INSTALL"
	UpdateActionDownloadInstallAllowDeferral  = "DOWNLOAD_INSTALL_ALLOW_DEFERRAL"
	UpdateActionDownloadInstallRestart        = "DOWNLOAD_INSTALL_RESTART"
	UpdateActionDownloadInstallSchedule       = "DOWNLOAD_INSTALL_SCHEDULE"
	UpdateActionUnknown                       = "UNKNOWN"
)

// ManagedSoftwareUpdatePlanVersionType represents the version type for a managed software update plan.
const (
	VersionTypeLatestMajor     = "LATEST_MAJOR"
	VersionTypeLatestMinor     = "LATEST_MINOR"
	VersionTypeLatestAny       = "LATEST_ANY"
	VersionTypeSpecificVersion = "SPECIFIC_VERSION"
	VersionTypeCustomVersion   = "CUSTOM_VERSION"
	VersionTypeUnknown         = "UNKNOWN"
)

// ManagedSoftwareUpdateStatusObjectType represents the object type for a managed software update status.
const (
	StatusObjectTypeComputer     = "COMPUTER"
	StatusObjectTypeMobileDevice = "MOBILE_DEVICE"
	StatusObjectTypeAppleTv      = "APPLE_TV"
)

// ManagedSoftwareUpdateStatusStatus represents the status of a managed software update.
const (
	UpdateStatusDownloading                   = "DOWNLOADING"
	UpdateStatusIdle                          = "IDLE"
	UpdateStatusInstalling                    = "INSTALLING"
	UpdateStatusInstalled                     = "INSTALLED"
	UpdateStatusError                         = "ERROR"
	UpdateStatusDownloadFailed                = "DOWNLOAD_FAILED"
	UpdateStatusDownloadRequiresComputer      = "DOWNLOAD_REQUIRES_COMPUTER"
	UpdateStatusDownloadInsufficientSpace     = "DOWNLOAD_INSUFFICIENT_SPACE"
	UpdateStatusDownloadInsufficientPower     = "DOWNLOAD_INSUFFICIENT_POWER"
	UpdateStatusDownloadInsufficientNetwork   = "DOWNLOAD_INSUFFICIENT_NETWORK"
	UpdateStatusInstallInsufficientSpace      = "INSTALL_INSUFFICIENT_SPACE"
	UpdateStatusInstallInsufficientPower      = "INSTALL_INSUFFICIENT_POWER"
	UpdateStatusInstallPhoneCallInProgress    = "INSTALL_PHONE_CALL_IN_PROGRESS"
	UpdateStatusInstallFailed                 = "INSTALL_FAILED"
	UpdateStatusUnknown                       = "UNKNOWN"
)

// MacOsManagedSoftwareUpdatePriority represents the priority of a macOS managed software update.
const (
	MacOsUpdatePriorityHigh = "HIGH"
	MacOsUpdatePriorityLow  = "LOW"
)

// MacOsManagedSoftwareUpdateUpdateAction represents the update action for a macOS managed software update.
const (
	MacOsUpdateActionDownloadAndInstall = "DOWNLOAD_AND_INSTALL"
	MacOsUpdateActionDownloadOnly       = "DOWNLOAD_ONLY"
)

// ManagedSoftwareUpdatePlanToggleStatusExitState represents the exit state of a plan toggle status.
const (
	ToggleExitStateUnknown   = "UNKNOWN"
	ToggleExitStateExecuting = "EXECUTING"
	ToggleExitStateCompleted = "COMPLETED"
	ToggleExitStateNoop      = "NOOP"
	ToggleExitStateFailed    = "FAILED"
	ToggleExitStateStopped   = "STOPPED"
)

// ManagedSoftwareUpdatePlanToggleStatusState represents the running state of a plan toggle status.
const (
	ToggleStateNotRunning = "NOT_RUNNING"
	ToggleStateRunning    = "RUNNING"
	ToggleStateNeverRan   = "NEVER_RAN"
)

// PlanGroupObjectType represents the object type for a plan group.
const (
	PlanGroupObjectTypeComputerGroup     = "COMPUTER_GROUP"
	PlanGroupObjectTypeMobileDeviceGroup = "MOBILE_DEVICE_GROUP"
)

// PlanDeviceObjectType represents the object type for a plan device.
const (
	PlanDeviceObjectTypeComputer     = "COMPUTER"
	PlanDeviceObjectTypeMobileDevice = "MOBILE_DEVICE"
	PlanDeviceObjectTypeAppleTv      = "APPLE_TV"
)
