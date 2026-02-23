package computer_inventory_collection

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ComputerInventoryCollectionServiceInterface defines the interface for Classic API computer inventory collection operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computerinventorycollection
	ComputerInventoryCollectionServiceInterface interface {
		// Get retrieves the computer inventory collection settings.
		Get(ctx context.Context) (*ResourceComputerInventoryCollection, *interfaces.Response, error)

		// Update updates the computer inventory collection settings.
		Update(ctx context.Context, settings *ResourceComputerInventoryCollection) (*interfaces.Response, error)
	}

	// Service handles communication with the computer inventory collection-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computerinventorycollection
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ComputerInventoryCollectionServiceInterface = (*Service)(nil)

// NewService returns a new computer inventory collection Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Computer Inventory Collection Operations
// -----------------------------------------------------------------------------

// Get retrieves the computer inventory collection settings.
// URL: GET /JSSResource/computerinventorycollection
// https://developer.jamf.com/jamf-pro/reference/computerinventorycollection
func (s *Service) Get(ctx context.Context) (*ResourceComputerInventoryCollection, *interfaces.Response, error) {
	var result ResourceComputerInventoryCollection

	endpoint := EndpointClassicComputerInventoryCollection

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

// Update updates the computer inventory collection settings.
// URL: PUT /JSSResource/computerinventorycollection
// https://developer.jamf.com/jamf-pro/reference/computerinventorycollection
func (s *Service) Update(ctx context.Context, settings *ResourceComputerInventoryCollection) (*interfaces.Response, error) {
	if settings == nil {
		return nil, fmt.Errorf("settings is required")
	}

	endpoint := EndpointClassicComputerInventoryCollection

	requestBody := struct {
		XMLName xml.Name `xml:"computer_inventory_collection"`
		*ResourceComputerInventoryCollection
	}{
		ResourceComputerInventoryCollection: settings,
	}

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, &requestBody, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
