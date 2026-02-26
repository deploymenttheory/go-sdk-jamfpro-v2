package inventory_information

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// InventoryInformationServiceInterface defines the interface for inventory information operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-inventory-information
	InventoryInformationServiceInterface interface {
		// GetV1 returns statistics about managed/unmanaged devices and computers in the inventory.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-inventory-information
		GetV1(ctx context.Context) (*ResourceInventoryInformation, *interfaces.Response, error)
	}

	// Service handles communication with the inventory information-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-inventory-information
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ InventoryInformationServiceInterface = (*Service)(nil)

// NewService returns a new inventory information Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Inventory Information Operations
// -----------------------------------------------------------------------------

// GetV1 returns statistics about managed/unmanaged devices and computers in the inventory.
// URL: GET /api/v1/inventory-information
// Response: object with managedComputers, unmanagedComputers, managedDevices, unmanagedDevices.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-inventory-information
func (s *Service) GetV1(ctx context.Context) (*ResourceInventoryInformation, *interfaces.Response, error) {
	endpoint := EndpointInventoryInformationV1

	var result ResourceInventoryInformation

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get inventory information: %w", err)
	}

	return &result, resp, nil
}
