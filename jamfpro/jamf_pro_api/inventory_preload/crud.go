package inventory_preload

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// InventoryPreloadServiceInterface defines the interface for inventory preload operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-inventory-preload-records
	InventoryPreloadServiceInterface interface {
		// CreateFromCSV creates inventory preload records from a CSV file (multipart/form-data).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-inventory-preload-csv
		CreateFromCSV(ctx context.Context, fileReader io.Reader, fileSize int64, fileName string) (CreateFromCSVResponse, *resty.Response, error)

		// CreateFromCSVFile opens the file at filePath and uploads it via CreateFromCSV.
		CreateFromCSVFile(ctx context.Context, filePath string) (CreateFromCSVResponse, *resty.Response, error)

		// GetCSVTemplate downloads the inventory preload CSV template as binary bytes.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-inventory-preload-csv-template
		GetCSVTemplate(ctx context.Context) ([]byte, *resty.Response, error)

		// ValidateCSV validates a CSV file without importing (multipart/form-data).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-inventory-preload-csv-validate
		ValidateCSV(ctx context.Context, fileReader io.Reader, fileSize int64, fileName string) (*CSVValidationSuccess, *resty.Response, error)

		// ValidateCSVFile opens the file at filePath and validates it via ValidateCSV.
		ValidateCSVFile(ctx context.Context, filePath string) (*CSVValidationSuccess, *resty.Response, error)

		// GetEAColumns returns extension attribute columns for inventory preload.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-inventory-preload-ea-columns
		GetEAColumns(ctx context.Context) (*ExtensionAttributeColumnResult, *resty.Response, error)

		// Export exports inventory preload records. Optional query params and body.
		// acceptType: constants.TextCSV or constants.ApplicationJSON.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-inventory-preload-export
		Export(ctx context.Context, rsqlQuery map[string]string, req *ExportRequest, acceptType string) ([]byte, *resty.Response, error)

		// ListHistory returns paginated inventory preload history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-inventory-preload-history
		ListHistory(ctx context.Context, rsqlQuery map[string]string) (*HistoryListResponse, *resty.Response, error)

		// AddHistoryNote adds a note to inventory preload history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-inventory-preload-history
		AddHistoryNote(ctx context.Context, req *AddHistoryNoteRequest) (*AddHistoryNoteResponse, *resty.Response, error)

		// ListRecords returns paginated inventory preload records.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-inventory-preload-records
		ListRecords(ctx context.Context, rsqlQuery map[string]string) (*RecordListResponse, *resty.Response, error)

		// CreateRecord creates a single inventory preload record (JSON).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-inventory-preload-records
		CreateRecord(ctx context.Context, record *InventoryPreloadRecord) (*CreateRecordResponse, *resty.Response, error)

		// DeleteAllRecords deletes all inventory preload records.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-inventory-preload-records-delete-all
		DeleteAllRecords(ctx context.Context) (*resty.Response, error)

		// GetRecordByID returns an inventory preload record by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-inventory-preload-records-id
		GetRecordByID(ctx context.Context, id string) (*InventoryPreloadRecord, *resty.Response, error)

		// UpdateRecord updates an inventory preload record by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-inventory-preload-records-id
		UpdateRecord(ctx context.Context, id string, record *InventoryPreloadRecord) (*InventoryPreloadRecord, *resty.Response, error)

		// DeleteRecord deletes an inventory preload record by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-inventory-preload-records-id
		DeleteRecord(ctx context.Context, id string) (*resty.Response, error)
	}

	// Service handles communication with the inventory preload methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-inventory-preload-records
	InventoryPreload struct {
		client transport.HTTPClient
	}
)

var _ InventoryPreloadServiceInterface = (*InventoryPreload)(nil)

func NewInventoryPreload(client transport.HTTPClient) *InventoryPreload {
	return &InventoryPreload{client: client}
}

// -----------------------------------------------------------------------------
// CSV operations
// -----------------------------------------------------------------------------

// CreateFromCSV creates inventory preload records from a CSV file.
// URL: POST /api/v2/inventory-preload/csv
func (s *InventoryPreload) CreateFromCSV(ctx context.Context, fileReader io.Reader, fileSize int64, fileName string) (CreateFromCSVResponse, *resty.Response, error) {
	if fileName == "" {
		fileName = "inventory-preload.csv"
	}
	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.MultipartFormData,
	}
	endpoint := constants.EndpointJamfProInventoryPreloadV2 + "/csv"
	var result CreateFromCSVResponse
	resp, err := s.client.PostMultipart(ctx, endpoint, "file", fileName, fileReader, fileSize, nil, headers, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// CreateFromCSVFile opens the file at filePath and uploads it via CreateFromCSV.
func (s *InventoryPreload) CreateFromCSVFile(ctx context.Context, filePath string) (CreateFromCSVResponse, *resty.Response, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("open CSV file: %w", err)
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return nil, nil, fmt.Errorf("stat CSV file: %w", err)
	}
	name := info.Name()
	if name == "" {
		name = "inventory-preload.csv"
	}
	return s.CreateFromCSV(ctx, f, info.Size(), name)
}

// GetCSVTemplate downloads the inventory preload CSV template.
// URL: GET /api/v2/inventory-preload/csv-template
func (s *InventoryPreload) GetCSVTemplate(ctx context.Context) ([]byte, *resty.Response, error) {
	endpoint := constants.EndpointJamfProInventoryPreloadV2 + "/csv-template"
	headers := map[string]string{"Accept": constants.TextCSV}
	resp, body, err := s.client.GetBytes(ctx, endpoint, nil, headers)
	if err != nil {
		return nil, resp, err
	}
	return body, resp, nil
}

// ValidateCSV validates a CSV file without importing.
// URL: POST /api/v2/inventory-preload/csv-validate
func (s *InventoryPreload) ValidateCSV(ctx context.Context, fileReader io.Reader, fileSize int64, fileName string) (*CSVValidationSuccess, *resty.Response, error) {
	if fileName == "" {
		fileName = "inventory-preload.csv"
	}
	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.MultipartFormData,
	}
	endpoint := constants.EndpointJamfProInventoryPreloadV2 + "/csv-validate"
	var result CSVValidationSuccess
	resp, err := s.client.PostMultipart(ctx, endpoint, "file", fileName, fileReader, fileSize, nil, headers, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// ValidateCSVFile opens the file at filePath and validates it via ValidateCSV.
func (s *InventoryPreload) ValidateCSVFile(ctx context.Context, filePath string) (*CSVValidationSuccess, *resty.Response, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("open CSV file: %w", err)
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return nil, nil, fmt.Errorf("stat CSV file: %w", err)
	}
	name := info.Name()
	if name == "" {
		name = "inventory-preload.csv"
	}
	return s.ValidateCSV(ctx, f, info.Size(), name)
}

// -----------------------------------------------------------------------------
// EA columns
// -----------------------------------------------------------------------------

// GetEAColumns returns extension attribute columns.
// URL: GET /api/v2/inventory-preload/ea-columns
func (s *InventoryPreload) GetEAColumns(ctx context.Context) (*ExtensionAttributeColumnResult, *resty.Response, error) {
	endpoint := constants.EndpointJamfProInventoryPreloadV2 + "/ea-columns"
	headers := map[string]string{"Accept": constants.ApplicationJSON}
	var result ExtensionAttributeColumnResult
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// -----------------------------------------------------------------------------
// Export
// -----------------------------------------------------------------------------

// Export exports inventory preload records.
// URL: POST /api/v2/inventory-preload/export
func (s *InventoryPreload) Export(ctx context.Context, rsqlQuery map[string]string, req *ExportRequest, acceptType string) ([]byte, *resty.Response, error) {
	endpoint := constants.EndpointJamfProInventoryPreloadV2 + "/export"
	if acceptType == "" {
		acceptType = constants.ApplicationJSON
	}
	headers := map[string]string{
		"Accept":       acceptType,
		"Content-Type": constants.ApplicationJSON,
	}
	var body any
	if req != nil {
		body = req
	}
	resp, err := s.client.PostWithQuery(ctx, endpoint, rsqlQuery, body, headers, nil)
	if err != nil {
		return nil, resp, fmt.Errorf("export inventory preload: %w", err)
	}
	return resp.Bytes(), resp, nil
}

// -----------------------------------------------------------------------------
// History
// -----------------------------------------------------------------------------

// ListHistory returns paginated inventory preload history.
// URL: GET /api/v2/inventory-preload/history
func (s *InventoryPreload) ListHistory(ctx context.Context, rsqlQuery map[string]string) (*HistoryListResponse, *resty.Response, error) {
	endpoint := constants.EndpointJamfProInventoryPreloadV2 + "/history"
	headers := map[string]string{"Accept": constants.ApplicationJSON}
	var result HistoryListResponse
	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// AddHistoryNote adds a note to inventory preload history.
// URL: POST /api/v2/inventory-preload/history
func (s *InventoryPreload) AddHistoryNote(ctx context.Context, req *AddHistoryNoteRequest) (*AddHistoryNoteResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request body is required")
	}
	endpoint := constants.EndpointJamfProInventoryPreloadV2 + "/history"
	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}
	var result AddHistoryNoteResponse
	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// -----------------------------------------------------------------------------
// Records CRUD
// -----------------------------------------------------------------------------

// ListRecords returns paginated inventory preload records.
// URL: GET /api/v2/inventory-preload/records
func (s *InventoryPreload) ListRecords(ctx context.Context, rsqlQuery map[string]string) (*RecordListResponse, *resty.Response, error) {
	endpoint := constants.EndpointJamfProInventoryPreloadV2 + "/records"

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	var result RecordListResponse

	mergePage := func(pageData []byte) error {
		var pageItems []InventoryPreloadRecord
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("list inventory preload records: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// CreateRecord creates a single inventory preload record.
// URL: POST /api/v2/inventory-preload/records
func (s *InventoryPreload) CreateRecord(ctx context.Context, record *InventoryPreloadRecord) (*CreateRecordResponse, *resty.Response, error) {
	if record == nil {
		return nil, nil, fmt.Errorf("record is required")
	}

	endpoint := constants.EndpointJamfProInventoryPreloadV2 + "/records"

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	var result CreateRecordResponse

	resp, err := s.client.Post(ctx, endpoint, record, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// DeleteAllRecords deletes all inventory preload records.
// URL: POST /api/v2/inventory-preload/records/delete-all
func (s *InventoryPreload) DeleteAllRecords(ctx context.Context) (*resty.Response, error) {
	endpoint := constants.EndpointJamfProInventoryPreloadV2 + "/records/delete-all"

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetRecordByID returns an inventory preload record by ID.
// URL: GET /api/v2/inventory-preload/records/{id}
func (s *InventoryPreload) GetRecordByID(ctx context.Context, id string) (*InventoryPreloadRecord, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("%s/records/%s", constants.EndpointJamfProInventoryPreloadV2, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	var result InventoryPreloadRecord

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UpdateRecord updates an inventory preload record by ID.
// URL: PUT /api/v2/inventory-preload/records/{id}
func (s *InventoryPreload) UpdateRecord(ctx context.Context, id string, record *InventoryPreloadRecord) (*InventoryPreloadRecord, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if record == nil {
		return nil, nil, fmt.Errorf("record is required")
	}
	endpoint := fmt.Sprintf("%s/records/%s", constants.EndpointJamfProInventoryPreloadV2, id)

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}
	var result InventoryPreloadRecord

	resp, err := s.client.Put(ctx, endpoint, record, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// DeleteRecord deletes an inventory preload record by ID.
// URL: DELETE /api/v2/inventory-preload/records/{id}
func (s *InventoryPreload) DeleteRecord(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("%s/records/%s", constants.EndpointJamfProInventoryPreloadV2, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
