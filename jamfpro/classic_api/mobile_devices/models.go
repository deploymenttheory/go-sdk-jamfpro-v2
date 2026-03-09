package mobile_devices

import (
	"encoding/xml"
)

// ListResponse is the response for List (GET /JSSResource/mobiledevices).
type ListResponse struct {
	XMLName xml.Name              `xml:"mobile_devices"`
	Size    int                   `xml:"size"`
	Results []MobileDeviceListItem `xml:"mobile_device"`
}

// MobileDeviceListItem represents a single mobile device item in the list.
type MobileDeviceListItem struct {
	ID              int    `xml:"id"`
	Name            string `xml:"name"`
	DeviceName      string `xml:"device_name"`
	UDID            string `xml:"udid"`
	SerialNumber    string `xml:"serial_number"`
	PhoneNumber     string `xml:"phone_number"`
	WifiMacAddress  string `xml:"wifi_mac_address"`
	Managed         bool   `xml:"managed"`
	Supervised      bool   `xml:"supervised"`
	Model           string `xml:"model"`
	ModelIdentifier string `xml:"model_identifier"`
	ModelDisplay    string `xml:"model_display"`
	Username        string `xml:"username"`
}

// ResponseMobileDevice is the main mobile device resource structure.
type ResponseMobileDevice struct {
	XMLName                 xml.Name                              `xml:"mobile_device"`
	General                 MobileDeviceSubsetGeneral               `xml:"general"`
	Location                MobileDeviceSubsetLocation              `xml:"location"`
	Purchasing              MobileDeviceSubsetPurchasing           `xml:"purchasing"`
	Applications            []MobileDeviceSubsetApplication        `xml:"applications>application"`
	SecurityObject          MobileDeviceSubsetSecurity             `xml:"security_object"`
	Network                 MobileDeviceSubsetNetwork              `xml:"network"`
	Certificates            []MobileDeviceSubsetCertificate        `xml:"certificates>certificate"`
	ConfigurationProfiles   []MobileDeviceSubsetConfigurationProfile `xml:"configuration_profiles>configuration_profile"`
	ProvisioningProfiles     []MobileDeviceSubsetProvisioningProfile  `xml:"provisioning_profiles>mobile_device_provisioning_profile"`
	MobileDeviceGroups      []MobileDeviceSubsetGroup              `xml:"mobile_device_groups>mobile_device_group"`
	ExtensionAttributes     []MobileDeviceSubsetExtensionAttribute `xml:"extension_attributes>extension_attribute"`
}

// MobileDeviceSubsetGeneral contains general mobile device information.
type MobileDeviceSubsetGeneral struct {
	ID                                 int    `xml:"id"`
	DisplayName                        string `xml:"display_name"`
	DeviceName                         string `xml:"device_name"`
	Name                               string `xml:"name"`
	AssetTag                           string `xml:"asset_tag"`
	LastInventoryUpdate                string `xml:"last_inventory_update"`
	LastInventoryUpdateEpoch           int64  `xml:"last_inventory_update_epoch"`
	LastInventoryUpdateUTC             string `xml:"last_inventory_update_utc"`
	Capacity                           int    `xml:"capacity"`
	CapacityMB                         int    `xml:"capacity_mb"`
	Available                          int    `xml:"available"`
	AvailableMB                        int    `xml:"available_mb"`
	PercentageUsed                     int    `xml:"percentage_used"`
	OSType                             string `xml:"os_type"`
	OSVersion                          string `xml:"os_version"`
	OSBuild                            string `xml:"os_build"`
	SerialNumber                       string `xml:"serial_number"`
	UDID                               string `xml:"udid"`
	InitialEntryDateEpoch              int64  `xml:"initial_entry_date_epoch"`
	InitialEntryDateUTC                string `xml:"initial_entry_date_utc"`
	PhoneNumber                        string `xml:"phone_number"`
	IPAddress                          string `xml:"ip_address"`
	WifiMacAddress                     string `xml:"wifi_mac_address"`
	BluetoothMacAddress                string `xml:"bluetooth_mac_address"`
	ModemFirmware                      string `xml:"modem_firmware"`
	Model                              string `xml:"model"`
	ModelIdentifier                    string `xml:"model_identifier"`
	ModelNumber                        string `xml:"model_number"`
	ModelDisplay                       string `xml:"model_display"`
	DeviceOwnershipLevel               string `xml:"device_ownership_level"`
	LastEnrollmentEpoch                int64  `xml:"last_enrollment_epoch"`
	LastEnrollmentUTC                  string `xml:"last_enrollment_utc"`
	Managed                            bool   `xml:"managed"`
	Supervised                         bool   `xml:"supervised"`
	ExchangeActiveSyncDeviceIdentifier string `xml:"exchange_activesync_device_identifier"`
	Shared                             string `xml:"shared"`
	Tethered                           string `xml:"tethered"`
	BatteryLevel                       int    `xml:"battery_level"`
	BLECapable                         bool   `xml:"ble_capable"`
	DeviceLocatorServiceEnabled        bool   `xml:"device_locator_service_enabled"`
	DoNotDisturbEnabled                bool   `xml:"do_not_disturb_enabled"`
	CloudBackupEnabled                 bool   `xml:"cloud_backup_enabled"`
	LastCloudBackupDateEpoch           int64  `xml:"last_cloud_backup_date_epoch"`
	LastCloudBackupDateUTC             string `xml:"last_cloud_backup_date_utc"`
	LocationServicesEnabled            bool   `xml:"location_services_enabled"`
	ItunesStoreAccountIsActive         bool   `xml:"itunes_store_account_is_active"`
	LastBackupTimeEpoch                int64  `xml:"last_backup_time_epoch"`
	LastBackupTimeUTC                  string `xml:"last_backup_time_utc"`
}

// MobileDeviceSubsetLocation contains location information.
type MobileDeviceSubsetLocation struct {
	Username     string `xml:"username"`
	RealName     string `xml:"realname"`
	EmailAddress string `xml:"email_address"`
	Position     string `xml:"position"`
	Phone        string `xml:"phone"`
	PhoneNumber  string `xml:"phone_number"`
	Department   string `xml:"department"`
	Building     string `xml:"building"`
	Room         int    `xml:"room"`
}

// MobileDeviceSubsetPurchasing contains purchasing information.
type MobileDeviceSubsetPurchasing struct {
	IsPurchased          bool   `xml:"is_purchased"`
	IsLeased             bool   `xml:"is_leased"`
	PONumber             string `xml:"po_number"`
	Vendor               string `xml:"vendor"`
	ApplecareID          string `xml:"applecare_id"`
	PurchasePrice        string `xml:"purchase_price"`
	PurchasingAccount    string `xml:"purchasing_account"`
	PODate               string `xml:"po_date"`
	PODateEpoch          int64  `xml:"po_date_epoch"`
	PODateUTC            string `xml:"po_date_utc"`
	WarrantyExpires      string `xml:"warranty_expires"`
	WarrantyExpiresEpoch int64  `xml:"warranty_expires_epoch"`
	WarrantyExpiresUTC   string `xml:"warranty_expires_utc"`
	LeaseExpires         string `xml:"lease_expires"`
	LeaseExpiresEpoch    int64  `xml:"lease_expires_epoch"`
	LeaseExpiresUTC      string `xml:"lease_expires_utc"`
	LifeExpectancy       int    `xml:"life_expectancy"`
	PurchasingContact    string `xml:"purchasing_contact"`
}

// MobileDeviceSubsetApplication contains application information.
type MobileDeviceSubsetApplication struct {
	ApplicationName    string `xml:"application_name"`
	ApplicationVersion string `xml:"application_version"`
	Identifier         string `xml:"identifier"`
}

// MobileDeviceSubsetSecurity contains security information.
type MobileDeviceSubsetSecurity struct {
	DataProtection                  bool    `xml:"data_protection"`
	BlockLevelEncryptionCapable     bool    `xml:"block_level_encryption_capable"`
	FileLevelEncryptionCapable      bool    `xml:"file_level_encryption_capable"`
	PasscodePresent                 bool    `xml:"passcode_present"`
	PasscodeCompliant               bool    `xml:"passcode_compliant"`
	PasscodeCompliantWithProfile    bool    `xml:"passcode_compliant_with_profile"`
	PasscodeLockGracePeriodEnforced string  `xml:"passcode_lock_grace_period_enforced"`
	HardwareEncryption              string  `xml:"hardware_encryption"`
	ActivationLockEnabled           bool    `xml:"activation_lock_enabled"`
	JailbreakDetected               string  `xml:"jailbreak_detected"`
	LostModeEnabled                 bool    `xml:"lost_mode_enabled"`
	LostModeEnforced                bool    `xml:"lost_mode_enforced"`
	LostModeEnableIssuedEpoch       int64   `xml:"lost_mode_enable_issued_epoch"`
	LostModeEnableIssuedUTC         string  `xml:"lost_mode_enable_issued_utc"`
	LostModeMessage                 string  `xml:"lost_mode_message"`
	LostModePhone                   string  `xml:"lost_mode_phone"`
	LostModeFootnote                string  `xml:"lost_mode_footnote"`
	LostLocationEpoch               int64   `xml:"lost_location_epoch"`
	LostLocationUTC                 string  `xml:"lost_location_utc"`
	LostLocationLatitude            float64 `xml:"lost_location_latitude"`
	LostLocationLongitude           float64 `xml:"lost_location_longitude"`
	LostLocationAltitude            float64 `xml:"lost_location_altitude"`
	LostLocationSpeed               float64 `xml:"lost_location_speed"`
	LostLocationCourse              float64 `xml:"lost_location_course"`
	LostLocationHorizontalAccuracy  float64 `xml:"lost_location_horizontal_accuracy"`
	LostLocationVerticalAccuracy    float64 `xml:"lost_location_vertical_accuracy"`
}

// MobileDeviceSubsetNetwork contains network information.
type MobileDeviceSubsetNetwork struct {
	HomeCarrierNetwork       string `xml:"home_carrier_network"`
	CellularTechnology       string `xml:"cellular_technology"`
	VoiceRoamingEnabled      string `xml:"voice_roaming_enabled"`
	IMEI                     string `xml:"imei"`
	ICCID                    string `xml:"iccid"`
	CurrentCarrierNetwork    string `xml:"current_carrier_network"`
	CarrierSettingsVersion   int    `xml:"carrier_settings_version"`
	CurrentMobileCountryCode int    `xml:"current_mobile_country_code"`
	CurrentMobileNetworkCode int    `xml:"current_mobile_network_code"`
	HomeMobileCountryCode    int    `xml:"home_mobile_country_code"`
	HomeMobileNetworkCode    int    `xml:"home_mobile_network_code"`
	DataRoamingEnabled       bool   `xml:"data_roaming_enabled"`
	PhoneNumber              string `xml:"phone_number"`
}

// MobileDeviceSubsetCertificate contains certificate information.
type MobileDeviceSubsetCertificate struct {
	CommonName string `xml:"common_name"`
	Identity   bool   `xml:"identity"`
}

// MobileDeviceSubsetConfigurationProfile contains configuration profile information.
type MobileDeviceSubsetConfigurationProfile struct {
	DisplayName string `xml:"display_name"`
	Version     int    `xml:"version"`
	Identifier  string `xml:"identifier"`
	UUID        string `xml:"uuid"`
}

// MobileDeviceSubsetProvisioningProfile contains provisioning profile information.
type MobileDeviceSubsetProvisioningProfile struct {
	DisplayName         string `xml:"display_name"`
	ExpirationDate      string `xml:"expiration_date"`
	ExpirationDateEpoch int64  `xml:"expiration_date_epoch"`
	ExpirationDateUTC   string `xml:"expiration_date_utc"`
	UUID                string `xml:"uuid"`
}

// MobileDeviceSubsetGroup contains mobile device group reference.
type MobileDeviceSubsetGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// MobileDeviceSubsetExtensionAttribute contains extension attribute information.
type MobileDeviceSubsetExtensionAttribute struct {
	ID    int    `xml:"id"`
	Name  string `xml:"name"`
	Type  string `xml:"type"`
	Value string `xml:"value"`
}
