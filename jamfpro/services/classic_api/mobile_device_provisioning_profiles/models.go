package mobile_device_provisioning_profiles

import (
	"encoding/xml"
)

// ListResponse is the response for List (GET /JSSResource/mobiledeviceprovisioningprofiles).
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceprovisioningprofiles
type ListResponse struct {
	XMLName  xml.Name                            `xml:"mobile_device_provisioning_profiles"`
	Size     int                                 `xml:"size"`
	Profiles []MobileDeviceProvisioningProfileListItem `xml:"mobile_device_provisioning_profile"`
}

// MobileDeviceProvisioningProfileListItem represents a single mobile device provisioning profile item in the list.
type MobileDeviceProvisioningProfileListItem struct {
	ID          int    `xml:"id"`
	Name        string `xml:"name"`
	DisplayName string `xml:"display_name"`
	UUID        string `xml:"uuid"`
}

// Resource represents the detailed structure of a mobile device provisioning profile.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceprovisioningprofilesbyid
type Resource struct {
	XMLName xml.Name `xml:"mobile_device_provisioning_profile"`
	General SubsetGeneral `xml:"general"`
}

// RequestResource is the body for creating or updating a mobile device provisioning profile.
// The target is specified via the URL path (id/0, id/N, name/X, or uuid/U for create/update).
type RequestResource struct {
	XMLName xml.Name `xml:"mobile_device_provisioning_profile"`
	General SubsetGeneral `xml:"general"`
}

// SubsetGeneral represents the general subset of a mobile device provisioning profile.
type SubsetGeneral struct {
	ID          int    `xml:"id,omitempty"`
	Name        string `xml:"name"`
	DisplayName string `xml:"display_name,omitempty"`
	UUID        string `xml:"uuid,omitempty"`
}
