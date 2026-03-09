package mobile_device_applications

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the mobile device applications-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/mobiledeviceapplications
	MobileDeviceApplications struct {
		client transport.HTTPClient
	}
)

// NewService returns a new mobile device applications Service backed by the provided HTTP client.
func NewMobileDeviceApplications(client transport.HTTPClient) *MobileDeviceApplications {
	return &MobileDeviceApplications{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Mobile Device Applications CRUD Operations
// -----------------------------------------------------------------------------

// List returns all mobile device applications.
//
// URL: GET /JSSResource/mobiledeviceapplications
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplications
func (s *MobileDeviceApplications) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	endpoint := constants.EndpointClassicMobileDeviceApplications

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

// GetByID returns the specified mobile device application by ID.
//
// URL: GET /JSSResource/mobiledeviceapplications/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplicationsbyid
func (s *MobileDeviceApplications) GetByID(ctx context.Context, id int) (*Resource, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device application ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMobileDeviceApplications, id)

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

// GetByName returns the specified mobile device application by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplicationsbyname
func (s *MobileDeviceApplications) GetByName(ctx context.Context, name string) (*Resource, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device application name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMobileDeviceApplications, name)

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

// GetByBundleID returns the specified mobile device application by bundle ID.
//
// URL: GET /JSSResource/mobiledeviceapplications/bundleid/{bundleid}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplicationsbybundleid
func (s *MobileDeviceApplications) GetByBundleID(ctx context.Context, bundleID string) (*Resource, *resty.Response, error) {
	if bundleID == "" {
		return nil, nil, fmt.Errorf("mobile device application bundle ID cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/bundleid/%s", constants.EndpointClassicMobileDeviceApplications, bundleID)

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

// GetByBundleIDAndVersion returns the specified mobile device application by bundle ID and version.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplicationsbybundleidandversion
func (s *MobileDeviceApplications) GetByBundleIDAndVersion(ctx context.Context, bundleID, version string) (*Resource, *resty.Response, error) {
	if bundleID == "" {
		return nil, nil, fmt.Errorf("mobile device application bundle ID cannot be empty")
	}
	if version == "" {
		return nil, nil, fmt.Errorf("mobile device application version cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/bundleid/%s/version/%s", constants.EndpointClassicMobileDeviceApplications, bundleID, version)

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

// GetByIDAndSubset returns a specific subset of a mobile device application by ID.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplicationsbyiddatasubset
func (s *MobileDeviceApplications) GetByIDAndSubset(ctx context.Context, id int, subset string) (*Resource, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device application ID must be a positive integer")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("mobile device application subset cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/id/%d/subset/%s", constants.EndpointClassicMobileDeviceApplications, id, subset)

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

// GetByNameAndSubset returns a specific subset of a mobile device application by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceapplicationsbynamedatasubset
func (s *MobileDeviceApplications) GetByNameAndSubset(ctx context.Context, name, subset string) (*Resource, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device application name cannot be empty")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("mobile device application subset cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", constants.EndpointClassicMobileDeviceApplications, name, subset)

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

// Create creates a new mobile device application.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledeviceapplicationbyid
func (s *MobileDeviceApplications) Create(ctx context.Context, req *Resource) (*CreateUpdateResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device application name is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicMobileDeviceApplications)

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

// UpdateByID updates the specified mobile device application by ID.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceapplicationbyid
func (s *MobileDeviceApplications) UpdateByID(ctx context.Context, id int, req *Resource) (*Resource, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device application ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device application name is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMobileDeviceApplications, id)

	var out Resource

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

// UpdateByName updates the specified mobile device application by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceapplicationbyname
func (s *MobileDeviceApplications) UpdateByName(ctx context.Context, name string, req *Resource) (*Resource, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device application name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device application name is required in request")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMobileDeviceApplications, name)

	var out Resource

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

// UpdateByBundleID updates the specified mobile device application by bundle ID.
//
// URL: PUT /JSSResource/mobiledeviceapplications/bundleid/{bundleid}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceapplicationbybundleid
func (s *MobileDeviceApplications) UpdateByBundleID(ctx context.Context, bundleID string, req *Resource) (*Resource, *resty.Response, error) {
	if bundleID == "" {
		return nil, nil, fmt.Errorf("mobile device application bundle ID cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device application name is required in request")
	}

	endpoint := fmt.Sprintf("%s/bundleid/%s", constants.EndpointClassicMobileDeviceApplications, bundleID)

	var out Resource

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

// UpdateByIDAndVersion updates the specified mobile device application by ID and version.
//
// URL: PUT /JSSResource/mobiledeviceapplications/id/{id}/version/{version}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceapplicationbyidandversion
func (s *MobileDeviceApplications) UpdateByIDAndVersion(ctx context.Context, id int, version string, req *Resource) (*Resource, *resty.Response, error) {
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

	endpoint := fmt.Sprintf("%s/id/%d/version/%s", constants.EndpointClassicMobileDeviceApplications, id, version)

	var out Resource

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

// DeleteByID removes the specified mobile device application by ID.
//
// URL: DELETE /JSSResource/mobiledeviceapplications/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceapplicationbyid
func (s *MobileDeviceApplications) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("mobile device application ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMobileDeviceApplications, id)

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

// DeleteByName removes the specified mobile device application by name.
//
// URL: DELETE /JSSResource/mobiledeviceapplications/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceapplicationbyname
func (s *MobileDeviceApplications) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("mobile device application name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMobileDeviceApplications, name)

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

// DeleteByBundleID removes the specified mobile device application by bundle ID.
//
// URL: DELETE /JSSResource/mobiledeviceapplications/bundleid/{bundleid}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceapplicationbybundleid
func (s *MobileDeviceApplications) DeleteByBundleID(ctx context.Context, bundleID string) (*resty.Response, error) {
	if bundleID == "" {
		return nil, fmt.Errorf("mobile device application bundle ID cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/bundleid/%s", constants.EndpointClassicMobileDeviceApplications, bundleID)

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

// DeleteByBundleIDAndVersion removes the specified mobile device application by bundle ID and version.
//
// URL: DELETE /JSSResource/mobiledeviceapplications/bundleid/{bundleid}/version/{version}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceapplicationbybundleidandversion
func (s *MobileDeviceApplications) DeleteByBundleIDAndVersion(ctx context.Context, bundleID, version string) (*resty.Response, error) {
	if bundleID == "" {
		return nil, fmt.Errorf("mobile device application bundle ID cannot be empty")
	}
	if version == "" {
		return nil, fmt.Errorf("mobile device application version cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/bundleid/%s/version/%s", constants.EndpointClassicMobileDeviceApplications, bundleID, version)

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
