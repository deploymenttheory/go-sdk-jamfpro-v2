package dock_items

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// DockItemsServiceInterface defines the interface for dock item operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-dock-items
	DockItemsServiceInterface interface {
		// GetByIDV1 returns the specified dock item by ID (Get specified Dock Item object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-dock-items-id
		GetByIDV1(ctx context.Context, id string) (*ResourceDockItem, *interfaces.Response, error)

		// CreateV1 creates a new dock item record (Create Dock Item record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-dock-items
		CreateV1(ctx context.Context, request *RequestDockItem) (*CreateResponse, *interfaces.Response, error)

		// UpdateByIDV1 updates the specified dock item by ID (Update specified Dock Item object).
		//
		// Returns the full updated dock item resource.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-dock-items-id
		UpdateByIDV1(ctx context.Context, id string, request *RequestDockItem) (*ResourceDockItem, *interfaces.Response, error)

		// DeleteByIDV1 removes the specified dock item by ID (Remove specified Dock Item record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-dock-items-id
		DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error)
	}

	// Service handles communication with the dock items-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-dock-items
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ DockItemsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Dock Items CRUD Operations
// -----------------------------------------------------------------------------

// GetByIDV1 returns the specified dock item by ID (Get specified Dock Item object).
// URL: GET /api/v1/dock-items/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-dock-items-id
func (s *Service) GetByIDV1(ctx context.Context, id string) (*ResourceDockItem, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("dock item ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointDockItemsV1, id)

	var result ResourceDockItem

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

// CreateV1 creates a new dock item record (Create Dock Item record).
// URL: POST /api/v1/dock-items
// Body: JSON with name, path, type
// https://developer.jamf.com/jamf-pro/reference/post_v1-dock-items
func (s *Service) CreateV1(ctx context.Context, request *RequestDockItem) (*CreateResponse, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := EndpointDockItemsV1

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

// UpdateByIDV1 updates the specified dock item by ID (Update specified Dock Item object).
// URL: PUT /api/v1/dock-items/{id}
// Body: JSON with name, path, type
// https://developer.jamf.com/jamf-pro/reference/put_v1-dock-items-id
func (s *Service) UpdateByIDV1(ctx context.Context, id string, request *RequestDockItem) (*ResourceDockItem, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointDockItemsV1, id)

	var result ResourceDockItem

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

// DeleteByIDV1 removes the specified dock item by ID (Remove specified Dock Item record).
// URL: DELETE /api/v1/dock-items/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-dock-items-id
func (s *Service) DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("dock item ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointDockItemsV1, id)

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
