package file_share_distribution_points

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the file-share-distribution-points-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/distributionpoints
	FileShareDistributionPoints struct {
		client client.Client
	}
)

// NewService returns a new file share distribution points Service backed by the provided HTTP client.
func NewFileShareDistributionPoints(client client.Client) *FileShareDistributionPoints {
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
	var out ListResponse

	endpoint := constants.EndpointClassicFileShareDistributionPoints

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&out).
		Get(endpoint)

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

	var out ResourceFileShareDistributionPoint

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicFileShareDistributionPoints, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&out).
		Get(endpoint)

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

	var out ResourceFileShareDistributionPoint

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicFileShareDistributionPoints, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&out).
		Get(endpoint)

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

	var out CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicFileShareDistributionPoints)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&out).
		Post(endpoint)

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

	var out CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicFileShareDistributionPoints, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&out).
		Put(endpoint)

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

	var out CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicFileShareDistributionPoints, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&out).
		Put(endpoint)

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

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicFileShareDistributionPoints, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

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

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicFileShareDistributionPoints, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
