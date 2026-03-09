package dock_items

// ResourceDockItem represents a dock item resource.
type ResourceDockItem struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Type     string `json:"type"` // App, File, or Folder
	Contents string `json:"contents"` // Read-only
}

// RequestDockItem is the body for creating and updating dock items.
type RequestDockItem struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"` // App, File, or Folder
}

// CreateResponse is the response for CreateDockItem.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}
