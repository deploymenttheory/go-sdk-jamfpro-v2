package mdm

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// MDMServiceInterface defines the interface for MDM command operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-mdm-commands
	MDMServiceInterface interface {
		// ListCommandsV2 retrieves information about MDM commands made by Jamf Pro.
		//
		// Supports optional RSQL filtering, pagination and sorting via rsqlQuery
		// (keys: filter, sort, page, page-size).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mdm-commands
		ListCommandsV2(ctx context.Context, rsqlQuery map[string]string) (*ListCommandsResponse, *resty.Response, error)

		// BlankPush sends an MDM blank push command to the specified devices.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-mdm-blank-push
		BlankPush(ctx context.Context, clientManagementIDs []string) (*BlankPushResponse, *resty.Response, error)

		// SendCommand sends an MDM command for creation and queuing.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-mdm-commands
		SendCommand(ctx context.Context, req *CommandRequest) (*CommandResponse, *resty.Response, error)

		// DeployPackage deploys a package using an MDM command.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-deploy-package
		DeployPackage(ctx context.Context, req *DeployPackageRequest) (*DeployPackageResponse, *resty.Response, error)

		// RenewProfile renews MDM profiles for the specified device UDIDs.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mdm-renew-profile
		RenewProfile(ctx context.Context, req *RenewProfileRequest) (*RenewProfileResponse, *resty.Response, error)
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

// ListCommandsV2 retrieves information about MDM commands made by Jamf Pro.
// URL: GET /api/v2/mdm/commands
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v2-mdm-commands
func (s *Service) ListCommandsV2(ctx context.Context, rsqlQuery map[string]string) (*ListCommandsResponse, *resty.Response, error) {
	endpoint := EndpointCommands

	var result ListCommandsResponse

	mergePage := func(pageData []byte) error {
		var pageItems []CommandInfo
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list MDM commands: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// BlankPush sends an MDM blank push command to the specified devices.
// URL: POST /api/v2/mdm/blank-push
// https://developer.jamf.com/jamf-pro/reference/post_v2-mdm-blank-push
func (s *Service) BlankPush(ctx context.Context, clientManagementIDs []string) (*BlankPushResponse, *resty.Response, error) {
	if len(clientManagementIDs) == 0 {
		return nil, nil, fmt.Errorf("clientManagementIDs is required and must not be empty")
	}

	reqBody := map[string][]string{"clientManagementIds": clientManagementIDs}
	var result BlankPushResponse

	endpoint := EndpointBlankPush
	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, reqBody, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// SendCommand sends an MDM command for creation and queuing.
// URL: POST /api/v2/mdm/commands
// https://developer.jamf.com/jamf-pro/reference/post_v2-mdm-commands
func (s *Service) SendCommand(ctx context.Context, req *CommandRequest) (*CommandResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CommandResponse

	endpoint := EndpointCommands
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

// DeployPackage deploys a package using an MDM command.
// URL: POST /api/v1/deploy-package?verbose=true
// https://developer.jamf.com/jamf-pro/reference/post_v1-deploy-package
func (s *Service) DeployPackage(ctx context.Context, req *DeployPackageRequest) (*DeployPackageResponse, *resty.Response, error) {
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
func (s *Service) RenewProfile(ctx context.Context, req *RenewProfileRequest) (*RenewProfileResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result RenewProfileResponse

	endpoint := EndpointProfileRenewal
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
