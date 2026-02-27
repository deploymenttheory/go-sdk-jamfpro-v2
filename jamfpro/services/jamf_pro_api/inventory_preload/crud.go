package inventory_preload

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// InventoryPreloadServiceInterface defines the interface for inventory preload operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-inventory-preload-records
	InventoryPreloadServiceInterface interface {
		// CreateFromCSV creates inventory preload records from a CSV file (multipart/form-data).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-inventory-preload-csv
		CreateFromCSV(ctx context.Context, fileReader io.Reader, fileSize int64, fileName string) (CreateFromCSVResponse, *interfaces.Response, error)

		// CreateFromCSVFile opens the file at filePath and uploads it via CreateFromCSV.
		CreateFromCSVFile(ctx context.Context, filePath string) (CreateFromCSVResponse, *interfaces.Response, error)

		// GetCSVTemplate downloads the inventory preload CSV template as binary bytes.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-inventory-preload-csv-template
		GetCSVTemplate(ctx context.Context) ([]byte, *interfaces.Response, error)

		// ValidateCSV validates a CSV file without importing (multipart/form-data).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-inventory-preload-csv-validate
		ValidateCSV(ctx context.Context, fileReader io.Reader, fileSize int64, fileName string) (*CSVValidationSuccess, *interfaces.Response, error)

		// ValidateCSVFile opens the file at filePath and validates it via ValidateCSV.
		ValidateCSVFile(ctx context.Context, filePath string) (*CSVValidationSuccess, *interfaces.Response, error)

		// GetEAColumns returns extension attribute columns for inventory preload.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-inventory-preload-ea-columns
		GetEAColumns(ctx context.Context) (*ExtensionAttributeColumnResult, *interfaces.Response, error)

		// Export exports inventory preload records. Optional query params and body.
		// acceptType: mime.TextCSV or mime.ApplicationJSON.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-inventory-preload-export
		Export(ctx context.Context, rsqlQuery map[string]string, req *ExportRequest, acceptType string) ([]byte, *interfaces.Response, error)

		// ListHistory returns paginated inventory preload history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-inventory-preload-history
		ListHistory(ctx context.Context, rsqlQuery map[string]string) (*HistoryListResponse, *interfaces.Response, error)

		// AddHistoryNote adds a note to inventory preload history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-inventory-preload-history
		AddHistoryNote(ctx context.Context, req *AddHistoryNoteRequest) (*AddHistoryNoteResponse, *interfaces.Response, error)

		// ListRecords returns paginated inventory preload records.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-inventory-preload-records
		ListRecords(ctx context.Context, rsqlQuery map[string]string) (*RecordListResponse, *interfaces.Response, error)

		// CreateRecord creates a single inventory preload record (JSON).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-inventory-preload-records
		CreateRecord(ctx context.Context, record *InventoryPreloadRecord) (*CreateRecordResponse, *interfaces.Response, error)

		// DeleteAllRecords deletes all inventory preload records.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-inventory-preload-records-delete-all
		DeleteAllRecords(ctx context.Context) (*interfaces.Response, error)

		// GetRecordByID returns an inventory preload record by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-inventory-preload-records-id
		GetRecordByID(ctx context.Context, id string) (*InventoryPreloadRecord, *interfaces.Response, error)

		// UpdateRecord updates an inventory preload record by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-inventory-preload-records-id
		UpdateRecord(ctx context.Context, id string, record *InventoryPreloadRecord) (*InventoryPreloadRecord, *interfaces.Response, error)

		// DeleteRecord deletes an inventory preload record by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-inventory-preload-records-id
		DeleteRecord(ctx context.Context, id string) (*interfaces.Response, error)
	}

	// Service handles communication with the inventory preload methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-inventory-preload-records
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ InventoryPreloadServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// CSV operations
// -----------------------------------------------------------------------------

// CreateFromCSV creates inventory preload records from a CSV file.
// URL: POST /api/v2/inventory-preload/csv
func (s *Service) CreateFromCSV(ctx context.Context, fileReader io.Reader, fileSize int64, fileName string) (CreateFromCSVResponse, *interfaces.Response, error) {
	if fileName == "" {
		fileName = "inventory-preload.csv"
	}
	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": "multipart/form-data",
	}
	endpoint := EndpointInventoryPreloadV2 + "/csv"
	var result CreateFromCSVResponse
	resp, err := s.client.PostMultipart(ctx, endpoint, "file", fileName, fileReader, fileSize, nil, headers, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// CreateFromCSVFile opens the file at filePath and uploads it via CreateFromCSV.
func (s *Service) CreateFromCSVFile(ctx context.Context, filePath string) (CreateFromCSVResponse, *interfaces.Response, error) {
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
func (s *Service) GetCSVTemplate(ctx context.Context) ([]byte, *interfaces.Response, error) {
	endpoint := EndpointInventoryPreloadV2 + "/csv-template"
	headers := map[string]string{"Accept": mime.TextCSV}
	resp, body, err := s.client.GetBytes(ctx, endpoint, nil, headers)
	if err != nil {
		return nil, resp, err
	}
	return body, resp, nil
}

// ValidateCSV validates a CSV file without importing.
// URL: POST /api/v2/inventory-preload/csv-validate
func (s *Service) ValidateCSV(ctx context.Context, fileReader io.Reader, fileSize int64, fileName string) (*CSVValidationSuccess, *interfaces.Response, error) {
	if fileName == "" {
		fileName = "inventory-preload.csv"
	}
	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": "multipart/form-data",
	}
	endpoint := EndpointInventoryPreloadV2 + "/csv-validate"
	var result CSVValidationSuccess
	resp, err := s.client.PostMultipart(ctx, endpoint, "file", fileName, fileReader, fileSize, nil, headers, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// ValidateCSVFile opens the file at filePath and validates it via ValidateCSV.
func (s *Service) ValidateCSVFile(ctx context.Context, filePath string) (*CSVValidationSuccess, *interfaces.Response, error) {
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
func (s *Service) GetEAColumns(ctx context.Context) (*ExtensionAttributeColumnResult, *interfaces.Response, error) {
	endpoint := EndpointInventoryPreloadV2 + "/ea-columns"
	headers := map[string]string{"Accept": mime.ApplicationJSON}
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
func (s *Service) Export(ctx context.Context, rsqlQuery map[string]string, req *ExportRequest, acceptType string) ([]byte, *interfaces.Response, error) {
	endpoint := EndpointInventoryPreloadV2 + "/export"
	if acceptType == "" {
		acceptType = mime.ApplicationJSON
	}
	headers := map[string]string{
		"Accept":       acceptType,
		"Content-Type": mime.ApplicationJSON,
	}
	var body any
	if req != nil {
		body = req
	}
	resp, err := s.client.PostWithQuery(ctx, endpoint, rsqlQuery, body, headers, nil)
	if err != nil {
		return nil, resp, fmt.Errorf("export inventory preload: %w", err)
	}
	return resp.Body, resp, nil
}

// -----------------------------------------------------------------------------
// History
// -----------------------------------------------------------------------------

// ListHistory returns paginated inventory preload history.
// URL: GET /api/v2/inventory-preload/history
func (s *Service) ListHistory(ctx context.Context, rsqlQuery map[string]string) (*HistoryListResponse, *interfaces.Response, error) {
	endpoint := EndpointInventoryPreloadV2 + "/history"
	headers := map[string]string{"Accept": mime.ApplicationJSON}
	var result HistoryListResponse
	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// AddHistoryNote adds a note to inventory preload history.
// URL: POST /api/v2/inventory-preload/history
func (s *Service) AddHistoryNote(ctx context.Context, req *AddHistoryNoteRequest) (*AddHistoryNoteResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request body is required")
	}
	endpoint := EndpointInventoryPreloadV2 + "/history"
	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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
func (s *Service) ListRecords(ctx context.Context, rsqlQuery map[string]string) (*RecordListResponse, *interfaces.Response, error) {
	endpoint := EndpointInventoryPreloadV2 + "/records"

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	var result RecordListResponse

	mergePage := func(pageData []byte) error {
		var pageResponse RecordListResponse
		if err := json.Unmarshal(pageData, &pageResponse); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResponse.Results...)
		result.TotalCount = pageResponse.TotalCount
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("list inventory preload records: %w", err)
	}
	return &result, resp, nil
}

// CreateRecord creates a single inventory preload record.
// URL: POST /api/v2/inventory-preload/records
func (s *Service) CreateRecord(ctx context.Context, record *InventoryPreloadRecord) (*CreateRecordResponse, *interfaces.Response, error) {
	if record == nil {
		return nil, nil, fmt.Errorf("record is required")
	}

	endpoint := EndpointInventoryPreloadV2 + "/records"

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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
func (s *Service) DeleteAllRecords(ctx context.Context) (*interfaces.Response, error) {
	endpoint := EndpointInventoryPreloadV2 + "/records/delete-all"

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetRecordByID returns an inventory preload record by ID.
// URL: GET /api/v2/inventory-preload/records/{id}
func (s *Service) GetRecordByID(ctx context.Context, id string) (*InventoryPreloadRecord, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("%s/records/%s", EndpointInventoryPreloadV2, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
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
func (s *Service) UpdateRecord(ctx context.Context, id string, record *InventoryPreloadRecord) (*InventoryPreloadRecord, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if record == nil {
		return nil, nil, fmt.Errorf("record is required")
	}
	endpoint := fmt.Sprintf("%s/records/%s", EndpointInventoryPreloadV2, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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
func (s *Service) DeleteRecord(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("%s/records/%s", EndpointInventoryPreloadV2, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
