package computer_history

import (
	"context"
	"fmt"
	"net/url"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ServiceInterface defines the interface for Classic API computer history operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computerhistory
	ServiceInterface interface {
		GetByID(ctx context.Context, id string) (*ResourceComputerHistory, *interfaces.Response, error)
		GetByIDAndSubset(ctx context.Context, id string, subset string) (*ResourceComputerHistory, *interfaces.Response, error)
		GetByName(ctx context.Context, name string) (*ResourceComputerHistory, *interfaces.Response, error)
		GetByNameAndSubset(ctx context.Context, name string, subset string) (*ResourceComputerHistory, *interfaces.Response, error)
		GetByUDID(ctx context.Context, udid string) (*ResourceComputerHistory, *interfaces.Response, error)
		GetByUDIDAndSubset(ctx context.Context, udid string, subset string) (*ResourceComputerHistory, *interfaces.Response, error)
		GetBySerialNumber(ctx context.Context, serialNumber string) (*ResourceComputerHistory, *interfaces.Response, error)
		GetBySerialNumberAndSubset(ctx context.Context, serialNumber string, subset string) (*ResourceComputerHistory, *interfaces.Response, error)
		GetByMACAddress(ctx context.Context, macAddress string) (*ResourceComputerHistory, *interfaces.Response, error)
		GetByMACAddressAndSubset(ctx context.Context, macAddress string, subset string) (*ResourceComputerHistory, *interfaces.Response, error)
	}

	// Service handles communication with the computer-history-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computerhistory
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ServiceInterface = (*Service)(nil)

// NewService returns a new computer history Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Computer History CRUD Operations
// -----------------------------------------------------------------------------

// GetByID retrieves computer history by ID.
// URL: GET /JSSResource/computerhistory/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyid
func (s *Service) GetByID(ctx context.Context, id string) (*ResourceComputerHistory, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("computer history ID cannot be empty")
	}
	return s.get(ctx, fmt.Sprintf("%s/id/%s", EndpointComputerHistory, url.PathEscape(id)))
}

// GetByIDAndSubset retrieves a subset of computer history by ID.
// URL: GET /JSSResource/computerhistory/id/{id}/subset/{subset}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyid
func (s *Service) GetByIDAndSubset(ctx context.Context, id string, subset string) (*ResourceComputerHistory, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("computer history ID cannot be empty")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("subset cannot be empty")
	}
	return s.get(ctx, fmt.Sprintf("%s/id/%s/subset/%s", EndpointComputerHistory, url.PathEscape(id), url.PathEscape(subset)))
}

// GetByName retrieves computer history by computer name.
// URL: GET /JSSResource/computerhistory/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyname
func (s *Service) GetByName(ctx context.Context, name string) (*ResourceComputerHistory, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("computer name cannot be empty")
	}
	return s.get(ctx, fmt.Sprintf("%s/name/%s", EndpointComputerHistory, url.PathEscape(name)))
}

// GetByNameAndSubset retrieves a subset of computer history by name.
// URL: GET /JSSResource/computerhistory/name/{name}/subset/{subset}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyname
func (s *Service) GetByNameAndSubset(ctx context.Context, name string, subset string) (*ResourceComputerHistory, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("computer name cannot be empty")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("subset cannot be empty")
	}
	return s.get(ctx, fmt.Sprintf("%s/name/%s/subset/%s", EndpointComputerHistory, url.PathEscape(name), url.PathEscape(subset)))
}

// GetByUDID retrieves computer history by UDID.
// URL: GET /JSSResource/computerhistory/udid/{udid}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyudid
func (s *Service) GetByUDID(ctx context.Context, udid string) (*ResourceComputerHistory, *interfaces.Response, error) {
	if udid == "" {
		return nil, nil, fmt.Errorf("UDID cannot be empty")
	}
	return s.get(ctx, fmt.Sprintf("%s/udid/%s", EndpointComputerHistory, url.PathEscape(udid)))
}

// GetByUDIDAndSubset retrieves a subset of computer history by UDID.
// URL: GET /JSSResource/computerhistory/udid/{udid}/subset/{subset}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyudid
func (s *Service) GetByUDIDAndSubset(ctx context.Context, udid string, subset string) (*ResourceComputerHistory, *interfaces.Response, error) {
	if udid == "" {
		return nil, nil, fmt.Errorf("UDID cannot be empty")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("subset cannot be empty")
	}
	return s.get(ctx, fmt.Sprintf("%s/udid/%s/subset/%s", EndpointComputerHistory, url.PathEscape(udid), url.PathEscape(subset)))
}

// GetBySerialNumber retrieves computer history by serial number.
// URL: GET /JSSResource/computerhistory/serialnumber/{serialNumber}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyserialnumber
func (s *Service) GetBySerialNumber(ctx context.Context, serialNumber string) (*ResourceComputerHistory, *interfaces.Response, error) {
	if serialNumber == "" {
		return nil, nil, fmt.Errorf("serial number cannot be empty")
	}
	return s.get(ctx, fmt.Sprintf("%s/serialnumber/%s", EndpointComputerHistory, url.PathEscape(serialNumber)))
}

// GetBySerialNumberAndSubset retrieves a subset of computer history by serial number.
// URL: GET /JSSResource/computerhistory/serialnumber/{serialNumber}/subset/{subset}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyserialnumber
func (s *Service) GetBySerialNumberAndSubset(ctx context.Context, serialNumber string, subset string) (*ResourceComputerHistory, *interfaces.Response, error) {
	if serialNumber == "" {
		return nil, nil, fmt.Errorf("serial number cannot be empty")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("subset cannot be empty")
	}
	return s.get(ctx, fmt.Sprintf("%s/serialnumber/%s/subset/%s", EndpointComputerHistory, url.PathEscape(serialNumber), url.PathEscape(subset)))
}

// GetByMACAddress retrieves computer history by MAC address.
// URL: GET /JSSResource/computerhistory/macaddress/{macAddress}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybymacaddress
func (s *Service) GetByMACAddress(ctx context.Context, macAddress string) (*ResourceComputerHistory, *interfaces.Response, error) {
	if macAddress == "" {
		return nil, nil, fmt.Errorf("MAC address cannot be empty")
	}
	return s.get(ctx, fmt.Sprintf("%s/macaddress/%s", EndpointComputerHistory, url.PathEscape(macAddress)))
}

// GetByMACAddressAndSubset retrieves a subset of computer history by MAC address.
// URL: GET /JSSResource/computerhistory/macaddress/{macAddress}/subset/{subset}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybymacaddress
func (s *Service) GetByMACAddressAndSubset(ctx context.Context, macAddress string, subset string) (*ResourceComputerHistory, *interfaces.Response, error) {
	if macAddress == "" {
		return nil, nil, fmt.Errorf("MAC address cannot be empty")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("subset cannot be empty")
	}
	return s.get(ctx, fmt.Sprintf("%s/macaddress/%s/subset/%s", EndpointComputerHistory, url.PathEscape(macAddress), url.PathEscape(subset)))
}

// get performs the GET request and unmarshals the response.
func (s *Service) get(ctx context.Context, endpoint string) (*ResourceComputerHistory, *interfaces.Response, error) {
	var out ResourceComputerHistory
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
