package directory_cache

type ResourceSyncStatus struct {
	State              string  `json:"state"`
	DirectoryServiceID *string `json:"directoryServiceId"`
	StartTime          *string `json:"startTime"`
	CurrentStep        *string `json:"currentStep"`
}
