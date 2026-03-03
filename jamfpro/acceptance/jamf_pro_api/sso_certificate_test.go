package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/sso_certificate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_SsoCertificate_get_v2(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.SsoCertificate
	ctx := context.Background()
	result, resp, err := svc.GetV2(ctx)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	if result != nil {
		_ = result.Keystore.Type
	}
}

func TestAcceptance_SsoCertificate_lifecycle(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.SsoCertificate
	ctx := context.Background()

	// Create a new certificate
	created, resp, err := svc.CreateV2(ctx)
	if err != nil {
		t.Skipf("Failed to create SSO certificate (may not be supported on this tenant): %v", err)
		return
	}
	require.NotNil(t, created)
	require.NotNil(t, resp)
	assert.Contains(t, []int{200, 201}, resp.StatusCode)
	assert.NotEmpty(t, created.Keystore.Key)

	// Cleanup: Delete the certificate at the end
	acc.Cleanup(t, func() {
		_, _ = svc.DeleteV2(ctx)
	})

	// Get the created certificate
	fetched, resp, err := svc.GetV2(ctx)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, fetched.Keystore.Key)

	// Download the certificate
	downloaded, resp, err := svc.DownloadV2(ctx)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	if len(downloaded) == 0 {
		t.Log("Warning: Downloaded certificate is empty")
	}

	// Delete the certificate
	resp, err = svc.DeleteV2(ctx)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestAcceptance_SsoCertificate_update_v2(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.SsoCertificate
	ctx := context.Background()

	// This test requires a valid keystore file, which we don't have in the test environment
	// Skip this test as it would require actual certificate data
	t.Skip("Skipping UpdateV2 test - requires valid keystore file data")

	updateReq := &sso_certificate.UpdateKeystoreRequest{
		KeystorePassword:  "test-password",
		KeystoreFile:      "base64-encoded-keystore-data",
		KeystoreFileName:  "test.p12",
		Type:              "PKCS12",
		KeystoreSetupType: "UPLOADED",
	}

	result, resp, err := svc.UpdateV2(ctx, updateReq)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestAcceptance_SsoCertificate_parse_v2(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.SsoCertificate
	ctx := context.Background()

	// This test requires a valid keystore file, which we don't have in the test environment
	// Skip this test as it would require actual certificate data
	t.Skip("Skipping ParseV2 test - requires valid keystore file data")

	parseReq := &sso_certificate.ParseKeystoreRequest{
		KeystoreFile:     "base64-encoded-keystore-data",
		KeystorePassword: "test-password",
		KeystoreFileName: "test.p12",
	}

	result, resp, err := svc.ParseV2(ctx, parseReq)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}
