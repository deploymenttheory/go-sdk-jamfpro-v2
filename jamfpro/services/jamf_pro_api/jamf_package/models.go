package jamf_package

// PackageArtifact represents a single package artifact in the Jamf package API responses.
// Used in both v1 (array element) and v2 (within artifacts array).
type PackageArtifact struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
	Version  string `json:"version"`
	Created  string `json:"created"`
	URL      string `json:"url"`
}

// ListV1Response is the response for GET /api/v1/jamf-package.
// Returns an array of package artifacts for the given application.
type ListV1Response []PackageArtifact

// ResourceJamfPackageV2 is the response for GET /api/v2/jamf-package.
// Returns a package object with display name, release history URL, and artifacts.
type ResourceJamfPackageV2 struct {
	DisplayName       string            `json:"displayName"`
	ReleaseHistoryURL string            `json:"releaseHistoryUrl"`
	Artifacts         []PackageArtifact `json:"artifacts"`
}
