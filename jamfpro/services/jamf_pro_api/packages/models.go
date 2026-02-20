package packages

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/normalization"
)

// ResourcePackage represents a package resource (get/list/update).
// Aligned with Jamf Pro API v1 packages schema.
type ResourcePackage struct {
	ID                   string `json:"id,omitempty"`
	PackageName          string `json:"packageName"`
	FileName             string `json:"fileName"`
	CategoryID           string `json:"categoryId"`
	Info                 string `json:"info,omitempty"`
	Notes                string `json:"notes,omitempty"`
	Priority             int    `json:"priority"`
	OSRequirements       string `json:"osRequirements,omitempty"`
	FillUserTemplate     *bool  `json:"fillUserTemplate,omitempty"`
	Indexed              *bool  `json:"indexed,omitempty"`
	FillExistingUsers    *bool  `json:"fillExistingUsers,omitempty"`
	SWU                  *bool  `json:"swu,omitempty"`
	RebootRequired       *bool  `json:"rebootRequired,omitempty"`
	SelfHealNotify       *bool  `json:"selfHealNotify,omitempty"`
	SelfHealingAction    string `json:"selfHealingAction,omitempty"`
	OSInstall            *bool  `json:"osInstall,omitempty"`
	SerialNumber         string `json:"serialNumber,omitempty"`
	ParentPackageID      string `json:"parentPackageId,omitempty"`
	BasePath             string `json:"basePath,omitempty"`
	SuppressUpdates      *bool  `json:"suppressUpdates,omitempty"`
	CloudTransferStatus  string `json:"cloudTransferStatus,omitempty"`
	IgnoreConflicts      *bool  `json:"ignoreConflicts,omitempty"`
	SuppressFromDock     *bool  `json:"suppressFromDock,omitempty"`
	SuppressEula         *bool  `json:"suppressEula,omitempty"`
	SuppressRegistration *bool  `json:"suppressRegistration,omitempty"`
	InstallLanguage      string `json:"installLanguage,omitempty"`
	MD5                  string `json:"md5,omitempty"`
	SHA256               string `json:"sha256,omitempty"`
	HashType             string `json:"hashType,omitempty"`
	HashValue            string `json:"hashValue,omitempty"`
	Size                 string `json:"size,omitempty"`
	OSInstallerVersion   string `json:"osInstallerVersion,omitempty"`
	Manifest             string `json:"manifest,omitempty"`
	ManifestFileName     string `json:"manifestFileName,omitempty"`
	Format               string `json:"format,omitempty"`
}

// ListResponse is the response for ListPackages.
type ListResponse struct {
	TotalCount int               `json:"totalCount"`
	Results    []ResourcePackage `json:"results"`
}

// RequestPackage is the body for creating packages (metadata only).
// Create is step one; file upload is separate via UploadV1.
// Required: PackageName, FileName, CategoryID, Priority, FillUserTemplate, RebootRequired,
// OSInstall, SuppressUpdates, SuppressFromDock, SuppressEula, SuppressRegistration.
type RequestPackage struct {
	PackageName          string `json:"packageName"`
	FileName             string `json:"fileName"`
	CategoryID           string `json:"categoryId"`
	Info                 string `json:"info,omitempty"`
	Notes                string `json:"notes,omitempty"`
	Priority             int    `json:"priority"`
	MD5                  string `json:"md5,omitempty"`
	SHA256               string `json:"sha256,omitempty"`
	OSRequirements       string `json:"osRequirements,omitempty"`
	FillUserTemplate     *bool  `json:"fillUserTemplate,omitempty"`
	FillExistingUsers    *bool  `json:"fillExistingUsers,omitempty"`
	RebootRequired       *bool  `json:"rebootRequired,omitempty"`
	OSInstall            *bool  `json:"osInstall,omitempty"`
	SuppressUpdates      *bool  `json:"suppressUpdates,omitempty"`
	SuppressFromDock     *bool  `json:"suppressFromDock,omitempty"`
	SuppressEula         *bool  `json:"suppressEula,omitempty"`
	SuppressRegistration *bool  `json:"suppressRegistration,omitempty"`
}

// CreateResponse is the response for CreatePackage.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// HistoryObject represents a package history entry.
type HistoryObject struct {
	ID       normalization.IDAsString `json:"id"`
	Username string                   `json:"username"`
	Date     string                   `json:"date"`
	Note     string                   `json:"note"`
	Details  string                   `json:"details"`
}

// HistoryResponse is the response for GetHistoryV1.
type HistoryResponse struct {
	TotalCount int             `json:"totalCount"`
	Results    []HistoryObject `json:"results"`
}

// AddHistoryNotesRequest is the body for AddHistoryNotesV1.
type AddHistoryNotesRequest struct {
	Note string `json:"note"`
}

// DeletePackagesByIDRequest is the body for DeletePackagesByIDV1 (delete multiple).
type DeletePackagesByIDRequest struct {
	IDs []string `json:"ids"`
}

// BoolPtr returns a pointer to b. Use for required *bool fields in RequestPackage.
func BoolPtr(b bool) *bool {
	return &b
}
