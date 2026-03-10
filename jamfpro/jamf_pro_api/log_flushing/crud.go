package log_flushing

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the log flushing-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-log-flushing
	LogFlushing struct {
		client client.Client
	}
)

// NewService creates a new log flushing service.
func NewLogFlushing(client client.Client) *LogFlushing {
	return &LogFlushing{client: client}
}

// GetSettingsV1 retrieves the current log flushing settings.
// URL: GET /api/v1/log-flushing
// https://developer.jamf.com/jamf-pro/reference/get_v1-log-flushing
func (s *LogFlushing) GetSettingsV1(ctx context.Context) (*ResourceLogFlushingSettings, *resty.Response, error) {
	endpoint := constants.EndpointJamfProLogFlushingV1

	var result ResourceLogFlushingSettings

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get log flushing settings: %w", err)
	}

	return &result, resp, nil
}

// ListTasksV1 retrieves all log flushing tasks.
// URL: GET /api/v1/log-flushing/task
// https://developer.jamf.com/jamf-pro/reference/get_v1-log-flushing-task
func (s *LogFlushing) ListTasksV1(ctx context.Context) ([]ResourceLogFlushingTask, *resty.Response, error) {
	endpoint := constants.EndpointJamfProLogFlushingV1 + "/task"

	var result []ResourceLogFlushingTask

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list log flushing tasks: %w", err)
	}

	return result, resp, nil
}

// GetTaskByIDV1 retrieves a specific log flushing task by ID.
// URL: GET /api/v1/log-flushing/task/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-log-flushing-task-id
func (s *LogFlushing) GetTaskByIDV1(ctx context.Context, id string) (*ResourceLogFlushingTask, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("task ID is required")
	}

	endpoint := fmt.Sprintf("%s/task/%s", constants.EndpointJamfProLogFlushingV1, id)

	var result ResourceLogFlushingTask

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get log flushing task %s: %w", id, err)
	}

	return &result, resp, nil
}

// QueueTaskV1 creates a new log flushing task.
// URL: POST /api/v1/log-flushing/task
// https://developer.jamf.com/jamf-pro/reference/post_v1-log-flushing-task
func (s *LogFlushing) QueueTaskV1(ctx context.Context, request *RequestLogFlushingTask) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("log flushing task request cannot be nil")
	}

	endpoint := constants.EndpointJamfProLogFlushingV1 + "/task"

	var result CreateResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to queue log flushing task: %w", err)
	}

	return &result, resp, nil
}

// DeleteTaskByIDV1 deletes a specific log flushing task by ID.
// URL: DELETE /api/v1/log-flushing/task/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-log-flushing-task-id
func (s *LogFlushing) DeleteTaskByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("task ID is required")
	}

	endpoint := fmt.Sprintf("%s/task/%s", constants.EndpointJamfProLogFlushingV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, fmt.Errorf("failed to delete log flushing task %s: %w", id, err)
	}

	return resp, nil
}
