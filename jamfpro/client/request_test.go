package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransport_Get_NilQuery_NilHeaders(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	var result map[string]string
	resp, err := tr.Get(ctx, "/api/x", nil, nil, &result)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_Get_EmptyQueryValuesSkipped(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	var result map[string]string
	resp, err := tr.Get(ctx, "/api/x", map[string]string{"a": "", "b": "val"}, nil, &result)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_Post_NilBody(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	var result map[string]string
	resp, err := tr.Post(ctx, "/api/x", nil, nil, &result)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_PostForm_NilFormData(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	var result map[string]string
	resp, err := tr.PostForm(ctx, "/api/x", nil, nil, &result)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_PostMultipart_NoFile(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	var result map[string]string
	resp, err := tr.PostMultipart(ctx, "/api/x", "", "", nil, 0, map[string]string{"k": "v"}, nil, nil, &result)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_Delete_NilrsqlQuery(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	resp, err := tr.Delete(ctx, "/api/x", nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_GetBytes_WithRsqlQuery(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	resp, body, err := tr.GetBytes(ctx, "/api/x", map[string]string{"filter": "name==\"test\""}, nil)
	require.NoError(t, err)
	require.NotNil(t, body)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_Put_NilBody(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	var result map[string]string
	resp, err := tr.Put(ctx, "/api/x", nil, nil, &result)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_Patch_NilBody(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	var result map[string]string
	resp, err := tr.Patch(ctx, "/api/x", nil, nil, &result)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_DeleteWithBody_NilBody(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	resp, err := tr.DeleteWithBody(ctx, "/api/x", nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
