package computer_extension_attributes

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the computer extension attributes-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes
	ComputerExtensionAttributes struct {
		client client.Client
	}
)

func NewComputerExtensionAttributes(client client.Client) *ComputerExtensionAttributes {
	return &ComputerExtensionAttributes{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Computer Extension Attributes CRUD Operations
// -----------------------------------------------------------------------------

// ListV1 returns all computer extension attribute objects.
// URL: GET /api/v1/computer-extension-attributes
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes
func (s *ComputerExtensionAttributes) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointJamfProComputerExtensionAttributesV1

	mergePage := func(pageData []byte) error {
		var items []ResourceComputerExtensionAttribute
		if err := json.Unmarshal(pageData, &items); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, items...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list computer extension attributes: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV1 returns the specified computer extension attribute by ID.
// URL: GET /api/v1/computer-extension-attributes/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-id
func (s *ComputerExtensionAttributes) GetByIDV1(ctx context.Context, id string) (*ResourceComputerExtensionAttribute, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("computer extension attribute ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProComputerExtensionAttributesV1, id)

	var result ResourceComputerExtensionAttribute

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateV1 creates a new computer extension attribute.
// URL: POST /api/v1/computer-extension-attributes
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-extension-attributes
func (s *ComputerExtensionAttributes) CreateV1(ctx context.Context, request *RequestComputerExtensionAttribute) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := constants.EndpointJamfProComputerExtensionAttributesV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV1 updates the specified computer extension attribute by ID.
// URL: PUT /api/v1/computer-extension-attributes/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-computer-extension-attributes-id
func (s *ComputerExtensionAttributes) UpdateByIDV1(ctx context.Context, id string, request *RequestComputerExtensionAttribute) (*ResourceComputerExtensionAttribute, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProComputerExtensionAttributesV1, id)

	var result ResourceComputerExtensionAttribute

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Put(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByIDV1 removes the specified computer extension attribute by ID.
// URL: DELETE /api/v1/computer-extension-attributes/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-computer-extension-attributes-id
func (s *ComputerExtensionAttributes) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("computer extension attribute ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProComputerExtensionAttributesV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteComputerExtensionAttributesByIDV1 deletes multiple computer extension attributes by their IDs.
// URL: POST /api/v1/computer-extension-attributes/delete-multiple
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-extension-attributes-delete-multiple
func (s *ComputerExtensionAttributes) DeleteComputerExtensionAttributesByIDV1(ctx context.Context, req *DeleteComputerExtensionAttributesByIDRequest) (*resty.Response, error) {
	if req == nil || len(req.IDs) == 0 {
		return nil, fmt.Errorf("ids are required")
	}

	endpoint := constants.EndpointJamfProComputerExtensionAttributesV1 + "/delete-multiple"

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(req).
		Post(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetHistoryByIDV1 returns the history for a computer extension attribute.
// URL: GET /api/v1/computer-extension-attributes/{id}/history
// Query params (optional): page, page-size, sort, filter (RSQL).
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-id-history
func (s *ComputerExtensionAttributes) GetHistoryByIDV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProComputerExtensionAttributesV1, id)

	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var items []HistoryItem
		if err := json.Unmarshal(pageData, &items); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, items...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get computer extension attribute history: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// AddHistoryNoteByIDV1 adds a note to the history for a computer extension attribute.
// URL: POST /api/v1/computer-extension-attributes/{id}/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-extension-attributes-id-history
func (s *ComputerExtensionAttributes) AddHistoryNoteByIDV1(ctx context.Context, id string, req *AddHistoryNoteRequest) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if req == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProComputerExtensionAttributesV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(req).
		Post(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ListTemplatesV1 returns all computer extension attribute templates.
// URL: GET /api/v1/computer-extension-attributes/templates
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-templates
func (s *ComputerExtensionAttributes) ListTemplatesV1(ctx context.Context, rsqlQuery map[string]string) (*TemplateListResponse, *resty.Response, error) {
	var result TemplateListResponse

	endpoint := constants.EndpointJamfProComputerExtensionAttributesV1 + "/templates"

	mergePage := func(pageData []byte) error {
		var items []ResourceComputerExtensionAttributeTemplate
		if err := json.Unmarshal(pageData, &items); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, items...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list computer extension attribute templates: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetTemplateByIDV1 returns the specified computer extension attribute template by ID.
// URL: GET /api/v1/computer-extension-attributes/templates/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-templates-id
func (s *ComputerExtensionAttributes) GetTemplateByIDV1(ctx context.Context, id string) (*ResourceComputerExtensionAttributeTemplate, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("template ID is required")
	}

	endpoint := fmt.Sprintf("%s/templates/%s", constants.EndpointJamfProComputerExtensionAttributesV1, id)

	var result ResourceComputerExtensionAttributeTemplate

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UploadV1 uploads a computer extension attribute (multipart/form-data).
// URL: POST /api/v1/computer-extension-attributes/upload
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-extension-attributes-upload
func (s *ComputerExtensionAttributes) UploadV1(ctx context.Context, fileReader io.Reader, fileSize int64, filename string) (*ResourceComputerExtensionAttribute, *resty.Response, error) {
	if fileReader == nil {
		return nil, nil, fmt.Errorf("file reader is required")
	}
	if filename == "" {
		return nil, nil, fmt.Errorf("filename is required")
	}

	endpoint := constants.EndpointJamfProComputerExtensionAttributesV1 + "/upload"

	var result ResourceComputerExtensionAttribute

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetMultipartFile("file", filename, fileReader, fileSize, nil).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetDataDependencyByIDV1 returns smart group/advanced search dependent objects for the specified computer extension attribute.
// URL: GET /api/v1/computer-extension-attributes/{id}/data-dependency
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-id-data-dependency
func (s *ComputerExtensionAttributes) GetDataDependencyByIDV1(ctx context.Context, id string) (*DataDependencyResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("computer extension attribute ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/data-dependency", constants.EndpointJamfProComputerExtensionAttributesV1, id)

	var result DataDependencyResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DownloadByIDV1 downloads the specified computer extension attribute in XML format.
// URL: GET /api/v1/computer-extension-attributes/{id}/download
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-id-download
func (s *ComputerExtensionAttributes) DownloadByIDV1(ctx context.Context, id string) ([]byte, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("computer extension attribute ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/download", constants.EndpointJamfProComputerExtensionAttributesV1, id)

	var result []byte

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
