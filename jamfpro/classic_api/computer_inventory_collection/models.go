package computer_inventory_collection

import "encoding/xml"

// ResourceComputerInventoryCollection represents the detailed information of inventory collection settings.
// Doc: https://developer.jamf.com/jamf-pro/reference/computerinventorycollection
type ResourceComputerInventoryCollection struct {
	XMLName xml.Name `xml:"computer_inventory_collection"`
	LocalUserAccounts             bool               `xml:"local_user_accounts"`
	HomeDirectorySizes            bool               `xml:"home_directory_sizes"`
	HiddenAccounts                bool               `xml:"hidden_accounts"`
	Printers                      bool               `xml:"printers"`
	ActiveServices                bool               `xml:"active_services"`
	MobileDeviceAppPurchasingInfo bool               `xml:"mobile_device_app_purchasing_info"`
	ComputerLocationInformation   bool               `xml:"computer_location_information"`
	PackageReceipts               bool               `xml:"package_receipts"`
	AvailableSoftwareUpdates      bool               `xml:"available_software_updates"`
	InclueApplications            bool               `xml:"inclue_applications"` // Typo preserved for backward compatibility
	InclueFonts                   bool               `xml:"inclue_fonts"`         // Typo preserved for backward compatibility
	IncluePlugins                 bool               `xml:"inclue_plugins"`      // Typo preserved for backward compatibility
	Applications                  []ApplicationEntry `xml:"applications>application,omitempty"`
	Fonts                         []FontEntry        `xml:"fonts>font,omitempty"`
	Plugins                       []PluginEntry      `xml:"plugins>plugin,omitempty"`
}

// ApplicationEntry wraps a single application entry.
// Preserves SDK v1 structure for backward compatibility.
// Embeds Application so <application><path>...</path><platform>...</platform></application> unmarshals correctly.
type ApplicationEntry struct {
	Application
}

// Application represents an application path and platform.
type Application struct {
	Path     string `xml:"path,omitempty"`
	Platform string `xml:"platform,omitempty"`
}

// FontEntry wraps a single font entry.
// Preserves SDK v1 structure for backward compatibility.
type FontEntry struct {
	Font
}

// Font represents a font path and platform.
type Font struct {
	Path     string `xml:"path,omitempty"`
	Platform string `xml:"platform,omitempty"`
}

// PluginEntry wraps a single plugin entry.
// Preserves SDK v1 structure for backward compatibility.
type PluginEntry struct {
	Plugin
}

// Plugin represents a plugin path and platform.
type Plugin struct {
	Path     string `xml:"path,omitempty"`
	Platform string `xml:"platform,omitempty"`
}
