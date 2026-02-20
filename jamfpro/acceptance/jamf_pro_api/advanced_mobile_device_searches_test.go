package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/advanced_mobile_device_searches"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_AdvancedMobileDeviceSearches_ListV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.AdvancedMobileDeviceSearches
	ctx := context.Background()

	result, resp, err := svc.ListV1(ctx, map[string]string{"page": "0", "page-size": "100"})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
}

func TestAcceptance_AdvancedMobileDeviceSearches_Lifecycle(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.AdvancedMobileDeviceSearches
	ctx := context.Background()

	name := fmt.Sprintf("acc-adv-md-search-%d", time.Now().UnixMilli())
	// Use valid criterion for mobile device searches
	falseBool := false
	siteId := "-1"
	search := &advanced_mobile_device_searches.ResourceAdvancedMobileDeviceSearch{
		Name:          name,
		Criteria:      []advanced_mobile_device_searches.CriteriaJamfProAPI{{Name: "Last Inventory Update", Priority: 0, AndOr: "and", SearchType: "more than x days ago", Value: "7", OpeningParen: &falseBool, ClosingParen: &falseBool}},
		DisplayFields: []string{"IP Address"},
		SiteId:        &siteId,
	}
	created, createResp, err := svc.CreateV1(ctx, search)
	require.NoError(t, err)
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	searchID := created.ID
	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = svc.DeleteByIDV1(cleanupCtx, searchID)
	})

	fetched, _, err := svc.GetByIDV1(ctx, searchID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, name, fetched.Name)

	delResp, err := svc.DeleteByIDV1(ctx, searchID)
	require.NoError(t, err)
	require.NotNil(t, delResp)
	assert.Equal(t, 204, delResp.StatusCode)
}
