package computer_inventory_collection

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the computer inventory collection-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computerinventorycollection
	ComputerInventoryCollection struct {
		client client.Client
	}
)

// NewService returns a new computer inventory collection Service backed by the provided HTTP client.
func NewComputerInventoryCollection(client client.Client) *ComputerInventoryCollection {
	return &ComputerInventoryCollection{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Computer Inventory Collection Operations
// -----------------------------------------------------------------------------

// Get retrieves the computer inventory collection settings.
// URL: GET /JSSResource/computerinventorycollection
// https://developer.jamf.com/jamf-pro/reference/computerinventorycollection
func (s *ComputerInventoryCollection) Get(ctx context.Context) (*ResourceComputerInventoryCollection, *resty.Response, error) {
	var result ResourceComputerInventoryCollection

	endpoint := constants.EndpointClassicComputerInventoryCollection

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

// Update updates the computer inventory collection settings.
// URL: PUT /JSSResource/computerinventorycollection
// https://developer.jamf.com/jamf-pro/reference/computerinventorycollection
func (s *ComputerInventoryCollection) Update(ctx context.Context, settings *ResourceComputerInventoryCollection) (*resty.Response, error) {
	if settings == nil {
		return nil, fmt.Errorf("settings is required")
	}

	endpoint := constants.EndpointClassicComputerInventoryCollection

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(settings).
		Put(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
