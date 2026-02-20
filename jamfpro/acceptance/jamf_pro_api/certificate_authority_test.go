package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_CertificateAuthority_GetActiveCertificateAuthorityV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.CertificateAuthority
	ctx := context.Background()

	result, resp, err := svc.GetActiveCertificateAuthorityV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, result.SubjectX500Principal)
}
