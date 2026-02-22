package jamfpro

import (
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/accounts"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/accounts_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/activation_code"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/advanced_computer_searches"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/advanced_user_searches"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/allowed_file_extensions"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/byoprofiles"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/classes"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/directory_bindings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/disk_encryption_configurations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/ibeacons"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/ldap_servers"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/network_segments"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/patch_external_sources"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/printers"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/removeable_mac_addresses"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/restricted_software"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/sites"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/software_update_servers"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/vpp_accounts"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/webhooks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/access_management_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/account_preferences"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/adcs_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/advanced_mobile_device_searches"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_integrations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_role_privileges"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_roles"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/app_installers"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/bookmarks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/buildings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cache_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/categories"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/certificate_authority"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/client_checkin"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_azure"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_idp"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_ldap"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_ldap_keystore"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_distribution_point"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_extension_attributes"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_inventory"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_inventory_collection_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_prestages"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/conditional_access"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/csa"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/declarative_device_management"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/departments"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/device_communication_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/device_enrollments"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/dock_items"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/dss_declarations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/engage"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/enrollment_customizations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/enrollment_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/icons"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_pro_information"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_pro_version"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/ldap"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/local_admin_password"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/locales"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/login_customization"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/adue_session_token_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/reenrollment"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/return_to_service"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/self_service_plus_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/self_service_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/service_discovery_enrollment"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/sso_certificate"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/sso_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/managed_software_updates"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mobile_device_extension_attributes"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mobile_device_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mobile_device_prestages"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/notifications"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/onboarding"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/packages"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/patch_policies"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/patch_software_title_configurations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/policy_properties"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/scripts"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/smtp_server"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/startup_status"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/time_zones"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/volume_purchasing_locations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/volume_purchasing_subscriptions"
	"go.uber.org/zap"
)

// Client is the main entry point for the Jamf Pro API SDK.
type Client struct {
	transport *client.Transport

	// Classic API services
	Accounts                     *accounts.Service
	AccountGroups                *accounts_groups.Service
	ActivationCode               *activation_code.Service
	AdvancedComputerSearches     *advanced_computer_searches.Service
	AdvancedUserSearches         *advanced_user_searches.Service
	AllowedFileExtensions        *allowed_file_extensions.Service
	BYOProfiles                  *byoprofiles.Service
	Classes                      *classes.Service
	DirectoryBindings            *directory_bindings.Service
	DiskEncryptionConfigurations *disk_encryption_configurations.Service
	IBeacons                     *ibeacons.Service
	LdapServers                  *ldap_servers.Service
	NetworkSegments              *network_segments.Service
	PatchExternalSources         *patch_external_sources.Service
	Printers                     *printers.Service
	RemoveableMacAddresses       *removeable_mac_addresses.Service
	RestrictedSoftware           *restricted_software.Service
	Sites                        *sites.Service
	SoftwareUpdateServers        *software_update_servers.Service
	VPPAccounts                  *vpp_accounts.Service
	Webhooks                     *webhooks.Service

	// Jamf Pro API services
	AccessManagementSettings        *access_management_settings.Service
	AccountPreferences              *account_preferences.Service
	AdcsSettings                    *adcs_settings.Service
	AdvancedMobileDeviceSearches    *advanced_mobile_device_searches.Service
	ApiIntegrations                 *api_integrations.Service
	APIRolePrivileges               *api_role_privileges.Service
	CertificateAuthority            *certificate_authority.Service
	APIRoles                        *api_roles.Service
	AppInstallers                   *app_installers.Service
	Bookmarks                       *bookmarks.Service
	Buildings                       *buildings.Service
	CacheSettings                   *cache_settings.Service
	Categories                      *categories.Service
	ClientCheckin                   *client_checkin.Service
	CloudAzure                      *cloud_azure.Service
	CloudIdp                        *cloud_idp.Service
	CloudLdap                       *cloud_ldap.Service
	CloudLdapKeystore               *cloud_ldap_keystore.Service
	CloudDistributionPoint          *cloud_distribution_point.Service
	ComputerExtensionAttributes            *computer_extension_attributes.Service
	ComputerInventory                      *computer_inventory.Service
	ComputerInventoryCollectionSettings    *computer_inventory_collection_settings.Service
	ComputerGroups                         *computer_groups.Service
	ComputerPrestages                      *computer_prestages.Service
	ConditionalAccess                      *conditional_access.Service
	Csa                                    *csa.Service
	DeclarativeDeviceManagement            *declarative_device_management.Service
	Departments                             *departments.Service
	DeviceCommunicationSettings             *device_communication_settings.Service
	DeviceEnrollments                       *device_enrollments.Service
	DockItems                       *dock_items.Service
	DSSDeclarations                 *dss_declarations.Service
	Engage                          *engage.Service
	EnrollmentCustomizations        *enrollment_customizations.Service
	EnrollmentSettings              *enrollment_settings.Service
	Groups                          *groups.Service
	Icons                           *icons.Service
	Ldap                            *ldap.Service
	LocalAdminPassword              *local_admin_password.Service
	LoginCustomization              *login_customization.Service
	ManagedSoftwareUpdates          *managed_software_updates.Service
	ServiceDiscoveryEnrollment      *service_discovery_enrollment.Service
	SelfServiceSettings              *self_service_settings.Service
	Reenrollment                     *reenrollment.Service
	AdueSessionTokenSettings         *adue_session_token_settings.Service
	SsoCertificate                   *sso_certificate.Service
	SsoSettings                      *sso_settings.Service
	JamfProInformation               *jamf_pro_information.Service
	JamfProVersion                  *jamf_pro_version.Service
	Locales                         *locales.Service
	MobileDeviceExtensionAttributes *mobile_device_extension_attributes.Service
	MobileDeviceGroups              *mobile_device_groups.Service
	MobileDevicePrestages           *mobile_device_prestages.Service
	Notifications                   *notifications.Service
	Onboarding                      *onboarding.Service
	Packages                        *packages.Service
	PatchPolicies                   *patch_policies.Service
	PatchSoftwareTitleConfigurations *patch_software_title_configurations.Service
	PolicyProperties                *policy_properties.Service
	ReturnToService                 *return_to_service.Service
	Scripts                         *scripts.Service
	SelfServicePlusSettings         *self_service_plus_settings.Service
	SMTPServer                      *smtp_server.Service
	StartupStatus                   *startup_status.Service
	TimeZones                       *time_zones.Service
	VolumePurchasingLocations       *volume_purchasing_locations.Service
	VolumePurchasingSubscriptions   *volume_purchasing_subscriptions.Service
}

// NewClient creates a new Jamf Pro API client.
func NewClient(authConfig *client.AuthConfig, options ...client.ClientOption) (*Client, error) {
	transport, err := client.NewTransport(authConfig, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create transport: %w", err)
	}
	c := &Client{
		transport:                       transport,
		Accounts:                        accounts.NewService(transport),
		AccountGroups:                   accounts_groups.NewService(transport),
		ActivationCode:                  activation_code.NewService(transport),
		AdvancedComputerSearches:        advanced_computer_searches.NewService(transport),
		AdvancedUserSearches:            advanced_user_searches.NewService(transport),
		AllowedFileExtensions:           allowed_file_extensions.NewService(transport),
		BYOProfiles:                     byoprofiles.NewService(transport),
		Classes:                         classes.NewService(transport),
		DirectoryBindings:               directory_bindings.NewService(transport),
		DiskEncryptionConfigurations:    disk_encryption_configurations.NewService(transport),
		IBeacons:                        ibeacons.NewService(transport),
		LdapServers:                     ldap_servers.NewService(transport),
		NetworkSegments:                 network_segments.NewService(transport),
		PatchExternalSources:            patch_external_sources.NewService(transport),
		Printers:                        printers.NewService(transport),
		RemoveableMacAddresses:          removeable_mac_addresses.NewService(transport),
		RestrictedSoftware:              restricted_software.NewService(transport),
		Sites:                           sites.NewService(transport),
		SoftwareUpdateServers:           software_update_servers.NewService(transport),
		VPPAccounts:                     vpp_accounts.NewService(transport),
		Webhooks:                        webhooks.NewService(transport),
		AccessManagementSettings:        access_management_settings.NewService(transport),
		AccountPreferences:              account_preferences.NewService(transport),
		AdcsSettings:                    adcs_settings.NewService(transport),
		AdvancedMobileDeviceSearches:    advanced_mobile_device_searches.NewService(transport),
		ApiIntegrations:                 api_integrations.NewService(transport),
		APIRolePrivileges:               api_role_privileges.NewService(transport),
		CertificateAuthority:            certificate_authority.NewService(transport),
		APIRoles:                        api_roles.NewService(transport),
		AppInstallers:                   app_installers.NewService(transport),
		Bookmarks:                       bookmarks.NewService(transport),
		Buildings:                       buildings.NewService(transport),
		CacheSettings:                   cache_settings.NewService(transport),
		Categories:                      categories.NewService(transport),
		ClientCheckin:                   client_checkin.NewService(transport),
		CloudAzure:                      cloud_azure.NewService(transport),
		CloudIdp:                        cloud_idp.NewService(transport),
		CloudLdap:                       cloud_ldap.NewService(transport),
		CloudLdapKeystore:               cloud_ldap_keystore.NewService(transport),
		CloudDistributionPoint:          cloud_distribution_point.NewService(transport),
		ComputerExtensionAttributes:            computer_extension_attributes.NewService(transport),
		ComputerInventory:                      computer_inventory.NewService(transport),
		ComputerInventoryCollectionSettings:    computer_inventory_collection_settings.NewService(transport),
		ComputerGroups:                         computer_groups.NewService(transport),
		ComputerPrestages:                      computer_prestages.NewService(transport),
		ConditionalAccess:                      conditional_access.NewService(transport),
		Csa:                                    csa.NewService(transport),
		DeclarativeDeviceManagement:            declarative_device_management.NewService(transport),
		Departments:                             departments.NewService(transport),
		DeviceCommunicationSettings:             device_communication_settings.NewService(transport),
		DeviceEnrollments:                       device_enrollments.NewService(transport),
		DockItems:                       dock_items.NewService(transport),
		DSSDeclarations:                 dss_declarations.NewService(transport),
		Engage:                          engage.NewService(transport),
		EnrollmentCustomizations:        enrollment_customizations.NewService(transport),
		EnrollmentSettings:              enrollment_settings.NewService(transport),
		Groups:                          groups.NewService(transport),
		Icons:                           icons.NewService(transport),
		Ldap:                            ldap.NewService(transport),
		LocalAdminPassword:              local_admin_password.NewService(transport),
		LoginCustomization:              login_customization.NewService(transport),
		ManagedSoftwareUpdates:          managed_software_updates.NewService(transport),
		ServiceDiscoveryEnrollment:      service_discovery_enrollment.NewService(transport),
		SelfServiceSettings:              self_service_settings.NewService(transport),
		Reenrollment:                     reenrollment.NewService(transport),
		AdueSessionTokenSettings:         adue_session_token_settings.NewService(transport),
		SsoCertificate:                  sso_certificate.NewService(transport),
		SsoSettings:                     sso_settings.NewService(transport),
		JamfProInformation:              jamf_pro_information.NewService(transport),
		JamfProVersion:                  jamf_pro_version.NewService(transport),
		Locales:                         locales.NewService(transport),
		MobileDeviceExtensionAttributes: mobile_device_extension_attributes.NewService(transport),
		MobileDeviceGroups:              mobile_device_groups.NewService(transport),
		MobileDevicePrestages:           mobile_device_prestages.NewService(transport),
		Notifications:                   notifications.NewService(transport),
		Onboarding:                      onboarding.NewService(transport),
		Packages:                        packages.NewService(transport),
		PatchPolicies:                   patch_policies.NewService(transport),
		PatchSoftwareTitleConfigurations: patch_software_title_configurations.NewService(transport),
		PolicyProperties:                policy_properties.NewService(transport),
		ReturnToService:                 return_to_service.NewService(transport),
		Scripts:                         scripts.NewService(transport),
		SelfServicePlusSettings:         self_service_plus_settings.NewService(transport),
		SMTPServer:                      smtp_server.NewService(transport),
		StartupStatus:                   startup_status.NewService(transport),
		TimeZones:                       time_zones.NewService(transport),
		VolumePurchasingLocations:       volume_purchasing_locations.NewService(transport),
		VolumePurchasingSubscriptions:   volume_purchasing_subscriptions.NewService(transport),
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

// EnableTracing enables OpenTelemetry HTTP tracing on the client's transport.
// Pass nil to use default OTel config (global tracer, "jamfpro-client" service name).
func (c *Client) EnableTracing(config *client.OTelConfig) error {
	return c.transport.EnableTracing(config)
}
