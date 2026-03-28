package computer_extension_attributes

// ComputerExtensionAttribute.dataType constants.
const (
	DataTypeString   = "STRING"
	DataTypeInteger  = "INTEGER"
	DataTypeDateTime = "DATE_TIME"
)

// ComputerExtensionAttribute.inputType constants.
const (
	InputTypeText   = "TEXT"
	InputTypePopup  = "POPUP"
	InputTypeScript = "SCRIPT"
	InputTypeLdap   = "LDAP"
)

// ComputerExtensionAttribute.inventoryDisplayType constants.
const (
	InventoryDisplayTypeGeneral              = "GENERAL"
	InventoryDisplayTypeHardware             = "HARDWARE"
	InventoryDisplayTypeOperatingSystem      = "OPERATING_SYSTEM"
	InventoryDisplayTypeUserAndLocation      = "USER_AND_LOCATION"
	InventoryDisplayTypePurchasing           = "PURCHASING"
	InventoryDisplayTypeExtensionAttributes  = "EXTENSION_ATTRIBUTES"
)

var validDataTypes = map[string]struct{}{
	DataTypeString:   {},
	DataTypeInteger:  {},
	DataTypeDateTime: {},
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
