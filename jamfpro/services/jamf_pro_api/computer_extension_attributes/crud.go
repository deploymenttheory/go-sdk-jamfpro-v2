package computer_extension_attributes

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"github.com/mitchellh/mapstructure"
)

type (
	// ComputerExtensionAttributesServiceInterface defines the interface for computer extension attribute operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes
	ComputerExtensionAttributesServiceInterface interface {
		// ListV1 returns all computer extension attribute objects (Get Computer Extension Attribute objects).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes
		ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetByIDV1 returns the specified computer extension attribute by ID (Get specified Computer Extension Attribute object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-id
		GetByIDV1(ctx context.Context, id string) (*ResourceComputerExtensionAttribute, *interfaces.Response, error)

		// CreateV1 creates a new computer extension attribute (Create Computer Extension Attribute record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-extension-attributes
		CreateV1(ctx context.Context, request *RequestComputerExtensionAttribute) (*CreateResponse, *interfaces.Response, error)

		// UpdateByIDV1 updates the specified computer extension attribute by ID (Update specified Computer Extension Attribute object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-computer-extension-attributes-id
		UpdateByIDV1(ctx context.Context, id string, request *RequestComputerExtensionAttribute) (*ResourceComputerExtensionAttribute, *interfaces.Response, error)

		// DeleteByIDV1 removes the specified computer extension attribute by ID (Remove specified Computer Extension Attribute record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-computer-extension-attributes-id
		DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error)

		// DeleteComputerExtensionAttributesByIDV1 deletes multiple computer extension attributes by their IDs (Delete multiple Computer Extension Attributes by their IDs).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-extension-attributes-delete-multiple
		DeleteComputerExtensionAttributesByIDV1(ctx context.Context, req *DeleteComputerExtensionAttributesByIDRequest) (*interfaces.Response, error)

		// GetHistoryByIDV1 returns the history for a computer extension attribute (Get Computer Extension Attribute History).
		//
		// Query params (optional, pass via rsqlQuery): page, page-size, sort, filter (RSQL).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-id-history
		GetHistoryByIDV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error)

		// AddHistoryNoteByIDV1 adds a note to the history for a computer extension attribute (Add Computer Extension Attribute History Note).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-extension-attributes-id-history
		AddHistoryNoteByIDV1(ctx context.Context, id string, req *AddHistoryNoteRequest) (*interfaces.Response, error)

		// ListTemplatesV1 returns all computer extension attribute templates.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-templates
		ListTemplatesV1(ctx context.Context, rsqlQuery map[string]string) (*TemplateListResponse, *interfaces.Response, error)

		// GetTemplateByIDV1 returns the specified computer extension attribute template by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-templates-id
		GetTemplateByIDV1(ctx context.Context, id string) (*ResourceComputerExtensionAttributeTemplate, *interfaces.Response, error)

		// UploadV1 uploads a computer extension attribute (multipart/form-data).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-extension-attributes-upload
		UploadV1(ctx context.Context, fileReader io.Reader, fileSize int64, filename string) (*ResourceComputerExtensionAttribute, *interfaces.Response, error)

		// GetDataDependencyByIDV1 returns smart group/advanced search dependent objects for the specified computer extension attribute.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-id-data-dependency
		GetDataDependencyByIDV1(ctx context.Context, id string) (*DataDependencyResponse, *interfaces.Response, error)

		// DownloadByIDV1 downloads the specified computer extension attribute in XML format.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-id-download
		DownloadByIDV1(ctx context.Context, id string) ([]byte, *interfaces.Response, error)
	}

	// Service handles communication with the computer extension attributes-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ComputerExtensionAttributesServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Computer Extension Attributes CRUD Operations
// -----------------------------------------------------------------------------

// ListV1 returns all computer extension attribute objects.
// URL: GET /api/v1/computer-extension-attributes
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes
func (s *Service) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointComputerExtensionAttributesV1

	mergePage := func(pageData []byte) error {
		var rawData map[string]any
		if err := json.Unmarshal(pageData, &rawData); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}

		if totalCount, ok := rawData["totalCount"].(float64); ok {
			result.TotalCount = int(totalCount)
		}

		if results, ok := rawData["results"].([]any); ok {
			for _, item := range results {
				var attr ResourceComputerExtensionAttribute
				if err := mapstructure.Decode(item, &attr); err != nil {
					return fmt.Errorf("failed to decode computer extension attribute: %w", err)
				}
				result.Results = append(result.Results, attr)
			}
		}

		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list computer extension attributes: %w", err)
	}

	return &result, resp, nil
}

// GetByIDV1 returns the specified computer extension attribute by ID.
// URL: GET /api/v1/computer-extension-attributes/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-id
func (s *Service) GetByIDV1(ctx context.Context, id string) (*ResourceComputerExtensionAttribute, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("computer extension attribute ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointComputerExtensionAttributesV1, id)

	var result ResourceComputerExtensionAttribute

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateV1 creates a new computer extension attribute.
// URL: POST /api/v1/computer-extension-attributes
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-extension-attributes
func (s *Service) CreateV1(ctx context.Context, request *RequestComputerExtensionAttribute) (*CreateResponse, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := EndpointComputerExtensionAttributesV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV1 updates the specified computer extension attribute by ID.
// URL: PUT /api/v1/computer-extension-attributes/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-computer-extension-attributes-id
func (s *Service) UpdateByIDV1(ctx context.Context, id string, request *RequestComputerExtensionAttribute) (*ResourceComputerExtensionAttribute, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointComputerExtensionAttributesV1, id)

	var result ResourceComputerExtensionAttribute

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByIDV1 removes the specified computer extension attribute by ID.
// URL: DELETE /api/v1/computer-extension-attributes/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-computer-extension-attributes-id
func (s *Service) DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("computer extension attribute ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointComputerExtensionAttributesV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteComputerExtensionAttributesByIDV1 deletes multiple computer extension attributes by their IDs.
// URL: POST /api/v1/computer-extension-attributes/delete-multiple
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-extension-attributes-delete-multiple
func (s *Service) DeleteComputerExtensionAttributesByIDV1(ctx context.Context, req *DeleteComputerExtensionAttributesByIDRequest) (*interfaces.Response, error) {
	if req == nil || len(req.IDs) == 0 {
		return nil, fmt.Errorf("ids are required")
	}

	endpoint := EndpointComputerExtensionAttributesV1 + "/delete-multiple"

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetHistoryByIDV1 returns the history for a computer extension attribute.
// URL: GET /api/v1/computer-extension-attributes/{id}/history
// Query params (optional): page, page-size, sort, filter (RSQL).
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-id-history
func (s *Service) GetHistoryByIDV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointComputerExtensionAttributesV1, id)

	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var rawData map[string]any
		if err := json.Unmarshal(pageData, &rawData); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}

		if totalCount, ok := rawData["totalCount"].(float64); ok {
			result.TotalCount = int(totalCount)
		}

		if results, ok := rawData["results"].([]any); ok {
			for _, item := range results {
				var history HistoryItem
				if err := mapstructure.Decode(item, &history); err != nil {
					return fmt.Errorf("failed to decode history item: %w", err)
				}
				result.Results = append(result.Results, history)
			}
		}

		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get computer extension attribute history: %w", err)
	}

	return &result, resp, nil
}

// AddHistoryNoteByIDV1 adds a note to the history for a computer extension attribute.
// URL: POST /api/v1/computer-extension-attributes/{id}/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-extension-attributes-id-history
func (s *Service) AddHistoryNoteByIDV1(ctx context.Context, id string, req *AddHistoryNoteRequest) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if req == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointComputerExtensionAttributesV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ListTemplatesV1 returns all computer extension attribute templates.
// URL: GET /api/v1/computer-extension-attributes/templates
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-templates
func (s *Service) ListTemplatesV1(ctx context.Context, rsqlQuery map[string]string) (*TemplateListResponse, *interfaces.Response, error) {
	var result TemplateListResponse

	endpoint := EndpointComputerExtensionAttributesV1 + "/templates"

	mergePage := func(pageData []byte) error {
		var rawData map[string]any
		if err := json.Unmarshal(pageData, &rawData); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}

		if totalCount, ok := rawData["totalCount"].(float64); ok {
			result.TotalCount = int(totalCount)
		}

		if results, ok := rawData["results"].([]any); ok {
			for _, item := range results {
				var template ResourceComputerExtensionAttributeTemplate
				if err := mapstructure.Decode(item, &template); err != nil {
					return fmt.Errorf("failed to decode template: %w", err)
				}
				result.Results = append(result.Results, template)
			}
		}

		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list computer extension attribute templates: %w", err)
	}

	return &result, resp, nil
}

// GetTemplateByIDV1 returns the specified computer extension attribute template by ID.
// URL: GET /api/v1/computer-extension-attributes/templates/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-templates-id
func (s *Service) GetTemplateByIDV1(ctx context.Context, id string) (*ResourceComputerExtensionAttributeTemplate, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("template ID is required")
	}

	endpoint := fmt.Sprintf("%s/templates/%s", EndpointComputerExtensionAttributesV1, id)

	var result ResourceComputerExtensionAttributeTemplate

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UploadV1 uploads a computer extension attribute (multipart/form-data).
// URL: POST /api/v1/computer-extension-attributes/upload
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-extension-attributes-upload
func (s *Service) UploadV1(ctx context.Context, fileReader io.Reader, fileSize int64, filename string) (*ResourceComputerExtensionAttribute, *interfaces.Response, error) {
	if fileReader == nil {
		return nil, nil, fmt.Errorf("file reader is required")
	}
	if filename == "" {
		return nil, nil, fmt.Errorf("filename is required")
	}

	endpoint := EndpointComputerExtensionAttributesV1 + "/upload"

	var result ResourceComputerExtensionAttribute

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.PostMultipart(ctx, endpoint, "file", filename, fileReader, fileSize, nil, headers, nil, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetDataDependencyByIDV1 returns smart group/advanced search dependent objects for the specified computer extension attribute.
// URL: GET /api/v1/computer-extension-attributes/{id}/data-dependency
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-id-data-dependency
func (s *Service) GetDataDependencyByIDV1(ctx context.Context, id string) (*DataDependencyResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("computer extension attribute ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/data-dependency", EndpointComputerExtensionAttributesV1, id)

	var result DataDependencyResponse

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DownloadByIDV1 downloads the specified computer extension attribute in XML format.
// URL: GET /api/v1/computer-extension-attributes/{id}/download
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-id-download
func (s *Service) DownloadByIDV1(ctx context.Context, id string) ([]byte, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("computer extension attribute ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/download", EndpointComputerExtensionAttributesV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationXML,
	}

	var result []byte

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
