package jamf_account_preferences

// ResourceAccountPreferences represents the Jamf Pro account preferences resource (GET/PATCH).
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-account-preferences
type ResourceAccountPreferences struct {
	Language                             string `json:"language"`
	DateFormat                           string `json:"dateFormat"`
	Timezone                             string `json:"timezone"`
	ResultsPerPage                       int    `json:"resultsPerPage"`
	UserInterfaceDisplayTheme            string `json:"userInterfaceDisplayTheme"`
	DisableRelativeDates                 bool   `json:"disableRelativeDates"`
	DisablePageLeaveCheck                bool   `json:"disablePageLeaveCheck"`
	DisableTablePagination               bool   `json:"disableTablePagination"`
	DisableShortcutsTooltips             bool   `json:"disableShortcutsTooltips"`
	ConfigProfilesSortingMethod          string `json:"configProfilesSortingMethod"`
	ComputerSearchMethod                 string `json:"computerSearchMethod"`
	ComputerApplicationSearchMethod      string `json:"computerApplicationSearchMethod"`
	ComputerApplicationUsageSearchMethod string `json:"computerApplicationUsageSearchMethod"`
	ComputerSoftwareUpdateSearchMethod   string `json:"computerSoftwareUpdateSearchMethod"`
	ComputerLocalUserAccountSearchMethod string `json:"computerLocalUserAccountSearchMethod"`
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
