package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHasNextPage(t *testing.T) {
	assert.False(t, HasNextPage(nil))
	assert.False(t, HasNextPage(&PaginationLinks{}))
	assert.False(t, HasNextPage(&PaginationLinks{Next: ""}))
	assert.True(t, HasNextPage(&PaginationLinks{Next: "https://example.com?page=2"}))
}

func TestExtractParamsFromURL(t *testing.T) {
	params, err := ExtractParamsFromURL("https://example.com/api?page=1&page-size=10&sort=name:asc")
	require.NoError(t, err)
	assert.Equal(t, "1", params["page"])
	assert.Equal(t, "10", params["page-size"])
	assert.Equal(t, "name:asc", params["sort"])
}

func TestExtractParamsFromURL_Invalid(t *testing.T) {
	_, err := ExtractParamsFromURL("://invalid")
	require.Error(t, err)
}

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

	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
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
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	mergePage := func([]byte) error { return assert.AnError }
	resp, err := tr.GetPaginated(ctx, "/api/v3/items", nil, nil, mergePage)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "merge page")
	assert.NotNil(t, resp)
}
