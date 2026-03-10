package macos_configuration_profiles

import (
	"encoding/xml"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"
)

// ListResponse is the response for List (GET /JSSResource/osxconfigurationprofiles).
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findosxconfigurationprofiles
type ListResponse struct {
	XMLName xml.Name           `xml:"os_x_configuration_profiles"`
	Size    int                `xml:"size,omitempty"`
	Results []ConfigurationProfileListItem `xml:"os_x_configuration_profile,omitempty"`
}

// ConfigurationProfileListItem represents a single macOS configuration profile item in the list.
type ConfigurationProfileListItem struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name"`
}

// Resource represents the detailed structure of a macOS configuration profile.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findosxconfigurationprofilesbyid
type Resource struct {
	XMLName     xml.Name                 `xml:"os_x_configuration_profile"`
	General     SubsetGeneral            `xml:"general"`
	Scope       *SubsetScope             `xml:"scope,omitempty"`
	SelfService *SubsetSelfService       `xml:"self_service,omitempty"`
}

// RequestResource is the body for creating or updating a macOS configuration profile.
// The target is specified via the URL path (id/0 for create, id/N or name/X for update).
type RequestResource struct {
	XMLName     xml.Name           `xml:"os_x_configuration_profile"`
	General     SubsetGeneral      `xml:"general"`
	Scope       *SubsetScope       `xml:"scope,omitempty"`
	SelfService *SubsetSelfService `xml:"self_service,omitempty"`
}

// CreateUpdateResponse represents the response from Create/Update operations.
type CreateUpdateResponse struct {
	XMLName xml.Name `xml:"os_x_configuration_profile"`
	ID      int      `xml:"id"`
}

// SubsetGeneral represents the general subset of a macOS configuration profile.
type SubsetGeneral struct {
	ID                 int                          `xml:"id,omitempty"`
	Name               string                       `xml:"name"`
	Description        string                       `xml:"description,omitempty"`
	Site               *models.SharedResourceSite   `xml:"site,omitempty"`
	Category           *models.SharedResourceCategory `xml:"category,omitempty"`
	DistributionMethod string                       `xml:"distribution_method,omitempty"`
	UserRemovable      bool                         `xml:"user_removable"`
	Level              string                       `xml:"level,omitempty"`
	UUID               string                       `xml:"uuid,omitempty"`
	RedeployOnUpdate   string                       `xml:"redeploy_on_update,omitempty"`
	Payloads            string                       `xml:"payloads,omitempty"`
}

// SubsetScope represents the scope subset of a macOS configuration profile.
type SubsetScope struct {
	AllComputers   bool                `xml:"all_computers"`
	AllJSSUsers    bool                `xml:"all_jss_users"`
	Computers      []ScopeComputer     `xml:"computers>computer,omitempty"`
	ComputerGroups []ScopeEntity       `xml:"computer_groups>computer_group,omitempty"`
	JSSUsers       []ScopeEntity       `xml:"jss_users>user,omitempty"`
	JSSUserGroups  []ScopeEntity       `xml:"jss_user_groups>user_group,omitempty"`
	Buildings      []ScopeEntity       `xml:"buildings>building,omitempty"`
	Departments    []ScopeEntity       `xml:"departments>department,omitempty"`
	Limitations    *SubsetLimitations  `xml:"limitations,omitempty"`
	Exclusions     *SubsetExclusions   `xml:"exclusions,omitempty"`
}

// SubsetSelfService represents the self-service subset of a macOS configuration profile.
type SubsetSelfService struct {
	SelfServiceDisplayName      string                        `xml:"self_service_display_name,omitempty"`
	InstallButtonText           string                        `xml:"install_button_text,omitempty"`
	SelfServiceDescription      string                        `xml:"self_service_description,omitempty"`
	ForceUsersToViewDescription bool                          `xml:"force_users_to_view_description"`
	SelfServiceIcon             models.SharedResourceSelfServiceIcon `xml:"self_service_icon,omitempty"`
	FeatureOnMainPage           bool                         `xml:"feature_on_main_page"`
	SelfServiceCategories       []SelfServiceCategory         `xml:"self_service_categories>category,omitempty"`
	Notification                string                        `xml:"notification,omitempty"`
	NotificationSubject         string                        `xml:"notification_subject,omitempty"`
	NotificationMessage         string                        `xml:"notification_message,omitempty"`
}

// SelfServiceCategory represents a self-service category.
type SelfServiceCategory struct {
	ID        int    `xml:"id,omitempty"`
	Name      string `xml:"name,omitempty"`
	DisplayIn bool   `xml:"display_in,omitempty"`
	FeatureIn bool   `xml:"feature_in,omitempty"`
}

// SubsetLimitations represents the limitations subset.
type SubsetLimitations struct {
	Users           []ScopeEntity       `xml:"users>user,omitempty"`
	UserGroups      []ScopeUserGroup    `xml:"user_groups>user_group,omitempty"`
	NetworkSegments []ScopeNetworkSegment `xml:"network_segments>network_segment,omitempty"`
	IBeacons        []ScopeEntity       `xml:"ibeacons>ibeacon,omitempty"`
}

// SubsetExclusions represents the exclusions subset.
type SubsetExclusions struct {
	Computers       []ScopeComputer       `xml:"computers>computer,omitempty"`
	ComputerGroups  []ScopeEntity         `xml:"computer_groups>computer_group,omitempty"`
	Users           []ScopeEntity         `xml:"users>user,omitempty"`
	UserGroups      []ScopeUserGroup      `xml:"user_groups>user_group,omitempty"`
	Buildings       []ScopeEntity         `xml:"buildings>building,omitempty"`
	Departments     []ScopeEntity         `xml:"departments>department,omitempty"`
	NetworkSegments []ScopeNetworkSegment `xml:"network_segments>network_segment,omitempty"`
	JSSUsers        []ScopeEntity         `xml:"jss_users>user,omitempty"`
	JSSUserGroups   []ScopeEntity         `xml:"jss_user_groups>user_group,omitempty"`
	IBeacons        []ScopeEntity         `xml:"ibeacons>ibeacon,omitempty"`
}

// ScopeComputer represents a computer in scope (includes UDID).
type ScopeComputer struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
	UDID string `xml:"udid,omitempty"`
}

// ScopeNetworkSegment represents a network segment in scope.
type ScopeNetworkSegment struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
	UID  string `xml:"uid,omitempty"`
}

// ScopeEntity represents a generic scope entity (id, name).
type ScopeEntity struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// ScopeUserGroup represents a user group in scope (id can be string for Jamf Pro).
type ScopeUserGroup struct {
	ID   string `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}
