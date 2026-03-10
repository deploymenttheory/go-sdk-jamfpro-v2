package enrollment_customizations

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"
)

type ListResponse struct {
	TotalCount int                               `json:"totalCount"`
	Results    []ResourceEnrollmentCustomization `json:"results"`
}

type ResourceEnrollmentCustomization struct {
	ID               string                 `json:"id"`
	SiteID           string                 `json:"siteId"`
	DisplayName      string                 `json:"displayName"`
	Description      string                 `json:"description"`
	BrandingSettings SubsetBrandingSettings `json:"enrollmentCustomizationBrandingSettings"`
}

type SubsetBrandingSettings struct {
	TextColor       string `json:"textColor"`
	ButtonColor     string `json:"buttonColor"`
	ButtonTextColor string `json:"buttonTextColor"`
	BackgroundColor string `json:"backgroundColor"`
	IconUrl         string `json:"iconUrl"`
}

type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

type ImageUploadResponse struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

// ResourceHistoryEntry is an alias to the shared history item struct.
type ResourceHistoryEntry = models.SharedHistoryItem

// HistoryResponse is an alias to the shared history response struct.
type HistoryResponse = models.SharedHistoryResponse

// RequestAddHistoryNotes is an alias to the shared history note request struct.
type RequestAddHistoryNotes = models.SharedHistoryNoteRequest

type ResponseAddHistoryNotes struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Date     string  `json:"date"`
	Note     string  `json:"note"`
	Details  *string `json:"details"`
}

type PrestagesResponse struct {
	Dependencies []ResourcePrestage `json:"dependencies"`
}

type ResourcePrestage struct {
	Name              string `json:"name"`
	HumanReadableName string `json:"humanReadableName"`
	Hyperlink         string `json:"hyperlink"`
}
