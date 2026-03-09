package network_segments

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// NetworkSegmentsServiceInterface defines the interface for Classic API network segment operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/networksegments
	NetworkSegmentsServiceInterface interface {
		// List returns all network segments.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findnetworksegments
		List(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified network segment by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findnetworksegmentsbyid
		GetByID(ctx context.Context, id int) (*ResourceNetworkSegment, *resty.Response, error)

		// GetByName returns the specified network segment by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findnetworksegmentsbyname
		GetByName(ctx context.Context, name string) (*ResourceNetworkSegment, *resty.Response, error)

		// Create creates a new network segment.
		//
		// Returns the created resource ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createnetworksegmentbyid
		Create(ctx context.Context, req *RequestNetworkSegment) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByID updates the specified network segment by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatenetworksegmentbyid
		UpdateByID(ctx context.Context, id int, req *RequestNetworkSegment) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByName updates the specified network segment by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatenetworksegmentbyname
		UpdateByName(ctx context.Context, name string, req *RequestNetworkSegment) (*CreateUpdateResponse, *resty.Response, error)

		// DeleteByID removes the specified network segment by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletenetworksegmentbyid
		DeleteByID(ctx context.Context, id int) (*resty.Response, error)

		// DeleteByName removes the specified network segment by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletenetworksegmentbyname
		DeleteByName(ctx context.Context, name string) (*resty.Response, error)
	}

	// Service handles communication with the network segment-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/networksegments
	NetworkSegments struct {
		client transport.HTTPClient
	}
)

var _ NetworkSegmentsServiceInterface = (*NetworkSegments)(nil)

// NewService returns a new network segments Service backed by the provided HTTP client.
func NewNetworkSegments(client transport.HTTPClient) *NetworkSegments {
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

	endpoint := EndpointClassicNetworkSegments

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
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

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicNetworkSegments, id)

	var result ResourceNetworkSegment

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
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

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicNetworkSegments, name)

	var result ResourceNetworkSegment

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
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

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicNetworkSegments)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
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

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicNetworkSegments, id)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
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

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicNetworkSegments, name)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
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

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicNetworkSegments, id)

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

// DeleteByName removes the specified network segment by name.
// URL: DELETE /JSSResource/networksegments/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deletenetworksegmentbyname
func (s *NetworkSegments) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("network segment name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicNetworkSegments, name)

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
