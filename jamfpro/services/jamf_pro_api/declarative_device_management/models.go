package declarative_device_management

// ResourceStatusItems represents the response containing status items for a device.
type ResourceStatusItems struct {
	StatusItems []StatusItem `json:"statusItems"`
}

// StatusItem represents a single status item in the status report.
type StatusItem struct {
	Key            string `json:"key"`
	Value          string `json:"value"`
	LastUpdateTime string `json:"lastUpdateTime"`
}
