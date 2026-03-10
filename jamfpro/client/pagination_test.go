package client

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/config"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransport_GetPaginated(t *testing.T) {
	pageCount := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/oauth/token" && r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"access_token":"t","expires_in":3600}`))
			return
		}
		if r.URL.Path == "/api/v3/items" {
			pageCount++
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			// Return one page only (totalCount 1, one result) so loop exits after first page
			_ = json.NewEncoder(w).Encode(map[string]any{
				"totalCount": 1,
				"results":    []any{map[string]string{"id": "1"}},
			})
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	var merged []json.RawMessage
	mergePage := func(pageData []byte) error {
		var arr []json.RawMessage
		if err := json.Unmarshal(pageData, &arr); err != nil {
			return err
		}
		merged = append(merged, arr...)
		return nil
	}

	resp, err := tr.GetPaginated(ctx, "/api/v3/items", nil, nil, mergePage)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Len(t, merged, 1)
}

func TestTransport_GetPaginated_MergeError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/oauth/token" && r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"access_token":"t","expires_in":3600}`))
			return
		}
		if r.URL.Path == "/api/v3/items" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"totalCount": 1,
				"results":    []any{map[string]string{"id": "1"}},
			})
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	mergePage := func([]byte) error { return assert.AnError }
	resp, err := tr.GetPaginated(ctx, "/api/v3/items", nil, nil, mergePage)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "merge page")
	assert.NotNil(t, resp)
}
