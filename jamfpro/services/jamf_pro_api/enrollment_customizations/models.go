package enrollment_customizations

type ListResponse struct {
	TotalCount int                               `json:"totalCount"`
	Results    []ResourceEnrollmentCustomization `json:"results"`
}

type ResourceEnrollmentCustomization struct {
	ID               string                                  `json:"id"`
	SiteID           string                                  `json:"siteId"`
	DisplayName      string                                  `json:"displayName"`
	Description      string                                  `json:"description"`
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

type HistoryResponse struct {
	TotalCount int                    `json:"totalCount"`
	Results    []ResourceHistoryEntry `json:"results"`
}

type ResourceHistoryEntry struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Date     string  `json:"date"`
	Note     string  `json:"note"`
	Details  *string `json:"details"`
}

type RequestAddHistoryNotes struct {
	Note string `json:"note"`
}

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
