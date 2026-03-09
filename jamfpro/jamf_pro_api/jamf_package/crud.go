package jamf_package

import (
	"context"
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// JamfPackageServiceInterface defines the interface for Jamf package operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-package
	JamfPackageServiceInterface interface {
		// ListV1 returns an array of packages for the given application (protect or connect).
		//
		// GET /api/v1/jamf-package?application={protect|connect}
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-package
		ListV1(ctx context.Context, application string) (ListV1Response, *resty.Response, error)

		// GetV2 returns the package object for the given application (protect or connect).
		//
		// GET /api/v2/jamf-package?application={protect|connect}
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-jamf-package
		GetV2(ctx context.Context, application string) (*ResourceJamfPackageV2, *resty.Response, error)
	}

	// Service handles communication with the Jamf package-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-package
	JamfPackage struct {
		client interfaces.HTTPClient
	}
)

var _ JamfPackageServiceInterface = (*JamfPackage)(nil)

func NewJamfPackage(client interfaces.HTTPClient) *JamfPackage {
	return &JamfPackage{client: client}
}

// validateApplication validates that application is "protect" or "connect".
func validateApplication(application string) error {
	app := strings.ToLower(strings.TrimSpace(application))
	if app != ApplicationProtect && app != ApplicationConnect {
		return fmt.Errorf("application must be %q or %q, got %q", ApplicationProtect, ApplicationConnect, application)
	}
	return nil
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Jamf Package Operations
// -----------------------------------------------------------------------------

// ListV1 returns an array of packages for the given application (protect or connect).
// URL: GET /api/v1/jamf-package?application={protect|connect}
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-package
func (s *JamfPackage) ListV1(ctx context.Context, application string) (ListV1Response, *resty.Response, error) {
	if err := validateApplication(application); err != nil {
		return nil, nil, fmt.Errorf("list jamf packages: %w", err)
	}

	var result ListV1Response

	rsqlQuery := map[string]string{
		"application": strings.ToLower(strings.TrimSpace(application)),
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, EndpointJamfPackageV1, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list jamf packages: %w", err)
	}

	return result, resp, nil
}

// GetV2 returns the package object for the given application (protect or connect).
// URL: GET /api/v2/jamf-package?application={protect|connect}
// https://developer.jamf.com/jamf-pro/reference/get_v2-jamf-package
func (s *JamfPackage) GetV2(ctx context.Context, application string) (*ResourceJamfPackageV2, *resty.Response, error) {
	if err := validateApplication(application); err != nil {
		return nil, nil, fmt.Errorf("get jamf package: %w", err)
	}

	var result ResourceJamfPackageV2

	rsqlQuery := map[string]string{
		"application": strings.ToLower(strings.TrimSpace(application)),
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, EndpointJamfPackageV2, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get jamf package: %w", err)
	}

	return &result, resp, nil
}
