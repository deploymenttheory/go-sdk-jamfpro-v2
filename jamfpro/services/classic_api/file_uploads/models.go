package file_uploads

// ResourceIDType represents the type of identifier being used (id or name).
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/fileuploads
type ResourceIDType string

const (
	// ResourceIDTypeID indicates the identifier is a numeric ID.
	ResourceIDTypeID ResourceIDType = "id"
	// ResourceIDTypeName indicates the identifier is a name string.
	ResourceIDTypeName ResourceIDType = "name"
)
