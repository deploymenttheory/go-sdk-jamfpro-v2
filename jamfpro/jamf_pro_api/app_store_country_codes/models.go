package app_store_country_codes

// AppStoreCountryCode represents a single country code entry for App Store locale.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/app-store-country-codes
type AppStoreCountryCode struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// ListResponse is the response for ListV1.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/app-store-country-codes
type ListResponse struct {
	CountryCodes []AppStoreCountryCode `json:"countryCodes"`
}
