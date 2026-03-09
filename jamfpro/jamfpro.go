package jamfpro

import (
	"fmt"

	classic_accounts "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/accounts"
	classic_accounts_groups "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/accounts_groups"
	classic_activation_code "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/activation_code"
	classic_advanced_computer_searches "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/advanced_computer_searches"
	classic_advanced_user_searches "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/advanced_user_searches"
	classic_allowed_file_extensions "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/allowed_file_extensions"
	classic_byoprofiles "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/byoprofiles"
	classic_classes "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/classes"
	classic_command_flush "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/command_flush"
	classic_computer_groups "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/computer_groups"
	classic_computer_history "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/computer_history"
	classic_computer_inventory_collection "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/computer_inventory_collection"
	classic_computer_invitations "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/computer_invitations"
	classic_computers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/computers"
	classic_directory_bindings "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/directory_bindings"
	classic_disk_encryption_configurations "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/disk_encryption_configurations"
	classic_dock_items "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/dock_items"
	classic_ebooks "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/ebooks"
	classic_file_share_distribution_points "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/file_share_distribution_points"
	classic_file_uploads "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/file_uploads"
	classic_ibeacons "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/ibeacons"
	classic_ldap_servers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/ldap_servers"
	classic_licensed_software "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/licensed_software"
	classic_mac_applications "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mac_applications"
	classic_macos_configuration_profiles "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/macos_configuration_profiles"
	classic_mobile_device_applications "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_device_applications"
	classic_mobile_device_configuration_profiles "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_device_configuration_profiles"
	classic_mobile_device_enrollment_profiles "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_device_enrollment_profiles"
	classic_mobile_device_groups "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_device_groups"
	classic_mobile_device_provisioning_profiles "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_device_provisioning_profiles"
	classic_mobile_devices "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_devices"
	classic_network_segments "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/network_segments"
	classic_patch_external_sources "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/patch_external_sources"
	classic_policies "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/policies"
	classic_printers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/printers"
	classic_removeable_mac_addresses "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/removeable_mac_addresses"
	classic_restricted_software "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/restricted_software"
	classic_sites "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/sites"
	classic_smart_user_groups "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/smart_user_groups"
	classic_software_update_servers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/software_update_servers"
	classic_static_user_groups "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/static_user_groups"
	classic_user_extension_attributes "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/user_extension_attributes"
	classic_users "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/users"
	classic_vpp_accounts "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/vpp_accounts"
	classic_vpp_assignments "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/vpp_assignments"
	classic_webhooks "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/webhooks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/config"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/access_management_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/account_preferences"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/accounts"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/activation_code"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/adcs_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/adue_session_token_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/advanced_mobile_device_searches"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/advanced_user_content_searches"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/api_authorization"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/api_integrations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/api_role_privileges"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/api_roles"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/apns_client_push_status"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/app_installers"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/app_request"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/app_store_country_codes"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/bookmarks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/branding"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/buildings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/cache_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/categories"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/certificate_authority"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/classic_ldap"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/client_checkin"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/cloud_azure"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/cloud_distribution_point"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/cloud_idp"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/cloud_information"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/cloud_ldap"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/cloud_ldap_keystore"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/computer_extension_attributes"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/computer_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/computer_inventory"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/computer_inventory_collection_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/computer_prestages"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/conditional_access"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/csa"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/declarative_device_management"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/departments"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/device_communication_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/device_enrollments"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/devices"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/digicert"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/distribution_point"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/dock_items"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/dss_declarations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/ebooks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/engage"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/enrollment"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/enrollment_customization_preview"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/enrollment_customizations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/enrollment_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/gsx_connection"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/health_check"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/icon"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/impact_alert_notification_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/inventory_information"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/inventory_preload"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/jamf_account_preferences"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/jamf_connect"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/jamf_management_framework"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/jamf_package"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/jamf_pro_information"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/jamf_pro_notifications"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/jamf_pro_server_url"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/jamf_pro_system_initialization"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/jamf_pro_user_account_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/jamf_pro_version"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/jamf_protect"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/jamf_remote_assist"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/jcds"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/ldap"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/local_admin_password"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/locales"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/log_flushing"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/login_customization"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/macos_configuration_profile_custom_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/managed_software_updates"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/mdm"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/mdm_renewal"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/mobile_device_apps"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/mobile_device_enrollment_profile"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/mobile_device_extension_attributes"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/mobile_device_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/mobile_device_prestages"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/notifications"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/oauth2_session_tokens"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/oidc"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/onboarding"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/packages"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/patch_management"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/patch_policies"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/patch_software_title_configurations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/policy_properties"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/reenrollment"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/return_to_service"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/scripts"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/self_service_branding_ios"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/self_service_branding_macos"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/self_service_branding_upload"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/self_service_plus_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/self_service_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/service_discovery_enrollment"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/sites"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/slasa"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/smart_computer_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/smart_mobile_device_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/smtp_server"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/sso_certificate"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/sso_failover"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/sso_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/startup_status"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/static_computer_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/static_mobile_device_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/time_zones"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/tomcat_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/user"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/venafi"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/volume_purchasing_locations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/volume_purchasing_subscriptions"
	"go.uber.org/zap"
)

// Re-export client configuration options.
var (
	WithBaseURL               = client.WithBaseURL
	WithTimeout               = client.WithTimeout
	WithRetryCount            = client.WithRetryCount
	WithRetryWaitTime         = client.WithRetryWaitTime
	WithRetryMaxWaitTime      = client.WithRetryMaxWaitTime
	WithLogger                = client.WithLogger
	WithDebug                 = client.WithDebug
	WithUserAgent             = client.WithUserAgent
	WithGlobalHeader          = client.WithGlobalHeader
	WithGlobalHeaders         = client.WithGlobalHeaders
	WithProxy                 = client.WithProxy
	WithTLSClientConfig       = client.WithTLSClientConfig
	WithTransport             = client.WithTransport
	WithInsecureSkipVerify    = client.WithInsecureSkipVerify
	WithMaxConcurrentRequests = client.WithMaxConcurrentRequests
	WithMandatoryRequestDelay = client.WithMandatoryRequestDelay
	WithTotalRetryDuration    = client.WithTotalRetryDuration
)

// Re-export pagination helper functions.
var (
	HasNextPage          = client.HasNextPage
	ExtractParamsFromURL = client.ExtractParamsFromURL
)

// Re-export configuration types and functions for convenience.
// Users should only need to import the jamfpro package.
type (
	AuthConfig      = config.AuthConfig
	ClientOption    = client.ClientOption
	PaginationLinks = client.PaginationLinks
)

// ClassicAPIClient groups all Classic API services.
type ClassicAPIClient struct {
	Accounts                          *classic_accounts.Accounts
	AccountGroups                     *classic_accounts_groups.AccountsGroups
	ActivationCode                    *classic_activation_code.ActivationCode
	AdvancedComputerSearches          *classic_advanced_computer_searches.AdvancedComputerSearches
	AdvancedUserSearches              *classic_advanced_user_searches.AdvancedUserSearches
	AllowedFileExtensions             *classic_allowed_file_extensions.AllowedFileExtensions
	BYOProfiles                       *classic_byoprofiles.Byoprofiles
	Classes                           *classic_classes.Classes
	CommandFlush                      *classic_command_flush.CommandFlush
	ComputerGroups                    *classic_computer_groups.ComputerGroups
	Computers                         *classic_computers.Computers
	ComputerHistory                   *classic_computer_history.ComputerHistory
	ComputerInvitations               *classic_computer_invitations.ComputerInvitations
	ComputerInventoryCollection       *classic_computer_inventory_collection.ComputerInventoryCollection
	DirectoryBindings                 *classic_directory_bindings.DirectoryBindings
	DiskEncryptionConfigurations      *classic_disk_encryption_configurations.DiskEncryptionConfigurations
	DockItems                         *classic_dock_items.DockItems
	Ebooks                            *classic_ebooks.Ebooks
	FileShareDistributionPoints       *classic_file_share_distribution_points.FileShareDistributionPoints
	FileUploads                       *classic_file_uploads.FileUploads
	IBeacons                          *classic_ibeacons.Ibeacons
	LdapServers                       *classic_ldap_servers.LdapServers
	LicensedSoftware                  *classic_licensed_software.LicensedSoftware
	MacApplications                   *classic_mac_applications.MacApplications
	MacOSConfigurationProfiles        *classic_macos_configuration_profiles.MacosConfigurationProfiles
	MobileDeviceApplications          *classic_mobile_device_applications.MobileDeviceApplications
	MobileDeviceConfigurationProfiles *classic_mobile_device_configuration_profiles.MobileDeviceConfigurationProfiles
	MobileDeviceEnrollmentProfiles    *classic_mobile_device_enrollment_profiles.MobileDeviceEnrollmentProfiles
	MobileDeviceGroups                *classic_mobile_device_groups.MobileDeviceGroups
	MobileDeviceProvisioningProfiles  *classic_mobile_device_provisioning_profiles.MobileDeviceProvisioningProfiles
	MobileDevices                     *classic_mobile_devices.MobileDevices
	NetworkSegments                   *classic_network_segments.NetworkSegments
	PatchExternalSources              *classic_patch_external_sources.PatchExternalSources
	Policies                          *classic_policies.Policies
	Printers                          *classic_printers.Printers
	RemoveableMacAddresses            *classic_removeable_mac_addresses.RemoveableMacAddresses
	RestrictedSoftware                *classic_restricted_software.RestrictedSoftware
	Sites                             *classic_sites.Sites
	SmartUserGroups                   *classic_smart_user_groups.SmartUserGroups
	SoftwareUpdateServers             *classic_software_update_servers.SoftwareUpdateServers
	StaticUserGroups                  *classic_static_user_groups.StaticUserGroups
	UserExtensionAttributes           *classic_user_extension_attributes.UserExtensionAttributes
	Users                             *classic_users.Users
	VppAccounts                       *classic_vpp_accounts.VppAccounts
	VppAssignments                    *classic_vpp_assignments.VppAssignments
	Webhooks                          *classic_webhooks.Webhooks
}

// JamfProAPIClient groups all Jamf Pro API services.
type JamfProAPIClient struct {
	AccessManagementSettings            *access_management_settings.AccessManagementSettings
	AccountPreferences                  *account_preferences.AccountPreferences
	Accounts                            *accounts.Accounts
	ActivationCode                      *activation_code.ActivationCode
	AdcsSettings                        *adcs_settings.AdcsSettings
	AdueSessionTokenSettings            *adue_session_token_settings.AdueSessionTokenSettings
	AdvancedMobileDeviceSearches        *advanced_mobile_device_searches.AdvancedMobileDeviceSearches
	AdvancedUserContentSearches         *advanced_user_content_searches.AdvancedUserContentSearches
	ApiAuthorization                    *api_authorization.ApiAuthorization
	ApiIntegrations                     *api_integrations.ApiIntegrations
	ApiRolePrivileges                   *api_role_privileges.ApiRolePrivileges
	ApiRoles                            *api_roles.ApiRoles
	ApnsClientPushStatus                *apns_client_push_status.ApnsClientPushStatus
	AppInstallers                       *app_installers.AppInstallers
	AppRequest                          *app_request.AppRequest
	AppStoreCountryCodes                *app_store_country_codes.AppStoreCountryCodes
	Bookmarks                           *bookmarks.Bookmarks
	Branding                            *branding.Branding
	Buildings                           *buildings.Buildings
	CacheSettings                       *cache_settings.CacheSettings
	Categories                          *categories.Categories
	CertificateAuthority                *certificate_authority.CertificateAuthority
	ClassicLdap                         *classic_ldap.ClassicLdap
	ClientCheckin                       *client_checkin.ClientCheckin
	CloudAzure                          *cloud_azure.CloudAzure
	CloudDistributionPoint              *cloud_distribution_point.CloudDistributionPoint
	CloudIdp                            *cloud_idp.CloudIdp
	CloudInformation                    *cloud_information.CloudInformation
	CloudLdap                           *cloud_ldap.CloudLdap
	CloudLdapKeystore                   *cloud_ldap_keystore.CloudLdapKeystore
	ComputerExtensionAttributes         *computer_extension_attributes.ComputerExtensionAttributes
	ComputerGroups                      *computer_groups.ComputerGroups
	ComputerInventory                   *computer_inventory.ComputerInventory
	ComputerInventoryCollectionSettings *computer_inventory_collection_settings.ComputerInventoryCollectionSettings
	ComputerPrestages                   *computer_prestages.ComputerPrestages
	ConditionalAccess                   *conditional_access.ConditionalAccess
	Csa                                 *csa.Csa
	DeclarativeDeviceManagement         *declarative_device_management.DeclarativeDeviceManagement
	Departments                         *departments.Departments
	DeviceCommunicationSettings         *device_communication_settings.DeviceCommunicationSettings
	DeviceEnrollments                   *device_enrollments.DeviceEnrollments
	Devices                             *devices.Devices
	Digicert                            *digicert.Digicert
	DistributionPoint                   *distribution_point.DistributionPoint
	DockItems                           *dock_items.DockItems
	DssDeclarations                     *dss_declarations.DssDeclarations
	Ebooks                              *ebooks.Ebooks
	Engage                              *engage.Engage
	Enrollment                          *enrollment.Enrollment
	EnrollmentCustomizationPreview      *enrollment_customization_preview.EnrollmentCustomizationPreview
	EnrollmentCustomizations            *enrollment_customizations.EnrollmentCustomizations
	EnrollmentSettings                  *enrollment_settings.EnrollmentSettings
	Groups                              *groups.Groups
	GsxConnection                       *gsx_connection.GsxConnection
	HealthCheck                         *health_check.HealthCheck
	Icon                                *icon.Icon
	ImpactAlertNotificationSettings     *impact_alert_notification_settings.ImpactAlertNotificationSettings
	InventoryInformation                *inventory_information.InventoryInformation
	InventoryPreload                    *inventory_preload.InventoryPreload
	JamfAccountPreferences              *jamf_account_preferences.JamfAccountPreferences
	JamfConnect                         *jamf_connect.JamfConnect
	JamfManagementFramework             *jamf_management_framework.JamfManagementFramework
	JamfPackage                         *jamf_package.JamfPackage
	JamfProInformation                  *jamf_pro_information.JamfProInformation
	JamfProNotifications                *jamf_pro_notifications.JamfProNotifications
	JamfProServerUrl                    *jamf_pro_server_url.JamfProServerUrl
	JamfProSystemInitialization         *jamf_pro_system_initialization.JamfProSystemInitialization
	JamfProUserAccountSettings          *jamf_pro_user_account_settings.JamfProUserAccountSettings
	JamfProVersion                      *jamf_pro_version.JamfProVersion
	JamfProtect                         *jamf_protect.JamfProtect
	JamfRemoteAssist                    *jamf_remote_assist.JamfRemoteAssist
	Jcds                                *jcds.Jcds
	Ldap                                *ldap.Ldap
	LocalAdminPassword                  *local_admin_password.LocalAdminPassword
	Locales                             *locales.Locales
	LogFlushing                         *log_flushing.LogFlushing
	LoginCustomization                  *login_customization.LoginCustomization
	MacosConfigProfileCustomSettings    *macos_configuration_profile_custom_settings.MacosConfigurationProfileCustomSettings
	ManagedSoftwareUpdates              *managed_software_updates.ManagedSoftwareUpdates
	Mdm                                 *mdm.Mdm
	MdmRenewal                          *mdm_renewal.MdmRenewal
	MobileDeviceApps                    *mobile_device_apps.MobileDeviceApps
	MobileDeviceEnrollmentProfile       *mobile_device_enrollment_profile.MobileDeviceEnrollmentProfile
	MobileDeviceExtensionAttributes     *mobile_device_extension_attributes.MobileDeviceExtensionAttributes
	MobileDeviceGroups                  *mobile_device_groups.MobileDeviceGroups
	MobileDevicePrestages               *mobile_device_prestages.MobileDevicePrestages
	Notifications                       *notifications.Notifications
	Oauth2SessionTokens                 *oauth2_session_tokens.Oauth2SessionTokens
	Oidc                                *oidc.Oidc
	Onboarding                          *onboarding.Onboarding
	Packages                            *packages.Packages
	PatchManagement                     *patch_management.PatchManagement
	PatchPolicies                       *patch_policies.PatchPolicies
	PatchSoftwareTitleConfigurations    *patch_software_title_configurations.PatchSoftwareTitleConfigurations
	PolicyProperties                    *policy_properties.PolicyProperties
	Reenrollment                        *reenrollment.Reenrollment
	ReturnToService                     *return_to_service.ReturnToService
	Scripts                             *scripts.Scripts
	SelfServiceBrandingIos              *self_service_branding_ios.SelfServiceBrandingIos
	SelfServiceBrandingMacos            *self_service_branding_macos.SelfServiceBrandingMacos
	SelfServiceBrandingUpload           *self_service_branding_upload.SelfServiceBrandingUpload
	SelfServicePlusSettings             *self_service_plus_settings.SelfServicePlusSettings
	SelfServiceSettings                 *self_service_settings.SelfServiceSettings
	ServiceDiscoveryEnrollment          *service_discovery_enrollment.ServiceDiscoveryEnrollment
	Sites                               *sites.Sites
	Slasa                               *slasa.Slasa
	SmartComputerGroups                 *smart_computer_groups.SmartComputerGroups
	SmartMobileDeviceGroups             *smart_mobile_device_groups.SmartMobileDeviceGroups
	SmtpServer                          *smtp_server.SmtpServer
	SsoCertificate                      *sso_certificate.SsoCertificate
	SsoFailover                         *sso_failover.SsoFailover
	SsoSettings                         *sso_settings.SsoSettings
	StartupStatus                       *startup_status.StartupStatus
	StaticComputerGroups                *static_computer_groups.StaticComputerGroups
	StaticMobileDeviceGroups            *static_mobile_device_groups.StaticMobileDeviceGroups
	TimeZones                           *time_zones.TimeZones
	TomcatSettings                      *tomcat_settings.TomcatSettings
	User                                *user.User
	Venafi                              *venafi.Venafi
	VolumePurchasingLocations           *volume_purchasing_locations.VolumePurchasingLocations
	VolumePurchasingSubscriptions       *volume_purchasing_subscriptions.VolumePurchasingSubscriptions
}

// Client is the main entry point for the Jamf Pro API SDK.
type Client struct {
	transport  *client.Transport
	ClassicAPI *ClassicAPIClient
	JamfProAPI *JamfProAPIClient
}

// NewClient creates a new Jamf Pro API client.
func NewClient(authConfig *AuthConfig, options ...ClientOption) (*Client, error) {
	transport, err := client.NewTransport(authConfig, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create transport: %w", err)
	}
	return &Client{
		transport:  transport,
		ClassicAPI: newClassicAPIClient(transport),
		JamfProAPI: newJamfProAPIClient(transport),
	}, nil
}

func newClassicAPIClient(transport *client.Transport) *ClassicAPIClient {
	return &ClassicAPIClient{
		Accounts:                          classic_accounts.NewAccounts(transport),
		AccountGroups:                     classic_accounts_groups.NewAccountsGroups(transport),
		ActivationCode:                    classic_activation_code.NewActivationCode(transport),
		AdvancedComputerSearches:          classic_advanced_computer_searches.NewAdvancedComputerSearches(transport),
		AdvancedUserSearches:              classic_advanced_user_searches.NewAdvancedUserSearches(transport),
		AllowedFileExtensions:             classic_allowed_file_extensions.NewAllowedFileExtensions(transport),
		BYOProfiles:                       classic_byoprofiles.NewByoprofiles(transport),
		Classes:                           classic_classes.NewClasses(transport),
		CommandFlush:                      classic_command_flush.NewCommandFlush(transport),
		ComputerGroups:                    classic_computer_groups.NewComputerGroups(transport),
		Computers:                         classic_computers.NewComputers(transport),
		ComputerHistory:                   classic_computer_history.NewComputerHistory(transport),
		ComputerInvitations:               classic_computer_invitations.NewComputerInvitations(transport),
		ComputerInventoryCollection:       classic_computer_inventory_collection.NewComputerInventoryCollection(transport),
		DirectoryBindings:                 classic_directory_bindings.NewDirectoryBindings(transport),
		DiskEncryptionConfigurations:      classic_disk_encryption_configurations.NewDiskEncryptionConfigurations(transport),
		DockItems:                         classic_dock_items.NewDockItems(transport),
		Ebooks:                            classic_ebooks.NewEbooks(transport),
		FileShareDistributionPoints:       classic_file_share_distribution_points.NewFileShareDistributionPoints(transport),
		FileUploads:                       classic_file_uploads.NewFileUploads(transport),
		IBeacons:                          classic_ibeacons.NewIbeacons(transport),
		LdapServers:                       classic_ldap_servers.NewLdapServers(transport),
		LicensedSoftware:                  classic_licensed_software.NewLicensedSoftware(transport),
		MacApplications:                   classic_mac_applications.NewMacApplications(transport),
		MacOSConfigurationProfiles:        classic_macos_configuration_profiles.NewMacosConfigurationProfiles(transport),
		MobileDeviceApplications:          classic_mobile_device_applications.NewMobileDeviceApplications(transport),
		MobileDeviceConfigurationProfiles: classic_mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(transport),
		MobileDeviceEnrollmentProfiles:    classic_mobile_device_enrollment_profiles.NewMobileDeviceEnrollmentProfiles(transport),
		MobileDeviceGroups:                classic_mobile_device_groups.NewMobileDeviceGroups(transport),
		MobileDeviceProvisioningProfiles:  classic_mobile_device_provisioning_profiles.NewMobileDeviceProvisioningProfiles(transport),
		MobileDevices:                     classic_mobile_devices.NewMobileDevices(transport),
		NetworkSegments:                   classic_network_segments.NewNetworkSegments(transport),
		PatchExternalSources:              classic_patch_external_sources.NewPatchExternalSources(transport),
		Policies:                          classic_policies.NewPolicies(transport),
		Printers:                          classic_printers.NewPrinters(transport),
		RemoveableMacAddresses:            classic_removeable_mac_addresses.NewRemoveableMacAddresses(transport),
		RestrictedSoftware:                classic_restricted_software.NewRestrictedSoftware(transport),
		Sites:                             classic_sites.NewSites(transport),
		SmartUserGroups:                   classic_smart_user_groups.NewSmartUserGroups(transport),
		SoftwareUpdateServers:             classic_software_update_servers.NewSoftwareUpdateServers(transport),
		StaticUserGroups:                  classic_static_user_groups.NewStaticUserGroups(transport),
		UserExtensionAttributes:           classic_user_extension_attributes.NewUserExtensionAttributes(transport),
		Users:                             classic_users.NewUsers(transport),
		VppAccounts:                       classic_vpp_accounts.NewVppAccounts(transport),
		VppAssignments:                    classic_vpp_assignments.NewVppAssignments(transport),
		Webhooks:                          classic_webhooks.NewWebhooks(transport),
	}
}

func newJamfProAPIClient(transport *client.Transport) *JamfProAPIClient {
	return &JamfProAPIClient{
		AccessManagementSettings:            access_management_settings.NewAccessManagementSettings(transport),
		AccountPreferences:                  account_preferences.NewAccountPreferences(transport),
		Accounts:                            accounts.NewAccounts(transport),
		ActivationCode:                      activation_code.NewActivationCode(transport),
		AdcsSettings:                        adcs_settings.NewAdcsSettings(transport),
		AdueSessionTokenSettings:            adue_session_token_settings.NewAdueSessionTokenSettings(transport),
		AdvancedMobileDeviceSearches:        advanced_mobile_device_searches.NewAdvancedMobileDeviceSearches(transport),
		AdvancedUserContentSearches:         advanced_user_content_searches.NewAdvancedUserContentSearches(transport),
		ApiAuthorization:                    api_authorization.NewApiAuthorization(transport),
		ApiIntegrations:                     api_integrations.NewApiIntegrations(transport),
		ApiRolePrivileges:                   api_role_privileges.NewApiRolePrivileges(transport),
		ApiRoles:                            api_roles.NewApiRoles(transport),
		ApnsClientPushStatus:                apns_client_push_status.NewApnsClientPushStatus(transport),
		AppInstallers:                       app_installers.NewAppInstallers(transport),
		AppRequest:                          app_request.NewAppRequest(transport),
		AppStoreCountryCodes:                app_store_country_codes.NewAppStoreCountryCodes(transport),
		Bookmarks:                           bookmarks.NewBookmarks(transport),
		Branding:                            branding.NewBranding(transport),
		Buildings:                           buildings.NewBuildings(transport),
		CacheSettings:                       cache_settings.NewCacheSettings(transport),
		Categories:                          categories.NewCategories(transport),
		CertificateAuthority:                certificate_authority.NewCertificateAuthority(transport),
		ClassicLdap:                         classic_ldap.NewClassicLdap(transport),
		ClientCheckin:                       client_checkin.NewClientCheckin(transport),
		CloudAzure:                          cloud_azure.NewCloudAzure(transport),
		CloudDistributionPoint:              cloud_distribution_point.NewCloudDistributionPoint(transport),
		CloudIdp:                            cloud_idp.NewCloudIdp(transport),
		CloudInformation:                    cloud_information.NewCloudInformation(transport),
		CloudLdap:                           cloud_ldap.NewCloudLdap(transport),
		CloudLdapKeystore:                   cloud_ldap_keystore.NewCloudLdapKeystore(transport),
		ComputerExtensionAttributes:         computer_extension_attributes.NewComputerExtensionAttributes(transport),
		ComputerGroups:                      computer_groups.NewComputerGroups(transport),
		ComputerInventory:                   computer_inventory.NewComputerInventory(transport),
		ComputerInventoryCollectionSettings: computer_inventory_collection_settings.NewComputerInventoryCollectionSettings(transport),
		ComputerPrestages:                   computer_prestages.NewComputerPrestages(transport),
		ConditionalAccess:                   conditional_access.NewConditionalAccess(transport),
		Csa:                                 csa.NewCsa(transport),
		DeclarativeDeviceManagement:         declarative_device_management.NewDeclarativeDeviceManagement(transport),
		Departments:                         departments.NewDepartments(transport),
		DeviceCommunicationSettings:         device_communication_settings.NewDeviceCommunicationSettings(transport),
		DeviceEnrollments:                   device_enrollments.NewDeviceEnrollments(transport),
		Devices:                             devices.NewDevices(transport),
		Digicert:                            digicert.NewDigicert(transport),
		DistributionPoint:                   distribution_point.NewDistributionPoint(transport),
		DockItems:                           dock_items.NewDockItems(transport),
		DssDeclarations:                     dss_declarations.NewDssDeclarations(transport),
		Ebooks:                              ebooks.NewEbooks(transport),
		Engage:                              engage.NewEngage(transport),
		Enrollment:                          enrollment.NewEnrollment(transport),
		EnrollmentCustomizationPreview:      enrollment_customization_preview.NewEnrollmentCustomizationPreview(transport),
		EnrollmentCustomizations:            enrollment_customizations.NewEnrollmentCustomizations(transport),
		EnrollmentSettings:                  enrollment_settings.NewEnrollmentSettings(transport),
		Groups:                              groups.NewGroups(transport),
		GsxConnection:                       gsx_connection.NewGsxConnection(transport),
		HealthCheck:                         health_check.NewHealthCheck(transport),
		Icon:                                icon.NewIcon(transport),
		ImpactAlertNotificationSettings:     impact_alert_notification_settings.NewImpactAlertNotificationSettings(transport),
		InventoryInformation:                inventory_information.NewInventoryInformation(transport),
		InventoryPreload:                    inventory_preload.NewInventoryPreload(transport),
		JamfAccountPreferences:              jamf_account_preferences.NewJamfAccountPreferences(transport),
		JamfConnect:                         jamf_connect.NewJamfConnect(transport),
		JamfManagementFramework:             jamf_management_framework.NewJamfManagementFramework(transport),
		JamfPackage:                         jamf_package.NewJamfPackage(transport),
		JamfProInformation:                  jamf_pro_information.NewJamfProInformation(transport),
		JamfProNotifications:                jamf_pro_notifications.NewJamfProNotifications(transport),
		JamfProServerUrl:                    jamf_pro_server_url.NewJamfProServerUrl(transport),
		JamfProSystemInitialization:         jamf_pro_system_initialization.NewJamfProSystemInitialization(transport),
		JamfProUserAccountSettings:          jamf_pro_user_account_settings.NewJamfProUserAccountSettings(transport),
		JamfProVersion:                      jamf_pro_version.NewJamfProVersion(transport),
		JamfProtect:                         jamf_protect.NewJamfProtect(transport),
		JamfRemoteAssist:                    jamf_remote_assist.NewJamfRemoteAssist(transport),
		Jcds:                                jcds.NewJcds(transport),
		Ldap:                                ldap.NewLdap(transport),
		LocalAdminPassword:                  local_admin_password.NewLocalAdminPassword(transport),
		Locales:                             locales.NewLocales(transport),
		LogFlushing:                         log_flushing.NewLogFlushing(transport),
		LoginCustomization:                  login_customization.NewLoginCustomization(transport),
		MacosConfigProfileCustomSettings:    macos_configuration_profile_custom_settings.NewMacosConfigurationProfileCustomSettings(transport),
		ManagedSoftwareUpdates:              managed_software_updates.NewManagedSoftwareUpdates(transport),
		Mdm:                                 mdm.NewMdm(transport),
		MdmRenewal:                          mdm_renewal.NewMdmRenewal(transport),
		MobileDeviceApps:                    mobile_device_apps.NewMobileDeviceApps(transport),
		MobileDeviceEnrollmentProfile:       mobile_device_enrollment_profile.NewMobileDeviceEnrollmentProfile(transport),
		MobileDeviceExtensionAttributes:     mobile_device_extension_attributes.NewMobileDeviceExtensionAttributes(transport),
		MobileDeviceGroups:                  mobile_device_groups.NewMobileDeviceGroups(transport),
		MobileDevicePrestages:               mobile_device_prestages.NewMobileDevicePrestages(transport),
		Notifications:                       notifications.NewNotifications(transport),
		Oauth2SessionTokens:                 oauth2_session_tokens.NewOauth2SessionTokens(transport),
		Oidc:                                oidc.NewOidc(transport),
		Onboarding:                          onboarding.NewOnboarding(transport),
		Packages:                            packages.NewPackages(transport),
		PatchManagement:                     patch_management.NewPatchManagement(transport),
		PatchPolicies:                       patch_policies.NewPatchPolicies(transport),
		PatchSoftwareTitleConfigurations:    patch_software_title_configurations.NewPatchSoftwareTitleConfigurations(transport),
		PolicyProperties:                    policy_properties.NewPolicyProperties(transport),
		Reenrollment:                        reenrollment.NewReenrollment(transport),
		ReturnToService:                     return_to_service.NewReturnToService(transport),
		Scripts:                             scripts.NewScripts(transport),
		SelfServiceBrandingIos:              self_service_branding_ios.NewSelfServiceBrandingIos(transport),
		SelfServiceBrandingMacos:            self_service_branding_macos.NewSelfServiceBrandingMacos(transport),
		SelfServiceBrandingUpload:           self_service_branding_upload.NewSelfServiceBrandingUpload(transport),
		SelfServicePlusSettings:             self_service_plus_settings.NewSelfServicePlusSettings(transport),
		SelfServiceSettings:                 self_service_settings.NewSelfServiceSettings(transport),
		ServiceDiscoveryEnrollment:          service_discovery_enrollment.NewServiceDiscoveryEnrollment(transport),
		Sites:                               sites.NewSites(transport),
		Slasa:                               slasa.NewSlasa(transport),
		SmartComputerGroups:                 smart_computer_groups.NewSmartComputerGroups(transport),
		SmartMobileDeviceGroups:             smart_mobile_device_groups.NewSmartMobileDeviceGroups(transport),
		SmtpServer:                          smtp_server.NewSmtpServer(transport),
		SsoCertificate:                      sso_certificate.NewSsoCertificate(transport),
		SsoFailover:                         sso_failover.NewSsoFailover(transport),
		SsoSettings:                         sso_settings.NewSsoSettings(transport),
		StartupStatus:                       startup_status.NewStartupStatus(transport),
		StaticComputerGroups:                static_computer_groups.NewStaticComputerGroups(transport),
		StaticMobileDeviceGroups:            static_mobile_device_groups.NewStaticMobileDeviceGroups(transport),
		TimeZones:                           time_zones.NewTimeZones(transport),
		TomcatSettings:                      tomcat_settings.NewTomcatSettings(transport),
		User:                                user.NewUser(transport),
		Venafi:                              venafi.NewVenafi(transport),
		VolumePurchasingLocations:           volume_purchasing_locations.NewVolumePurchasingLocations(transport),
		VolumePurchasingSubscriptions:       volume_purchasing_subscriptions.NewVolumePurchasingSubscriptions(transport),
	}
}

// NewClientFromEnv creates a new client using environment variables.
// Required: INSTANCE_DOMAIN, AUTH_METHOD; for oauth2: CLIENT_ID, CLIENT_SECRET; for basic: BASIC_AUTH_USERNAME, BASIC_AUTH_PASSWORD.
func NewClientFromEnv(options ...ClientOption) (*Client, error) {
	authConfig := config.AuthConfigFromEnv()
	if err := authConfig.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config from env: %w", err)
	}
	return NewClient(authConfig, options...)
}

// GetLogger returns the configured zap logger.
func (c *Client) GetLogger() *zap.Logger {
	return c.transport.GetLogger()
}

// GetTransport returns the underlying HTTP transport.
func (c *Client) GetTransport() *client.Transport {
	return c.transport
}

// LoadAuthConfigFromFile loads authentication configuration from a JSON file.
func LoadAuthConfigFromFile(path string) (*AuthConfig, error) {
	return config.LoadAuthConfigFromFile(path)
}

// AuthConfigFromEnv builds AuthConfig from environment variables.
func AuthConfigFromEnv() *AuthConfig {
	return config.AuthConfigFromEnv()
}
