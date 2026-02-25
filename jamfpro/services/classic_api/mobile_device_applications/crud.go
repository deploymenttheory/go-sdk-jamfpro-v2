package mobile_device_applications

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ServiceInterface defines the interface for Classic API mobile device application operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/mobiledeviceapplications
	ServiceInterface interface {
		// List returns all mobile device applications.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplications
		List(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetByID returns the specified mobile device application by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplicationsbyid
		GetByID(ctx context.Context, id int) (*Resource, *interfaces.Response, error)

		// GetByName returns the specified mobile device application by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplicationsbyname
		GetByName(ctx context.Context, name string) (*Resource, *interfaces.Response, error)

		// GetByBundleID returns the specified mobile device application by bundle ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplicationsbybundleid
		GetByBundleID(ctx context.Context, bundleID string) (*Resource, *interfaces.Response, error)

		// GetByBundleIDAndVersion returns the specified mobile device application by bundle ID and version.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplicationsbybundleidandversion
		GetByBundleIDAndVersion(ctx context.Context, bundleID, version string) (*Resource, *interfaces.Response, error)

		// GetByIDAndSubset returns a specific subset of a mobile device application by ID.
		// Subset values: General, Scope, SelfService, VPP, AppConfiguration.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplicationsbyiddatasubset
		GetByIDAndSubset(ctx context.Context, id int, subset string) (*Resource, *interfaces.Response, error)

		// GetByNameAndSubset returns a specific subset of a mobile device application by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplicationsbynamedatasubset
		GetByNameAndSubset(ctx context.Context, name, subset string) (*Resource, *interfaces.Response, error)

		// Create creates a new mobile device application.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledeviceapplicationbyid
		Create(ctx context.Context, req *Resource) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateByID updates the specified mobile device application by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceapplicationbyid
		UpdateByID(ctx context.Context, id int, req *Resource) (*Resource, *interfaces.Response, error)

		// UpdateByName updates the specified mobile device application by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceapplicationbyname
		UpdateByName(ctx context.Context, name string, req *Resource) (*Resource, *interfaces.Response, error)

		// UpdateByBundleID updates the specified mobile device application by bundle ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceapplicationbybundleid
		UpdateByBundleID(ctx context.Context, bundleID string, req *Resource) (*Resource, *interfaces.Response, error)

		// UpdateByIDAndVersion updates the specified mobile device application by ID and version.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceapplicationbyidandversion
		UpdateByIDAndVersion(ctx context.Context, id int, version string, req *Resource) (*Resource, *interfaces.Response, error)

		// DeleteByID removes the specified mobile device application by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceapplicationbyid
		DeleteByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteByName removes the specified mobile device application by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceapplicationbyname
		DeleteByName(ctx context.Context, name string) (*interfaces.Response, error)

		// DeleteByBundleID removes the specified mobile device application by bundle ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceapplicationbybundleid
		DeleteByBundleID(ctx context.Context, bundleID string) (*interfaces.Response, error)

		// DeleteByBundleIDAndVersion removes the specified mobile device application by bundle ID and version.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceapplicationbybundleidandversion
		DeleteByBundleIDAndVersion(ctx context.Context, bundleID, version string) (*interfaces.Response, error)
	}

	// Service handles communication with the mobile device applications-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/mobiledeviceapplications
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ServiceInterface = (*Service)(nil)

// NewService returns a new mobile device applications Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Mobile Device Applications CRUD Operations
// -----------------------------------------------------------------------------

// List returns all mobile device applications.
//
// URL: GET /JSSResource/mobiledeviceapplications
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplications
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	endpoint := EndpointMobileDeviceApplications

	var out ListResponse

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

// GetByID returns the specified mobile device application by ID.
//
// URL: GET /JSSResource/mobiledeviceapplications/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplicationsbyid
func (s *Service) GetByID(ctx context.Context, id int) (*Resource, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device application ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointMobileDeviceApplications, id)

	var out Resource

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

// GetByName returns the specified mobile device application by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplicationsbyname
func (s *Service) GetByName(ctx context.Context, name string) (*Resource, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device application name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointMobileDeviceApplications, name)

	var out Resource

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

// GetByBundleID returns the specified mobile device application by bundle ID.
//
// URL: GET /JSSResource/mobiledeviceapplications/bundleid/{bundleid}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplicationsbybundleid
func (s *Service) GetByBundleID(ctx context.Context, bundleID string) (*Resource, *interfaces.Response, error) {
	if bundleID == "" {
		return nil, nil, fmt.Errorf("mobile device application bundle ID cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/bundleid/%s", EndpointMobileDeviceApplications, bundleID)

	var out Resource

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

// GetByBundleIDAndVersion returns the specified mobile device application by bundle ID and version.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplicationsbybundleidandversion
func (s *Service) GetByBundleIDAndVersion(ctx context.Context, bundleID, version string) (*Resource, *interfaces.Response, error) {
	if bundleID == "" {
		return nil, nil, fmt.Errorf("mobile device application bundle ID cannot be empty")
	}
	if version == "" {
		return nil, nil, fmt.Errorf("mobile device application version cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/bundleid/%s/version/%s", EndpointMobileDeviceApplications, bundleID, version)

	var out Resource

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

// GetByIDAndSubset returns a specific subset of a mobile device application by ID.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplicationsbyiddatasubset
func (s *Service) GetByIDAndSubset(ctx context.Context, id int, subset string) (*Resource, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device application ID must be a positive integer")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("mobile device application subset cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/id/%d/subset/%s", EndpointMobileDeviceApplications, id, subset)

	var out Resource

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

// GetByNameAndSubset returns a specific subset of a mobile device application by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplicationsbynamedatasubset
func (s *Service) GetByNameAndSubset(ctx context.Context, name, subset string) (*Resource, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device application name cannot be empty")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("mobile device application subset cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", EndpointMobileDeviceApplications, name, subset)

	var out Resource

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

// Create creates a new mobile device application.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledeviceapplicationbyid
func (s *Service) Create(ctx context.Context, req *Resource) (*CreateUpdateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device application name is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointMobileDeviceApplications)

	var out CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// UpdateByID updates the specified mobile device application by ID.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceapplicationbyid
func (s *Service) UpdateByID(ctx context.Context, id int, req *Resource) (*Resource, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device application ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device application name is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointMobileDeviceApplications, id)

	var out Resource

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// UpdateByName updates the specified mobile device application by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceapplicationbyname
func (s *Service) UpdateByName(ctx context.Context, name string, req *Resource) (*Resource, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device application name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device application name is required in request")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointMobileDeviceApplications, name)

	var out Resource

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// UpdateByBundleID updates the specified mobile device application by bundle ID.
//
// URL: PUT /JSSResource/mobiledeviceapplications/bundleid/{bundleid}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceapplicationbybundleid
func (s *Service) UpdateByBundleID(ctx context.Context, bundleID string, req *Resource) (*Resource, *interfaces.Response, error) {
	if bundleID == "" {
		return nil, nil, fmt.Errorf("mobile device application bundle ID cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device application name is required in request")
	}

	endpoint := fmt.Sprintf("%s/bundleid/%s", EndpointMobileDeviceApplications, bundleID)

	var out Resource

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// UpdateByIDAndVersion updates the specified mobile device application by ID and version.
//
// URL: PUT /JSSResource/mobiledeviceapplications/id/{id}/version/{version}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceapplicationbyidandversion
func (s *Service) UpdateByIDAndVersion(ctx context.Context, id int, version string, req *Resource) (*Resource, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device application ID must be a positive integer")
	}
	if version == "" {
		return nil, nil, fmt.Errorf("mobile device application version cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device application name is required in request")
	}

	endpoint := fmt.Sprintf("%s/id/%d/version/%s", EndpointMobileDeviceApplications, id, version)

	var out Resource

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// DeleteByID removes the specified mobile device application by ID.
//
// URL: DELETE /JSSResource/mobiledeviceapplications/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceapplicationbyid
func (s *Service) DeleteByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("mobile device application ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointMobileDeviceApplications, id)

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

// DeleteByName removes the specified mobile device application by name.
//
// URL: DELETE /JSSResource/mobiledeviceapplications/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceapplicationbyname
func (s *Service) DeleteByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("mobile device application name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointMobileDeviceApplications, name)

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

// DeleteByBundleID removes the specified mobile device application by bundle ID.
//
// URL: DELETE /JSSResource/mobiledeviceapplications/bundleid/{bundleid}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceapplicationbybundleid
func (s *Service) DeleteByBundleID(ctx context.Context, bundleID string) (*interfaces.Response, error) {
	if bundleID == "" {
		return nil, fmt.Errorf("mobile device application bundle ID cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/bundleid/%s", EndpointMobileDeviceApplications, bundleID)

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

// DeleteByBundleIDAndVersion removes the specified mobile device application by bundle ID and version.
//
// URL: DELETE /JSSResource/mobiledeviceapplications/bundleid/{bundleid}/version/{version}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceapplicationbybundleidandversion
func (s *Service) DeleteByBundleIDAndVersion(ctx context.Context, bundleID, version string) (*interfaces.Response, error) {
	if bundleID == "" {
		return nil, fmt.Errorf("mobile device application bundle ID cannot be empty")
	}
	if version == "" {
		return nil, fmt.Errorf("mobile device application version cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/bundleid/%s/version/%s", EndpointMobileDeviceApplications, bundleID, version)

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
