package computer_inventory_collection_settings

// ResourceComputerInventoryCollectionSettings represents the computer inventory collection settings.
type ResourceComputerInventoryCollectionSettings struct {
	ComputerInventoryCollectionPreferences Preferences    `json:"computerInventoryCollectionPreferences"`
	ApplicationPaths                       []PathResponse `json:"applicationPaths"`
}

// Preferences represents the computer inventory collection preferences.
type Preferences struct {
	MonitorApplicationUsage                      bool `json:"monitorApplicationUsage"`
	IncludePackages                              bool `json:"includePackages"`
	IncludeSoftwareUpdates                       bool `json:"includeSoftwareUpdates"`
	IncludeSoftwareId                            bool `json:"includeSoftwareId"`
	IncludeAccounts                              bool `json:"includeAccounts"`
	CalculateSizes                               bool `json:"calculateSizes"`
	IncludeHiddenAccounts                        bool `json:"includeHiddenAccounts"`
	IncludePrinters                              bool `json:"includePrinters"`
	IncludeServices                              bool `json:"includeServices"`
	CollectSyncedMobileDeviceInfo                bool `json:"collectSyncedMobileDeviceInfo"`
	UpdateLdapInfoOnComputerInventorySubmissions bool `json:"updateLdapInfoOnComputerInventorySubmissions"`
	MonitorBeacons                               bool `json:"monitorBeacons"`
	AllowChangingUserAndLocation                 bool `json:"allowChangingUserAndLocation"`
	UseUnixUserPaths                             bool `json:"useUnixUserPaths"`
	CollectUnmanagedCertificates                 bool `json:"collectUnmanagedCertificates"`
}

// PathResponse represents an application path in the response.
type PathResponse struct {
	ID   string `json:"id"`
	Path string `json:"path"`
}

// CustomPathRequest represents the request body for creating a custom path.
type CustomPathRequest struct {
	Scope string `json:"scope"`
	Path  string `json:"path"`
}

// CustomPathResponse represents the response after creating a custom path.
type CustomPathResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}
