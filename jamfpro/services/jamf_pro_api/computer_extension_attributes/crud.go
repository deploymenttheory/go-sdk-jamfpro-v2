package computer_extension_attributes

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ComputerExtensionAttributesServiceInterface defines the interface for computer extension attribute operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes
	ComputerExtensionAttributesServiceInterface interface {
		// ListComputerExtensionAttributesV1 returns all computer extension attribute objects (Get Computer Extension Attribute objects).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes
		ListComputerExtensionAttributesV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetComputerExtensionAttributeByIDV1 returns the specified computer extension attribute by ID (Get specified Computer Extension Attribute object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-id
		GetComputerExtensionAttributeByIDV1(ctx context.Context, id string) (*ResourceComputerExtensionAttribute, *interfaces.Response, error)

		// CreateComputerExtensionAttributeV1 creates a new computer extension attribute (Create Computer Extension Attribute record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-extension-attributes
		CreateComputerExtensionAttributeV1(ctx context.Context, req *RequestComputerExtensionAttribute) (*CreateResponse, *interfaces.Response, error)

		// UpdateComputerExtensionAttributeByIDV1 updates the specified computer extension attribute by ID (Update specified Computer Extension Attribute object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-computer-extension-attributes-id
		UpdateComputerExtensionAttributeByIDV1(ctx context.Context, id string, req *RequestComputerExtensionAttribute) (*ResourceComputerExtensionAttribute, *interfaces.Response, error)

		// DeleteComputerExtensionAttributeByIDV1 removes the specified computer extension attribute by ID (Remove specified Computer Extension Attribute record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-computer-extension-attributes-id
		DeleteComputerExtensionAttributeByIDV1(ctx context.Context, id string) (*interfaces.Response, error)

		// DeleteComputerExtensionAttributesByIDV1 deletes multiple computer extension attributes by their IDs (Delete multiple Computer Extension Attributes by their IDs).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-extension-attributes-delete-multiple
		DeleteComputerExtensionAttributesByIDV1(ctx context.Context, req *DeleteComputerExtensionAttributesByIDRequest) (*interfaces.Response, error)
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

// ListComputerExtensionAttributesV1 returns all computer extension attribute objects.
// URL: GET /api/v1/computer-extension-attributes
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes
func (s *Service) ListComputerExtensionAttributesV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointComputerExtensionAttributesV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetComputerExtensionAttributeByIDV1 returns the specified computer extension attribute by ID.
// URL: GET /api/v1/computer-extension-attributes/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-id
func (s *Service) GetComputerExtensionAttributeByIDV1(ctx context.Context, id string) (*ResourceComputerExtensionAttribute, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("computer extension attribute ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointComputerExtensionAttributesV1, id)

	var result ResourceComputerExtensionAttribute

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateComputerExtensionAttributeV1 creates a new computer extension attribute.
// URL: POST /api/v1/computer-extension-attributes
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-extension-attributes
func (s *Service) CreateComputerExtensionAttributeV1(ctx context.Context, req *RequestComputerExtensionAttribute) (*CreateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := EndpointComputerExtensionAttributesV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateComputerExtensionAttributeByIDV1 updates the specified computer extension attribute by ID.
// URL: PUT /api/v1/computer-extension-attributes/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-computer-extension-attributes-id
func (s *Service) UpdateComputerExtensionAttributeByIDV1(ctx context.Context, id string, req *RequestComputerExtensionAttribute) (*ResourceComputerExtensionAttribute, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointComputerExtensionAttributesV1, id)

	var result ResourceComputerExtensionAttribute

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteComputerExtensionAttributeByIDV1 removes the specified computer extension attribute by ID.
// URL: DELETE /api/v1/computer-extension-attributes/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-computer-extension-attributes-id
func (s *Service) DeleteComputerExtensionAttributeByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("computer extension attribute ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointComputerExtensionAttributesV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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
