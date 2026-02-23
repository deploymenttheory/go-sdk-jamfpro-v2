package gsx_connection

// ResourceGSXConnection represents the GSX connection settings.
type ResourceGSXConnection struct {
	Enabled          bool        `json:"enabled"`
	Username         string      `json:"username"`
	ServiceAccountNo string      `json:"serviceAccountNo"`
	ShipToNo         string      `json:"shipToNo"`
	GsxKeystore      GsxKeystore `json:"gsxKeystore"`
}

// GsxKeystore represents the GSX keystore details.
type GsxKeystore struct {
	Name            string `json:"name"`
	ExpirationEpoch int64  `json:"expirationEpoch"`
	ErrorMessage    string `json:"errorMessage"`
}

// HistoryResponse is the response for GetHistoryV1.
type HistoryResponse struct {
	TotalCount int             `json:"totalCount"`
	Results    []HistoryObject `json:"results"`
}

// HistoryObject represents a single GSX connection history entry.
type HistoryObject struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}
