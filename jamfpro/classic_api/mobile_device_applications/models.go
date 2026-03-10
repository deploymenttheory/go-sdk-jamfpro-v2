package mobile_device_applications

import (
	"encoding/xml"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"
)

// ListResponse is the response for List (GET /JSSResource/mobiledeviceapplications).
type ListResponse struct {
	XMLName xml.Name   `xml:"mobile_device_applications"`
	Size    int        `xml:"size"`
	Results []ListItem `xml:"mobile_device_application"`
}

// ListItem represents a single mobile device application item in the list.
type ListItem struct {
	ID          int    `xml:"id"`
	Name        string `xml:"name"`
	DisplayName string `xml:"display_name"`
	BundleID    string `xml:"bundle_id"`
	Version     string `xml:"version"`
	InternalApp *bool  `xml:"internal_app"`
}

// Resource represents the detailed structure of a mobile device application response.
type Resource struct {
	XMLName         xml.Name                 `xml:"mobile_device_application"`
	General         SubsetGeneral            `xml:"general"`
	Scope           SubsetScope              `xml:"scope"`
	SelfService     SubsetSelfService        `xml:"self_service"`
	VPP             *SubsetVPP               `xml:"vpp,omitempty"`
	AppConfiguration *SubsetAppConfiguration `xml:"app_configuration,omitempty"`
}

// SubsetGeneral represents the general information of a mobile device application.
type SubsetGeneral struct {
	ID                               int                          `xml:"id,omitempty"`
	Name                             string                       `xml:"name"`
	DisplayName                      string                       `xml:"display_name"`
	Description                      string                       `xml:"description,omitempty"`
	BundleID                         string                       `xml:"bundle_id"`
	Version                          string                       `xml:"version"`
	InternalApp                      *bool                        `xml:"internal_app,omitempty"`
	OsType                           string                       `xml:"os_type,omitempty"`
	Category                         *models.SharedResourceCategory `xml:"category"`
	IPA                              SubsetIPA                    `xml:"ipa,omitempty"`
	Icon                             SubsetIcon                   `xml:"icon"`
	ProvisioningProfile              int                          `xml:"mobile_device_provisioning_profile,omitempty"`
	ITunesStoreURL                   string                       `xml:"itunes_store_url,omitempty"`
	MakeAvailableAfterInstall        bool                         `xml:"make_available_after_install,omitempty"`
	ITunesCountryRegion              string                       `xml:"itunes_country_region,omitempty"`
	ITunesSyncTime                   int                          `xml:"itunes_sync_time,omitempty"`
	DeploymentType                   string                       `xml:"deployment_type,omitempty"`
	DeployAutomatically              *bool                        `xml:"deploy_automatically,omitempty"`
	DeployAsManagedApp               *bool                        `xml:"deploy_as_managed_app,omitempty"`
	RemoveAppWhenMDMProfileIsRemoved *bool                        `xml:"remove_app_when_mdm_profile_is_removed,omitempty"`
	PreventBackupOfAppData           *bool                        `xml:"prevent_backup_of_app_data,omitempty"`
	AllowUserToDelete                *bool                        `xml:"allow_user_to_delete,omitempty"`
	KeepDescriptionAndIconUpToDate   *bool                        `xml:"keep_description_and_icon_up_to_date,omitempty"`
	KeepAppUpdatedOnDevices          *bool                        `xml:"keep_app_updated_on_devices,omitempty"`
	Free                             *bool                        `xml:"free,omitempty"`
	TakeOverManagement               *bool                        `xml:"take_over_management,omitempty"`
	HostExternally                   *bool                        `xml:"host_externally,omitempty"`
	ExternalURL                      string                       `xml:"external_url,omitempty"`
	Site                             *models.SharedResourceSite   `xml:"site"`
}

// SubsetIPA represents IPA configuration for mobile device applications.
type SubsetIPA struct {
	Name string `xml:"name,omitempty"`
	URI  string `xml:"uri,omitempty"`
	Data string `xml:"data,omitempty"`
}

// SubsetSelfService represents self-service configuration for mobile device applications.
type SubsetSelfService struct {
	SelfServiceDescription string                              `xml:"self_service_description,omitempty"`
	SelfServiceIcon       SubsetIcon                          `xml:"self_service_icon,omitempty"`
	FeatureOnMainPage     *bool                               `xml:"feature_on_main_page,omitempty"`
	SelfServiceCategories []models.SharedResourceSelfServiceCategory `xml:"self_service_categories>category,omitempty"`
	Notification          *bool                               `xml:"notification,omitempty"`
	NotificationSubject   string                              `xml:"notification_subject,omitempty"`
	NotificationMessage   string                              `xml:"notification_message,omitempty"`
}

// SubsetVPP represents VPP configuration for mobile device applications.
type SubsetVPP struct {
	AssignVPPDeviceBasedLicenses *bool `xml:"assign_vpp_device_based_licenses,omitempty"`
	VPPAdminAccountID           int   `xml:"vpp_admin_account_id,omitempty"`
}

// SubsetAppConfiguration represents app configuration preferences.
type SubsetAppConfiguration struct {
	Preferences string `xml:"preferences"`
}

// SubsetScope represents the scope of a mobile device application.
type SubsetScope struct {
	AllMobileDevices   *bool                 `xml:"all_mobile_devices,omitempty"`
	AllJSSUsers        *bool                 `xml:"all_jss_users,omitempty"`
	MobileDevices      []ScopeMobileDevice   `xml:"mobile_devices>mobile_device,omitempty"`
	Buildings          []ScopeEntity         `xml:"buildings>building,omitempty"`
	Departments        []ScopeEntity         `xml:"departments>department,omitempty"`
	MobileDeviceGroups []ScopeEntity         `xml:"mobile_device_groups>mobile_device_group,omitempty"`
	JSSUsers           []ScopeEntity         `xml:"jss_users>user,omitempty"`
	JSSUserGroups      []ScopeEntity         `xml:"jss_user_groups>user_group,omitempty"`
	Limitations        SubsetLimitation      `xml:"limitations,omitempty"`
	Exclusions         SubsetExclusion      `xml:"exclusions,omitempty"`
}

// SubsetLimitation represents limitations within the scope.
type SubsetLimitation struct {
	Users           []ScopeEntity         `xml:"users>user,omitempty"`
	UserGroups      []ScopeEntity         `xml:"user_groups>user_group,omitempty"`
	NetworkSegments []ScopeNetworkSegment `xml:"network_segments>network_segment,omitempty"`
}

// SubsetExclusion represents exclusions within the scope.
type SubsetExclusion struct {
	MobileDevices      []ScopeMobileDevice   `xml:"mobile_devices>mobile_device,omitempty"`
	Buildings          []ScopeEntity         `xml:"buildings>building,omitempty"`
	Users              []ScopeEntity         `xml:"users>user,omitempty"`
	UserGroups         []ScopeEntity         `xml:"user_groups>user_group,omitempty"`
	Departments        []ScopeEntity         `xml:"departments>department,omitempty"`
	MobileDeviceGroups []ScopeEntity         `xml:"mobile_device_groups>mobile_device_group,omitempty"`
	NetworkSegments    []ScopeNetworkSegment `xml:"network_segments>network_segment,omitempty"`
	JSSUsers           []ScopeEntity         `xml:"jss_users>user,omitempty"`
	JSSUserGroups      []ScopeEntity         `xml:"jss_user_groups>user_group,omitempty"`
}

// SubsetIcon represents an icon for mobile device applications.
type SubsetIcon struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name"`
	URI  string `xml:"uri,omitempty"`
	Data string `xml:"data,omitempty"`
}

// ScopeMobileDevice represents a mobile device in scope.
type ScopeMobileDevice struct {
	ID             int    `xml:"id"`
	Name           string `xml:"name,omitempty"`
	UDID           string `xml:"udid,omitempty"`
	WifiMacAddress string `xml:"wifi_mac_address,omitempty"`
}

// ScopeEntity represents a generic entity in scope (building, department, user, etc.).
type ScopeEntity struct {
	ID   int    `xml:"id"`
	Name string `xml:"name,omitempty"`
}

// ScopeNetworkSegment represents a network segment in scope.
type ScopeNetworkSegment struct {
	ScopeEntity
	Name string `xml:"name,omitempty"`
}

// CreateUpdateResponse represents the response from Create operations.
type CreateUpdateResponse struct {
	XMLName xml.Name `xml:"mobile_device_application"`
	ID      int      `xml:"id"`
}
