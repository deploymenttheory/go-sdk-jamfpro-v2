package sites

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// SitesServiceInterface defines the interface for sites operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-sites
	SitesServiceInterface interface {
		// ListV1 returns all sites.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-sites
		ListV1(ctx context.Context) ([]ResourceSite, *resty.Response, error)

		// GetObjectsByIDV1 returns paginated objects for a site.
		//
		// Query params: page, page-size, sort, filter (RSQL)
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-sites-id-objects
		GetObjectsByIDV1(ctx context.Context, id string, rsqlQuery map[string]string) (*ObjectsListResponse, *resty.Response, error)
	}

	// Service handles communication with the sites-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-sites
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ SitesServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// ListV1 returns all sites.
// URL: GET /api/v1/sites
func (s *Service) ListV1(ctx context.Context) ([]ResourceSite, *resty.Response, error) {
	endpoint := EndpointSitesV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
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
func (s *Service) GetObjectsByIDV1(ctx context.Context, id string, rsqlQuery map[string]string) (*ObjectsListResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/objects", EndpointSitesV1, id)

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
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get site objects: %w", err)
	}

	result.TotalCount = len(result.Results)
	return &result, resp, nil
}
