package dock_items

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the dock items-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-dock-items
	DockItems struct {
		client transport.HTTPClient
	}
)

func NewDockItems(client transport.HTTPClient) *DockItems {
	return &DockItems{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Dock Items CRUD Operations
// -----------------------------------------------------------------------------

// GetByIDV1 returns the specified dock item by ID (Get specified Dock Item object).
// URL: GET /api/v1/dock-items/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-dock-items-id
func (s *DockItems) GetByIDV1(ctx context.Context, id string) (*ResourceDockItem, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("dock item ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProDockItemsV1, id)

	var result ResourceDockItem

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *DockItems) CreateV1(ctx context.Context, request *RequestDockItem) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := constants.EndpointJamfProDockItemsV1

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
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
func (s *DockItems) UpdateByIDV1(ctx context.Context, id string, request *RequestDockItem) (*ResourceDockItem, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProDockItemsV1, id)

	var result ResourceDockItem

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
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
func (s *DockItems) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("dock item ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProDockItemsV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
