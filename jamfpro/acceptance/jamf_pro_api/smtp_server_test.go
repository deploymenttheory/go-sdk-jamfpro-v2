package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/smtp_server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_SMTPServer_get_v2(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.JamfProAPI.SmtpServer
	ctx := context.Background()

	result, resp, err := svc.GetV2(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestAcceptance_SmtpServer_GetAllowedAuthTypesV2(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.JamfProAPI.SmtpServer
	ctx := context.Background()

	result, resp, err := svc.GetAllowedAuthTypesV2(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())

	// Availability is instance-controlled, so the set can legitimately be any
	// subset of the documented enum -- but every entry must be a known value.
	known := []string{
		smtp_server.AuthenticationTypeNone,
		smtp_server.AuthenticationTypeBasic,
		smtp_server.AuthenticationTypeGraphApi,
		smtp_server.AuthenticationTypeGoogleMail,
	}
	for _, authType := range result.AllowedAuthenticationTypes {
		assert.Contains(t, known, authType, "unexpected SMTP authentication type returned by the server")
	}
}
