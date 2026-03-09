package sites

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/sites/mocks"
	"github.com/stretchr/testify/require"
)

func TestUnit_Sites_ListV1_Success(t *testing.T) {
	mock := mocks.NewSitesMock()
	mock.RegisterListV1Mock()
	svc := NewSites(mock)

	result, resp, err := svc.ListV1(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	require.Len(t, result, 3)
	require.Equal(t, "1", result[0].ID)
	require.Equal(t, "Default Site", result[0].Name)
	require.Equal(t, "2", result[1].ID)
	require.Equal(t, "Branch Office", result[1].Name)
}

func TestUnit_Sites_ListV1_NoMockRegistered(t *testing.T) {
	mock := mocks.NewSitesMock()
	svc := NewSites(mock)

	result, resp, err := svc.ListV1(context.Background())

	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 404, resp.StatusCode())
}

func TestUnit_Sites_GetObjectsByIDV1_Success(t *testing.T) {
	mock := mocks.NewSitesMock()
	mock.RegisterGetObjectsByIDV1Mock()
	svc := NewSites(mock)

	rsqlQuery := map[string]string{
		"page":      "0",
		"page-size": "100",
		"sort":      "objectType:asc",
	}

	result, resp, err := svc.GetObjectsByIDV1(context.Background(), "1", rsqlQuery)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	require.Equal(t, 3, result.TotalCount)
	require.Len(t, result.Results, 3)
	require.Equal(t, "1", result.Results[0].SiteID)
	require.Equal(t, "Computer", result.Results[0].ObjectType)
	require.Equal(t, "101", result.Results[0].ObjectID)
	require.Equal(t, rsqlQuery, mock.LastRSQLQuery)
}

func TestUnit_Sites_GetObjectsByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewSitesMock()
	svc := NewSites(mock)

	result, resp, err := svc.GetObjectsByIDV1(context.Background(), "", nil)

	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "id is required")
}

func TestUnit_Sites_GetObjectsByIDV1_NoMockRegistered(t *testing.T) {
	mock := mocks.NewSitesMock()
	svc := NewSites(mock)

	result, resp, err := svc.GetObjectsByIDV1(context.Background(), "999", nil)

	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 404, resp.StatusCode())
}
