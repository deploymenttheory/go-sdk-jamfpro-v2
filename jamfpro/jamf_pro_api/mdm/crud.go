package mdm

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the MDM command methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-mdm-commands
	Mdm struct {
		client client.Client
	}
)

// NewService returns a new MDM service.
func NewMdm(client client.Client) *Mdm {
	return &Mdm{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - MDM Commands
// -----------------------------------------------------------------------------

// ListCommandsV2 retrieves information about MDM commands made by Jamf Pro.
// URL: GET /api/v2/mdm/commands
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v2-mdm-commands
func (s *Mdm) ListCommandsV2(ctx context.Context, rsqlQuery map[string]string) (*ListCommandsResponse, *resty.Response, error) {
	endpoint := constants.EndpointJamfProCommands

	var result ListCommandsResponse

	mergePage := func(pageData []byte) error {
		var pageItems []CommandInfo
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list MDM commands: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// BlankPush sends an MDM blank push command to the specified devices.
// URL: POST /api/v2/mdm/blank-push
// https://developer.jamf.com/jamf-pro/reference/post_v2-mdm-blank-push
func (s *Mdm) BlankPush(ctx context.Context, clientManagementIDs []string) (*BlankPushResponse, *resty.Response, error) {
	if len(clientManagementIDs) == 0 {
		return nil, nil, fmt.Errorf("clientManagementIDs is required and must not be empty")
	}

	reqBody := map[string][]string{"clientManagementIds": clientManagementIDs}
	var result BlankPushResponse

	endpoint := constants.EndpointJamfProBlankPush

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(reqBody).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// SendCommand sends an MDM command for creation and queuing.
// URL: POST /api/v2/mdm/commands
// https://developer.jamf.com/jamf-pro/reference/post_v2-mdm-commands
func (s *Mdm) SendCommand(ctx context.Context, req *CommandRequest) (*CommandResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CommandResponse

	endpoint := constants.EndpointJamfProCommands

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(req).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeployPackage deploys a package using an MDM command.
// URL: POST /api/v1/deploy-package?verbose=true
// https://developer.jamf.com/jamf-pro/reference/post_v1-deploy-package
func (s *Mdm) DeployPackage(ctx context.Context, req *DeployPackageRequest) (*DeployPackageResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := constants.EndpointJamfProDeployPackage + "?verbose=true"
	var result DeployPackageResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(req).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// RenewProfile renews MDM profiles for the specified device UDIDs.
// URL: POST /api/v1/mdm/renew-profile
// https://developer.jamf.com/jamf-pro/reference/post_v1-mdm-renew-profile
func (s *Mdm) RenewProfile(ctx context.Context, req *RenewProfileRequest) (*RenewProfileResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result RenewProfileResponse

	endpoint := constants.EndpointJamfProProfileRenewal

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(req).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
