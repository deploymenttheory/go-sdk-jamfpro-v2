package jamfpro

import (
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/allowed_file_extensions"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/directory_bindings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/disk_encryption_configurations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/ibeacons"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/network_segments"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/patch_external_sources"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/printers"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/sites"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/software_update_servers"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/vpp_accounts"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/webhooks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/advanced_mobile_device_searches"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_role_privileges"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_roles"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/app_installers"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/bookmarks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/client_checkin"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/icons"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_pro_information"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_pro_version"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/buildings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/categories"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_distribution_point"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_extension_attributes"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/departments"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/dock_items"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/locales"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mobile_device_extension_attributes"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mobile_device_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/notifications"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/packages"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/scripts"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/smtp_server"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/time_zones"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/volume_purchasing_locations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/volume_purchasing_subscriptions"
	"go.uber.org/zap"
)

// Client is the main entry point for the Jamf Pro API SDK.
type Client struct {
	transport *client.Transport

	// Classic API services
	AllowedFileExtensions          *allowed_file_extensions.Service
	DirectoryBindings              *directory_bindings.Service
	DiskEncryptionConfigurations   *disk_encryption_configurations.Service
	IBeacons                       *ibeacons.Service
	NetworkSegments                *network_segments.Service
	PatchExternalSources           *patch_external_sources.Service
	Printers                       *printers.Service
	Sites                          *sites.Service
	SoftwareUpdateServers          *software_update_servers.Service
	VPPAccounts                    *vpp_accounts.Service
	Webhooks                       *webhooks.Service

	// Jamf Pro API services
	AdvancedMobileDeviceSearches   *advanced_mobile_device_searches.Service
	APIRolePrivileges              *api_role_privileges.Service
	APIRoles                       *api_roles.Service
	AppInstallers                  *app_installers.Service
	Bookmarks                      *bookmarks.Service
	Buildings                      *buildings.Service
	Categories                     *categories.Service
	ClientCheckin                  *client_checkin.Service
	CloudDistributionPoint         *cloud_distribution_point.Service
	ComputerExtensionAttributes    *computer_extension_attributes.Service
	ComputerGroups                 *computer_groups.Service
	Departments                    *departments.Service
	DockItems                      *dock_items.Service
	Icons                          *icons.Service
	JamfProInformation             *jamf_pro_information.Service
	JamfProVersion                 *jamf_pro_version.Service
	Locales                        *locales.Service
	MobileDeviceExtensionAttributes *mobile_device_extension_attributes.Service
	MobileDeviceGroups             *mobile_device_groups.Service
	Notifications                  *notifications.Service
	Packages                       *packages.Service
	Scripts                        *scripts.Service
	SMTPServer                     *smtp_server.Service
	TimeZones                      *time_zones.Service
	VolumePurchasingLocations      *volume_purchasing_locations.Service
	VolumePurchasingSubscriptions  *volume_purchasing_subscriptions.Service
}

// NewClient creates a new Jamf Pro API client.
func NewClient(authConfig *client.AuthConfig, options ...client.ClientOption) (*Client, error) {
	transport, err := client.NewTransport(authConfig, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create transport: %w", err)
	}
	c := &Client{
		transport:                       transport,
		AllowedFileExtensions:          allowed_file_extensions.NewService(transport),
		DirectoryBindings:              directory_bindings.NewService(transport),
		DiskEncryptionConfigurations:   disk_encryption_configurations.NewService(transport),
		IBeacons:                       ibeacons.NewService(transport),
		NetworkSegments:                network_segments.NewService(transport),
		PatchExternalSources:           patch_external_sources.NewService(transport),
		Printers:                       printers.NewService(transport),
		Sites:                          sites.NewService(transport),
		SoftwareUpdateServers:          software_update_servers.NewService(transport),
		VPPAccounts:                    vpp_accounts.NewService(transport),
		Webhooks:                       webhooks.NewService(transport),
		AdvancedMobileDeviceSearches:   advanced_mobile_device_searches.NewService(transport),
		APIRolePrivileges:              api_role_privileges.NewService(transport),
		APIRoles:                       api_roles.NewService(transport),
		AppInstallers:                  app_installers.NewService(transport),
		Bookmarks:                      bookmarks.NewService(transport),
		Buildings:                      buildings.NewService(transport),
		Categories:                     categories.NewService(transport),
		ClientCheckin:                  client_checkin.NewService(transport),
		CloudDistributionPoint:         cloud_distribution_point.NewService(transport),
		ComputerExtensionAttributes:     computer_extension_attributes.NewService(transport),
		ComputerGroups:                 computer_groups.NewService(transport),
		Departments:                    departments.NewService(transport),
		DockItems:                      dock_items.NewService(transport),
		Icons:                          icons.NewService(transport),
		JamfProInformation:             jamf_pro_information.NewService(transport),
		JamfProVersion:                 jamf_pro_version.NewService(transport),
		Locales:                        locales.NewService(transport),
		MobileDeviceExtensionAttributes: mobile_device_extension_attributes.NewService(transport),
		MobileDeviceGroups:             mobile_device_groups.NewService(transport),
		Notifications:                  notifications.NewService(transport),
		Packages:                       packages.NewService(transport),
		Scripts:                        scripts.NewService(transport),
		SMTPServer:                     smtp_server.NewService(transport),
		TimeZones:                      time_zones.NewService(transport),
		VolumePurchasingLocations:      volume_purchasing_locations.NewService(transport),
		VolumePurchasingSubscriptions:  volume_purchasing_subscriptions.NewService(transport),
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
