package mdm

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// MDMServiceInterface defines the interface for MDM command operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-mdm-commands
	MDMServiceInterface interface {
		// BlankPush sends an MDM blank push command to the specified devices.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-mdm-blank-push
		BlankPush(ctx context.Context, clientManagementIDs []string) (*BlankPushResponse, *interfaces.Response, error)

		// SendCommand sends an MDM command for creation and queuing.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-mdm-commands
		SendCommand(ctx context.Context, req *CommandRequest) (*CommandResponse, *interfaces.Response, error)

		// DeployPackage deploys a package using an MDM command.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-deploy-package
		DeployPackage(ctx context.Context, req *DeployPackageRequest) (*DeployPackageResponse, *interfaces.Response, error)

		// RenewProfile renews MDM profiles for the specified device UDIDs.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mdm-renew-profile
		RenewProfile(ctx context.Context, req *RenewProfileRequest) (*RenewProfileResponse, *interfaces.Response, error)
	}

	// Service handles communication with the MDM command methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-mdm-commands
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ MDMServiceInterface = (*Service)(nil)

// NewService returns a new MDM service.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - MDM Commands
// -----------------------------------------------------------------------------

// BlankPush sends an MDM blank push command to the specified devices.
// URL: POST /api/v2/mdm/blank-push
// https://developer.jamf.com/jamf-pro/reference/post_v2-mdm-blank-push
func (s *Service) BlankPush(ctx context.Context, clientManagementIDs []string) (*BlankPushResponse, *interfaces.Response, error) {
	if len(clientManagementIDs) == 0 {
		return nil, nil, fmt.Errorf("clientManagementIDs is required and must not be empty")
	}

	reqBody := map[string][]string{"clientManagementIds": clientManagementIDs}
	var result BlankPushResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, EndpointBlankPush, reqBody, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// SendCommand sends an MDM command for creation and queuing.
// URL: POST /api/v2/mdm/commands
// https://developer.jamf.com/jamf-pro/reference/post_v2-mdm-commands
func (s *Service) SendCommand(ctx context.Context, req *CommandRequest) (*CommandResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CommandResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, EndpointCommands, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeployPackage deploys a package using an MDM command.
// URL: POST /api/v1/deploy-package?verbose=true
// https://developer.jamf.com/jamf-pro/reference/post_v1-deploy-package
func (s *Service) DeployPackage(ctx context.Context, req *DeployPackageRequest) (*DeployPackageResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := EndpointDeployPackage + "?verbose=true"
	var result DeployPackageResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// RenewProfile renews MDM profiles for the specified device UDIDs.
// URL: POST /api/v1/mdm/renew-profile
// https://developer.jamf.com/jamf-pro/reference/post_v1-mdm-renew-profile
func (s *Service) RenewProfile(ctx context.Context, req *RenewProfileRequest) (*RenewProfileResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result RenewProfileResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, EndpointProfileRenewal, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
