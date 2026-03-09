package mobile_device_groups

import (
	"encoding/xml"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
)

// ResourceMobileDeviceGroup represents a Jamf Pro Classic API mobile device group resource.
type ResourceMobileDeviceGroup struct {
	XMLName                 xml.Name                       `xml:"mobile_device_group"`
	ID                      int                            `xml:"id,omitempty"`
	Name                    string                         `xml:"name"`
	IsSmart                 bool                           `xml:"is_smart"`
	Criteria                *CriteriaContainer             `xml:"criteria,omitempty"`
	Site                    *shared.SharedResourceSite     `xml:"site,omitempty"`
	MobileDevices           []MobileDeviceSubsetItem       `xml:"mobile_devices>mobile_device,omitempty"`
	MobileDeviceAdditions   []MobileDeviceSubsetItem       `xml:"mobile_device_additions>mobile_device,omitempty"`
	MobileDeviceDeletions   []MobileDeviceSubsetItem       `xml:"mobile_device_deletions>mobile_device,omitempty"`
}

// ListResponse is the response for List (GET /JSSResource/mobiledevicegroups).
type ListResponse struct {
	XMLName xml.Name               `xml:"mobile_device_groups"`
	Size    int                    `xml:"size"`
	Results []MobileDeviceGroupItem `xml:"mobile_device_group"`
}

// MobileDeviceGroupItem represents a single mobile device group item in the list.
type MobileDeviceGroupItem struct {
	ID      int    `xml:"id"`
	Name    string `xml:"name"`
	IsSmart bool   `xml:"is_smart"`
}

// RequestMobileDeviceGroup is the body for creating or updating a mobile device group.
// The ID field is not included; the target is specified via the URL path.
type RequestMobileDeviceGroup struct {
	XMLName               xml.Name                   `xml:"mobile_device_group"`
	Name                  string                     `xml:"name"`
	IsSmart               bool                       `xml:"is_smart"`
	Site                  *shared.SharedResourceSite  `xml:"site,omitempty"`
	Criteria              *CriteriaContainer         `xml:"criteria,omitempty"`
	MobileDevices         []MobileDeviceSubsetItem   `xml:"mobile_devices>mobile_device,omitempty"`
	MobileDeviceAdditions []MobileDeviceSubsetItem   `xml:"mobile_device_additions>mobile_device,omitempty"`
	MobileDeviceDeletions []MobileDeviceSubsetItem   `xml:"mobile_device_deletions>mobile_device,omitempty"`
}

// CreateUpdateResponse represents the response from Create/Update operations.
// The Classic API returns only the ID for these operations.
type CreateUpdateResponse struct {
	XMLName xml.Name `xml:"mobile_device_group"`
	ID      int      `xml:"id"`
}

// CriteriaContainer wraps the criteria for smart mobile device groups.
type CriteriaContainer struct {
	Size      int                          `xml:"size,omitempty"`
	Criterion []shared.SharedSubsetCriteria `xml:"criterion,omitempty"`
}

// MobileDeviceSubsetItem represents a single mobile device within a group.
type MobileDeviceSubsetItem struct {
	ID             int    `xml:"id"`
	Name           string `xml:"name"`
	MacAddress     string `xml:"mac_address,omitempty"`
	UDID           string `xml:"udid"`
	WifiMacAddress string `xml:"wifi_mac_address,omitempty"`
	SerialNumber   string `xml:"serial_number,omitempty"`
}
