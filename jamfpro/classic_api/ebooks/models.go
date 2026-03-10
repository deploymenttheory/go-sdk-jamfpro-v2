package ebooks

import (
	"encoding/xml"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"
)

// ListResponse is the response for List (GET /JSSResource/ebooks).
type ListResponse struct {
	XMLName xml.Name       `xml:"ebooks"`
	Size    int            `xml:"size"`
	Results []EbookListItem `xml:"ebook"`
}

// EbookListItem represents a single ebook item in the list.
type EbookListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Resource represents the detailed structure of an ebook response.
type Resource struct {
	XMLName     xml.Name              `xml:"ebook"`
	General     SubsetGeneral         `xml:"general"`
	Scope       SubsetScope           `xml:"scope"`
	SelfService SubsetSelfService     `xml:"self_service"`
}

// SubsetGeneral represents the general information of an ebook.
type SubsetGeneral struct {
	ID              int                            `xml:"id"`
	Name            string                         `xml:"name"`
	Author          string                         `xml:"author"`
	Version         string                         `xml:"version"`
	Free            bool                           `xml:"free"`
	URL             string                         `xml:"url"`
	DeploymentType  string                         `xml:"deployment_type"`
	FileType        string                         `xml:"file_type"`
	DeployAsManaged bool                           `xml:"deploy_as_managed"`
	Category        *models.SharedResourceCategory `xml:"category,omitempty"`
	SelfServiceIcon models.SharedResourceSelfServiceIcon `xml:"self_service_icon"`
	Site            models.SharedResourceSite      `xml:"site"`
}

// SubsetScope represents the scope of an ebook.
type SubsetScope struct {
	AllComputers       bool                          `xml:"all_computers"`
	AllMobileDevices   bool                          `xml:"all_mobile_devices"`
	AllJSSUsers        bool                          `xml:"all_jss_users"`
	Computers          []ScopeComputer               `xml:"computers>computer"`
	ComputerGroups     []ScopeComputerGroup          `xml:"computer_groups>computer_group"`
	MobileDevices      []ScopeMobileDevice           `xml:"mobile_devices>mobile_device"`
	MobileDeviceGroups []ScopeMobileDeviceGroup      `xml:"mobile_device_groups>mobile_device_group"`
	Buildings          []ScopeBuilding               `xml:"buildings>building"`
	Departments        []ScopeDepartment             `xml:"departments>department"`
	JSSUsers           []ScopeUser                   `xml:"jss_users>user"`
	JSSUserGroups      []ScopeUserGroup              `xml:"jss_user_groups>user_group"`
	Classes            []ScopeClass                  `xml:"classes>class"`
	Limitations        SubsetScopeLimitations        `xml:"limitations"`
	Exclusions         SubsetScopeExclusions         `xml:"exclusions"`
}

// SubsetScopeLimitations represents limitations within the scope.
type SubsetScopeLimitations struct {
	NetworkSegments []ScopeNetworkSegment `xml:"network_segments>network_segment"`
	Users           []ScopeUser           `xml:"users>user"`
	UserGroups      []ScopeUserGroup     `xml:"user_groups>user_group"`
}

// ScopeNetworkSegment represents a network segment in scope.
type ScopeNetworkSegment struct {
	ID   int    `xml:"id"`
	UID  string `xml:"uid,omitempty"`
	Name string `xml:"name"`
}

// SubsetScopeExclusions represents exclusions within the scope.
type SubsetScopeExclusions struct {
	Computers          []ScopeComputer          `xml:"computers>computer"`
	ComputerGroups     []ScopeComputerGroup     `xml:"computer_groups>computer_group"`
	MobileDevices      []ScopeMobileDevice      `xml:"mobile_devices>mobile_device"`
	MobileDeviceGroups []ScopeMobileDeviceGroup `xml:"mobile_device_groups>mobile_device_group"`
	Buildings          []ScopeBuilding          `xml:"buildings>building"`
	Departments        []ScopeDepartment        `xml:"departments>department"`
	JSSUsers           []ScopeUser              `xml:"jss_users>user"`
	JSSUserGroups      []ScopeUserGroup         `xml:"jss_user_groups>user_group"`
}

// ScopeClass represents a class within the scope.
type ScopeClass struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// SubsetSelfService represents self-service configuration.
type SubsetSelfService struct {
	SelfServiceDisplayName      string                        `xml:"self_service_display_name"`
	InstallButtonText           string                        `xml:"install_button_text"`
	SelfServiceDescription      string                        `xml:"self_service_description"`
	ForceUsersToViewDescription bool                          `xml:"force_users_to_view_description"`
	SelfServiceIcon             models.SharedResourceSelfServiceIcon `xml:"self_service_icon"`
	FeatureOnMainPage           bool                          `xml:"feature_on_main_page"`
	SelfServiceCategories       SelfServiceCategories         `xml:"self_service_categories"`
	Notification                bool                          `xml:"notification"`
	NotificationSubject         string                        `xml:"notification_subject"`
	NotificationMessage         string                        `xml:"notification_message"`
}

// SelfServiceCategories represents categories within self-service.
type SelfServiceCategories struct {
	Category []SelfServiceCategoryItem `xml:"category"`
}

// SelfServiceCategoryItem represents a single category in self-service.
type SelfServiceCategoryItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ScopeComputer represents a computer in scope.
type ScopeComputer struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
	UDID string `xml:"udid"`
}

// ScopeComputerGroup represents a computer group in scope.
type ScopeComputerGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ScopeMobileDevice represents a mobile device in scope.
type ScopeMobileDevice struct {
	ID             int    `xml:"id"`
	Name           string `xml:"name"`
	UDID           string `xml:"udid"`
	WiFiMacAddress string `xml:"wifi_mac_address"`
}

// ScopeMobileDeviceGroup represents a mobile device group in scope.
type ScopeMobileDeviceGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ScopeBuilding represents a building in scope.
type ScopeBuilding struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ScopeDepartment represents a department in scope.
type ScopeDepartment struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ScopeUser represents a user in scope.
type ScopeUser struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ScopeUserGroup represents a user group in scope.
type ScopeUserGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// CreateUpdateResponse represents the response from Create/Update operations.
type CreateUpdateResponse struct {
	XMLName xml.Name `xml:"ebook"`
	ID      int      `xml:"id"`
}
