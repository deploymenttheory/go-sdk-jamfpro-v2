package gsx_connection

import "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"

// ResourceGSXConnection represents the GSX connection settings.
type ResourceGSXConnection struct {
	Enabled          bool        `json:"enabled"`
	Username         string      `json:"username"`
	ServiceAccountNo string      `json:"serviceAccountNo"`
	ShipToNo         string      `json:"shipToNo,omitempty"`
	Token            string      `json:"token,omitempty"`
	GsxKeystore      GsxKeystore `json:"gsxKeystore"`
}

// GsxKeystore represents the GSX keystore details.
type GsxKeystore struct {
	Name             string `json:"name"`
	ExpirationEpoch  int64  `json:"expirationEpoch,omitempty"`
	ErrorMessage     string `json:"errorMessage,omitempty"`
	KeystoreBytes    string `json:"keystoreBytes,omitempty"`
	KeystorePassword string `json:"keystorePassword,omitempty"`
}

// HistoryObject is an alias to the shared history item struct with string IDs.
type HistoryObject = models.SharedHistoryItemString

// HistoryResponse is an alias to the shared history response struct with string IDs.
type HistoryResponse = models.SharedHistoryResponseString

// AddHistoryNoteRequest is an alias to the shared history note request struct.
type AddHistoryNoteRequest = models.SharedHistoryNoteRequest

// AddHistoryNoteResponse represents the response after adding a history note.
type AddHistoryNoteResponse = models.SharedHistoryNoteResponse
