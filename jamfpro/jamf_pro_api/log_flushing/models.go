package log_flushing

// ResourceLogFlushingSettings represents the log flushing settings configuration.
type ResourceLogFlushingSettings struct {
	RetentionPolicies []ResourceLogRetentionPolicy `json:"retentionPolicies"` // List of log retention policies
	HourOfDay         int                          `json:"hourOfDay"`         // Hour of day (0-23) when log flushing occurs
}

// ResourceLogRetentionPolicy represents a single log retention policy.
type ResourceLogRetentionPolicy struct {
	DisplayName         string `json:"displayName"`         // Human-readable name of the log type
	Qualifier           string `json:"qualifier"`           // Log qualifier (e.g., "JAMFSoftwareServer", "ApacheAccess")
	RetentionPeriod     int    `json:"retentionPeriod"`     // Number of retention period units
	RetentionPeriodUnit string `json:"retentionPeriodUnit"` // Unit of retention period (e.g., "Days", "Weeks", "Months")
}

// RequestLogFlushingTask represents a request to create a log flushing task.
type RequestLogFlushingTask struct {
	Qualifier           string `json:"qualifier"`           // Log qualifier to flush
	RetentionPeriod     int    `json:"retentionPeriod"`     // Number of retention period units
	RetentionPeriodUnit string `json:"retentionPeriodUnit"` // Unit of retention period
}

// CreateResponse represents the response after creating a log flushing task.
type CreateResponse struct {
	ID   string `json:"id"`   // Unique identifier of the created task
	Href string `json:"href"` // API href to the created task
}

// ResourceLogFlushingTask represents a log flushing task with full details.
type ResourceLogFlushingTask struct {
	ID                  string `json:"id"`                  // Unique identifier of the task
	Qualifier           string `json:"qualifier"`           // Log qualifier being flushed
	RetentionPeriod     int    `json:"retentionPeriod"`     // Number of retention period units
	RetentionPeriodUnit string `json:"retentionPeriodUnit"` // Unit of retention period
	State               string `json:"state"`               // Task state (e.g., "QUEUED", "RUNNING", "COMPLETED")
}
