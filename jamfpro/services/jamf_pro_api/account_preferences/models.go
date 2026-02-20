package account_preferences

// ResourceAccountPreferencesV2 represents the account preferences resource (get/update).
type ResourceAccountPreferencesV2 struct {
	Language                             string `json:"language"`
	DateFormat                           string `json:"dateFormat"`
	Timezone                             string `json:"timezone"`
	DisableRelativeDates                 bool   `json:"disableRelativeDates"`
	DisablePageLeaveCheck                bool   `json:"disablePageLeaveCheck"`
	DisableShortcutsTooltips             bool   `json:"disableShortcutsTooltips"`
	DisableTablePagination               bool   `json:"disableTablePagination"`
	ConfigProfilesSortingMethod          string `json:"configProfilesSortingMethod"`
	ResultsPerPage                       int    `json:"resultsPerPage"`
	UserInterfaceDisplayTheme            string `json:"userInterfaceDisplayTheme"`
	ComputerSearchMethod                 string `json:"computerSearchMethod"`
	ComputerApplicationSearchMethod      string `json:"computerApplicationSearchMethod"`
	ComputerApplicationUsageSearchMethod string `json:"computerApplicationUsageSearchMethod"`
	ComputerFontSearchMethod             string `json:"computerFontSearchMethod"`
	ComputerPluginSearchMethod           string `json:"computerPluginSearchMethod"`
	ComputerLocalUserAccountSearchMethod string `json:"computerLocalUserAccountSearchMethod"`
	ComputerSoftwareUpdateSearchMethod   string `json:"computerSoftwareUpdateSearchMethod"`
	ComputerPackageReceiptSearchMethod   string `json:"computerPackageReceiptSearchMethod"`
	ComputerPrinterSearchMethod          string `json:"computerPrinterSearchMethod"`
	ComputerPeripheralSearchMethod       string `json:"computerPeripheralSearchMethod"`
	ComputerServiceSearchMethod          string `json:"computerServiceSearchMethod"`
	MobileDeviceSearchMethod             string `json:"mobileDeviceSearchMethod"`
	MobileDeviceAppSearchMethod          string `json:"mobileDeviceAppSearchMethod"`
	UserSearchMethod                     string `json:"userSearchMethod"`
	UserAllContentSearchMethod           string `json:"userAllContentSearchMethod"`
	UserMobileDeviceAppSearchMethod      string `json:"userMobileDeviceAppSearchMethod"`
	UserMacAppStoreAppSearchMethod       string `json:"userMacAppStoreAppSearchMethod"`
	UserEbookSearchMethod                string `json:"userEbookSearchMethod"`
}
