package licensed_software

import (
	"encoding/xml"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

// ListResponse is the response for List (GET /JSSResource/licensedsoftware).
// The API returns a root element "licensed_software" containing repeated "licensed_software" items.
type ListResponse struct {
	XMLName xml.Name              `xml:"licensed_software"`
	Size    int                   `xml:"size,omitempty"`
	Results []LicensedSoftwareItem `xml:"licensed_software"`
}

// LicensedSoftwareItem represents a single licensed software item in the list.
type LicensedSoftwareItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Resource represents the detailed structure of a licensed software resource.
type Resource struct {
	XMLName             xml.Name                    `xml:"licensed_software"`
	General             SubsetGeneral                `xml:"general"`
	SoftwareDefinitions []SubsetSoftwareDefinition   `xml:"software_definitions>definition"`
	FontDefinitions     []SubsetFontDefinition      `xml:"font_definitions>definition"`
	PluginDefinitions   []SubsetPluginDefinition   `xml:"plugin_definitions>definition"`
	Licenses            []SubsetLicenseContainer    `xml:"licenses>license"`
}

// SubsetGeneral represents the general information of licensed software.
type SubsetGeneral struct {
	ID                                 int                      `xml:"id"`
	Name                               string                   `xml:"name"`
	Publisher                          string                   `xml:"publisher"`
	Platform                           string                   `xml:"platform"`
	SendEmailOnViolation               bool                     `xml:"send_email_on_violation"`
	RemoveTitlesFromInventoryReports   bool                     `xml:"remove_titles_from_inventory_reports"`
	ExcludeTitlesPurchasedFromAppStore bool                     `xml:"exclude_titles_purchased_from_app_store"`
	Notes                              string                   `xml:"notes"`
	Site                               shared.SharedResourceSite `xml:"site"`
}

// SubsetSoftwareDefinition represents a software definition.
type SubsetSoftwareDefinition struct {
	CompareType string `xml:"compare_type"`
	Name        string `xml:"name"`
	Version     int    `xml:"version"`
}

// SubsetFontDefinition represents a font definition.
type SubsetFontDefinition struct {
	CompareType string `xml:"compare_type"`
	Name        string `xml:"name"`
	Version     int    `xml:"version"`
}

// SubsetPluginDefinition represents a plugin definition.
type SubsetPluginDefinition struct {
	CompareType string `xml:"compare_type"`
	Name        string `xml:"name"`
	Version     int    `xml:"version"`
}

// SubsetLicenseContainer wraps a license with size.
type SubsetLicenseContainer struct {
	Size    int              `xml:"size"`
	License SubsetLicense    `xml:"license"`
}

// SubsetLicense represents license details.
type SubsetLicense struct {
	SerialNumber1    string                      `xml:"serial_number_1"`
	SerialNumber2    string                      `xml:"serial_number_2"`
	OrganizationName string                      `xml:"organization_name"`
	RegisteredTo    string                      `xml:"registered_to"`
	LicenseType     string                      `xml:"license_type"`
	LicenseCount    int                         `xml:"license_count"`
	Notes           string                      `xml:"notes"`
	Purchasing      SubsetLicensePurchasing     `xml:"purchasing"`
	Attachments     []SubsetLicenseAttachment   `xml:"attachments>attachment"`
}

// SubsetLicensePurchasing represents purchasing information for a license.
type SubsetLicensePurchasing struct {
	IsPerpetual         bool   `xml:"is_perpetual"`
	IsAnnual            bool   `xml:"is_annual"`
	PONumber            string `xml:"po_number"`
	Vendor              string `xml:"vendor"`
	PurchasePrice       string `xml:"purchase_price"`
	PurchasingAccount   string `xml:"purchasing_account"`
	PODate              string `xml:"po_date"`
	PODateEpoch         int64  `xml:"po_date_epoch"`
	PODateUTC           string `xml:"po_date_utc"`
	LicenseExpires      string `xml:"license_expires"`
	LicenseExpiresEpoch int64  `xml:"license_expires_epoch"`
	LicenseExpiresUTC   string `xml:"license_expires_utc"`
	LifeExpectancy      int    `xml:"life_expectancy"`
	PurchasingContact   string `xml:"purchasing_contact"`
}

// SubsetLicenseAttachment represents an attachment on a license.
type SubsetLicenseAttachment struct {
	ID       int    `xml:"id"`
	Filename string `xml:"filename"`
	URI      string `xml:"uri"`
}

// CreateUpdateResponse represents the response from Create/Update operations.
type CreateUpdateResponse struct {
	XMLName xml.Name `xml:"licensed_software"`
	ID      int      `xml:"id"`
}
