package computer_history

import (
	"context"
	"fmt"
	"net/url"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// ServiceInterface defines the interface for Classic API computer history operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computerhistory
	ServiceInterface interface {
		// GetByID retrieves computer history by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyid
		GetByID(ctx context.Context, id string) (*ResourceComputerHistory, *resty.Response, error)

		// GetByIDAndSubset retrieves a subset of computer history by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyid
		GetByIDAndSubset(ctx context.Context, id string, subset string) (*ResourceComputerHistory, *resty.Response, error)

		// GetByName retrieves computer history by computer name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyname
		GetByName(ctx context.Context, name string) (*ResourceComputerHistory, *resty.Response, error)

		// GetByNameAndSubset retrieves a subset of computer history by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyname
		GetByNameAndSubset(ctx context.Context, name string, subset string) (*ResourceComputerHistory, *resty.Response, error)

		// GetByUDID retrieves computer history by UDID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyudid
		GetByUDID(ctx context.Context, udid string) (*ResourceComputerHistory, *resty.Response, error)

		// GetByUDIDAndSubset retrieves a subset of computer history by UDID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyudid
		GetByUDIDAndSubset(ctx context.Context, udid string, subset string) (*ResourceComputerHistory, *resty.Response, error)

		// GetBySerialNumber retrieves computer history by serial number.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyserialnumber
		GetBySerialNumber(ctx context.Context, serialNumber string) (*ResourceComputerHistory, *resty.Response, error)

		// GetBySerialNumberAndSubset retrieves a subset of computer history by serial number.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyserialnumber
		GetBySerialNumberAndSubset(ctx context.Context, serialNumber string, subset string) (*ResourceComputerHistory, *resty.Response, error)

		// GetByMACAddress retrieves computer history by MAC address.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybymacaddress
		GetByMACAddress(ctx context.Context, macAddress string) (*ResourceComputerHistory, *resty.Response, error)

		// GetByMACAddressAndSubset retrieves a subset of computer history by MAC address.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybymacaddress
		GetByMACAddressAndSubset(ctx context.Context, macAddress string, subset string) (*ResourceComputerHistory, *resty.Response, error)
	}

	// Service handles communication with the computer-history-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computerhistory
	ComputerHistory struct {
		client interfaces.HTTPClient
	}
)

var _ ServiceInterface = (*ComputerHistory)(nil)

// NewService returns a new computer history Service backed by the provided HTTP client.
func NewComputerHistory(client interfaces.HTTPClient) *ComputerHistory {
	return &ComputerHistory{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Computer History CRUD Operations
// -----------------------------------------------------------------------------

// GetByID retrieves computer history by ID.
// URL: GET /JSSResource/computerhistory/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyid
func (s *ComputerHistory) GetByID(ctx context.Context, id string) (*ResourceComputerHistory, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("computer history ID cannot be empty")
	}
	return s.get(ctx, fmt.Sprintf("%s/id/%s", EndpointComputerHistory, url.PathEscape(id)))
}

// GetByIDAndSubset retrieves a subset of computer history by ID.
// URL: GET /JSSResource/computerhistory/id/{id}/subset/{subset}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyid
func (s *ComputerHistory) GetByIDAndSubset(ctx context.Context, id string, subset string) (*ResourceComputerHistory, *resty.Response, error) {
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
func (s *ComputerHistory) GetByName(ctx context.Context, name string) (*ResourceComputerHistory, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("computer name cannot be empty")
	}
	return s.get(ctx, fmt.Sprintf("%s/name/%s", EndpointComputerHistory, url.PathEscape(name)))
}

// GetByNameAndSubset retrieves a subset of computer history by name.
// URL: GET /JSSResource/computerhistory/name/{name}/subset/{subset}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyname
func (s *ComputerHistory) GetByNameAndSubset(ctx context.Context, name string, subset string) (*ResourceComputerHistory, *resty.Response, error) {
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
func (s *ComputerHistory) GetByUDID(ctx context.Context, udid string) (*ResourceComputerHistory, *resty.Response, error) {
	if udid == "" {
		return nil, nil, fmt.Errorf("UDID cannot be empty")
	}
	return s.get(ctx, fmt.Sprintf("%s/udid/%s", EndpointComputerHistory, url.PathEscape(udid)))
}

// GetByUDIDAndSubset retrieves a subset of computer history by UDID.
// URL: GET /JSSResource/computerhistory/udid/{udid}/subset/{subset}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyudid
func (s *ComputerHistory) GetByUDIDAndSubset(ctx context.Context, udid string, subset string) (*ResourceComputerHistory, *resty.Response, error) {
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
func (s *ComputerHistory) GetBySerialNumber(ctx context.Context, serialNumber string) (*ResourceComputerHistory, *resty.Response, error) {
	if serialNumber == "" {
		return nil, nil, fmt.Errorf("serial number cannot be empty")
	}
	return s.get(ctx, fmt.Sprintf("%s/serialnumber/%s", EndpointComputerHistory, url.PathEscape(serialNumber)))
}

// GetBySerialNumberAndSubset retrieves a subset of computer history by serial number.
// URL: GET /JSSResource/computerhistory/serialnumber/{serialNumber}/subset/{subset}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyserialnumber
func (s *ComputerHistory) GetBySerialNumberAndSubset(ctx context.Context, serialNumber string, subset string) (*ResourceComputerHistory, *resty.Response, error) {
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
func (s *ComputerHistory) GetByMACAddress(ctx context.Context, macAddress string) (*ResourceComputerHistory, *resty.Response, error) {
	if macAddress == "" {
		return nil, nil, fmt.Errorf("MAC address cannot be empty")
	}
	return s.get(ctx, fmt.Sprintf("%s/macaddress/%s", EndpointComputerHistory, url.PathEscape(macAddress)))
}

// GetByMACAddressAndSubset retrieves a subset of computer history by MAC address.
// URL: GET /JSSResource/computerhistory/macaddress/{macAddress}/subset/{subset}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybymacaddress
func (s *ComputerHistory) GetByMACAddressAndSubset(ctx context.Context, macAddress string, subset string) (*ResourceComputerHistory, *resty.Response, error) {
	if macAddress == "" {
		return nil, nil, fmt.Errorf("MAC address cannot be empty")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("subset cannot be empty")
	}
	return s.get(ctx, fmt.Sprintf("%s/macaddress/%s/subset/%s", EndpointComputerHistory, url.PathEscape(macAddress), url.PathEscape(subset)))
}

// get performs the GET request and unmarshals the response.
func (s *ComputerHistory) get(ctx context.Context, endpoint string) (*ResourceComputerHistory, *resty.Response, error) {
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
