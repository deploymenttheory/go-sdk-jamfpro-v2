package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/bookmarks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_Bookmarks_List(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Bookmarks
	ctx := context.Background()

	result, resp, err := svc.ListV1(ctx, map[string]string{"page": "0", "page-size": "100"})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
}

func TestAcceptance_Bookmarks_Lifecycle(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Bookmarks
	ctx := context.Background()

	name := fmt.Sprintf("acc-bookmark-%d", time.Now().UnixMilli())
	displayInBrowser := true
	bm := &bookmarks.ResourceBookmark{
		Name:             name,
		URL:              "https://example.com",
		SiteID:           "-1",
		IconID:           "0",
		Priority:         1,
		DisplayInBrowser: &displayInBrowser,
	}
	created, createResp, err := svc.CreateV1(ctx, bm)
	require.NoError(t, err)
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	id := created.ID
	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = svc.DeleteByIDV1(cleanupCtx, id)
	})

	fetched, _, err := svc.GetByIDV1(ctx, id)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, name, fetched.Name)

	delResp, err := svc.DeleteByIDV1(ctx, id)
	require.NoError(t, err)
	require.NotNil(t, delResp)
	assert.Equal(t, 204, delResp.StatusCode)
}
