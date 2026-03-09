package file_share_distribution_points

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// FileShareDistributionPointsServiceInterface defines the interface for Classic API file share distribution point operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/distributionpoints
	FileShareDistributionPointsServiceInterface interface {
		// List returns all file share distribution points.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/distributionpoints
		List(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified file share distribution point by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finddistributionpointsbyid
		GetByID(ctx context.Context, id int) (*ResourceFileShareDistributionPoint, *resty.Response, error)

		// GetByName returns the specified file share distribution point by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finddistributionpointsbyname
		GetByName(ctx context.Context, name string) (*ResourceFileShareDistributionPoint, *resty.Response, error)

		// Create creates a new file share distribution point.
		//
		// Returns the created distribution point ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createdistributionpointbyid
		Create(ctx context.Context, req *RequestFileShareDistributionPoint) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByID updates the specified file share distribution point by ID.
		//
		// Returns the updated distribution point ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatedistributionpointbyid
		UpdateByID(ctx context.Context, id int, req *RequestFileShareDistributionPoint) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByName updates the specified file share distribution point by name.
		//
		// Returns the updated distribution point ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatedistributionpointbyname
		UpdateByName(ctx context.Context, name string, req *RequestFileShareDistributionPoint) (*CreateUpdateResponse, *resty.Response, error)

		// DeleteByID removes the specified file share distribution point by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletedistributionpointbyid
		DeleteByID(ctx context.Context, id int) (*resty.Response, error)

		// DeleteByName removes the specified file share distribution point by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletedistributionpointbyname
		DeleteByName(ctx context.Context, name string) (*resty.Response, error)
	}

	// Service handles communication with the file-share-distribution-points-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/distributionpoints
	FileShareDistributionPoints struct {
		client transport.HTTPClient
	}
)

var _ FileShareDistributionPointsServiceInterface = (*FileShareDistributionPoints)(nil)

// NewService returns a new file share distribution points Service backed by the provided HTTP client.
func NewFileShareDistributionPoints(client transport.HTTPClient) *FileShareDistributionPoints {
	return &FileShareDistributionPoints{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - File Share Distribution Points CRUD Operations
// -----------------------------------------------------------------------------

// List returns all file share distribution points.
// URL: GET /JSSResource/distributionpoints
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/distributionpoints
func (s *FileShareDistributionPoints) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	endpoint := EndpointFileShareDistributionPoints

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

// GetByID returns the specified file share distribution point by ID.
// URL: GET /JSSResource/distributionpoints/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finddistributionpointsbyid
func (s *FileShareDistributionPoints) GetByID(ctx context.Context, id int) (*ResourceFileShareDistributionPoint, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("distribution point ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointFileShareDistributionPoints, id)

	var out ResourceFileShareDistributionPoint

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

// GetByName returns the specified file share distribution point by name.
// URL: GET /JSSResource/distributionpoints/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finddistributionpointsbyname
func (s *FileShareDistributionPoints) GetByName(ctx context.Context, name string) (*ResourceFileShareDistributionPoint, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("distribution point name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointFileShareDistributionPoints, name)

	var out ResourceFileShareDistributionPoint

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

// Create creates a new file share distribution point.
//
// Returns the created distribution point ID only (Classic API behavior).
// URL: POST /JSSResource/distributionpoints/id/0
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createdistributionpointbyid
func (s *FileShareDistributionPoints) Create(ctx context.Context, req *RequestFileShareDistributionPoint) (*CreateUpdateResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("distribution point name is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointFileShareDistributionPoints)

	var out CreateUpdateResponse

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

// UpdateByID updates the specified file share distribution point by ID.
//
// Returns the updated distribution point ID only (Classic API behavior).
// URL: PUT /JSSResource/distributionpoints/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatedistributionpointbyid
func (s *FileShareDistributionPoints) UpdateByID(ctx context.Context, id int, req *RequestFileShareDistributionPoint) (*CreateUpdateResponse, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("distribution point ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("distribution point name is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointFileShareDistributionPoints, id)

	var out CreateUpdateResponse

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

// UpdateByName updates the specified file share distribution point by name.
//
// Returns the updated distribution point ID only (Classic API behavior).
// URL: PUT /JSSResource/distributionpoints/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatedistributionpointbyname
func (s *FileShareDistributionPoints) UpdateByName(ctx context.Context, name string, req *RequestFileShareDistributionPoint) (*CreateUpdateResponse, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("distribution point name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("distribution point name is required in request")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointFileShareDistributionPoints, name)

	var out CreateUpdateResponse

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

// DeleteByID removes the specified file share distribution point by ID.
// URL: DELETE /JSSResource/distributionpoints/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletedistributionpointbyid
func (s *FileShareDistributionPoints) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("distribution point ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointFileShareDistributionPoints, id)

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

// DeleteByName removes the specified file share distribution point by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletedistributionpointbyname
func (s *FileShareDistributionPoints) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("distribution point name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointFileShareDistributionPoints, name)

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
