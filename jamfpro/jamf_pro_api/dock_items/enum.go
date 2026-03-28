package dock_items

// DockItemType enum values for ResourceDockItem.Type / RequestDockItem.Type.
const (
	DockItemTypeApp    = "APP"
	DockItemTypeFile   = "FILE"
	DockItemTypeFolder = "FOLDER"
)

// validDockItemTypes is the set of accepted Type values.
var validDockItemTypes = map[string]struct{}{
	DockItemTypeApp:    {},
	DockItemTypeFile:   {},
	DockItemTypeFolder: {},
}
