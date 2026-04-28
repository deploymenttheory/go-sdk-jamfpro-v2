package cloud_information

// ResourceCloudInformation represents information related to cloud setup.
type ResourceCloudInformation struct {
	CloudInstance                  bool `json:"cloudInstance"`
	RampInstance                   bool `json:"rampInstance"`
	GovCloudInstance               bool `json:"govCloudInstance"`
	ManagedServiceProviderInstance bool `json:"managedServiceProviderInstance"`
	FedRampInstance                bool `json:"fedRampInstance"`
	FipsEnabled                    bool `json:"fipsEnabled"`
	HighComplianceInstance         bool `json:"highComplianceInstance"`
}
