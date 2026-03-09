package dock_items

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// ServiceInterface defines the interface for Classic API dock item operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/dockitems
	ServiceInterface interface {
		// List returns all dock items.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finddockitems
		List(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified dock item by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finddockitemsbyid
		GetByID(ctx context.Context, id int) (*Resource, *resty.Response, error)

		// GetByName returns the specified dock item by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finddockitemsbyname
		GetByName(ctx context.Context, name string) (*Resource, *resty.Response, error)

		// Create creates a new dock item.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createdockitembyid
		Create(ctx context.Context, req *Request) (*Resource, *resty.Response, error)

		// UpdateByID updates the specified dock item by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatedockitembyid
		UpdateByID(ctx context.Context, id int, req *Request) (*Resource, *resty.Response, error)

		// UpdateByName updates the specified dock item by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatedockitembyname
		UpdateByName(ctx context.Context, name string, req *Request) (*Resource, *resty.Response, error)

		// DeleteByID removes the specified dock item by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletedockitembyid
		DeleteByID(ctx context.Context, id int) (*resty.Response, error)

		// DeleteByName removes the specified dock item by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletedockitembyname
		DeleteByName(ctx context.Context, name string) (*resty.Response, error)
	}

	// Service handles communication with the dock-items-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/dockitems
	DockItems struct {
		client interfaces.HTTPClient
	}
)

var _ ServiceInterface = (*DockItems)(nil)

// NewService returns a new dock items Service backed by the provided HTTP client.
func NewDockItems(client interfaces.HTTPClient) *DockItems {
	return &DockItems{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Dock Items CRUD Operations
// -----------------------------------------------------------------------------

// List returns all dock items.
// URL: GET /JSSResource/dockitems
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finddockitems
func (s *DockItems) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	endpoint := EndpointDockItems

	var out ListResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// GetByID returns the specified dock item by ID.
// URL: GET /JSSResource/dockitems/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finddockitemsbyid
func (s *DockItems) GetByID(ctx context.Context, id int) (*Resource, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("dock item ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointDockItems, id)

	var out Resource

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// GetByName returns the specified dock item by name.
// URL: GET /JSSResource/dockitems/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finddockitemsbyname
func (s *DockItems) GetByName(ctx context.Context, name string) (*Resource, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("dock item name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointDockItems, name)

	var out Resource

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Create creates a new dock item.
// URL: POST /JSSResource/dockitems/id/0
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createdockitembyid
func (s *DockItems) Create(ctx context.Context, req *Request) (*Resource, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("dock item name is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointDockItems)

	var out Resource

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// UpdateByID updates the specified dock item by ID.
// URL: PUT /JSSResource/dockitems/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatedockitembyid
func (s *DockItems) UpdateByID(ctx context.Context, id int, req *Request) (*Resource, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("dock item ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("dock item name is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointDockItems, id)

	var out Resource

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// UpdateByName updates the specified dock item by name.
// URL: PUT /JSSResource/dockitems/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatedockitembyname
func (s *DockItems) UpdateByName(ctx context.Context, name string, req *Request) (*Resource, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("dock item name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("dock item name is required in request")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointDockItems, name)

	var out Resource

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// DeleteByID removes the specified dock item by ID.
// URL: DELETE /JSSResource/dockitems/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletedockitembyid
func (s *DockItems) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("dock item ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointDockItems, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// DeleteByName removes the specified dock item by name.
// URL: DELETE /JSSResource/dockitems/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletedockitembyname
func (s *DockItems) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("dock item name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointDockItems, name)

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
