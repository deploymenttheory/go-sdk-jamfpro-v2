package action_logs

import (
	"context"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/action_logs/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnit_ActionLogs_ReadV1(t *testing.T) {
	mock := mocks.NewActionLogsMock()
	mock.Register("GET", "/api/v1/action-logs", 200, "validate_list.json")
	mock.Register("GET", "/api/v1/action-logs/1", 200, "validate_get.json")
	mock.Register("GET", "/api/v1/action-logs/1/devices", 200, "validate_devices.json")
	svc := NewActionLogs(mock)
	list, resp, err := svc.ListV1(context.Background(), map[string]string{"filter": "status==COMPLETED"})
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode())
	require.Equal(t, 2, list.TotalCount)
	require.Equal(t, "status==COMPLETED", mock.LastRSQLQuery["filter"])
	detail, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode())
	require.Equal(t, 1, detail.StatusSummary.Completed)
	devices, resp, err := svc.GetDevicesV1(context.Background(), "1", nil)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode())
	require.Equal(t, 1, devices.TotalCount)
	require.Equal(t, "COMPLETED", devices.Results[0].Outcome)
}
func TestUnit_ActionLogs_ErrorsV1(t *testing.T) {
	mock := mocks.NewActionLogsMock()
	svc := NewActionLogs(mock)
	mock.RegisterError("GET", "/api/v1/action-logs", 400, "error_sort.json", "Invalid sort")
	list, resp, err := svc.ListV1(context.Background(), map[string]string{"sort": "invalid:asc"})
	require.Error(t, err)
	require.Nil(t, list)
	require.Equal(t, 400, resp.StatusCode())
	mock.RegisterError("GET", "/api/v1/action-logs/1", 404, "error_not_found.json", "Not found")
	detail, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.Error(t, err)
	require.Nil(t, detail)
	require.Equal(t, 404, resp.StatusCode())
	_, resp, err = svc.GetByIDV1(context.Background(), "")
	require.Error(t, err)
	require.Nil(t, resp)
	_, resp, err = svc.GetDevicesV1(context.Background(), "", nil)
	require.Error(t, err)
	require.Nil(t, resp)
	mock.RegisterError("GET", "/api/v1/action-logs/1/devices", 500, "", "Server error")
	_, resp, err = svc.GetDevicesV1(context.Background(), "1", nil)
	require.Error(t, err)
	require.Equal(t, 500, resp.StatusCode())
}
