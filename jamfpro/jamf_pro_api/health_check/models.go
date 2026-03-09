package health_check

// No response body; 200 indicates healthy.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-health-check

// ResourceHealthStatus represents request acceptance ratios for each concurrency group and time window.
// Returned by GetHealthStatusV1. Only available in Jamf Cloud; returns 404 on non-cloud nodes.
type ResourceHealthStatus struct {
	API        ResourceHealthStatusMetric `json:"api"`
	UI         ResourceHealthStatusMetric `json:"ui"`
	Enrollment ResourceHealthStatusMetric `json:"enrollment"`
	Device     ResourceHealthStatusMetric `json:"device"`
	Default    ResourceHealthStatusMetric `json:"default"`
}

// ResourceHealthStatusMetric holds acceptance ratio metrics for a time window.
// Values are between 0 and 1 (1 = all requests accepted, 0 = all denied).
type ResourceHealthStatusMetric struct {
	ThirtySeconds  float64 `json:"thirtySeconds"`
	OneMinute      float64 `json:"oneMinute"`
	FiveMinutes    float64 `json:"fiveMinutes"`
	FifteenMinutes float64 `json:"fifteenMinutes"`
	ThirtyMinutes  float64 `json:"thirtyMinutes"`
}
