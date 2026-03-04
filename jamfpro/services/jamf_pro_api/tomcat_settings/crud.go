package tomcat_settings

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// TomcatSettingsServiceInterface defines the interface for Tomcat settings operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_settings-issue-tomcat-ssl-certificate
	TomcatSettingsServiceInterface interface {
		// IssueTomcatSslCertificate generates an SSL certificate via Jamf CA.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_settings-issue-tomcat-ssl-certificate
		IssueTomcatSslCertificate(ctx context.Context) (*resty.Response, error)
	}

	// Service handles communication with the Tomcat settings methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_settings-issue-tomcat-ssl-certificate
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ TomcatSettingsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// IssueTomcatSslCertificate generates an SSL certificate via Jamf CA.
// URL: POST /api/settings/issueTomcatSslCertificate
func (s *Service) IssueTomcatSslCertificate(ctx context.Context) (*resty.Response, error) {
	endpoint := EndpointIssueTomcatSslCertificate

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
