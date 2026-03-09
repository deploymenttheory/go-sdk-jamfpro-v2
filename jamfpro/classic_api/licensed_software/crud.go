package licensed_software

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// ServiceInterface defines the interface for Classic API licensed software operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/licensedsoftware
	ServiceInterface interface {
		// List returns all licensed software.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findlicensedsoftware
		List(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified licensed software by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findlicensedsoftwarebyid
		GetByID(ctx context.Context, id int) (*Resource, *resty.Response, error)

		// GetByName returns the specified licensed software by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findlicensedsoftwarebyname
		GetByName(ctx context.Context, name string) (*Resource, *resty.Response, error)

		// Create creates a new licensed software item.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createlicensedsoftwarebyid
		Create(ctx context.Context, req *Resource) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByID updates the specified licensed software by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatelicensedsoftwarebyid
		UpdateByID(ctx context.Context, id int, req *Resource) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByName updates the specified licensed software by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatelicensedsoftwarebyname
		UpdateByName(ctx context.Context, name string, req *Resource) (*CreateUpdateResponse, *resty.Response, error)

		// DeleteByID removes the specified licensed software by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletelicensedsoftwarebyid
		DeleteByID(ctx context.Context, id int) (*resty.Response, error)

		// DeleteByName removes the specified licensed software by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletelicensedsoftwarebyname
		DeleteByName(ctx context.Context, name string) (*resty.Response, error)
	}

	// Service handles communication with the licensed-software-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/licensedsoftware
	LicensedSoftware struct {
		client transport.HTTPClient
	}
)

var _ ServiceInterface = (*LicensedSoftware)(nil)

// NewService returns a new licensed software Service backed by the provided HTTP client.
func NewLicensedSoftware(client transport.HTTPClient) *LicensedSoftware {
	return &LicensedSoftware{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Licensed Software CRUD Operations
// -----------------------------------------------------------------------------

// List returns all licensed software.
//
// URL: GET /JSSResource/licensedsoftware
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findlicensedsoftware
func (s *LicensedSoftware) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	endpoint := constants.EndpointClassicLicensedSoftware

	var out ListResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// GetByID returns the specified licensed software by ID.
//
// URL: GET /JSSResource/licensedsoftware/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findlicensedsoftwarebyid
func (s *LicensedSoftware) GetByID(ctx context.Context, id int) (*Resource, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("licensed software ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicLicensedSoftware, id)

	var out Resource

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// GetByName returns the specified licensed software by name.
//
// URL: GET /JSSResource/licensedsoftware/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findlicensedsoftwarebyname
func (s *LicensedSoftware) GetByName(ctx context.Context, name string) (*Resource, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("licensed software name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicLicensedSoftware, name)

	var out Resource

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Create creates a new licensed software item.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createlicensedsoftwarebyid
func (s *LicensedSoftware) Create(ctx context.Context, req *Resource) (*CreateUpdateResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("licensed software name is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicLicensedSoftware)

	var out CreateUpdateResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// UpdateByID updates the specified licensed software by ID.
//
// URL: PUT /JSSResource/licensedsoftware/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatelicensedsoftwarebyid
func (s *LicensedSoftware) UpdateByID(ctx context.Context, id int, req *Resource) (*CreateUpdateResponse, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("licensed software ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("licensed software name is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicLicensedSoftware, id)

	var out CreateUpdateResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// UpdateByName updates the specified licensed software by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatelicensedsoftwarebyname
func (s *LicensedSoftware) UpdateByName(ctx context.Context, name string, req *Resource) (*CreateUpdateResponse, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("licensed software name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("licensed software name is required in request")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicLicensedSoftware, name)

	var out CreateUpdateResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// DeleteByID removes the specified licensed software by ID.
//
// URL: DELETE /JSSResource/licensedsoftware/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletelicensedsoftwarebyid
func (s *LicensedSoftware) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("licensed software ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicLicensedSoftware, id)

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// DeleteByName removes the specified licensed software by name.
//
// URL: DELETE /JSSResource/licensedsoftware/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletelicensedsoftwarebyname
func (s *LicensedSoftware) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("licensed software name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicLicensedSoftware, name)

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
