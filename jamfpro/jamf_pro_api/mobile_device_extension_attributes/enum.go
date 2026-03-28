package mobile_device_extension_attributes

// MobileDeviceExtensionAttributeType represents the data type of a mobile device extension attribute.
const (
	ExtensionAttributeTypeString  = "STRING"
	ExtensionAttributeTypeInteger = "INTEGER"
	ExtensionAttributeTypeDate    = "DATE"
)

// MobileDeviceExtensionAttribute.inputType constants.
const (
	InputTypeText   = "TEXT"
	InputTypePopup  = "POPUP"
	InputTypeScript = "SCRIPT"
	InputTypeLdap   = "LDAP"
)

// MobileDeviceExtensionAttribute.inventoryDisplayType constants.
const (
	InventoryDisplayTypeGeneral             = "GENERAL"
	InventoryDisplayTypeHardware            = "HARDWARE"
	InventoryDisplayTypeOperatingSystem     = "OPERATING_SYSTEM"
	InventoryDisplayTypeUserAndLocation     = "USER_AND_LOCATION"
	InventoryDisplayTypePurchasing          = "PURCHASING"
	InventoryDisplayTypeExtensionAttributes = "EXTENSION_ATTRIBUTES"
)

var validDataTypes = map[string]struct{}{
	ExtensionAttributeTypeString:  {},
	ExtensionAttributeTypeInteger: {},
	ExtensionAttributeTypeDate:    {},
}

var validInputTypes = map[string]struct{}{
	InputTypeText:   {},
	InputTypePopup:  {},
	InputTypeScript: {},
	InputTypeLdap:   {},
}

var validInventoryDisplayTypes = map[string]struct{}{
	InventoryDisplayTypeGeneral:             {},
	InventoryDisplayTypeHardware:            {},
	InventoryDisplayTypeOperatingSystem:     {},
	InventoryDisplayTypeUserAndLocation:     {},
	InventoryDisplayTypePurchasing:          {},
	InventoryDisplayTypeExtensionAttributes: {},
}
