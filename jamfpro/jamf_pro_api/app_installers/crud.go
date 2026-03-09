package app_installers

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the app installers-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: Undocumented
	AppInstallers struct {
		client transport.HTTPClient
	}
)

func NewAppInstallers(client transport.HTTPClient) *AppInstallers {
	return &AppInstallers{client: client}
}

// ListTitlesV1 returns all app installer titles.
// URL: GET /api/v1/app-installers/titles
// Jamf Pro API docs: Undocumented
func (s *AppInstallers) ListTitlesV1(ctx context.Context, rsqlQuery map[string]string) (*ListTitlesResponse, *resty.Response, error) {
	var result ListTitlesResponse

	endpoint := constants.EndpointJamfProAppInstallersTitlesV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetTitleByIDV1 returns the specified app installer title by ID.
// URL: GET /api/v1/app-installers/titles/{id}
// Jamf Pro API docs: Undocumented-id
func (s *AppInstallers) GetTitleByIDV1(ctx context.Context, id string) (*ResourceJamfAppCatalogAppInstaller, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("title ID is required")
	}
	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAppInstallersTitlesV1, id)
	var result ResourceJamfAppCatalogAppInstaller

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// ListDeploymentsV1 returns all app installer deployments.
// URL: GET /api/v1/app-installers/deployments
// Jamf Pro API docs: Undocumented
func (s *AppInstallers) ListDeploymentsV1(ctx context.Context, rsqlQuery map[string]string) (*ListDeploymentsResponse, *resty.Response, error) {
	var result ListDeploymentsResponse

	endpoint := constants.EndpointJamfProAppInstallersDeploymentsV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetDeploymentByIDV1 returns the specified deployment by ID.
// URL: GET /api/v1/app-installers/deployments/{id}
// Jamf Pro API docs: Undocumented-id
func (s *AppInstallers) GetDeploymentByIDV1(ctx context.Context, id string) (*ResourceJamfAppCatalogDeployment, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("deployment ID is required")
	}
	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAppInstallersDeploymentsV1, id)
	var result ResourceJamfAppCatalogDeployment

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateDeploymentV1 creates a new app installer deployment.
// URL: POST /api/v1/app-installers/deployments
// Jamf Pro API docs: Undocumented
func (s *AppInstallers) CreateDeploymentV1(ctx context.Context, request *RequestDeployment) (*CreateDeploymentResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateDeploymentResponse

	endpoint := constants.EndpointJamfProAppInstallersDeploymentsV1

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateDeploymentByIDV1 updates the specified deployment by ID.
// URL: PUT /api/v1/app-installers/deployments/{id}
// Jamf Pro API docs: Undocumented
func (s *AppInstallers) UpdateDeploymentByIDV1(ctx context.Context, id string, request *RequestDeployment) (*ResourceJamfAppCatalogDeployment, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAppInstallersDeploymentsV1, id)

	var result ResourceJamfAppCatalogDeployment

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteDeploymentByIDV1 removes the specified deployment by ID.
// URL: DELETE /api/v1/app-installers/deployments/{id}
// Jamf Pro API docs: Undocumented
func (s *AppInstallers) DeleteDeploymentByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("deployment ID is required")
	}
	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAppInstallersDeploymentsV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
