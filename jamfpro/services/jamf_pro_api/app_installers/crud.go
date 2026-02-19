package app_installers

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

type (
	// AppInstallersServiceInterface defines the interface for app installer operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-app-installers-titles
	AppInstallersServiceInterface interface {
		// ListTitlesV1 returns all app installer titles (Get App Installer Title objects).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-app-installers-titles
		ListTitlesV1(ctx context.Context, queryParams map[string]string) (*ListTitlesResponse, *interfaces.Response, error)

		// GetTitleByIDV1 returns the specified app installer title by ID (Get specified App Installer Title object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-app-installers-titles-id
		GetTitleByIDV1(ctx context.Context, id string) (*ResourceJamfAppCatalogAppInstaller, *interfaces.Response, error)

		// ListDeploymentsV1 returns all app installer deployments (Get App Installer Deployment objects).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-app-installers-deployments
		ListDeploymentsV1(ctx context.Context, queryParams map[string]string) (*ListDeploymentsResponse, *interfaces.Response, error)

		// GetDeploymentByIDV1 returns the specified deployment by ID (Get specified App Installer Deployment object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-app-installers-deployments-id
		GetDeploymentByIDV1(ctx context.Context, id string) (*ResourceJamfAppCatalogDeployment, *interfaces.Response, error)

		// CreateDeploymentV1 creates a new app installer deployment (Create App Installer Deployment record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-app-installers-deployments
		CreateDeploymentV1(ctx context.Context, req *RequestDeployment) (*CreateDeploymentResponse, *interfaces.Response, error)

		// UpdateDeploymentByIDV1 updates the specified deployment by ID (Update specified App Installer Deployment object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-app-installers-deployments-id
		UpdateDeploymentByIDV1(ctx context.Context, id string, req *RequestDeployment) (*ResourceJamfAppCatalogDeployment, *interfaces.Response, error)

		// DeleteDeploymentByIDV1 removes the specified deployment by ID (Remove specified App Installer Deployment record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-app-installers-deployments-id
		DeleteDeploymentByIDV1(ctx context.Context, id string) (*interfaces.Response, error)
	}

	// Service handles communication with the app installers-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-app-installers-titles
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ AppInstallersServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// ListTitlesV1 returns all app installer titles.
// URL: GET /api/v1/app-installers/titles
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-app-installers-titles
func (s *Service) ListTitlesV1(ctx context.Context, queryParams map[string]string) (*ListTitlesResponse, *interfaces.Response, error) {
	var result ListTitlesResponse
	resp, err := s.client.Get(ctx, EndpointTitlesV1, queryParams, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetTitleByIDV1 returns the specified app installer title by ID.
// URL: GET /api/v1/app-installers/titles/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-app-installers-titles-id
func (s *Service) GetTitleByIDV1(ctx context.Context, id string) (*ResourceJamfAppCatalogAppInstaller, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("title ID is required")
	}
	endpoint := fmt.Sprintf("%s/%s", EndpointTitlesV1, id)
	var result ResourceJamfAppCatalogAppInstaller
	resp, err := s.client.Get(ctx, endpoint, nil, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// ListDeploymentsV1 returns all app installer deployments.
// URL: GET /api/v1/app-installers/deployments
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-app-installers-deployments
func (s *Service) ListDeploymentsV1(ctx context.Context, queryParams map[string]string) (*ListDeploymentsResponse, *interfaces.Response, error) {
	var result ListDeploymentsResponse
	resp, err := s.client.Get(ctx, EndpointDeploymentsV1, queryParams, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetDeploymentByIDV1 returns the specified deployment by ID.
// URL: GET /api/v1/app-installers/deployments/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-app-installers-deployments-id
func (s *Service) GetDeploymentByIDV1(ctx context.Context, id string) (*ResourceJamfAppCatalogDeployment, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("deployment ID is required")
	}
	endpoint := fmt.Sprintf("%s/%s", EndpointDeploymentsV1, id)
	var result ResourceJamfAppCatalogDeployment
	resp, err := s.client.Get(ctx, endpoint, nil, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// CreateDeploymentV1 creates a new app installer deployment.
// URL: POST /api/v1/app-installers/deployments
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-app-installers-deployments
func (s *Service) CreateDeploymentV1(ctx context.Context, req *RequestDeployment) (*CreateDeploymentResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	var result CreateDeploymentResponse
	resp, err := s.client.Post(ctx, EndpointDeploymentsV1, req, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UpdateDeploymentByIDV1 updates the specified deployment by ID.
// URL: PUT /api/v1/app-installers/deployments/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-app-installers-deployments-id
func (s *Service) UpdateDeploymentByIDV1(ctx context.Context, id string, req *RequestDeployment) (*ResourceJamfAppCatalogDeployment, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	endpoint := fmt.Sprintf("%s/%s", EndpointDeploymentsV1, id)
	var result ResourceJamfAppCatalogDeployment
	resp, err := s.client.Put(ctx, endpoint, req, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// DeleteDeploymentByIDV1 removes the specified deployment by ID.
// URL: DELETE /api/v1/app-installers/deployments/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-app-installers-deployments-id
func (s *Service) DeleteDeploymentByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("deployment ID is required")
	}
	endpoint := fmt.Sprintf("%s/%s", EndpointDeploymentsV1, id)
	resp, err := s.client.Delete(ctx, endpoint, nil, shared.JSONHeaders(), nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
