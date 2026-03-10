package sites

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the sites-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-sites
	Sites struct {
		client client.Client
	}
)

func NewSites(client client.Client) *Sites {
	return &Sites{client: client}
}

// ListV1 returns all sites.
// URL: GET /api/v1/sites
func (s *Sites) ListV1(ctx context.Context) ([]ResourceSite, *resty.Response, error) {
	endpoint := constants.EndpointJamfProSitesV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	var result []ResourceSite
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list sites: %w", err)
	}

	return result, resp, nil
}

// GetObjectsByIDV1 returns paginated objects for a site.
// URL: GET /api/v1/sites/{id}/objects
func (s *Sites) GetObjectsByIDV1(ctx context.Context, id string, rsqlQuery map[string]string) (*ObjectsListResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/objects", constants.EndpointJamfProSitesV1, id)

	var result ObjectsListResponse

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceSiteObject
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get site objects: %w", err)
	}

	result.TotalCount = len(result.Results)
	return &result, resp, nil
}
