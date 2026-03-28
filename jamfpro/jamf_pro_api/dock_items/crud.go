package dock_items

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the dock items-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-dock-items
	DockItems struct {
		client client.Client
	}
)

func NewDockItems(client client.Client) *DockItems {
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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
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

	if _, ok := validDockItemTypes[request.Type]; !ok {
		return nil, nil, fmt.Errorf("invalid type %q: must be one of APP, FILE, FOLDER", request.Type)
	}

	var result CreateResponse

	endpoint := constants.EndpointJamfProDockItemsV1

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

	if _, ok := validDockItemTypes[request.Type]; !ok {
		return nil, nil, fmt.Errorf("invalid type %q: must be one of APP, FILE, FOLDER", request.Type)
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProDockItemsV1, id)

	var result ResourceDockItem

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

// DeleteByIDV1 removes the specified dock item by ID (Remove specified Dock Item record).
// URL: DELETE /api/v1/dock-items/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-dock-items-id
func (s *DockItems) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("dock item ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProDockItemsV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
