package inventory_preload

import "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"

// InventoryPreloadRecord represents an inventory preload record.
// DeviceType: "Computer", "Mobile Device", or "Unknown".
type InventoryPreloadRecord struct {
	ID                   string                         `json:"id,omitempty"`
	SerialNumber         string                         `json:"serialNumber"`
	DeviceType           string                         `json:"deviceType"` // Computer, Mobile Device, Unknown
	Username             *string                        `json:"username,omitempty"`
	FullName             *string                        `json:"fullName,omitempty"`
	EmailAddress         *string                        `json:"emailAddress,omitempty"`
	PhoneNumber          *string                        `json:"phoneNumber,omitempty"`
	Position             *string                        `json:"position,omitempty"`
	Department           *string                        `json:"department,omitempty"`
	Building             *string                        `json:"building,omitempty"`
	Room                 *string                        `json:"room,omitempty"`
	PONumber             *string                        `json:"poNumber,omitempty"`
	PODate               *string                        `json:"poDate,omitempty"`
	WarrantyExpiration   *string                        `json:"warrantyExpiration,omitempty"`
	AppleCareID         *string                        `json:"appleCareId,omitempty"`
	LifeExpectancy       *string                        `json:"lifeExpectancy,omitempty"`
	PurchasePrice        *string                        `json:"purchasePrice,omitempty"`
	PurchasingContact    *string                        `json:"purchasingContact,omitempty"`
	PurchasingAccount    *string                        `json:"purchasingAccount,omitempty"`
	LeaseExpiration      *string                        `json:"leaseExpiration,omitempty"`
	BarCode1             *string                        `json:"barCode1,omitempty"`
	BarCode2             *string                        `json:"barCode2,omitempty"`
	AssetTag             *string                        `json:"assetTag,omitempty"`
	Vendor               *string                        `json:"vendor,omitempty"`
	ExtensionAttributes  []InventoryPreloadExtensionAttr `json:"extensionAttributes,omitempty"`
}

// InventoryPreloadExtensionAttr represents an extension attribute on an inventory preload record.
type InventoryPreloadExtensionAttr struct {
	Name  string `json:"name"`
	Value string `json:"value,omitempty"`
}

// RecordListResponse is the paginated response for listing records.
type RecordListResponse struct {
	TotalCount int                      `json:"totalCount"`
	Results    []InventoryPreloadRecord `json:"results"`
}

// CreateRecordResponse is the response for creating a record (JSON or CSV).
type CreateRecordResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// CreateFromCSVResponse is the response for creating records from CSV (array of CreateRecordResponse).
type CreateFromCSVResponse []CreateRecordResponse

// CSVValidationSuccess is the response when CSV validation succeeds.
type CSVValidationSuccess struct {
	RecordCount int `json:"recordCount"`
}

// CSVValidationError represents a single CSV validation error.
type CSVValidationError struct {
	Line    int    `json:"line,omitempty"`
	Column  string `json:"column,omitempty"`
	Message string `json:"message,omitempty"`
}

// CSVValidationErrorCause represents the cause of a validation error.
type CSVValidationErrorCause struct {
	Line    int    `json:"line,omitempty"`
	Column  string `json:"column,omitempty"`
	Message string `json:"message,omitempty"`
}

// InvalidCSVResponse is the response when CSV validation fails.
type InvalidCSVResponse struct {
	HTTPStatus int                    `json:"httpsStatus,omitempty"`
	Errors     []CSVValidationError   `json:"errors,omitempty"`
}

// ExtensionAttributeColumn represents an EA column for inventory preload.
type ExtensionAttributeColumn struct {
	Name     string `json:"name"`
	FullName string `json:"fullName"`
}

// ExtensionAttributeColumnResult is the response for GET /ea-columns.
type ExtensionAttributeColumnResult struct {
	TotalCount int                         `json:"totalCount"`
	Results    []ExtensionAttributeColumn  `json:"results"`
}

// HistoryObject is an alias to the shared history item struct.
type HistoryObject = shared.SharedHistoryItem

// HistoryListResponse is an alias to the shared history response struct.
type HistoryListResponse = shared.SharedHistoryResponse

// AddHistoryNoteRequest is an alias to the shared history note request struct.
type AddHistoryNoteRequest = shared.SharedHistoryNoteRequest

// AddHistoryNoteResponse is the response for adding a history note.
type AddHistoryNoteResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// ExportRequest is the optional request body for export.
// Overrides query parameters when URI would exceed ~2k characters.
type ExportRequest struct {
	Page     *int     `json:"page,omitempty"`
	PageSize *int     `json:"pageSize,omitempty"`
	Sort     []string `json:"sort,omitempty"`
	Filter   *string  `json:"filter,omitempty"`
	Fields   []string `json:"fields,omitempty"`
}
