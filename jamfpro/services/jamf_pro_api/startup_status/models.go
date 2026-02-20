package startup_status

// ResourceStartupStatusV1 represents the Jamf Pro server startup status response.
type ResourceStartupStatusV1 struct {
	Step                    string `json:"step"`
	StepCode                string `json:"stepCode"`
	StepParam               string `json:"stepParam"`
	Percentage              int    `json:"percentage"`
	Warning                 string `json:"warning"`
	WarningCode             string `json:"warningCode"`
	WarningParam            string `json:"warningParam"`
	Error                   string `json:"error"`
	ErrorCode               string `json:"errorCode"`
	SetupAssistantNecessary bool   `json:"setupAssistantNecessary"`
}
