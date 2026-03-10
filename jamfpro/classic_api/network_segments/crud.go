package network_segments

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the network segment-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/networksegments
	NetworkSegments struct {
		client client.Client
	}
)

// NewService returns a new network segments Service backed by the provided HTTP client.
func NewNetworkSegments(client client.Client) *NetworkSegments {
	return &NetworkSegments{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Network Segments CRUD Operations
// -----------------------------------------------------------------------------

// List returns all network segments.
// URL: GET /JSSResource/networksegments
// https://developer.jamf.com/jamf-pro/reference/findnetworksegments
func (s *NetworkSegments) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointClassicNetworkSegments

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByID returns the specified network segment by ID.
// URL: GET /JSSResource/networksegments/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findnetworksegmentsbyid
func (s *NetworkSegments) GetByID(ctx context.Context, id int) (*ResourceNetworkSegment, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("network segment ID must be a positive integer")
	}

	var result ResourceNetworkSegment

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicNetworkSegments, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByName returns the specified network segment by name.
// URL: GET /JSSResource/networksegments/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findnetworksegmentsbyname
func (s *NetworkSegments) GetByName(ctx context.Context, name string) (*ResourceNetworkSegment, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("network segment name is required")
	}

	var result ResourceNetworkSegment

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicNetworkSegments, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// Create creates a new network segment.
// URL: POST /JSSResource/networksegments/id/0
// Returns the created resource ID.
// https://developer.jamf.com/jamf-pro/reference/createnetworksegmentbyid
func (s *NetworkSegments) Create(ctx context.Context, req *RequestNetworkSegment) (*CreateUpdateResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicNetworkSegments)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&result).
		Post(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByID updates the specified network segment by ID.
// URL: PUT /JSSResource/networksegments/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updatenetworksegmentbyid
func (s *NetworkSegments) UpdateByID(ctx context.Context, id int, req *RequestNetworkSegment) (*CreateUpdateResponse, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("network segment ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicNetworkSegments, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&result).
		Put(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByName updates the specified network segment by name.
// URL: PUT /JSSResource/networksegments/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updatenetworksegmentbyname
func (s *NetworkSegments) UpdateByName(ctx context.Context, name string, req *RequestNetworkSegment) (*CreateUpdateResponse, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("network segment name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicNetworkSegments, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&result).
		Put(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByID removes the specified network segment by ID.
// URL: DELETE /JSSResource/networksegments/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deletenetworksegmentbyid
func (s *NetworkSegments) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("network segment ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicNetworkSegments, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByName removes the specified network segment by name.
// URL: DELETE /JSSResource/networksegments/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deletenetworksegmentbyname
func (s *NetworkSegments) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("network segment name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicNetworkSegments, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
