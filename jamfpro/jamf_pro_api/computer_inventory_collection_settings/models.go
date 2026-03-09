package computer_inventory_collection_settings

type ResourceComputerInventoryCollectionSettings struct {
	ComputerInventoryCollectionPreferences SubsetPreferences    `json:"computerInventoryCollectionPreferences"`
	ApplicationPaths                       []SubsetPathResponse `json:"applicationPaths"`
}

type SubsetPreferences struct {
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

type SubsetPathResponse struct {
	ID   string `json:"id"`
	Path string `json:"path"`
}

type SubsetPathItem struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

type RequestCustomPath struct {
	Scope string `json:"scope"`
	Path  string `json:"path"`
}
