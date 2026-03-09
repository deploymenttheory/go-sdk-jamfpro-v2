package mobile_device_configuration_profiles

import (
	"encoding/xml"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
)

// ListResponse is the response for List (GET /JSSResource/mobiledeviceconfigurationprofiles).
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceconfigurationprofiles
type ListResponse struct {
	XMLName             xml.Name                        `xml:"configuration_profiles"`
	ConfigurationProfiles []ConfigurationProfileListItem `xml:"configuration_profile"`
}

// ConfigurationProfileListItem represents a single mobile device configuration profile item in the list.
type ConfigurationProfileListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Resource represents the detailed structure of a mobile device configuration profile.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceconfigurationprofilesbyid
type Resource struct {
	XMLName     xml.Name          `xml:"configuration_profile"`
	General     SubsetGeneral     `xml:"general"`
	Scope       *SubsetScope      `xml:"scope,omitempty"`
	SelfService *SubsetSelfService `xml:"self_service,omitempty"`
}

// RequestResource is the body for creating or updating a mobile device configuration profile.
// The target is specified via the URL path (id/0 for create, id/N or name/X for update).
type RequestResource struct {
	XMLName     xml.Name           `xml:"configuration_profile"`
	General     SubsetGeneral      `xml:"general"`
	Scope       *SubsetScope       `xml:"scope,omitempty"`
	SelfService *SubsetSelfService `xml:"self_service,omitempty"`
}

// CreateUpdateResponse represents the response from Create/Update operations.
type CreateUpdateResponse struct {
	XMLName xml.Name `xml:"configuration_profile"`
	ID      int      `xml:"id"`
}

// SubsetGeneral represents the general subset of a mobile device configuration profile.
type SubsetGeneral struct {
	ID                            int                              `xml:"id,omitempty"`
	Name                          string                           `xml:"name"`
	Description                   string                           `xml:"description,omitempty"`
	Level                         string                           `xml:"level,omitempty"`
	Site                          *shared.SharedResourceSite       `xml:"site,omitempty"`
	Category                      *shared.SharedResourceCategory    `xml:"category,omitempty"`
	UUID                          string                           `xml:"uuid,omitempty"`
	DeploymentMethod              string                           `xml:"deployment_method,omitempty"`
	RedeployOnUpdate              string                           `xml:"redeploy_on_update,omitempty"`
	RedeployDaysBeforeCertExpires int                              `xml:"redeploy_days_before_certificate_expires,omitempty"`
	Payloads                      string                           `xml:"payloads,omitempty"`
}

// SubsetScope represents the scope subset of a mobile device configuration profile.
type SubsetScope struct {
	AllMobileDevices   bool                `xml:"all_mobile_devices,omitempty"`
	AllJSSUsers        bool                `xml:"all_jss_users,omitempty"`
	MobileDevices      []ScopeMobileDevice `xml:"mobile_devices>mobile_device,omitempty"`
	MobileDeviceGroups []ScopeEntity      `xml:"mobile_device_groups>mobile_device_group,omitempty"`
	JSSUsers           []ScopeEntity      `xml:"jss_users>user,omitempty"`
	JSSUserGroups      []ScopeEntity      `xml:"jss_user_groups>user_group,omitempty"`
	Buildings          []ScopeEntity      `xml:"buildings>building,omitempty"`
	Departments        []ScopeEntity      `xml:"departments>department,omitempty"`
	Limitations        *SubsetLimitation  `xml:"limitations,omitempty"`
	Exclusions         *SubsetExclusion   `xml:"exclusions,omitempty"`
}

// SubsetSelfService represents the self-service subset of a mobile device configuration profile.
type SubsetSelfService struct {
	SelfServiceDescription string                        `xml:"self_service_description,omitempty"`
	SecurityName           *SubsetSelfServiceSecurityName `xml:"security_name,omitempty"`
	SelfServiceIcon        shared.SharedResourceSelfServiceIcon `xml:"self_service_icon,omitempty"`
	FeatureOnMainPage      bool                          `xml:"feature_on_main_page,omitempty"`
	SelfServiceCategories  []shared.SharedResourceSelfServiceCategory `xml:"self_service_categories>category,omitempty"`
}

// SubsetSelfServiceSecurityName represents the security name for self-service removal.
type SubsetSelfServiceSecurityName struct {
	RemovalDisallowed string `xml:"removal_disallowed,omitempty"`
}

// SubsetLimitation represents the limitations subset.
type SubsetLimitation struct {
	NetworkSegments []ScopeNetworkSegment `xml:"network_segments>network_segment,omitempty"`
	Users           []ScopeEntity         `xml:"users>user,omitempty"`
	UserGroups      []ScopeEntity         `xml:"user_groups>user_group,omitempty"`
	Ibeacons        []ScopeEntity         `xml:"ibeacons>ibeacon,omitempty"`
}

// SubsetExclusion represents the exclusions subset.
type SubsetExclusion struct {
	MobileDevices      []ScopeMobileDevice   `xml:"mobile_devices>mobile_device,omitempty"`
	MobileDeviceGroups []ScopeEntity         `xml:"mobile_device_groups>mobile_device_group,omitempty"`
	Users              []ScopeEntity         `xml:"users>user,omitempty"`
	UserGroups         []ScopeEntity         `xml:"user_groups>user_group,omitempty"`
	Buildings          []ScopeEntity         `xml:"buildings>building,omitempty"`
	Departments        []ScopeEntity         `xml:"departments>department,omitempty"`
	NetworkSegments    []ScopeNetworkSegment `xml:"network_segments>network_segment,omitempty"`
	JSSUsers           []ScopeEntity         `xml:"jss_users>user,omitempty"`
	JSSUserGroups      []ScopeEntity         `xml:"jss_user_groups>user_group,omitempty"`
	IBeacons           []ScopeEntity         `xml:"ibeacons>ibeacon,omitempty"`
}

// ScopeMobileDevice represents a mobile device in scope.
type ScopeMobileDevice struct {
	ID             int    `xml:"id"`
	Name           string `xml:"name,omitempty"`
	UDID           string `xml:"udid,omitempty"`
	WifiMacAddress string `xml:"wifi_mac_address,omitempty"`
}

// ScopeEntity represents a generic scope entity (id, name).
type ScopeEntity struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// ScopeNetworkSegment represents a network segment in scope.
type ScopeNetworkSegment struct {
	ScopeEntity
	UID string `xml:"uid,omitempty"`
}
