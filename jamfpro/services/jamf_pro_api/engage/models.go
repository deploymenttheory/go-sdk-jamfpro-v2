package engage

type ResourceEngageSettings struct {
	IsEnabled bool `json:"isEnabled"`
}

type HistoryResponse struct {
	TotalCount int                  `json:"totalCount"`
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
