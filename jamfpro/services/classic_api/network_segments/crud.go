package network_segments

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

type (
	// NetworkSegmentsServiceInterface defines the interface for Classic API network segment operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/networksegments
	NetworkSegmentsServiceInterface interface {
		// ListNetworkSegments returns all network segments.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findallnetworksegments
		ListNetworkSegments(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetNetworkSegmentByID returns the specified network segment by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findnetworksegmentsbyid
		GetNetworkSegmentByID(ctx context.Context, id int) (*ResourceNetworkSegment, *interfaces.Response, error)

		// GetNetworkSegmentByName returns the specified network segment by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findnetworksegmentsbyname
		GetNetworkSegmentByName(ctx context.Context, name string) (*ResourceNetworkSegment, *interfaces.Response, error)

		// CreateNetworkSegment creates a new network segment.
		//
		// Returns the created resource ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createnetworksegmentbyid
		CreateNetworkSegment(ctx context.Context, req *RequestNetworkSegment) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateNetworkSegmentByID updates the specified network segment by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatenetworksegmentbyid
		UpdateNetworkSegmentByID(ctx context.Context, id int, req *RequestNetworkSegment) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateNetworkSegmentByName updates the specified network segment by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatenetworksegmentbyname
		UpdateNetworkSegmentByName(ctx context.Context, name string, req *RequestNetworkSegment) (*CreateUpdateResponse, *interfaces.Response, error)

		// DeleteNetworkSegmentByID removes the specified network segment by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletenetworksegmentbyid
		DeleteNetworkSegmentByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteNetworkSegmentByName removes the specified network segment by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletenetworksegmentbyname
		DeleteNetworkSegmentByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the network segment-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/networksegments
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ NetworkSegmentsServiceInterface = (*Service)(nil)

// NewService returns a new network segments Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Network Segments CRUD Operations
// -----------------------------------------------------------------------------

// ListNetworkSegments returns all network segments.
// URL: GET /JSSResource/networksegments
// https://developer.jamf.com/jamf-pro/reference/findallnetworksegments
func (s *Service) ListNetworkSegments(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	resp, err := s.client.Get(ctx, EndpointClassicNetworkSegments, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetNetworkSegmentByID returns the specified network segment by ID.
// URL: GET /JSSResource/networksegments/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findnetworksegmentsbyid
func (s *Service) GetNetworkSegmentByID(ctx context.Context, id int) (*ResourceNetworkSegment, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("network segment ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicNetworkSegments, id)

	var result ResourceNetworkSegment

	resp, err := s.client.Get(ctx, endpoint, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetNetworkSegmentByName returns the specified network segment by name.
// URL: GET /JSSResource/networksegments/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findnetworksegmentsbyname
func (s *Service) GetNetworkSegmentByName(ctx context.Context, name string) (*ResourceNetworkSegment, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("network segment name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicNetworkSegments, name)

	var result ResourceNetworkSegment

	resp, err := s.client.Get(ctx, endpoint, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateNetworkSegment creates a new network segment.
// URL: POST /JSSResource/networksegments/id/0
// Returns the created resource ID.
// https://developer.jamf.com/jamf-pro/reference/createnetworksegmentbyid
func (s *Service) CreateNetworkSegment(ctx context.Context, req *RequestNetworkSegment) (*CreateUpdateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicNetworkSegments)

	var result CreateUpdateResponse

	resp, err := s.client.Post(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateNetworkSegmentByID updates the specified network segment by ID.
// URL: PUT /JSSResource/networksegments/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updatenetworksegmentbyid
func (s *Service) UpdateNetworkSegmentByID(ctx context.Context, id int, req *RequestNetworkSegment) (*CreateUpdateResponse, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("network segment ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicNetworkSegments, id)

	var result CreateUpdateResponse

	resp, err := s.client.Put(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateNetworkSegmentByName updates the specified network segment by name.
// URL: PUT /JSSResource/networksegments/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updatenetworksegmentbyname
func (s *Service) UpdateNetworkSegmentByName(ctx context.Context, name string, req *RequestNetworkSegment) (*CreateUpdateResponse, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("network segment name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicNetworkSegments, name)

	var result CreateUpdateResponse

	resp, err := s.client.Put(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteNetworkSegmentByID removes the specified network segment by ID.
// URL: DELETE /JSSResource/networksegments/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deletenetworksegmentbyid
func (s *Service) DeleteNetworkSegmentByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("network segment ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicNetworkSegments, id)

	resp, err := s.client.Delete(ctx, endpoint, nil, shared.XMLHeaders(), nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteNetworkSegmentByName removes the specified network segment by name.
// URL: DELETE /JSSResource/networksegments/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deletenetworksegmentbyname
func (s *Service) DeleteNetworkSegmentByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("network segment name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicNetworkSegments, name)

	resp, err := s.client.Delete(ctx, endpoint, nil, shared.XMLHeaders(), nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
