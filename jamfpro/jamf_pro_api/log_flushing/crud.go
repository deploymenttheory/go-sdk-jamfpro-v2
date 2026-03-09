package log_flushing

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

// ServiceInterface defines the interface for log flushing operations.
//
// Log flushing allows administrators to manage Jamf Pro log retention and
// schedule log cleanup tasks to maintain system performance and comply with
// data retention policies.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-log-flushing
type ServiceInterface interface {
	// GetSettingsV1 retrieves the current log flushing settings.
	//
	// Returns retention policies for different log types and the scheduled
	// hour of day when automatic log flushing occurs.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-log-flushing
	GetSettingsV1(ctx context.Context) (*ResourceLogFlushingSettings, *resty.Response, error)

	// ListTasksV1 retrieves all log flushing tasks.
	//
	// Returns a list of queued, running, and completed log flushing tasks.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-log-flushing-task
	ListTasksV1(ctx context.Context) ([]ResourceLogFlushingTask, *resty.Response, error)

	// GetTaskByIDV1 retrieves a specific log flushing task by ID.
	//
	// Returns detailed information about a single log flushing task including
	// its current state and configuration.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-log-flushing-task-id
	GetTaskByIDV1(ctx context.Context, id string) (*ResourceLogFlushingTask, *resty.Response, error)

	// QueueTaskV1 creates a new log flushing task.
	//
	// Queues a task to flush logs matching the specified qualifier and retention
	// period. The task will be processed according to Jamf Pro's task scheduling.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-log-flushing-task
	QueueTaskV1(ctx context.Context, request *RequestLogFlushingTask) (*CreateResponse, *resty.Response, error)

	// DeleteTaskByIDV1 deletes a specific log flushing task by ID.
	//
	// Removes a queued or completed log flushing task. Running tasks may not
	// be deleted and will return an error.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-log-flushing-task-id
	DeleteTaskByIDV1(ctx context.Context, id string) (*resty.Response, error)
}

type (
	// Service handles communication with the log flushing-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-log-flushing
	LogFlushing struct {
		client transport.HTTPClient
	}
)

var _ ServiceInterface = (*LogFlushing)(nil)

// NewService creates a new log flushing service.
func NewLogFlushing(client transport.HTTPClient) *LogFlushing {
	return &LogFlushing{client: client}
}

// GetSettingsV1 retrieves the current log flushing settings.
// URL: GET /api/v1/log-flushing
// https://developer.jamf.com/jamf-pro/reference/get_v1-log-flushing
func (s *LogFlushing) GetSettingsV1(ctx context.Context) (*ResourceLogFlushingSettings, *resty.Response, error) {
	endpoint := EndpointLogFlushingV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	var result ResourceLogFlushingSettings
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get log flushing settings: %w", err)
	}

	return &result, resp, nil
}

// ListTasksV1 retrieves all log flushing tasks.
// URL: GET /api/v1/log-flushing/task
// https://developer.jamf.com/jamf-pro/reference/get_v1-log-flushing-task
func (s *LogFlushing) ListTasksV1(ctx context.Context) ([]ResourceLogFlushingTask, *resty.Response, error) {
	endpoint := EndpointLogFlushingV1 + "/task"

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	var result []ResourceLogFlushingTask
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
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

	endpoint := fmt.Sprintf("%s/task/%s", EndpointLogFlushingV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	var result ResourceLogFlushingTask
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
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

	endpoint := EndpointLogFlushingV1 + "/task"

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	var result CreateResponse
	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
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

	endpoint := fmt.Sprintf("%s/task/%s", EndpointLogFlushingV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to delete log flushing task %s: %w", id, err)
	}

	return resp, nil
}
