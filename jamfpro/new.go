package jamfpro

import (
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	classic_accounts "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/accounts"
	classic_accounts_groups "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/accounts_groups"
	classic_activation_code "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/activation_code"
	classic_advanced_computer_searches "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/advanced_computer_searches"
	classic_advanced_user_searches "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/advanced_user_searches"
	classic_allowed_file_extensions "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/allowed_file_extensions"
	classic_byoprofiles "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/byoprofiles"
	classic_classes "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/classes"
	classic_command_flush "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/command_flush"
	classic_computer_groups "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/computer_groups"
	classic_computer_history "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/computer_history"
	classic_computer_inventory_collection "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/computer_inventory_collection"
	classic_computer_invitations "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/computer_invitations"
	classic_computers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/computers"
	classic_directory_bindings "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/directory_bindings"
	classic_disk_encryption_configurations "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/disk_encryption_configurations"
	classic_dock_items "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/dock_items"
	classic_ebooks "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/ebooks"
	classic_file_share_distribution_points "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/file_share_distribution_points"
	classic_file_uploads "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/file_uploads"
	classic_ibeacons "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/ibeacons"
	classic_ldap_servers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/ldap_servers"
	classic_licensed_software "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/licensed_software"
	classic_mac_applications "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mac_applications"
	classic_macos_configuration_profiles "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/macos_configuration_profiles"
	classic_mobile_device_applications "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_device_applications"
	classic_mobile_device_configuration_profiles "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_device_configuration_profiles"
	classic_mobile_device_enrollment_profiles "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_device_enrollment_profiles"
	classic_mobile_device_groups "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_device_groups"
	classic_mobile_device_provisioning_profiles "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_device_provisioning_profiles"
	classic_mobile_devices "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_devices"
	classic_network_segments "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/network_segments"
	classic_patch_external_sources "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/patch_external_sources"
	classic_policies "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/policies"
	classic_printers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/printers"
	classic_removeable_mac_addresses "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/removeable_mac_addresses"
	classic_restricted_software "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/restricted_software"
	classic_sites "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/sites"
	classic_software_update_servers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/software_update_servers"
	classic_user_extension_attributes "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/user_extension_attributes"
	classic_usergroups "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/usergroups"
	classic_users "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/users"
	classic_vpp_accounts "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/vpp_accounts"
	classic_vpp_assignments "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/vpp_assignments"
	classic_webhooks "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/webhooks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/access_management_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/account_preferences"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/accounts"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/activation_code"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/adcs_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/adue_session_token_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/advanced_mobile_device_searches"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/advanced_user_content_searches"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_authorization"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_integrations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_role_privileges"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_roles"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/apns_client_push_status"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/app_installers"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/app_request"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/app_store_country_codes"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/bookmarks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/branding"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/buildings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cache_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/categories"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/certificate_authority"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/classic_ldap"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/client_checkin"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_azure"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_distribution_point"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_idp"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_information"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_ldap"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_ldap_keystore"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_extension_attributes"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_inventory"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_inventory_collection_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_prestages"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/conditional_access"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/csa"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/declarative_device_management"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/departments"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/digicert"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/devices"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/device_communication_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/device_enrollments"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/distribution_point"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/dock_items"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/dss_declarations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/ebooks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/engage"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/enrollment"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/enrollment_customization_preview"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/enrollment_customizations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/enrollment_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/gsx_connection"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/icon"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/impact_alert_notification_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/inventory_information"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/inventory_preload"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_connect"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_management_framework"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_package"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_pro_information"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_pro_notifications"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_pro_server_url"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_pro_system_initialization"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_pro_version"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_remote_assist"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jcds"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/ldap"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/local_admin_password"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/locales"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/log_flushing"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/login_customization"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/macos_configuration_profile_custom_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/managed_software_updates"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mdm"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mdm_renewal"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mobile_device_apps"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mobile_device_extension_attributes"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mobile_device_enrollment_profile"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mobile_device_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mobile_device_prestages"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/notifications"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/oauth2_session_tokens"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/oidc"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/onboarding"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/packages"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/patch_management"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/patch_policies"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/patch_software_title_configurations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/policy_properties"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/reenrollment"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/return_to_service"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/scripts"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/self_service_branding_ios"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/self_service_branding_macos"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/self_service_branding_upload"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/self_service_plus_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/self_service_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/service_discovery_enrollment"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/slasa"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/smart_computer_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/smart_mobile_device_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/smtp_server"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/sites"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/sso_certificate"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/sso_failover"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/sso_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/startup_status"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/static_computer_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/static_mobile_device_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/time_zones"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/tomcat_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/user"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/venafi"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/volume_purchasing_locations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/volume_purchasing_subscriptions"
	"go.uber.org/zap"
)

// Client is the main entry point for the Jamf Pro API SDK.
type Client struct {
	transport *client.Transport

	// Classic API services
	ClassicAccounts                          *classic_accounts.Service
	ClassicAccountGroups                     *classic_accounts_groups.Service
	ClassicActivationCode                    *classic_activation_code.Service
	ClassicAdvancedComputerSearches          *classic_advanced_computer_searches.Service
	ClassicAdvancedUserSearches              *classic_advanced_user_searches.Service
	ClassicAllowedFileExtensions             *classic_allowed_file_extensions.Service
	ClassicBYOProfiles                       *classic_byoprofiles.Service
	ClassicClasses                           *classic_classes.Service
	ClassicCommandFlush                      *classic_command_flush.Service
	ClassicComputerGroups                    *classic_computer_groups.Service
	ClassicComputers                         *classic_computers.Service
	ClassicComputerHistory                   *classic_computer_history.Service
	ClassicComputerInvitations               *classic_computer_invitations.Service
	ClassicComputerInventoryCollection       *classic_computer_inventory_collection.Service
	ClassicDirectoryBindings                 *classic_directory_bindings.Service
	ClassicDockItems                         *classic_dock_items.Service
	ClassicEbooks                            *classic_ebooks.Service
	ClassicDiskEncryptionConfigurations      *classic_disk_encryption_configurations.Service
	ClassicFileShareDistributionPoints       *classic_file_share_distribution_points.Service
	ClassicFileUploads                       *classic_file_uploads.Service
	ClassicIBeacons                          *classic_ibeacons.Service
	ClassicLdapServers                       *classic_ldap_servers.Service
	ClassicLicensedSoftware                  *classic_licensed_software.Service
	ClassicMacApplications                   *classic_mac_applications.Service
	ClassicMacOSConfigurationProfiles        *classic_macos_configuration_profiles.Service
	ClassicMobileDeviceApplications          *classic_mobile_device_applications.Service
	ClassicMobileDeviceConfigurationProfiles *classic_mobile_device_configuration_profiles.Service
	ClassicMobileDeviceEnrollmentProfiles    *classic_mobile_device_enrollment_profiles.Service
	ClassicMobileDeviceGroups                *classic_mobile_device_groups.Service
	ClassicMobileDeviceProvisioningProfiles  *classic_mobile_device_provisioning_profiles.Service
	ClassicMobileDevices                     *classic_mobile_devices.Service
	ClassicNetworkSegments                   *classic_network_segments.Service
	ClassicPatchExternalSources              *classic_patch_external_sources.Service
	ClassicPolicies                          *classic_policies.Service
	ClassicPrinters                          *classic_printers.Service
	ClassicRemoveableMacAddresses            *classic_removeable_mac_addresses.Service
	ClassicRestrictedSoftware                *classic_restricted_software.Service
	ClassicSites                             *classic_sites.Service
	ClassicSoftwareUpdateServers             *classic_software_update_servers.Service
	ClassicUsers                             *classic_users.Service
	ClassicUserExtensionAttributes           *classic_user_extension_attributes.Service
	ClassicUserGroups                        *classic_usergroups.Service
	ClassicVPPAccounts                       *classic_vpp_accounts.Service
	ClassicVPPAssignments                    *classic_vpp_assignments.Service
	ClassicWebhooks                          *classic_webhooks.Service

	// Jamf Pro API services
	AccessManagementSettings            *access_management_settings.Service
	AccountPreferences                  *account_preferences.Service
	Accounts                            *accounts.Service
	ActivationCode                      *activation_code.Service
	APNSClientPushStatus                *apns_client_push_status.Service
	AdcsSettings                        *adcs_settings.Service
	AdvancedMobileDeviceSearches        *advanced_mobile_device_searches.Service
	AdvancedUserContentSearches         *advanced_user_content_searches.Service
	ApiIntegrations                     *api_integrations.Service
	APIRolePrivileges                   *api_role_privileges.Service
	APIRoles                            *api_roles.Service
	ApiAuthorization                    *api_authorization.Service
	AppRequest                          *app_request.Service
	AppStoreCountryCodes                *app_store_country_codes.Service
	CertificateAuthority                *certificate_authority.Service
	ClassicLdap                         *classic_ldap.Service
	AppInstallers                       *app_installers.Service
	Bookmarks                           *bookmarks.Service
	Branding                            *branding.Service
	Buildings                           *buildings.Service
	CacheSettings                       *cache_settings.Service
	Categories                          *categories.Service
	ClientCheckin                       *client_checkin.Service
	CloudAzure                          *cloud_azure.Service
	CloudIdp                            *cloud_idp.Service
	CloudInformation                    *cloud_information.Service
	CloudLdap                           *cloud_ldap.Service
	CloudLdapKeystore                   *cloud_ldap_keystore.Service
	CloudDistributionPoint              *cloud_distribution_point.Service
	ComputerExtensionAttributes         *computer_extension_attributes.Service
	ComputerInventory                   *computer_inventory.Service
	ComputerInventoryCollectionSettings *computer_inventory_collection_settings.Service
	ComputerGroups                      *computer_groups.Service
	ComputerPrestages                   *computer_prestages.Service
	ConditionalAccess                   *conditional_access.Service
	Csa                                 *csa.Service
	DeclarativeDeviceManagement         *declarative_device_management.Service
	Departments                         *departments.Service
	Devices                             *devices.Service
	Digicert                            *digicert.Service
	DeviceCommunicationSettings         *device_communication_settings.Service
	DeviceEnrollments                   *device_enrollments.Service
	DistributionPoint                   *distribution_point.Service
	DockItems                           *dock_items.Service
	DSSDeclarations                     *dss_declarations.Service
	Ebooks                              *ebooks.Service
	Engage                              *engage.Service
	Enrollment                          *enrollment.Service
	EnrollmentCustomizationPreview      *enrollment_customization_preview.Service
	EnrollmentCustomizations            *enrollment_customizations.Service
	EnrollmentSettings                  *enrollment_settings.Service
	Groups                              *groups.Service
	GSXConnection                       *gsx_connection.Service
	Icon                                *icon.Service
		ImpactAlertNotificationSettings     *impact_alert_notification_settings.Service
		InventoryInformation                *inventory_information.Service
		InventoryPreload                     *inventory_preload.Service
		JamfConnect                         *jamf_connect.Service
	JamfManagementFramework            *jamf_management_framework.Service
	JamfPackage                         *jamf_package.Service
	JCDS                                *jcds.Service
	Ldap                                *ldap.Service
	LocalAdminPassword                  *local_admin_password.Service
	LoginCustomization                  *login_customization.Service
	LogFlushing                         *log_flushing.Service
	MacOSConfigProfileCustomSettings    *macos_configuration_profile_custom_settings.Service
	ManagedSoftwareUpdates              *managed_software_updates.Service
	MDM                                 *mdm.Service
	MDMRenewal                           *mdm_renewal.Service
	ServiceDiscoveryEnrollment          *service_discovery_enrollment.Service
	SelfServiceBrandingIOS              *self_service_branding_ios.Service
	SelfServiceBrandingMacOS            *self_service_branding_macos.Service
	SelfServiceBrandingUpload           *self_service_branding_upload.Service
	SelfServiceSettings                 *self_service_settings.Service
	SLASA                               *slasa.Service
	Reenrollment                        *reenrollment.Service
	AdueSessionTokenSettings            *adue_session_token_settings.Service
	SsoCertificate                      *sso_certificate.Service
	SsoFailover                         *sso_failover.Service
	SsoSettings                         *sso_settings.Service
	JamfProInformation                  *jamf_pro_information.Service
	JamfProNotifications                *jamf_pro_notifications.Service
	JamfProServerURL                    *jamf_pro_server_url.Service
	JamfProSystemInitialization         *jamf_pro_system_initialization.Service
	JamfProVersion                      *jamf_pro_version.Service
	JamfRemoteAssist                    *jamf_remote_assist.Service
	Locales                             *locales.Service
	MobileDeviceApps                    *mobile_device_apps.Service
	MobileDeviceEnrollmentProfile       *mobile_device_enrollment_profile.Service
	MobileDeviceExtensionAttributes     *mobile_device_extension_attributes.Service
	MobileDeviceGroups                  *mobile_device_groups.Service
	StaticMobileDeviceGroups            *static_mobile_device_groups.Service
	MobileDevicePrestages               *mobile_device_prestages.Service
	Notifications                       *notifications.Service
	OAuth2SessionTokens                 *oauth2_session_tokens.Service
	OIDC                                *oidc.Service
	Onboarding                          *onboarding.Service
	Packages                            *packages.Service
	PatchManagement                     *patch_management.Service
	PatchPolicies                       *patch_policies.Service
	PatchSoftwareTitleConfigurations    *patch_software_title_configurations.Service
	PolicyProperties                    *policy_properties.Service
	ReturnToService                     *return_to_service.Service
	Scripts                             *scripts.Service
	SmartComputerGroups                 *smart_computer_groups.Service
	SmartMobileDeviceGroups             *smart_mobile_device_groups.Service
	StaticComputerGroups                *static_computer_groups.Service
	SelfServicePlusSettings             *self_service_plus_settings.Service
	SMTPServer                          *smtp_server.Service
	Sites                               *sites.Service
	StartupStatus                       *startup_status.Service
	TimeZones                           *time_zones.Service
	TomcatSettings                      *tomcat_settings.Service
	User                                *user.Service
	Venafi                              *venafi.Service
	VolumePurchasingLocations           *volume_purchasing_locations.Service
	VolumePurchasingSubscriptions       *volume_purchasing_subscriptions.Service
}

// NewClient creates a new Jamf Pro API client.
func NewClient(authConfig *client.AuthConfig, options ...client.ClientOption) (*Client, error) {
	transport, err := client.NewTransport(authConfig, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create transport: %w", err)
	}
	c := &Client{
		transport: transport,

		// Classic API services
		ClassicAccounts:                          classic_accounts.NewService(transport),
		ClassicAccountGroups:                     classic_accounts_groups.NewService(transport),
		ClassicActivationCode:                    classic_activation_code.NewService(transport),
		ClassicAdvancedComputerSearches:          classic_advanced_computer_searches.NewService(transport),
		ClassicAdvancedUserSearches:              classic_advanced_user_searches.NewService(transport),
		ClassicAllowedFileExtensions:             classic_allowed_file_extensions.NewService(transport),
		ClassicBYOProfiles:                       classic_byoprofiles.NewService(transport),
		ClassicClasses:                           classic_classes.NewService(transport),
		ClassicCommandFlush:                      classic_command_flush.NewService(transport),
		ClassicComputerGroups:                    classic_computer_groups.NewService(transport),
		ClassicComputers:                         classic_computers.NewService(transport),
		ClassicComputerHistory:                   classic_computer_history.NewService(transport),
		ClassicComputerInvitations:               classic_computer_invitations.NewService(transport),
		ClassicComputerInventoryCollection:       classic_computer_inventory_collection.NewService(transport),
		ClassicDirectoryBindings:                 classic_directory_bindings.NewService(transport),
		ClassicDockItems:                         classic_dock_items.NewService(transport),
		ClassicEbooks:                            classic_ebooks.NewService(transport),
		ClassicDiskEncryptionConfigurations:      classic_disk_encryption_configurations.NewService(transport),
		ClassicFileShareDistributionPoints:       classic_file_share_distribution_points.NewService(transport),
		ClassicFileUploads:                       classic_file_uploads.NewService(transport),
		ClassicIBeacons:                          classic_ibeacons.NewService(transport),
		ClassicLdapServers:                       classic_ldap_servers.NewService(transport),
		ClassicLicensedSoftware:                  classic_licensed_software.NewService(transport),
		ClassicMacApplications:                   classic_mac_applications.NewService(transport),
		ClassicMacOSConfigurationProfiles:        classic_macos_configuration_profiles.NewService(transport),
		ClassicMobileDeviceApplications:          classic_mobile_device_applications.NewService(transport),
		ClassicMobileDeviceConfigurationProfiles: classic_mobile_device_configuration_profiles.NewService(transport),
		ClassicMobileDeviceEnrollmentProfiles:    classic_mobile_device_enrollment_profiles.NewService(transport),
		ClassicMobileDeviceGroups:                classic_mobile_device_groups.NewService(transport),
		ClassicMobileDeviceProvisioningProfiles:  classic_mobile_device_provisioning_profiles.NewService(transport),
		ClassicMobileDevices:                     classic_mobile_devices.NewService(transport),
		ClassicNetworkSegments:                   classic_network_segments.NewService(transport),
		ClassicPatchExternalSources:              classic_patch_external_sources.NewService(transport),
		ClassicPolicies:                          classic_policies.NewService(transport),
		ClassicPrinters:                          classic_printers.NewService(transport),
		ClassicRemoveableMacAddresses:            classic_removeable_mac_addresses.NewService(transport),
		ClassicRestrictedSoftware:                classic_restricted_software.NewService(transport),
		ClassicSites:                             classic_sites.NewService(transport),
		ClassicSoftwareUpdateServers:             classic_software_update_servers.NewService(transport),
		ClassicUsers:                             classic_users.NewService(transport),
		ClassicUserExtensionAttributes:           classic_user_extension_attributes.NewService(transport),
		ClassicUserGroups:                        classic_usergroups.NewService(transport),
		ClassicVPPAccounts:                       classic_vpp_accounts.NewService(transport),
		ClassicVPPAssignments:                    classic_vpp_assignments.NewService(transport),
		ClassicWebhooks:                          classic_webhooks.NewService(transport),

		// Jamf Pro API services
		AccessManagementSettings:            access_management_settings.NewService(transport),
		AccountPreferences:                  account_preferences.NewService(transport),
		Accounts:                            accounts.NewService(transport),
		ActivationCode:                      activation_code.NewService(transport),
		APNSClientPushStatus:                apns_client_push_status.NewService(transport),
		AdcsSettings:                        adcs_settings.NewService(transport),
		AdvancedMobileDeviceSearches:        advanced_mobile_device_searches.NewService(transport),
		AdvancedUserContentSearches:         advanced_user_content_searches.NewService(transport),
		ApiIntegrations:                     api_integrations.NewService(transport),
		APIRolePrivileges:                   api_role_privileges.NewService(transport),
		APIRoles:                            api_roles.NewService(transport),
		AppRequest:                          app_request.NewService(transport),
		AppStoreCountryCodes:                app_store_country_codes.NewService(transport),
		CertificateAuthority:                certificate_authority.NewService(transport),
		ClassicLdap:                         classic_ldap.NewService(transport),
		AppInstallers:                       app_installers.NewService(transport),
		Bookmarks:                           bookmarks.NewService(transport),
		Branding:                            branding.NewService(transport),
		Buildings:                           buildings.NewService(transport),
		CacheSettings:                       cache_settings.NewService(transport),
		Categories:                          categories.NewService(transport),
		ClientCheckin:                       client_checkin.NewService(transport),
		CloudAzure:                          cloud_azure.NewService(transport),
		CloudIdp:                            cloud_idp.NewService(transport),
		CloudInformation:                    cloud_information.NewService(transport),
		CloudLdap:                           cloud_ldap.NewService(transport),
		CloudLdapKeystore:                   cloud_ldap_keystore.NewService(transport),
		CloudDistributionPoint:              cloud_distribution_point.NewService(transport),
		ComputerExtensionAttributes:         computer_extension_attributes.NewService(transport),
		ComputerInventory:                   computer_inventory.NewService(transport),
		ComputerInventoryCollectionSettings: computer_inventory_collection_settings.NewService(transport),
		ComputerGroups:                      computer_groups.NewService(transport),
		ComputerPrestages:                   computer_prestages.NewService(transport),
		ConditionalAccess:                   conditional_access.NewService(transport),
		Csa:                                 csa.NewService(transport),
		DeclarativeDeviceManagement:         declarative_device_management.NewService(transport),
		Departments:                         departments.NewService(transport),
		Devices:                             devices.NewService(transport),
		Digicert:                            digicert.NewService(transport),
		DeviceCommunicationSettings:         device_communication_settings.NewService(transport),
		DeviceEnrollments:                   device_enrollments.NewService(transport),
		DistributionPoint:                   distribution_point.NewService(transport),
		DockItems:                           dock_items.NewService(transport),
		DSSDeclarations:                     dss_declarations.NewService(transport),
		Ebooks:                              ebooks.NewService(transport),
		Engage:                              engage.NewService(transport),
		Enrollment:                          enrollment.NewService(transport),
		EnrollmentCustomizationPreview:      enrollment_customization_preview.NewService(transport),
		EnrollmentCustomizations:            enrollment_customizations.NewService(transport),
		EnrollmentSettings:                  enrollment_settings.NewService(transport),
		Groups:                              groups.NewService(transport),
		GSXConnection:                       gsx_connection.NewService(transport),
		Icon:                                icon.NewService(transport),
		ImpactAlertNotificationSettings:     impact_alert_notification_settings.NewService(transport),
		InventoryInformation:                inventory_information.NewService(transport),
		InventoryPreload:                     inventory_preload.NewService(transport),
		JamfConnect:                         jamf_connect.NewService(transport),
		JamfManagementFramework:            jamf_management_framework.NewService(transport),
		JamfPackage:                         jamf_package.NewService(transport),
		JCDS:                                jcds.NewService(transport),
		Ldap:                                ldap.NewService(transport),
		LocalAdminPassword:                  local_admin_password.NewService(transport),
		LoginCustomization:                  login_customization.NewService(transport),
		LogFlushing:                         log_flushing.NewService(transport),
		MacOSConfigProfileCustomSettings:    macos_configuration_profile_custom_settings.NewService(transport),
		ManagedSoftwareUpdates:              managed_software_updates.NewService(transport),
		MDM:                                 mdm.NewService(transport),
		MDMRenewal:                           mdm_renewal.NewService(transport),
		ServiceDiscoveryEnrollment:          service_discovery_enrollment.NewService(transport),
		SelfServiceBrandingIOS:              self_service_branding_ios.NewService(transport),
		SelfServiceBrandingMacOS:            self_service_branding_macos.NewService(transport),
		SelfServiceBrandingUpload:           self_service_branding_upload.NewService(transport),
		SelfServiceSettings:                 self_service_settings.NewService(transport),
		SLASA:                               slasa.NewService(transport),
		Reenrollment:                        reenrollment.NewService(transport),
		AdueSessionTokenSettings:            adue_session_token_settings.NewService(transport),
		SsoCertificate:                      sso_certificate.NewService(transport),
		SsoFailover:                         sso_failover.NewService(transport),
		SsoSettings:                         sso_settings.NewService(transport),
		JamfProInformation:                  jamf_pro_information.NewService(transport),
		JamfProNotifications:                jamf_pro_notifications.NewService(transport),
		JamfProServerURL:                    jamf_pro_server_url.NewService(transport),
		JamfProSystemInitialization:         jamf_pro_system_initialization.NewService(transport),
		JamfProVersion:                      jamf_pro_version.NewService(transport),
		JamfRemoteAssist:                    jamf_remote_assist.NewService(transport),
		Locales:                             locales.NewService(transport),
		MobileDeviceApps:                    mobile_device_apps.NewService(transport),
		MobileDeviceEnrollmentProfile:       mobile_device_enrollment_profile.NewService(transport),
		MobileDeviceExtensionAttributes:     mobile_device_extension_attributes.NewService(transport),
		MobileDeviceGroups:                  mobile_device_groups.NewService(transport),
		StaticMobileDeviceGroups:            static_mobile_device_groups.NewService(transport),
		MobileDevicePrestages:               mobile_device_prestages.NewService(transport),
		Notifications:                       notifications.NewService(transport),
		OAuth2SessionTokens:                 oauth2_session_tokens.NewService(transport),
		OIDC:                                oidc.NewService(transport),
		Onboarding:                          onboarding.NewService(transport),
		Packages:                            packages.NewService(transport),
		PatchManagement:                     patch_management.NewService(transport),
		PatchPolicies:                       patch_policies.NewService(transport),
		PatchSoftwareTitleConfigurations:    patch_software_title_configurations.NewService(transport),
		PolicyProperties:                    policy_properties.NewService(transport),
		ReturnToService:                     return_to_service.NewService(transport),
		Scripts:                             scripts.NewService(transport),
		SmartComputerGroups:                 smart_computer_groups.NewService(transport),
		SmartMobileDeviceGroups:             smart_mobile_device_groups.NewService(transport),
		StaticComputerGroups:                static_computer_groups.NewService(transport),
		SelfServicePlusSettings:             self_service_plus_settings.NewService(transport),
		SMTPServer:                          smtp_server.NewService(transport),
		Sites:                               sites.NewService(transport),
		StartupStatus:                       startup_status.NewService(transport),
		TimeZones:                           time_zones.NewService(transport),
		TomcatSettings:                      tomcat_settings.NewService(transport),
		User:                                user.NewService(transport),
		Venafi:                              venafi.NewService(transport),
		VolumePurchasingLocations:           volume_purchasing_locations.NewService(transport),
		VolumePurchasingSubscriptions:       volume_purchasing_subscriptions.NewService(transport),
	}
	return c, nil
}

// NewClientFromEnv creates a new client using environment variables.
// Required: INSTANCE_DOMAIN, AUTH_METHOD; for oauth2: CLIENT_ID, CLIENT_SECRET; for basic: BASIC_AUTH_USERNAME, BASIC_AUTH_PASSWORD.
func NewClientFromEnv(options ...client.ClientOption) (*Client, error) {
	authConfig := client.AuthConfigFromEnv()
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

// EnableTracing enables OpenTelemetry HTTP tracing on the client's transport.
// Pass nil to use default OTel config (global tracer, "jamfpro-client" service name).
func (c *Client) EnableTracing(config *client.OTelConfig) error {
	return c.transport.EnableTracing(config)
}
