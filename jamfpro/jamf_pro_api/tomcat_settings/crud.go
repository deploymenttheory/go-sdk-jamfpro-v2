package tomcat_settings

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Tomcat settings methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_settings-issue-tomcat-ssl-certificate
	TomcatSettings struct {
		client transport.HTTPClient
	}
)

func NewTomcatSettings(client transport.HTTPClient) *TomcatSettings {
	return &TomcatSettings{client: client}
}

// IssueTomcatSslCertificate generates an SSL certificate via Jamf CA.
// URL: POST /api/settings/issueTomcatSslCertificate
func (s *TomcatSettings) IssueTomcatSslCertificate(ctx context.Context) (*resty.Response, error) {
	endpoint := constants.EndpointJamfProIssueTomcatSslCertificate

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
