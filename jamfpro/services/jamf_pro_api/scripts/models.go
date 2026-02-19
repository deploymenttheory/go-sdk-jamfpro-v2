package scripts

// ResourceScript represents a script resource returned by the Jamf Pro API.
type ResourceScript struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	CategoryName   string `json:"categoryName,omitempty"`
	CategoryId     string `json:"categoryId,omitempty"`
	Info           string `json:"info,omitempty"`
	Notes          string `json:"notes,omitempty"`
	OSRequirements string `json:"osRequirements,omitempty"`
	Priority       string `json:"priority,omitempty"`
	ScriptContents string `json:"scriptContents,omitempty"`
	Parameter4     string `json:"parameter4,omitempty"`
	Parameter5     string `json:"parameter5,omitempty"`
	Parameter6     string `json:"parameter6,omitempty"`
	Parameter7     string `json:"parameter7,omitempty"`
	Parameter8     string `json:"parameter8,omitempty"`
	Parameter9     string `json:"parameter9,omitempty"`
	Parameter10    string `json:"parameter10,omitempty"`
	Parameter11    string `json:"parameter11,omitempty"`
}

// ListResponse is the paginated response for ListScripts.
type ListResponse struct {
	TotalCount int              `json:"totalCount"`
	Results    []ResourceScript `json:"results"`
}

// RequestScript is the request body for creating or updating a script.
type RequestScript struct {
	Name           string `json:"name"`
	CategoryName   string `json:"categoryName,omitempty"`
	CategoryId     string `json:"categoryId,omitempty"`
	Info           string `json:"info,omitempty"`
	Notes          string `json:"notes,omitempty"`
	OSRequirements string `json:"osRequirements,omitempty"`
	Priority       string `json:"priority,omitempty"`
	ScriptContents string `json:"scriptContents,omitempty"`
	Parameter4     string `json:"parameter4,omitempty"`
	Parameter5     string `json:"parameter5,omitempty"`
	Parameter6     string `json:"parameter6,omitempty"`
	Parameter7     string `json:"parameter7,omitempty"`
	Parameter8     string `json:"parameter8,omitempty"`
	Parameter9     string `json:"parameter9,omitempty"`
	Parameter10    string `json:"parameter10,omitempty"`
	Parameter11    string `json:"parameter11,omitempty"`
}

// CreateResponse is the response for CreateScript (POST /api/v1/scripts → 201).
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// HistoryObject represents a single script history entry.
type HistoryObject struct {
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

// ScriptHistoryResponse is the response for GetScriptHistory.
type ScriptHistoryResponse struct {
	TotalCount int             `json:"totalCount"`
	Results    []HistoryObject `json:"results"`
}

// AddScriptHistoryNotesRequest is the request body for AddScriptHistoryNotes.
type AddScriptHistoryNotesRequest struct {
	Note string `json:"note"`
}
