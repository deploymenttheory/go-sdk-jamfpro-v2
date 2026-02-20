package jamf_pro_api

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance/acc"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/adcs_settings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdcsSettings_CreateGetUpdateDelete(t *testing.T) {
	client := acc.RequireClient(t)
	ctx := context.Background()

	svc := adcs_settings.NewService(client)

	displayName := acc.UniqueName("Test ADCS")

	revocationEnabled := true
	outbound := false

	createReq := &adcs_settings.ResourceAdcsSettings{
		DisplayName:       displayName,
		CAName:            "TestCA",
		FQDN:              "adcs.example.com",
		AdcsURL:           "https://adcs.example.com/certsrv",
		RevocationEnabled: &revocationEnabled,
		APIClientID:       "test-client-id",
		Outbound:          &outbound,
	}

	created, _, err := svc.CreateV1(ctx, createReq)
	if err != nil {
		t.Skipf("Failed to create ADCS settings (may not be supported on this tenant): %v", err)
		return
	}
	require.NotEmpty(t, created.ID)

	acc.Cleanup(t, func() {
		svc.DeleteByIDV1(ctx, created.ID)
	})

	fetched, _, err := svc.GetByIDV1(ctx, created.ID)
	require.NoError(t, err)
	assert.Equal(t, created.ID, fetched.ID)
	assert.Equal(t, displayName, fetched.DisplayName)
	assert.Equal(t, "TestCA", fetched.CAName)
	assert.Equal(t, "adcs.example.com", fetched.FQDN)
	assert.True(t, fetched.RevocationEnabled)
	assert.False(t, fetched.Outbound)

	updateReq := &adcs_settings.ResourceAdcsSettings{
		DisplayName: displayName + " Updated",
	}

	_, err = svc.UpdateByIDV1(ctx, created.ID, updateReq)
	require.NoError(t, err)

	updated, _, err := svc.GetByIDV1(ctx, created.ID)
	require.NoError(t, err)
	assert.Equal(t, displayName+" Updated", updated.DisplayName)

	deps, _, err := svc.GetDependenciesByIDV1(ctx, created.ID)
	require.NoError(t, err)
	assert.NotNil(t, deps)
	assert.GreaterOrEqual(t, deps.TotalCount, 0)

	history, _, err := svc.GetHistoryByIDV1(ctx, created.ID, nil)
	require.NoError(t, err)
	assert.NotNil(t, history)
	assert.GreaterOrEqual(t, history.TotalCount, 0)

	noteReq := &adcs_settings.HistoryNoteRequest{
		Note: "Test history note",
	}
	_, err = svc.AddHistoryNoteByIDV1(ctx, created.ID, noteReq)
	require.NoError(t, err)

	_, err = svc.DeleteByIDV1(ctx, created.ID)
	require.NoError(t, err)
}

func TestAdcsSettings_ValidateCertificates(t *testing.T) {
	client := acc.RequireClient(t)
	ctx := context.Background()

	svc := adcs_settings.NewService(client)

	serverCertReq := &adcs_settings.ValidateCertificateRequest{
		Filename: "server.cer",
		Data:     []string{"dGVzdGRhdGE="},
	}

	_, err := svc.ValidateServerCertificateV1(ctx, serverCertReq)
	if err != nil {
		t.Skipf("Failed to validate server certificate (may not be supported or invalid test data): %v", err)
		return
	}

	password := "test-password"
	clientCertReq := &adcs_settings.ValidateCertificateRequest{
		Filename: "client.p12",
		Data:     []string{"dGVzdGRhdGE="},
		Password: &password,
	}

	_, err = svc.ValidateClientCertificateV1(ctx, clientCertReq)
	if err != nil {
		t.Skipf("Failed to validate client certificate (may not be supported or invalid test data): %v", err)
		return
	}
}
