package mac_applications

import (
	"encoding/xml"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

// ListResponse is the response for List (GET /JSSResource/macapplications).
type ListResponse struct {
	XMLName xml.Name   `xml:"mac_applications"`
	Size    int        `xml:"size"`
	Results []ListItem `xml:"mac_application"`
}

// ListItem represents a single Mac application item in the list.
type ListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Resource represents the detailed structure of a Mac application response.
type Resource struct {
	XMLName     xml.Name              `xml:"mac_application"`
	General     SubsetGeneral         `xml:"general"`
	Scope       SubsetScope           `xml:"scope"`
	SelfService SubsetSelfService     `xml:"self_service"`
	VPP         *SubsetVPP            `xml:"vpp,omitempty"`
}

// SubsetGeneral represents the general information of a Mac application.
type SubsetGeneral struct {
	ID             int                          `xml:"id"`
	Name           string                       `xml:"name"`
	Version        string                       `xml:"version"`
	IsFree         *bool                        `xml:"is_free"`
	BundleID       string                       `xml:"bundle_id"`
	URL            string                       `xml:"url"`
	Category       *shared.SharedResourceCategory `xml:"category,omitempty"`
	Site           *shared.SharedResourceSite    `xml:"site,omitempty"`
	DeploymentType string                       `xml:"deployment_type"`
}

// SubsetScope represents the scope of a Mac application.
type SubsetScope struct {
	AllComputers   *bool                 `xml:"all_computers"`
	AllJSSUsers    *bool                 `xml:"all_jss_users"`
	Buildings      []ScopeBuilding       `xml:"buildings>building"`
	Departments    []ScopeDepartment     `xml:"departments>department"`
	Computers      []ScopeComputer       `xml:"computers>computer"`
	ComputerGroups []ScopeComputerGroup  `xml:"computer_groups>computer_group"`
	JSSUsers       []ScopeUser           `xml:"jss_users>user"`
	JSSUserGroups  []ScopeJSSUserGroup   `xml:"jss_user_groups>user_group"`
	Limitations    ScopeLimitations      `xml:"limitations"`
	Exclusions     ScopeExclusions       `xml:"exclusions"`
}

// ScopeLimitations represents limitations within the scope.
type ScopeLimitations struct {
	Users           []ScopeUser           `xml:"users>user"`
	UserGroups      []ScopeUserGroup      `xml:"user_groups>user_group"`
	NetworkSegments []ScopeNetworkSegment `xml:"network_segments>network_segment"`
}

// ScopeExclusions represents exclusions within the scope.
type ScopeExclusions struct {
	Buildings       []ScopeBuilding       `xml:"buildings>building"`
	Departments     []ScopeDepartment     `xml:"departments>department"`
	Users           []ScopeUser           `xml:"users>user"`
	UserGroups      []ScopeUserGroup      `xml:"user_groups>user_group"`
	NetworkSegments []ScopeNetworkSegment `xml:"network_segments>network_segment"`
	Computers       []ScopeComputer       `xml:"computers>computer"`
	ComputerGroups  []ScopeComputerGroup  `xml:"computer_groups>computer_group"`
	JSSUsers        []ScopeUser           `xml:"jss_users>user"`
	JSSUserGroups  []ScopeJSSUserGroup   `xml:"jss_user_groups>user_group"`
}

// SubsetSelfService represents self-service configuration.
type SubsetSelfService struct {
	InstallButtonText           string                         `xml:"install_button_text"`
	SelfServiceDescription      string                         `xml:"self_service_description"`
	ForceUsersToViewDescription *bool                          `xml:"force_users_to_view_description"`
	SelfServiceIcon             shared.SharedResourceSelfServiceIcon `xml:"self_service_icon"`
	FeatureOnMainPage           *bool                          `xml:"feature_on_main_page"`
	SelfServiceCategories       []SelfServiceCategory          `xml:"self_service_categories>category"`
	Notification                string                         `xml:"notification"`
	NotificationSubject         string                         `xml:"notification_subject"`
	NotificationMessage         string                         `xml:"notification_message"`
}

// SelfServiceCategory represents a category in self-service.
type SelfServiceCategory struct {
	ID        int    `xml:"id"`
	Name      string `xml:"name"`
	DisplayIn *bool  `xml:"display_in"`
	FeatureIn *bool  `xml:"feature_in"`
}

// SubsetVPP represents VPP configuration.
type SubsetVPP struct {
	AssignVPPDeviceBasedLicenses *bool `xml:"assign_vpp_device_based_licenses"`
	VPPAdminAccountID            int   `xml:"vpp_admin_account_id"`
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

// ScopeUser represents a user in scope.
type ScopeUser struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ScopeUserGroup represents a user group in scope (limitations/exclusions).
type ScopeUserGroup struct {
	ID   string `xml:"id"`
	Name string `xml:"name"`
}

// ScopeJSSUserGroup represents a JSS user group in scope.
type ScopeJSSUserGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ScopeNetworkSegment represents a network segment in scope.
type ScopeNetworkSegment struct {
	ID   int    `xml:"id"`
	UID  string `xml:"uid,omitempty"`
	Name string `xml:"name"`
}

// CreateUpdateResponse represents the response from Create/Update operations.
type CreateUpdateResponse struct {
	XMLName xml.Name `xml:"mac_application"`
	ID      int      `xml:"id"`
}
